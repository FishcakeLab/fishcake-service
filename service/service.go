package service

import (
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/service/activity_service"
	"github.com/FishcakeLab/fishcake-service/service/rpc_service"
)

type service struct {
	Db                  *database.DB
	Cfg                 *config.Config
	ActivityInfoService activity_service.ActivityInfoService
	RpcService          rpc_service.RpcService
}

var BaseService *service

func NewBaseService(db *database.DB, cfg *config.Config) {
	BaseService = &service{
		Db:                  db,
		Cfg:                 cfg,
		ActivityInfoService: activity_service.NewActivityInfoService(db),
		RpcService:          rpc_service.NewRpcService(cfg.PolygonRpc),
	}
}
