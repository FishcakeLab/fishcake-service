package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	dbEvent "github.com/FishcakeLab/fishcake-service/database/event"
	"github.com/FishcakeLab/fishcake-service/database/mining_record"
	"github.com/FishcakeLab/fishcake-service/event/polygon/abi"
	"github.com/FishcakeLab/fishcake-service/event/polygon/unpack"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repairStats struct {
	Batches               int
	EventsSeen            int
	ActivityFinishRecords int
	MintBoosterRecords    int
}

func (s *repairStats) Add(other repairStats) {
	s.Batches += other.Batches
	s.EventsSeen += other.EventsSeen
	s.ActivityFinishRecords += other.ActivityFinishRecords
	s.MintBoosterRecords += other.MintBoosterRecords
}

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))

	configPath := flag.String("config", "./config.yaml", "Path to config file")
	fromBlock := flag.Uint64("from-block", 0, "Start block for replay; defaults to event_start_block or start_block")
	toBlock := flag.Uint64("to-block", 0, "End block for replay; defaults to latest indexed block")
	batchSize := flag.Uint64("batch-size", 5000, "Block batch size")
	apply := flag.Bool("apply", false, "Write changes to the database")
	reset := flag.Bool("reset", false, "Delete mining_record before replay")
	flag.Parse()

	if *batchSize == 0 {
		log.Error("batch-size must be > 0")
		os.Exit(1)
	}

	cfg, err := config.New(*configPath)
	if err != nil {
		log.Error("failed to load config", "err", err)
		os.Exit(1)
	}

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Error("failed to connect database", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	rawDB, err := openRawDB(cfg)
	if err != nil {
		log.Error("failed to connect raw database", "err", err)
		os.Exit(1)
	}

	merchantAddress := strings.TrimSpace(cfg.ContractInfo.FishcakeEventManager)
	if merchantAddress == "" {
		log.Error("missing contract_info.fishcake_event_manager in config")
		os.Exit(1)
	}

	nftManagerAddress := strings.TrimSpace(cfg.ContractInfo.NFTManager)
	if nftManagerAddress == "" {
		log.Error("missing contract_info.nft_manager in config")
		os.Exit(1)
	}

	resolvedFrom, resolvedTo, err := resolveBlockRange(db, cfg, *fromBlock, *toBlock)
	if err != nil {
		log.Error("failed to resolve block range", "err", err)
		os.Exit(1)
	}

	if *apply && *reset {
		log.Info("resetting mining_record table before replay")
		if err := resetMiningRecordTable(rawDB); err != nil {
			log.Error("failed to reset mining_record table", "err", err)
			os.Exit(1)
		}
	}

	stats := repairStats{}
	for start := resolvedFrom; start <= resolvedTo; {
		end := start + *batchSize - 1
		if end > resolvedTo {
			end = resolvedTo
		}

		batchStats, err := replayBatch(rawDB, db, merchantAddress, nftManagerAddress, start, end, *apply)
		if err != nil {
			log.Error("replay batch failed", "fromBlock", start, "toBlock", end, "err", err)
			os.Exit(1)
		}
		batchStats.Batches = 1
		stats.Add(batchStats)

		log.Info(
			"repair batch complete",
			"fromBlock", start,
			"toBlock", end,
			"eventsSeen", batchStats.EventsSeen,
			"activityFinishRecords", batchStats.ActivityFinishRecords,
			"mintBoosterRecords", batchStats.MintBoosterRecords,
			"apply", *apply,
		)

		start = end + 1
	}

	mode := "DRY-RUN"
	if *apply {
		mode = "APPLY"
	}

	log.Info(
		"repair mining record complete",
		"mode", mode,
		"fromBlock", resolvedFrom,
		"toBlock", resolvedTo,
		"batches", stats.Batches,
		"eventsSeen", stats.EventsSeen,
		"activityFinishRecords", stats.ActivityFinishRecords,
		"mintBoosterRecords", stats.MintBoosterRecords,
	)
}

func resolveBlockRange(db *database.DB, cfg *config.Config, fromFlag, toFlag uint64) (uint64, uint64, error) {
	from := fromFlag
	if from == 0 {
		if cfg.EventStartBlock != 0 {
			from = cfg.EventStartBlock
		} else {
			from = cfg.StartBlock
		}
	}

	to := toFlag
	if to == 0 {
		latestHeader, err := db.Blocks.LatestBlockHeader()
		if err != nil {
			return 0, 0, err
		}
		if latestHeader == nil || latestHeader.Number == nil {
			return 0, 0, fmt.Errorf("latest indexed block header not found")
		}
		to = latestHeader.Number.Uint64()
	}

	if from > to {
		return 0, 0, fmt.Errorf("from-block %d is greater than to-block %d", from, to)
	}

	return from, to, nil
}

func openRawDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", cfg.DbHost, cfg.DbName)
	if cfg.DbPort != 0 {
		dsn += fmt.Sprintf(" port=%d", cfg.DbPort)
	}
	if cfg.DbUser != "" {
		dsn += fmt.Sprintf(" user=%s", cfg.DbUser)
	}
	if cfg.DbPassword != "" {
		dsn += fmt.Sprintf(" password=%s", cfg.DbPassword)
	}

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func resetMiningRecordTable(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM mining_record").Error; err != nil {
			return fmt.Errorf("delete mining_record: %w", err)
		}
		return nil
	})
}

