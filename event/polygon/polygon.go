package polygon

import (
	"context"
	"fmt"
	"math/big"
	"sync"
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
	loopInterval        time.Duration
	resourceCtx         context.Context
	resourceCancel      context.CancelFunc
	tasks               tasks.Group
	db                  *database.DB
	epoch               uint64
	eventStartBlock     uint64
	startHeight         *big.Int
	contracts           []string
	aliContentClient    *green.Client
	cfg                 *config.Config
	client              node.EthClient
	nativeCurrentHeight *big.Int
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
	pp.tasks.Go(func() error {
		for range tickerEventOn1.C {
			err := pp.onNativeToken()
			if err != nil {
				log.Warn("native token scan failed", "err", err)
				time.Sleep(time.Second * 3)
				continue
			}
		}
		return nil
	})
	return nil
}

// 记录当前扫描到的区块
var nativeScanLock sync.Mutex

func (pp *PolygonEventProcessor) onNativeToken() error {
	nativeScanLock.Lock()
	defer nativeScanLock.Unlock()

	// 第一次初始化 currentHeight
	if pp.nativeCurrentHeight == nil {
		pp.nativeCurrentHeight = big.NewInt(int64(pp.cfg.EventStartBlock))
		log.Info("init native scanner start block", "height", pp.nativeCurrentHeight)
	}

	// 每次 tick 只处理一个区块（防止卡死）
	height := new(big.Int).Set(pp.nativeCurrentHeight)

	block, err := pp.client.BlockByNumber(height)
	if err != nil {
		log.Warn("native: fetch block failed", "height", height, "err", err)
		return nil // 不退出任务
	}
	if block == nil {
		log.Warn("native: block is nil", "height", height)
		return nil
	}

	log.Info("native: success getting block", "height", height, "txs", len(block.Transactions()))

	for _, tx := range block.Transactions() {
		if tx.Value().Cmp(big.NewInt(0)) <= 0 {
			continue
		}

		signer := types.LatestSignerForChainID(tx.ChainId()) // 找出正确的签名规则
		from, err := types.Sender(signer, tx)                // 根据交易内容和签名恢复发送者地址
		if err != nil {
			log.Warn("native: failed to get sender", "txHash", tx.Hash(), "err", err)
			continue
		}

		to := "ContractCreation"
		if tx.To() != nil {
			to = tx.To().Hex()
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
			Description:  "Polygon Native Token Received",
			Timestamp:    uint64(block.Time()),
		}

		if err := pp.db.Transaction(func(txdb *database.DB) error {
			if err := txdb.TokenSentDB.StoreTokenSent(tokenSent); err != nil {
				return err
			}
			if err := txdb.TokenReceivedDB.StoreTokenReceived(tokenReceived); err != nil {
				return err
			}
			return nil
		}); err != nil {
			log.Error("native: store transfer fail", "err", err)
		}
	}

	// 完成一个区块后 +1
	pp.nativeCurrentHeight = new(big.Int).Add(pp.nativeCurrentHeight, big.NewInt(1))

	log.Info("native: finished block", "nextHeight", pp.nativeCurrentHeight)

	return nil
}

// func (pp *PolygonEventProcessor) onNativeToken(startBlock uint64) error {
// 	//  再处理原生代币转账
// 	for height := new(big.Int).Set(fromHeight); height.Cmp(toHeight) <= 0; height.Add(height, big.NewInt(1)) {
// 		var block *types.Block

// 		for retry := 0; retry < 3; retry++ {
// 			block1, err := pp.client.BlockByNumber(height)
// 			if err != nil {
// 				log.Warn("fetch block failed", "height", height, "retry", retry, "err", err)
// 				time.Sleep(time.Second * 2)
// 				continue
// 			}
// 			if block1 == nil {
// 				log.Warn("block is nil, will retry", "height", height, "retry", retry)
// 				time.Sleep(time.Second)
// 				continue
// 			}
// 			block = block1
// 			log.Info("Success getting whole block")
// 			break // 成功拿到 block
// 		}
// 		if block == nil {
// 			log.Error("skip block after 3 retries", "height", height)
// 			// 可以记录到 missed_blocks 表里，下次同步补偿
// 			continue
// 		}

// 		for _, tx := range block.Transactions() {
// 			if tx.Value().Cmp(big.NewInt(0)) > 0 {

// 				signer := types.LatestSignerForChainID(tx.ChainId()) // 找出正确的签名规则

