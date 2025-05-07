package polygon

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/green"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/FishcakeLab/fishcake-service/common/tasks"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/block_listener"
	"github.com/FishcakeLab/fishcake-service/database/event"
	"github.com/FishcakeLab/fishcake-service/event/polygon/abi"
	"github.com/FishcakeLab/fishcake-service/event/polygon/unpack"
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
}

func NewEventProcessor(db *database.DB, loopInterval time.Duration, contracts []string,
	startHeight uint64, eventStartBlock uint64, epoch uint64, shutdown context.CancelCauseFunc,
	aliConfig config.AliConfig) (*PolygonEventProcessor, error) {
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
		startHeight:      new(big.Int).SetUint64(startHeight),
		eventStartBlock:  eventStartBlock,
		contracts:        contracts,
		loopInterval:     loopInterval,
		epoch:            epoch,
		aliContentClient: nil,
	}, nil
}

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
	if pp.startHeight == nil {
		lastListenBlock, err := pp.db.BlockListener.GetLastBlockNumber()
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

		pp.startHeight = lastListenBlock.BlockNumber
		if pp.startHeight.Cmp(big.NewInt(int64(pp.eventStartBlock))) == -1 {
			pp.startHeight = big.NewInt(int64(pp.eventStartBlock))
		}
	} else {
		pp.startHeight = new(big.Int).Add(pp.startHeight, bigint.One)
		log.Info("pp.start height", "pp.startHeight", pp.startHeight)

	}
	fromHeight := pp.startHeight
	toHeight := new(big.Int).Add(fromHeight, big.NewInt(int64(pp.epoch)))

	log.Info("Handle event start and end block", "start", fromHeight, "end", toHeight)

	latestBlockHeader, err := pp.db.Blocks.LatestBlockHeader()
	if err != nil {
		pp.startHeight = new(big.Int).Sub(pp.startHeight, bigint.One)
		return err
	}
	if latestBlockHeader == nil {
		pp.startHeight = new(big.Int).Sub(pp.startHeight, bigint.One)
		return nil
	}
	if latestBlockHeader.Number.Cmp(fromHeight) == -1 {
		fromHeight = latestBlockHeader.Number
		toHeight = latestBlockHeader.Number
	} else {
		if latestBlockHeader.Number.Cmp(toHeight) == -1 {
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
	return nil
}

func (pp *PolygonEventProcessor) eventUnpack(event event.ContractEvent) error {
	merchantAbi, _ := abi.FishcakeEventManagerMetaData.GetAbi()
	nftTokenAbi, _ := abi.NftManagerMetaData.GetAbi()
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
	}
	return nil
}
