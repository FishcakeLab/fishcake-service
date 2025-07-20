package queue_transaction

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/tasks"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/wallet"
	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"github.com/FishcakeLab/fishcake-service/service"
	"github.com/FishcakeLab/fishcake-service/synchronizer/retry"
)

const (
	UnSendQueueTx  = 0
	UnGotReceiptTx = 1
)

type QueueTxProcessor struct {
	baseService    *service.Service
	db             *database.DB
	tasks          tasks.Group
	resourceCtx    context.Context
	resourceCancel context.CancelFunc
}

func NewQueueTxProcessor(db *database.DB, cfg *config.Config, shutdown context.CancelCauseFunc) (*QueueTxProcessor, error) {
	baseService := service.NewBaseService(db, cfg)

	resCtx, resCancel := context.WithCancel(context.Background())

	queueTxWorkerPro := QueueTxProcessor{
		baseService: baseService,
		db:          db,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in queue tx processor: %w", err))
		}},
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
	}
	return &queueTxWorkerPro, nil
}

func (qt *QueueTxProcessor) QueueTxStart() error {
	tickerRun := time.NewTicker(time.Second * 2)
	qt.tasks.Go(func() error {
		for range tickerRun.C {
			log.Info("send queue transaction list")
			err := qt.ProcessSendQueueTx()
			if err != nil {
				log.Error("handle queue transaction list fail", "err", err)
				return err
			}
		}

		for range tickerRun.C {
			log.Info("fetch queue transaction receipt")
			err := qt.AfterSentQueueTx()
			if err != nil {
				log.Error("handle queue transaction receipt fail", "err", err)
				return err
			}
		}

		return nil
	})
	return nil
}

func (qt *QueueTxProcessor) ProcessSendQueueTx() error {
	unhandledTxList, err := qt.db.QueueTxDB.QueryRawTxInfoByStatus(UnSendQueueTx)
	if err != nil {
		log.Error("query unhandled transaction list fail", "err", err)
		return err
	}

	var handledTxList []wallet.QueueTx
	for _, unhandledTx := range unhandledTxList {
		reqTx := &account.SendTxRequest{
			Chain:   "Polygon",
			Network: "mainnet",
			RawTx:   unhandledTx.RawTx,
		}

		sendTx, errSentTx := qt.baseService.RpcService.SendTx(context.Background(), reqTx)
		if errSentTx != nil {
			log.Error("RPC send tx error: %v", err)
			unhandledTx.Result = fmt.Sprintf("RPC send tx error: %v", err)
			unhandledTx.Status = 3
		}
		if sendTx != nil {
			log.Info("send tx success", "txHash", sendTx.TxHash)
			unhandledTx.TransactionHash = sendTx.TxHash
			unhandledTx.Timestamp = uint64(time.Now().Unix())
			unhandledTx.Result = "success to send tx"
			unhandledTx.Status = 1
		} else {
			continue
		}
		handledTxList = append(handledTxList, unhandledTx)
	}
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](qt.resourceCtx, 10, retryStrategy, func() (interface{}, error) {
		if err := qt.db.Transaction(func(tx *database.DB) error {
			if len(handledTxList) > 0 {
				err := qt.db.QueueTxDB.MarkedTxToSentOrSuccess(handledTxList)
				if err != nil {
					log.Error("Market tx to send fail", "err", err)
					return err
				}
			}
			return nil
		}); err != nil {
			log.Info("unable to persist batch", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

func (qt *QueueTxProcessor) AfterSentQueueTx() error {
	unhandledTxList, err := qt.db.QueueTxDB.QueryRawTxInfoByStatus(UnGotReceiptTx)
	if err != nil {
		log.Error("query unhandled transaction list fail", "err", err)
		return err
	}

	var handledTxList []wallet.QueueTx
	for _, unhandledTx := range unhandledTxList {
		reqTx := &account.TxHashRequest{
			Chain:   "Polygon",
			Network: "mainnet",
			Hash:    unhandledTx.TransactionHash,
		}

		fetchTx, errFetchTx := qt.baseService.RpcService.GetTxByHash(context.Background(), reqTx)
		if errFetchTx != nil {
			log.Error("RPC send tx error: %v", err)
			continue
		}
		if fetchTx != nil {
			if fetchTx.Tx.Status == account.TxStatus_Success {
				unhandledTx.Status = 2
			} else {
				unhandledTx.Status = 3
			}
		} else {
			continue
		}
		handledTxList = append(handledTxList, unhandledTx)
	}
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](qt.resourceCtx, 10, retryStrategy, func() (interface{}, error) {
		if err := qt.db.Transaction(func(tx *database.DB) error {
			if len(handledTxList) > 0 {
				err := qt.db.QueueTxDB.MarkedTxToSentOrSuccess(handledTxList)
				if err != nil {
					log.Error("Market tx to success fail", "err", err)
					return err
				}
			}
			return nil
		}); err != nil {
			log.Info("unable to persist batch", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}