func replayBatch(rawDB *gorm.DB, db *database.DB, merchantAddress, nftManagerAddress string, fromBlock, toBlock uint64, apply bool) (repairStats, error) {
	stats := repairStats{}

	merchantEvents, err := fetchContractEventsDirect(rawDB, merchantAddress, fromBlock, toBlock)
	if err != nil {
		return stats, err
	}

	for _, event := range merchantEvents {
		handled, err := replayMerchantEvent(db, event, apply, &stats)
		if err != nil {
			return stats, err
		}
		if handled {
			stats.EventsSeen++
		}
	}

	nftEvents, err := fetchContractEventsDirect(rawDB, nftManagerAddress, fromBlock, toBlock)
	if err != nil {
		return stats, err
	}

	for _, event := range nftEvents {
		handled, err := replayNftEvent(db, event, apply, &stats)
		if err != nil {
			return stats, err
		}
		if handled {
			stats.EventsSeen++
		}
	}

	return stats, nil
}

func fetchContractEventsDirect(rawDB *gorm.DB, contractAddress string, fromBlock, toBlock uint64) ([]dbEvent.ContractEvent, error) {
	var events []dbEvent.ContractEvent

	err := rawDB.Table("contract_events").
		Where("contract_address ILIKE ?", common.HexToAddress(contractAddress).Hex()).
		Where("block_number >= ? AND block_number <= ?", new(big.Int).SetUint64(fromBlock), new(big.Int).SetUint64(toBlock)).
		Order("block_number ASC, log_index ASC").
		Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func replayMerchantEvent(db *database.DB, event dbEvent.ContractEvent, apply bool, stats *repairStats) (bool, error) {
	merchantAbi, _ := abi.FishcakeEventManagerMetaData.GetAbi()

	if event.EventSignature.String() != merchantAbi.Events["ActivityFinish"].ID.String() {
		return false, nil
	}

	if apply {
		if err := storeActivityFinishRecord(db, event); err != nil {
			return false, err
		}
	}
	stats.ActivityFinishRecords++
	return true, nil
}

func replayNftEvent(db *database.DB, event dbEvent.ContractEvent, apply bool, stats *repairStats) (bool, error) {
	nftAbi, _ := abi.NftManagerMetaData.GetAbi()

	if event.EventSignature.String() != nftAbi.Events["MintBoosterNFT"].ID.String() {
		return false, nil
	}

	if apply {
		if err := storeMintBoosterRecord(db, event); err != nil {
			return false, err
		}
	}
	stats.MintBoosterRecords++
	return true, nil
}

func storeActivityFinishRecord(db *database.DB, event dbEvent.ContractEvent) error {
	uEvent, err := unpack.MerchantUnpack.ParseActivityFinish(*event.RLPLog)
	if err != nil {
		return fmt.Errorf("parse ActivityFinish tx=%s logIndex=%d: %w", event.TransactionHash.Hex(), event.RLPLog.Index, err)
	}

	activityInfo := db.ActivityInfoDB.ActivityInfo(int(uEvent.ActivityId.Int64()))
	activityID := uEvent.ActivityId.Int64()
	record := mining_record.MiningRecord{
		Address:          activityInfo.BusinessAccount,
		RecordType:       mining_record.RecordTypeActivityFinish,
		MinedAmountDelta: new(big.Int).Set(uEvent.MinedAmount),
		PowerIncrease:    new(big.Int).Set(uEvent.MinedAmount),
		PowerDecrease:    big.NewInt(0),
		ActivityID:       &activityID,
		Description:      "ActivityFinish reward",
		Timestamp:        event.Timestamp,
		TxHash:           event.TransactionHash.Hex(),
		LogIndex:         uint(event.RLPLog.Index),
		BlockNumber:      event.BlockNumber.Uint64(),
	}

	return db.MiningRecordDB.Store(record)
}

func storeMintBoosterRecord(db *database.DB, event dbEvent.ContractEvent) error {
	uEvent, err := unpack.NftTokenUnpack.ParseMintBoosterNFT(*event.RLPLog)
	if err != nil {
		return fmt.Errorf("parse MintBoosterNFT tx=%s logIndex=%d: %w", event.TransactionHash.Hex(), event.RLPLog.Index, err)
	}

	address := strings.ToLower(uEvent.Miner.Hex())
	tokenID := uEvent.TokenId.Int64()
	record := mining_record.MiningRecord{
		Address:          address,
		RecordType:       mining_record.RecordTypeMintBoosterNFT,
		MinedAmountDelta: big.NewInt(0),
		PowerIncrease:    big.NewInt(0),
		PowerDecrease:    new(big.Int).Set(uEvent.UsedFishCakePower),
		TokenID:          &tokenID,
		Description:      "MintBoosterNFT power spent",
		Timestamp:        uint64(event.Timestamp),
		TxHash:           event.TransactionHash.Hex(),
		LogIndex:         uint(event.RLPLog.Index),
		BlockNumber:      event.BlockNumber.Uint64(),
	}

	return db.MiningRecordDB.Store(record)
}
