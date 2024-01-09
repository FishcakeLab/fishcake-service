package service

import (
	"github.com/FishcakeLab/fishcake-service/api/service/activity_service"
	"github.com/FishcakeLab/fishcake-service/api/service/rpc_service"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
)

type service struct {
	db                  *database.DB
	ActivityInfoService activity_service.ActivityInfoService
	RpcService          rpc_service.RpcService
}

var BaseService *service

func NewBaseService(db *database.DB, cfg *config.Config) {
	BaseService = &service{
		db:                  db,
		ActivityInfoService: activity_service.NewActivityInfoService(db),
		RpcService:          rpc_service.NewRpcService(cfg.PolygonRpc),
	}
}
