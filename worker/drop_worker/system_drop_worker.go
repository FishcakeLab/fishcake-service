package drop_worker

import (
	"context"
	"fmt"
	"github.com/FishcakeLab/fishcake-service/synchronizer/node"
	"github.com/ethereum/go-ethereum/common"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/tasks"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/service"
)

type SystemDropWorkerProcessor struct {
	baseService *service.Service
	db          *database.DB
	tasks       tasks.Group
	ethClient   node.EthClient
}

func NewSystemDropWorkerProcessor(db *database.DB, cfg *config.Config, shutdown context.CancelCauseFunc, client node.EthClient) (*SystemDropWorkerProcessor, error) {
	baseService := service.NewBaseService(db, cfg)
	systemDropWorkerProcessor := SystemDropWorkerProcessor{
		baseService: baseService,
		db:          db,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in systm drop worker processor: %w", err))
		}},
		ethClient: client,
	}
	return &systemDropWorkerProcessor, nil
}

func (dp *SystemDropWorkerProcessor) SystemDropWorkerStart() error {
	tickerRun := time.NewTicker(time.Second * 15)
	dp.tasks.Go(func() error {
		for range tickerRun.C {
			log.Info("system drop service start.....")
			unDropConfirmAddresses := dp.db.SystemDropInfoDB.UnDropConfirmAddresses()
			if unDropConfirmAddresses != nil {
				for _, value := range unDropConfirmAddresses {
					receipt, err := dp.ethClient.TxReceiptByHash(common.HexToHash(value.TransactionHash))
					if err != nil {
						log.Error("fetch tx receipt by hash fail", "err", err)
						continue
					}
					if receipt.Status != uint64(1) {
						continue
					}

					if err := dp.db.SystemDropInfoDB.UpdateStatus(value.Id); err != nil {
						log.Error("update status failed", "id", value.Id, "err", err)
						continue
					}
					log.Info("unDropConfirm Addresses Update Status success", "hash", value.TransactionHash)
				}
			}
		}
		return nil
	})
	return nil
}
