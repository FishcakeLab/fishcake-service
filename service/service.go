package service

import (
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/service/activity_service"
	"github.com/FishcakeLab/fishcake-service/service/contract_info"
	"github.com/FishcakeLab/fishcake-service/service/drop_service"
	"github.com/FishcakeLab/fishcake-service/service/nft_service"
	"github.com/FishcakeLab/fishcake-service/service/rpc_service"
)

type service struct {
	Db                     *database.DB
	Cfg                    *config.Config
	ActivityInfoService    activity_service.ActivityInfoService
	ActivityInfoExtService activity_service.ActivityInfoExtService
	RpcService             rpc_service.RpcService
	NftService             nft_service.NftService
	DropService            drop_service.DropService
	ContractInfoService    contract_info.ContractInfoService
}

var BaseService *service

func NewBaseService(db *database.DB, cfg *config.Config) {
	BaseService = &service{
		Db:                     db,
		Cfg:                    cfg,
		ActivityInfoService:    activity_service.NewActivityInfoService(db),
		ActivityInfoExtService: activity_service.NewActivityInfoExtService(db),
		RpcService:             rpc_service.NewRpcService(cfg.RpcUrl),
		NftService:             nft_service.NewNftService(db),
		DropService:            drop_service.NewDropService(db),
		ContractInfoService:    contract_info.NewContractInfoService(cfg),
	}
}
