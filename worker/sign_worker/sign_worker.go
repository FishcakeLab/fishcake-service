package sign_worker

import (
	"context"
	"fmt"
	"github.com/FishcakeLab/fishcake-service/synchronizer/node"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/tasks"
	"github.com/FishcakeLab/fishcake-service/database"
)

type SignWorkerProcessor struct {
	db        *database.DB
	tasks     tasks.Group
	ethClient node.EthClient
}

func NewSignWorkerProcessor(db *database.DB, shutdown context.CancelCauseFunc, client node.EthClient) (*SignWorkerProcessor, error) {
	signWorkerProcessor := SignWorkerProcessor{
		db: db,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker processor: %w", err))
		}},
		ethClient: client,
	}
	return &signWorkerProcessor, nil
}

func (pr *SignWorkerProcessor) WorkerStart() error {
	tickerRun := time.NewTicker(time.Second * 1)
	pr.tasks.Go(func() error {
		for range tickerRun.C {
			log.Info("offline sign service start.....")
			pendingSigns := pr.db.SignStatusDB.GetSignsByStatus(0)
			if pendingSigns != nil {
				for _, value := range pendingSigns {
					txHash, err := pr.ethClient.SendTx(value.Signature)
					if err != nil {
						log.Error("send offline sign fail", "err", err)
						if err := pr.db.SignStatusDB.UpdateSignStatus(value.Id, 3, txHash, "send offline sign fail"); err != nil {
							log.Error("pending signs update status failed", "id", value.Id, "err", err)
							continue
						}
					}
					if err := pr.db.SignStatusDB.UpdateSignStatus(value.Id, 1, txHash, "send offline sign success"); err != nil {
						log.Error("pending signs update status failed", "id", value.Id, "err", err)
						continue
					}
					log.Info("pending signs Update Status success", "hash", value.Signature)
				}
			}
		}
		return nil
	})
	return nil
}
