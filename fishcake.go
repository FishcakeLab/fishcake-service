package fishcake_service

import (
	"context"
	"fmt"
	"github.com/FishcakeLab/fishcake-service/api/activity_info"
	"github.com/FishcakeLab/fishcake-service/common/logs"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/gin-gonic/gin"
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

func (f *FishCake) newApi(cfg *config.Config, db *database.DB) error {
	gin.ForceConsoleColor()
	gin.DefaultWriter = logs.MyLogWriter()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	activity_info.ActivityInfoApi(r, db)

	port := fmt.Sprintf(":%d", cfg.HttpPort)
	r.Run(port)
	return nil
}
