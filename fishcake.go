package fishcake_service

import (
	"context"
	"fmt"
	"github.com/FishcakeLab/fishcake-service/api/notification"

	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"

	"github.com/FishcakeLab/fishcake-service/api/activity_info"
	"github.com/FishcakeLab/fishcake-service/api/chain_info"
	"github.com/FishcakeLab/fishcake-service/api/contract_info"
	"github.com/FishcakeLab/fishcake-service/api/drop_info"
	"github.com/FishcakeLab/fishcake-service/api/nft_info"
	"github.com/FishcakeLab/fishcake-service/api/wallet_info"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	"github.com/FishcakeLab/fishcake-service/common/middleware"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/event/polygon"
	"github.com/FishcakeLab/fishcake-service/service"
	"github.com/FishcakeLab/fishcake-service/synchronizer"
	"github.com/FishcakeLab/fishcake-service/synchronizer/node"
	"github.com/FishcakeLab/fishcake-service/worker/clean_data_worker"
	"github.com/FishcakeLab/fishcake-service/worker/drop_worker"
)

type FishCake struct {
}

func (f *FishCake) Start(ctx context.Context) error {
	return nil
}

func (f *FishCake) Stop(ctx context.Context) error {
	return nil
}

func (f *FishCake) Stopped() bool {
	return true
}

func NewFishCake(cfg *config.Config, db *database.DB) *FishCake {
	f := &FishCake{}
	err := f.newApi(cfg, db)
	if err != nil {
		return nil
	}
	return f
}

func NewIndex(ctx *cli.Context, cfg *config.Config, db *database.DB, shutdown context.CancelCauseFunc) *FishCake {
	f := &FishCake{}
	err := f.newIndex(ctx, cfg, db, shutdown)
	if err != nil {
		return nil
	}
	err = f.newEvent(cfg, db, shutdown)
	if err != nil {
		return nil
	}
	return f
}

func (f *FishCake) newApi(cfg *config.Config, db *database.DB) error {

	service.NewApiBaseService(db, cfg)

	gin.ForceConsoleColor()
	r := gin.Default()

	r.Use(errors_h.Recover)
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	activity_info.ActivityInfoApi(r)
	activity_info.ActivityInfoExtApi(r)
	chain_info.ChainInfoApi(r)
	nft_info.NftInfoApi(r)
	drop_info.DropInfoApi(r)
	contract_info.ContractInfoApi(r)
	wallet_info.WalletInfoApi(r)
	notification.NotificationApi(r)
	port := fmt.Sprintf(":%d", cfg.HttpPort)
	err := r.Run(port)
	if err != nil {
		log.Error("new api fail", "err", err)
		return err
	}
	return nil
}

func (f *FishCake) newIndex(ctx *cli.Context, cfg *config.Config, db *database.DB, shutdown context.CancelCauseFunc) error {
	chainId, _ := strconv.ParseUint(cfg.PolygonChainId, 10, 64)
	syncConfig := &synchronizer.Config{
		LoopIntervalMsec:  1,
		HeaderBufferSize:  100,
		StartHeight:       new(big.Int).SetUint64(cfg.StartBlock),
		ConfirmationDepth: big.NewInt(50),
		ChainId:           uint(chainId),
	}
	client, _ := node.DialEthClient(ctx.Context, cfg.PolygonRpc)
	syncer, _ := synchronizer.NewSynchronizer(cfg, syncConfig, db, client, shutdown)
	worker, _ := clean_data_worker.NewWorkerProcessor(db, shutdown)
	dropWorker, _ := drop_worker.NewDropWorkerProcessor(db, cfg, shutdown)
	systemDropWorker, _ := drop_worker.NewSystemDropWorkerProcessor(db, cfg, shutdown, client)
	err := syncer.Start()
	if err != nil {
		log.Error("failed to start synchronizer:", err)
		return err
	}
	err = worker.WorkerStart()
	if err != nil {
		log.Error("failed to start clean data worker:", err)
		return err
	}
	err = dropWorker.DropWorkerStart()
	if err != nil {
		log.Error("failed to start drop  worker:", err)
		return err
	}
	err = systemDropWorker.SystemDropWorkerStart()
	if err != nil {
		log.Error("failed to start system drop worker:", err)
		return err
	}
	return nil
}

func (f *FishCake) newEvent(cfg *config.Config, db *database.DB, shutdown context.CancelCauseFunc) error {
	var epoch uint64 = 10_000
	var loopInterval time.Duration = time.Second * 2
	eventProcessor, _ := polygon.NewEventProcessor(db, loopInterval, cfg.Contracts,
		cfg.StartBlock, cfg.EventStartBlock, epoch, shutdown, cfg.AliConfig)
	err := eventProcessor.Start()
	if err != nil {
		log.Error("failed to start eventProcessor:", err)
		return err
	}
	return nil
}