// 				from, err := types.Sender(signer, tx) // 根据交易内容和签名恢复发送者地址
// 				if err != nil {
// 					log.Warn("failed to get sender", "txHash", tx.Hash(), "err", err)
// 					continue
// 				}

// 				to := ""
// 				if tx.To() != nil {
// 					to = tx.To().Hex()
// 				} else {
// 					to = "ContractCreation"
// 				}

// 				tokenSent := token_transfer.TokenSent{
// 					Address:      from.Hex(),
// 					TokenAddress: "0x0000000000000000000000000000000000001010", // Polygon 原生代币地址
// 					Amount:       tx.Value(),
// 					Description:  "Polygon Native Token Transfer",
// 					Timestamp:    uint64(block.Time()),
// 				}

// 				tokenReceived := token_transfer.TokenReceived{
// 					Address:      to,
// 					TokenAddress: "0x0000000000000000000000000000000000001010", // Polygon 原生代币地址
// 					Amount:       tx.Value(),
// 					Description:  "Polygon Native Token Received",
// 					Timestamp:    uint64(block.Time()),
// 				}

// 				if err := pp.db.Transaction(func(tx *database.DB) error {

// 					if err := tx.TokenSentDB.StoreTokenSent(tokenSent); err != nil {
// 						return err
// 					}
// 					if err := tx.TokenReceivedDB.StoreTokenReceived(tokenReceived); err != nil {
// 						return err
// 					}
// 					return nil
// 				}); err != nil {
// 					return err
// 				}
// 			}
// 		}
// 	}
// }

func (pp *PolygonEventProcessor) onData() error {
	log.Info("onData triggered tick") // 第一行，证明 for-loop 触发了
	defer log.Info("onData exited", "startHeight", pp.startHeight)

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
		if pp.startHeight.Cmp(big.NewInt(0)) == 0 {
			pp.startHeight = big.NewInt(int64(pp.cfg.StartBlock))
		}

		log.Info("config StartBlock", "StartBlock", pp.startHeight)
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
	log.Info("LatestBlockHeader got: ", "LatestBlockHeader", pp.startHeight)
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
		eventsFetchErr := pp.eventsFetch(fromHeight, toHeight, tx) // *** 传入事务tx ***
		if eventsFetchErr != nil {
			return eventsFetchErr
		}
		lastBlock := block_listener.BlockListener{
			BlockNumber: toHeight,
		}
		updateErr := tx.BlockListener.SaveOrUpdateLastBlockNumber(lastBlock)
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

func (pp *PolygonEventProcessor) eventsFetch(fromHeight, toHeight *big.Int, tx *database.DB) error {

	//  先处理合约事件
	contracts := pp.contracts
	for _, contract := range contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(contract)}
		events, err := pp.db.ContractEvent.ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
		if err != nil {
			log.Info("failed to index ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := pp.eventUnpack(contractEvent, tx)
			if unpackErr != nil {
				log.Info("failed to index events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}

	return nil
}

func (pp *PolygonEventProcessor) eventUnpack(event event.ContractEvent, tx *database.DB) error {

	transferSig := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	if event.EventSignature == transferSig {
		address := event.RLPLog.Address.Hex()
		if address != "0x84eBc138F4Ab844A3050a6059763D269dC9951c6" && address != "0xc2132D05D31c914a87C6611C10748AEb04B58e8F" {
			log.Info("skipping transfer event for non fcc/USDT", "address", address)
			return nil
		} // 筛选出 FCC 和 USDT 合约地址的 Transfer 事件
		err := unpack.Transfer(event, tx, address)
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
		err := unpack.ActivityAdd(event, tx)
		return err
	case merchantAbi.Events["ActivityFinish"].ID.String():
		err := unpack.ActivityFinish(event, tx)
		return err
	case nftTokenAbi.Events["CreateNFT"].ID.String():
		err := unpack.MintNft(event, tx)
		return err

	case nftTokenAbi.Events["MintBoosterNFT"].ID.String():
		err := unpack.MintBoosterNft(event, tx)
		return err

	case merchantAbi.Events["Drop"].ID.String():
		err := unpack.Drop(event, tx)
		return err

	case stakeAbi.Events["StakeHolderDepositStaking"].ID.String():
		err := unpack.StakeHolderDepositStaking(event, tx)
		return err

	case stakeAbi.Events["StakeHolderWithdrawStaking"].ID.String():
		err := unpack.StakeHolderWithdrawStaking(event, tx)
		return err
	}
	return nil
}
