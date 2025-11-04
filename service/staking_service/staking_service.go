package staking_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/stake"
)

// StakingInfoService 定义 stake 信息业务层接口
type StakingInfoService interface {
	// 获取 stake 排行榜（总量或月度）
	GetStakeRank(monthFilter bool) ([]stake.StakeRank, error)

	// 获取已领取奖励排行榜（status = 1）
	GetClaimedRank(monthFilter bool) ([]stake.ClaimedRank, error)

	// 获取总奖励排行榜（status = 0 + status = 1）
	GetTotalRewardRank(monthFilter bool) ([]stake.TotalRewardRank, error)
}

// stakeInfoService 结构体实现
type stakeInfoService struct {
	Db *database.DB
}

// NewStakingInfoService 创建新的实例
func NewStakingInfoService(db *database.DB) StakingInfoService {
	return &stakeInfoService{Db: db}
}

// GetStakeRank 查询 stake 排名
func (s *stakeInfoService) GetStakeRank(monthFilter bool) ([]stake.StakeRank, error) {
	ranks, err := s.Db.StakingDB.GetStakeRank(monthFilter)
	if err != nil {
		return nil, err
	}
	return ranks, nil
}

// GetClaimedRank 查询已领取奖励排名（status = 1）
func (s *stakeInfoService) GetClaimedRank(monthFilter bool) ([]stake.ClaimedRank, error) {
	ranks, err := s.Db.StakingDB.GetClaimedRank(monthFilter)
	if err != nil {
		return nil, err
	}
	return ranks, nil
}

// GetTotalRewardRank 查询总奖励排名（包含 status = 0 和 status = 1）
func (s *stakeInfoService) GetTotalRewardRank(monthFilter bool) ([]stake.TotalRewardRank, error) {
	ranks, err := s.Db.StakingDB.GetTotalRewardRank(monthFilter)
	if err != nil {
		return nil, err
	}
	return ranks, nil
}
