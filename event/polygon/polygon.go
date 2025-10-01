package polygon

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/green"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/FishcakeLab/fishcake-service/common/tasks"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/block_listener"
	"github.com/FishcakeLab/fishcake-service/database/event"
	"github.com/FishcakeLab/fishcake-service/database/token_transfer"
	"github.com/FishcakeLab/fishcake-service/event/polygon/abi"
	"github.com/FishcakeLab/fishcake-service/event/polygon/unpack"
	"github.com/FishcakeLab/fishcake-service/synchronizer/node"
)

type PolygonEventProcessor struct {
	loopInterval     time.Duration
	resourceCtx      context.Context
	resourceCancel   context.CancelFunc
	tasks            tasks.Group
	db               *database.DB
	epoch            uint64
	eventStartBlock  uint64
	startHeight      *big.Int
	contracts        []string
	aliContentClient *green.Client
	cfg              *config.Config
	client           node.EthClient
}

func NewEventProcessor(db *database.DB, cfg *config.Config, client node.EthClient, loopInterval time.Duration, epoch uint64, shutdown context.CancelCauseFunc,
) (*PolygonEventProcessor, error) {

	var contracts []string = cfg.Contracts

	eventStartBlock := cfg.EventStartBlock
	// aliConfig := cfg.AliConfig

	resCtx, resCancel := context.WithCancel(context.Background())
	//aliContentClient, createGreenClientErr := green.NewClientWithAccessKey(
	//	aliConfig.RegionId,
	//	aliConfig.AccessKeyId,
	//	aliConfig.AccessKeySecret)
	//if createGreenClientErr != nil {
	//	log.Info("failed to create green client", "err", createGreenClientErr)
	//	// handle exceptions
	//	panic(createGreenClientErr)
	//}
	return &PolygonEventProcessor{
		db:             db,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error processor: %w", err))
		}},
		startHeight:      nil,
		eventStartBlock:  eventStartBlock,
		contracts:        contracts,
		loopInterval:     loopInterval,
		epoch:            epoch,
		aliContentClient: nil,
		cfg:              cfg,
		client:           client,
	}, nil
}

// func NewEventProcessor1(db *database.DB, loopInterval time.Duration, contracts []string,
// 	startHeight uint64, eventStartBlock uint64, epoch uint64, shutdown context.CancelCauseFunc,
// 	aliConfig config.AliConfig) (*PolygonEventProcessor, error) {
// 	resCtx, resCancel := context.WithCancel(context.Background())
// 	//aliContentClient, createGreenClientErr := green.NewClientWithAccessKey(
// 	//	aliConfig.RegionId,
// 	//	aliConfig.AccessKeyId,
// 	//	aliConfig.AccessKeySecret)
// 	//if createGreenClientErr != nil {
// 	//	log.Info("failed to create green client", "err", createGreenClientErr)
// 	//	// handle exceptions
// 	//	panic(createGreenClientErr)
// 	//}
// 	return &PolygonEventProcessor{
// 		db:             db,
// 		resourceCtx:    resCtx,
// 		resourceCancel: resCancel,
// 		tasks: tasks.Group{HandleCrit: func(err error) {
// 			shutdown(fmt.Errorf("critical error processor: %w", err))
// 		}},
// 		startHeight:      new(big.Int).SetUint64(startHeight),
// 		eventStartBlock:  eventStartBlock,
// 		contracts:        contracts,
// 		loopInterval:     loopInterval,
// 		epoch:            epoch,
// 		aliContentClient: nil,
// 	}, nil
// }

func (pp *PolygonEventProcessor) Start() error {
	tickerEventOn1 := time.NewTicker(pp.loopInterval)
	log.Info("starting polygon bridge processor...")
	pp.tasks.Go(func() error {
		for range tickerEventOn1.C {
			err := pp.onData()
			if err != nil {
				log.Info("no more l1 etl updates. shutting down l1 task")
				continue
			}
		}
		return nil
	})
	return nil
}

func (pp *PolygonEventProcessor) onData() error {
	if pp.startHeight == nil { // 如果设置为空
		lastListenBlock, err := pp.db.BlockListener.GetLastBlockNumber() // 获取上次监听的区块高度
		if err != nil {
			log.Info("failed to get last block heard", "err", err)
			return err
		}

		if lastListenBlock == nil {
			lastListenBlock = &block_listener.BlockListener{
				BlockNumber: big.NewInt(0),
			}
		}
		log.Info("last Listen block", "lastListenBlock", lastListenBlock.BlockNumber)

		pp.startHeight = lastListenBlock.BlockNumber // 把起始高度设置为上次监听的区块高度
		if pp.startHeight == big.NewInt(0) {
			pp.startHeight = big.NewInt(int64(pp.cfg.StartBlock))
		}
		// 如果起始高度（上次监听的区块高度）小于配置的事件起始高度，则设置为配置的起始高度
		if pp.startHeight.Cmp(big.NewInt(int64(pp.eventStartBlock))) == -1 {
			pp.startHeight = big.NewInt(int64(pp.eventStartBlock))
		}
	} else {
		// 如果配置不为空，直接+1
		pp.startHeight = new(big.Int).Add(pp.startHeight, bigint.One)
		log.Info("pp.start height", "pp.startHeight", pp.startHeight)

	}
	fromHeight := pp.startHeight
	toHeight := new(big.Int).Add(fromHeight, big.NewInt(int64(pp.epoch)))

	log.Info("Handle event start and end block", "start", fromHeight, "end", toHeight)

	latestBlockHeader, err := pp.db.Blocks.LatestBlockHeader() // 获取数据库拉取的最新的区块头
	if err != nil {
		pp.startHeight = new(big.Int).Sub(pp.startHeight, bigint.One) // 如果数据库获取失败，回退1
		return err
	}
	if latestBlockHeader == nil {
		pp.startHeight = new(big.Int).Sub(pp.startHeight, bigint.One) // 如果没有最新区块头，回退1
		return nil
	}
	if latestBlockHeader.Number.Cmp(fromHeight) == -1 { // 如果数据库的最新区块头小于起始高度，起始高度和结束高度都设置为最新区块头
		fromHeight = latestBlockHeader.Number
		toHeight = latestBlockHeader.Number
	} else {
		if latestBlockHeader.Number.Cmp(toHeight) == -1 { // 如果数据库的最新区块头处于起始高度和结束高度之间，则把结束高度设置为最新区块头
			toHeight = latestBlockHeader.Number
		}
	}
	if err := pp.db.Transaction(func(tx *database.DB) error {
		eventsFetchErr := pp.eventsFetch(fromHeight, toHeight)
		if eventsFetchErr != nil {
			return eventsFetchErr
		}
		lastBlock := block_listener.BlockListener{
			BlockNumber: toHeight,
		}
		updateErr := pp.db.BlockListener.SaveOrUpdateLastBlockNumber(lastBlock)
		if updateErr != nil {
			log.Info("update last block err :", updateErr)
			return updateErr
		}
		return nil
	}); err != nil {
		pp.startHeight = new(big.Int).Sub(pp.startHeight, bigint.One)
		return err
	}
	pp.startHeight = toHeight
	return nil
}

