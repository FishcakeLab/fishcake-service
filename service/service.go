package service

import (
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/service/activity_service"
	"github.com/FishcakeLab/fishcake-service/service/contract_info"
	"github.com/FishcakeLab/fishcake-service/service/dapplink_service"
	"github.com/FishcakeLab/fishcake-service/service/drop_service"
	"github.com/FishcakeLab/fishcake-service/service/nft_service"
	"github.com/FishcakeLab/fishcake-service/service/reward_service"
	"github.com/FishcakeLab/fishcake-service/service/rpc_service"
	"github.com/FishcakeLab/fishcake-service/service/wallet_service"
	"github.com/ethereum/go-ethereum/log"
)

var BaseService *Service

type Service struct {
	Db                     *database.DB
	Cfg                    *config.Config
	ActivityInfoService    activity_service.ActivityInfoService
	ActivityInfoExtService activity_service.ActivityInfoExtService
	RpcService             rpc_service.RpcService
	NftService             nft_service.NftService
	DropService            drop_service.DropService
	ContractInfoService    contract_info.ContractInfoService
	RewardService          reward_service.RewardService
	WalletService          wallet_service.WalletService
	DappLinkService        dapplink_service.DappLinkService
}

func NewBaseService(db *database.DB, cfg *config.Config) *Service {
	dapplinkService, err := dapplink_service.NewDappLinkService(cfg)
	if err != nil {
		log.Error("new dappplink service fail", "err", err)
		return nil
	}
	return &Service{
		Db:                     db,
		Cfg:                    cfg,
		ActivityInfoService:    activity_service.NewActivityInfoService(db),
		ActivityInfoExtService: activity_service.NewActivityInfoExtService(db),
		RpcService:             rpc_service.NewRpcService(cfg.RpcUrl),
		NftService:             nft_service.NewNftService(db),
		DropService:            drop_service.NewDropService(db),
		ContractInfoService:    contract_info.NewContractInfoService(cfg),
		RewardService:          reward_service.NewRewardService(cfg),
		WalletService:          wallet_service.NewWalletService(db),
		DappLinkService:        *dapplinkService,
	}
}

func NewApiBaseService(db *database.DB, cfg *config.Config) {
	BaseService = NewBaseService(db, cfg)
}
