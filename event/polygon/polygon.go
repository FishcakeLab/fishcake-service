package polygon

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/green"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/FishcakeLab/fishcake-service/common/tasks"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/block_listener"
	"github.com/FishcakeLab/fishcake-service/database/event"
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
	log.Info("polygon event processor contracts:", "contracts", contracts, " length of contracts:", len(contracts))

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

	log.Info("polygon event processor created", "eventStartBlock", eventStartBlock)

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

func (pp *PolygonEventProcessor) Start() error {
	log.Info("starting polygon bridge processor...")
	pp.tasks.Go(func() error {
		ticker := time.NewTicker(pp.loopInterval)
		defer ticker.Stop()
		for range ticker.C {
			err := pp.onData()
			if err != nil {
				log.Info("no more l1 etl updates. skipping this tick")
				continue
			}
		}
		return nil
	})

	return nil
}

func (pp *PolygonEventProcessor) onData() error {
	log.Info("onData triggered tick") // 第一行，证明 for-loop 触发了
	defer log.Info("onData exited", "startHeight", pp.startHeight)

	if pp.startHeight == nil { // 如果设置为空
		lastListenBlock, err := pp.db.BlockListener.GetLastBlockNumber("event") // 获取上次监听的区块高度
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
			pp.startHeight = big.NewInt(int64(pp.cfg.StartBlock)) // 如果上次监听的区块高度为0，则设置为配置的起始高度
		}

		// 如果起始高度（上次监听的区块高度）小于配置的事件起始高度，则设置为配置的起始高度
		// if pp.startHeight.Cmp(big.NewInt(int64(pp.eventStartBlock))) == -1 {
		// 	pp.startHeight = big.NewInt(int64(pp.eventStartBlock))
		// }

		if int64(pp.eventStartBlock) != 0 {
			pp.startHeight = big.NewInt(int64(pp.eventStartBlock))
		}

		log.Info("config StartBlock", "StartBlock", pp.startHeight)
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
		log.Info("eventsFetch Action 27c1049", "fromHeight", fromHeight, "toHeight", toHeight)
		eventsFetchErr := pp.eventsFetch(fromHeight, toHeight, tx) // *** 传入事务tx ***
		if eventsFetchErr != nil {
			return eventsFetchErr
		}
		lastBlock := block_listener.BlockListener{
			BlockNumber: toHeight,
			ConfName:    "event",
		}
		updateErr := tx.BlockListener.SaveOrUpdateLastBlockNumber(lastBlock)
		if updateErr != nil {
			log.Info("update last block err :", updateErr)
			return updateErr
		}
		log.Info("updated last listened block to ", "block", toHeight)
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
			log.Warn("failed to index ContractEventsWithFilter ", "err", err)
			return err
		}
		log.Info("fetched contract events", "count", len(events), "contract", contract, "fromHeight", fromHeight, "toHeight", toHeight)

		for _, contractEvent := range events {
			unpackErr := pp.eventUnpack(contractEvent, tx)
			if unpackErr != nil {
				log.Warn("failed to index events", "unpackErr", unpackErr)
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
		if address != "0x84eBc138F4Ab844A3050a6059763D269dC9951c6" {
			log.Info("skipping transfer event for non fcc", "address", address)
			return nil
		} // 筛选出 FCC 合约地址的 Transfer 事件
		err := unpack.Transfer(event, tx, address)
		if err != nil {
			log.Error("failed to unpack transfer event", "err", err)
			return err
		}
		return nil
	}

	log.Info("====== Start parse ourself contracts events===========")

	merchantAbi, _ := abi.FishcakeEventManagerMetaData.GetAbi()
	nftTokenAbi, _ := abi.NftManagerMetaData.GetAbi()
	stakeAbi, _ := abi.StakingManagerMetaData.GetAbi()

	switch event.EventSignature.String() {
	case merchantAbi.Events["ActivityAdd"].ID.String():
		err := unpack.ActivityAdd(event, tx)
		if err != nil {
			log.Error("failed to unpack ActivityAdd event", "err", err)
			return err
		}
		return nil

	case merchantAbi.Events["ActivityFinish"].ID.String():
		err := unpack.ActivityFinish(event, tx)
		if err != nil {
			log.Error("failed to unpack ActivityFinish event", "err", err)
			return err
		}
		return nil

	case nftTokenAbi.Events["CreateNFT"].ID.String():
		err := unpack.MintNft(event, tx)
		if err != nil {
			log.Error("failed to unpack CreateNFT event", "err", err)
			return err
		}
		return nil

	case nftTokenAbi.Events["MintBoosterNFT"].ID.String():
		err := unpack.MintBoosterNft(event, tx)
		if err != nil {
			log.Error("failed to unpack MintBoosterNFT event", "err", err)
			return err
		}
		return nil

	case merchantAbi.Events["Drop"].ID.String():
		err := unpack.Drop(event, tx)
		if err != nil {
			log.Error("failed to unpack Drop event", "err", err)
			return err
		}
		return nil

	case stakeAbi.Events["StakeHolderDepositStaking"].ID.String():
		err := unpack.StakeHolderDepositStaking(event, tx)
		if err != nil {
			log.Error("failed to unpack StakeHolderDepositStaking event", "err", err)
			return err
		}
		return nil

	case stakeAbi.Events["StakeHolderWithdrawStaking"].ID.String():
		err := unpack.StakeHolderWithdrawStaking(event, tx)
		if err != nil {
			log.Error("failed to unpack StakeHolderWithdrawStaking event", "err", err)
			return err
		}
		return nil
	}

	return nil
}
