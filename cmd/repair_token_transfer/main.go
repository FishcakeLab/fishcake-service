package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	dbEvent "github.com/FishcakeLab/fishcake-service/database/event"
	"github.com/FishcakeLab/fishcake-service/database/token_transfer"
	"github.com/FishcakeLab/fishcake-service/event/polygon/abi"
	"github.com/FishcakeLab/fishcake-service/event/polygon/unpack"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repairStats struct {
	Batches                int
	EventsSeen             int
	ActivityAddSent        int
	ActivityFinishReceived int
	DropReceived           int
	FccTransferSent        int
	FccTransferReceived    int
}

func (s *repairStats) Add(other repairStats) {
	s.Batches += other.Batches
	s.EventsSeen += other.EventsSeen
	s.ActivityAddSent += other.ActivityAddSent
	s.ActivityFinishReceived += other.ActivityFinishReceived
	s.DropReceived += other.DropReceived
	s.FccTransferSent += other.FccTransferSent
	s.FccTransferReceived += other.FccTransferReceived
}

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))

	configPath := flag.String("config", "./config.yaml", "Path to config file")
	fromBlock := flag.Uint64("from-block", 0, "Start block for replay; defaults to event_start_block or start_block")
	toBlock := flag.Uint64("to-block", 0, "End block for replay; defaults to latest indexed block")
	batchSize := flag.Uint64("batch-size", 5000, "Block batch size")
	apply := flag.Bool("apply", false, "Write changes to the database")
	reset := flag.Bool("reset", false, "Delete token_sent and token_received before replay")
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

	fccAddress := strings.TrimSpace(cfg.FCC)
	if fccAddress == "" {
		log.Error("missing fcc in config")
		os.Exit(1)
	}

	resolvedFrom, resolvedTo, err := resolveBlockRange(db, cfg, *fromBlock, *toBlock)
	if err != nil {
		log.Error("failed to resolve block range", "err", err)
		os.Exit(1)
	}

	if *apply && *reset {
		log.Info("resetting token transfer tables before replay")
		if err := resetTokenTransferTables(rawDB); err != nil {
			log.Error("failed to reset token transfer tables", "err", err)
			os.Exit(1)
		}
	}

	stats := repairStats{}
	for start := resolvedFrom; start <= resolvedTo; {
		end := start + *batchSize - 1
		if end > resolvedTo {
			end = resolvedTo
		}

		batchStats, err := replayBatch(rawDB, db, merchantAddress, fccAddress, start, end, *apply)
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
			"activityAddSent", batchStats.ActivityAddSent,
			"activityFinishReceived", batchStats.ActivityFinishReceived,
			"dropReceived", batchStats.DropReceived,
			"fccTransferSent", batchStats.FccTransferSent,
			"fccTransferReceived", batchStats.FccTransferReceived,
			"apply", *apply,
		)

		start = end + 1
	}

	mode := "DRY-RUN"
	if *apply {
		mode = "APPLY"
	}

	log.Info(
		"repair token transfer complete",
		"mode", mode,
		"fromBlock", resolvedFrom,
		"toBlock", resolvedTo,
		"batches", stats.Batches,
		"eventsSeen", stats.EventsSeen,
		"activityAddSent", stats.ActivityAddSent,
		"activityFinishReceived", stats.ActivityFinishReceived,
		"dropReceived", stats.DropReceived,
		"fccTransferSent", stats.FccTransferSent,
		"fccTransferReceived", stats.FccTransferReceived,
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

func resetTokenTransferTables(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM token_sent").Error; err != nil {
			return fmt.Errorf("delete token_sent: %w", err)
		}
		if err := tx.Exec("DELETE FROM token_received").Error; err != nil {
			return fmt.Errorf("delete token_received: %w", err)
		}
		return nil
	})
}

