package fishcake_service

import (
	"context"
	"fmt"
	"github.com/FishcakeLab/fishcake-service/api/activity_info"
	"github.com/FishcakeLab/fishcake-service/api/chain_info"
	"github.com/FishcakeLab/fishcake-service/api/drop_info"
	"github.com/FishcakeLab/fishcake-service/api/nft_info"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	"github.com/FishcakeLab/fishcake-service/common/logs"
	"github.com/FishcakeLab/fishcake-service/common/middleware"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/event/polygon"
	"github.com/FishcakeLab/fishcake-service/service"
	"github.com/FishcakeLab/fishcake-service/synchronizer"
	"github.com/FishcakeLab/fishcake-service/synchronizer/node"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"log"
	"math/big"
	"strconv"
	"time"
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
	f.newApi(cfg, db)
	return f
}

func NewIndex(ctx *cli.Context, cfg *config.Config, db *database.DB, shutdown context.CancelCauseFunc) *FishCake {
	f := &FishCake{}
	//f.newIndex(ctx, cfg, db, shutdown)
	f.newEvent(cfg, db, shutdown)
	return f
}

func (f *FishCake) newApi(cfg *config.Config, db *database.DB) error {

	// init base service
	service.NewBaseService(db, cfg)

	gin.ForceConsoleColor()
	gin.DefaultWriter = logs.MyLogWriter()

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

	port := fmt.Sprintf(":%d", cfg.HttpPort)
	r.Run(port)
	return nil
}

func (f *FishCake) newIndex(ctx *cli.Context, cfg *config.Config, db *database.DB, shutdown context.CancelCauseFunc) error {
	chainId, _ := strconv.ParseUint(cfg.PolygonChainId, 10, 64)
	syncConfig := &synchronizer.Config{
		LoopIntervalMsec:  1,
		HeaderBufferSize:  500,
		StartHeight:       new(big.Int).SetUint64(cfg.StartBlock),
		ConfirmationDepth: big.NewInt(50),
		ChainId:           uint(chainId),
	}
	client, _ := node.DialEthClient(ctx.Context, cfg.PolygonRpc)
	syncer, _ := synchronizer.NewSynchronizer(cfg, syncConfig, db, client, shutdown)
	err := syncer.Start()
	if err != nil {
		log.Println("failed to start synchronizer:", err)
		return err
	}
	return nil
}

func (f *FishCake) newEvent(cfg *config.Config, db *database.DB, shutdown context.CancelCauseFunc) error {
	var epoch uint64 = 10_000
	var loopInterval time.Duration = time.Second * 2
	eventProcessor, _ := polygon.NewEventProcessor(db, loopInterval, cfg.Contracts, cfg.StartBlock, cfg.EventStartBlock, epoch, shutdown)
	err := eventProcessor.Start()
	if err != nil {
		log.Println("failed to start eventProcessor:", err)
		return err
	}
	return nil
}
