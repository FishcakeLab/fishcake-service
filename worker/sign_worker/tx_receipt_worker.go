package sign_worker

import (
	"context"
	"fmt"
	"github.com/FishcakeLab/fishcake-service/synchronizer/node"
	"github.com/ethereum/go-ethereum/common"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/tasks"
	"github.com/FishcakeLab/fishcake-service/database"
)

type TxReceiptWorkerProcessor struct {
	db        *database.DB
	tasks     tasks.Group
	ethClient node.EthClient
}

func NewTxReceiptWorkerProcessor(db *database.DB, shutdown context.CancelCauseFunc, client node.EthClient) (*TxReceiptWorkerProcessor, error) {
	txReceiptWorkerProcessor := TxReceiptWorkerProcessor{
		db: db,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker processor: %w", err))
		}},
		ethClient: client,
	}
	return &txReceiptWorkerProcessor, nil
}

func (pr *TxReceiptWorkerProcessor) WorkerStart() error {
	tickerRun := time.NewTicker(time.Second * 5)
	pr.tasks.Go(func() error {
		for range tickerRun.C {
			log.Info("transaction receipt service start.....")
			completeSendSigns := pr.db.SignStatusDB.GetSignsByStatus(1)
			if completeSendSigns != nil {
				for _, value := range completeSendSigns {
					receipt, err := pr.ethClient.TxReceiptByHash(common.HexToHash(value.TransactionHash))
					if err != nil {
						log.Error("transaction receipt fail", "err", err)
						if err := pr.db.SignStatusDB.UpdateSignStatus(value.Id, 3, value.TransactionHash, "transaction receipt fail"); err != nil {
							log.Error("transaction receipt failed", "id", value.Id, "err", err)
							continue
						}
					}
					if receipt.Status != uint64(1) {
						if err := pr.db.SignStatusDB.UpdateSignStatus(value.Id, 3, value.TransactionHash, "transaction receipt fail"); err != nil {
							log.Error("transaction receipt failed", "id", value.Id, "err", err)
							continue
						}
					}

					if err := pr.db.SignStatusDB.UpdateSignStatus(value.Id, 2, value.TransactionHash, "transaction receipt success"); err != nil {
						log.Error("transaction receipt failed", "id", value.Id, "err", err)
						continue
					}
					log.Info("transaction receipt Status success", "hash", value.Signature)
				}
			}
		}
		return nil
	})
	return nil
}
