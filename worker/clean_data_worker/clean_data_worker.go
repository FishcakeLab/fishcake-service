package clean_data_worker

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/tasks"
	"github.com/FishcakeLab/fishcake-service/database"
)

type WorkerProcessor struct {
	db    *database.DB
	tasks tasks.Group
}

func NewWorkerProcessor(db *database.DB, shutdown context.CancelCauseFunc) (*WorkerProcessor, error) {
	workerProcessor := WorkerProcessor{
		db: db,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker processor: %w", err))
		}},
	}
	return &workerProcessor, nil
}

func (b *WorkerProcessor) WorkerStart() error {
	tickerRun := time.NewTicker(time.Second * 5)
	b.tasks.Go(func() error {
		for range tickerRun.C {
			log.Info("clean table ")
			err := b.db.Blocks.CleanBlockHerders()
			if err != nil {
				log.Error("clean table by chainId", "err ", err)
				continue
			}
		}
		return nil
	})
	return nil
}