func replayBatch(rawDB *gorm.DB, db *database.DB, merchantAddress, fccAddress string, fromBlock, toBlock uint64, apply bool) (repairStats, error) {
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

	fccEvents, err := fetchContractEventsDirect(rawDB, fccAddress, fromBlock, toBlock)
	if err != nil {
		return stats, err
	}

	for _, event := range fccEvents {
		handled, err := replayFccTransferEvent(db, event, fccAddress, apply, &stats)
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

	switch event.EventSignature.String() {
	case merchantAbi.Events["ActivityAdd"].ID.String():
		if apply {
			if err := storeActivityAddTokenSent(db, event); err != nil {
				return false, err
			}
		}
		stats.ActivityAddSent++
		return true, nil
	case merchantAbi.Events["ActivityFinish"].ID.String():
		if apply {
			if err := storeActivityFinishTokenReceived(db, event); err != nil {
				return false, err
			}
		}
		stats.ActivityFinishReceived++
		return true, nil
	case merchantAbi.Events["Drop"].ID.String():
		if apply {
			if err := storeDropTokenReceived(db, event); err != nil {
				return false, err
			}
		}
		stats.DropReceived++
		return true, nil
	default:
		return false, nil
	}
}

func replayFccTransferEvent(db *database.DB, event dbEvent.ContractEvent, fccAddress string, apply bool, stats *repairStats) (bool, error) {
	transferSig := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	if event.EventSignature != transferSig {
		return false, nil
	}

	if apply {
		if err := storeFccTransfer(db, event, fccAddress); err != nil {
			return false, err
		}
	}
	stats.FccTransferSent++
	stats.FccTransferReceived++
	return true, nil
}

func storeActivityAddTokenSent(db *database.DB, event dbEvent.ContractEvent) error {
	uEvent, err := unpack.MerchantUnpack.ParseActivityAdd(*event.RLPLog)
	if err != nil {
		return fmt.Errorf("parse ActivityAdd tx=%s logIndex=%d: %w", event.TransactionHash.Hex(), event.RLPLog.Index, err)
	}

	record := token_transfer.TokenSent{
		Address:      uEvent.Who.String(),
		TokenAddress: uEvent.TokenContractAddr.String(),
		Amount:       uEvent.TotalDropAmts,
		Description:  fmt.Sprintf("ActivityAdd: %s", uEvent.ActivityContent),
		Timestamp:    event.Timestamp,
		TxHash:       event.TransactionHash.Hex(),
		LogIndex:     uint(event.RLPLog.Index),
	}

	return db.TokenSentDB.StoreTokenSent(record)
}

func storeActivityFinishTokenReceived(db *database.DB, event dbEvent.ContractEvent) error {
	uEvent, err := unpack.MerchantUnpack.ParseActivityFinish(*event.RLPLog)
	if err != nil {
		return fmt.Errorf("parse ActivityFinish tx=%s logIndex=%d: %w", event.TransactionHash.Hex(), event.RLPLog.Index, err)
	}

	activityInfo := db.ActivityInfoDB.ActivityInfo(int(uEvent.ActivityId.Int64()))
	record := token_transfer.TokenReceived{
		Address:      activityInfo.BusinessAccount,
		TokenAddress: uEvent.TokenContractAddr.String(),
		Amount:       uEvent.ReturnAmount,
		Description:  fmt.Sprintf("ActivityFinish Return: %s", activityInfo.ActivityContent),
		Timestamp:    event.Timestamp,
		TxHash:       event.TransactionHash.Hex(),
		LogIndex:     uint(event.RLPLog.Index),
	}

	return db.TokenReceivedDB.StoreTokenReceived(record)
}

func storeDropTokenReceived(db *database.DB, event dbEvent.ContractEvent) error {
	uEvent, err := unpack.MerchantUnpack.ParseDrop(*event.RLPLog)
	if err != nil {
		return fmt.Errorf("parse Drop tx=%s logIndex=%d: %w", event.TransactionHash.Hex(), event.RLPLog.Index, err)
	}

	activityInfo := db.ActivityInfoDB.ActivityInfo(int(uEvent.ActivityId.Int64()))
	record := token_transfer.TokenReceived{
		Address:      uEvent.Who.String(),
		TokenAddress: activityInfo.TokenContractAddr,
		Amount:       uEvent.DropAmt,
		Description:  fmt.Sprintf("Drop Receive: %s", activityInfo.ActivityContent),
		Timestamp:    event.Timestamp,
		TxHash:       event.TransactionHash.Hex(),
		LogIndex:     uint(event.RLPLog.Index),
	}

	return db.TokenReceivedDB.StoreTokenReceived(record)
}

func storeFccTransfer(db *database.DB, event dbEvent.ContractEvent, fccAddress string) error {
	from := event.RLPLog.Topics[1].Hex()
	to := event.RLPLog.Topics[2].Hex()
	value := new(big.Int).SetBytes(event.RLPLog.Data)

	fishcakeEventManager := strings.ToLower("0x2CAf752814f244b3778e30c27051cc6B45CB1fc9")
	if strings.ToLower(from) == fishcakeEventManager || strings.ToLower(to) == fishcakeEventManager {
		return nil
	}

	tokenSent := token_transfer.TokenSent{
		Address:      from,
		TokenAddress: fccAddress,
		Amount:       value,
		Description:  "FCC Transfer Sent",
		Timestamp:    event.Timestamp,
		TxHash:       event.TransactionHash.Hex(),
		LogIndex:     uint(event.RLPLog.Index),
	}
	if err := db.TokenSentDB.StoreTokenSent(tokenSent); err != nil {
		return err
	}

	tokenReceived := token_transfer.TokenReceived{
		Address:      to,
		TokenAddress: fccAddress,
		Amount:       value,
		Description:  "FCC Transfer Received",
		Timestamp:    event.Timestamp,
		TxHash:       event.TransactionHash.Hex(),
		LogIndex:     uint(event.RLPLog.Index),
	}
	return db.TokenReceivedDB.StoreTokenReceived(tokenReceived)
}
