package contract_info

import (
	"github.com/FishcakeLab/fishcake-service/config"
)

type ContractInfoService interface {
	ContractInfo() config.ContractInfo
}

type contractInfoService struct {
	cfg *config.Config
}

func NewContractInfoService(cfg *config.Config) ContractInfoService {
	return &contractInfoService{
		cfg: cfg,
	}
}

func (cis *contractInfoService) ContractInfo() config.ContractInfo {
	info := cis.cfg.ContractInfo
	return info
}