func (pp *PolygonEventProcessor) eventsFetch(fromHeight, toHeight *big.Int) error {
	contracts := pp.contracts
	for _, contract := range contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(contract)}
		events, err := pp.db.ContractEvent.ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
		if err != nil {
			log.Info("failed to index ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := pp.eventUnpack(contractEvent)
			if unpackErr != nil {
				log.Info("failed to index events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}

	// 2. 再处理原生代币转账
	for height := new(big.Int).Set(fromHeight); height.Cmp(toHeight) <= 0; height.Add(height, big.NewInt(1)) {
		block, err := pp.client.BlockByNumber(height)
		if err != nil {
			log.Error("failed to fetch block", "height", height, "err", err)
			return err
		}

		for _, tx := range block.Transactions() {
			if tx.Value().Cmp(big.NewInt(0)) > 0 {

				signer := types.LatestSignerForChainID(tx.ChainId())

				from, err := types.Sender(signer, tx)
				if err != nil {
					log.Warn("failed to get sender", "txHash", tx.Hash(), "err", err)
					continue
				}

				to := ""
				if tx.To() != nil {
					to = tx.To().Hex()
				} else {
					to = "ContractCreation"
				}

				tokenSent := token_transfer.TokenSent{
					Address:      from.Hex(),
					TokenAddress: "0x0000000000000000000000000000000000001010", // Polygon 原生代币地址
					Amount:       tx.Value(),
					Description:  "Polygon Native Token Transfer",
					Timestamp:    uint64(block.Time()),
				}

				tokenReceived := token_transfer.TokenReceived{
					Address:      to,
					TokenAddress: "0x0000000000000000000000000000000000001010", // Polygon 原生代币地址
					Amount:       tx.Value(),
					Description:  "Polygon Native Token Transfer",
					Timestamp:    uint64(block.Time()),
				}

				if err := pp.db.Transaction(func(tx *database.DB) error {

					if err := pp.db.TokenSentDB.StoreTokenSent(tokenSent); err != nil {
						return err
					}
					if err := pp.db.TokenReceivedDB.StoreTokenReceived(tokenReceived); err != nil {
						return err
					}
					return nil
				}); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (pp *PolygonEventProcessor) eventUnpack(event event.ContractEvent) error {

	transferSig := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	if event.EventSignature == transferSig {
		address := event.RLPLog.Address.Hex()
		if address != "0x84eBc138F4Ab844A3050a6059763D269dC9951c6" && address != "0xc2132D05D31c914a87C6611C10748AEb04B58e8F" {
			log.Info("skipping transfer event for non fcc/USDT", "address", address)
			return nil
		}
		err := unpack.Transfer(event, pp.db, address)
		if err != nil {
			log.Info("failed to unpack transfer event", "err", err)
			return err
		}
		return nil
	}

	merchantAbi, _ := abi.FishcakeEventManagerMetaData.GetAbi()
	nftTokenAbi, _ := abi.NftManagerMetaData.GetAbi()
	stakeAbi, _ := abi.StakingManagerMetaData.GetAbi()
	switch event.EventSignature.String() {
	case merchantAbi.Events["ActivityAdd"].ID.String():
		err := unpack.ActivityAdd(event, pp.db)
		return err
	case merchantAbi.Events["ActivityFinish"].ID.String():
		err := unpack.ActivityFinish(event, pp.db)
		return err
	case nftTokenAbi.Events["CreateNFT"].ID.String():
		err := unpack.MintNft(event, pp.db)
		return err
	case merchantAbi.Events["Drop"].ID.String():
		err := unpack.Drop(event, pp.db)
		return err

	case stakeAbi.Events["StakeHolderDepositStaking"].ID.String():
		err := unpack.StakeHolderDepositStaking(event, pp.db)
		return err

	case stakeAbi.Events["StakeHolderWithdrawStaking"].ID.String():
		err := unpack.StakeHolderWithdrawStaking(event, pp.db)
		return err
	}
	return nil
}
