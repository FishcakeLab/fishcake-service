package fishcake_service

import (
	"context"
	"fmt"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/utils/database"
	"github.com/FishcakeLab/fishcake-service/utils/mylogs"
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

func NewFishCake(cfg *config.Config) *FishCake {
	f := &FishCake{}
	database.NewOrm(cfg)
	f.newApi(cfg)
	return f
}

func (f *FishCake) newApi(cfg *config.Config) error {
	gin.ForceConsoleColor()
	gin.DefaultWriter = mylogs.MyLogWriter()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	port := fmt.Sprintf(":%d", cfg.HttpPort)
	r.Run(port)
	return nil
}
