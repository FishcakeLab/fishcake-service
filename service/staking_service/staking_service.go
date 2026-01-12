package staking_service

import (
	"math/big"

	"github.com/FishcakeLab/fishcake-service/database"
)

// StakingInfoService 定义 stake 信息业务层接口
type StakingInfoService interface {
	// 获取 stake 排行榜（总量或月度）
	GetStakeRank(monthFilter bool) ([]StakeRankResponse, error)

	// 获取已领取奖励排行榜（status = 1）
	GetClaimedRank(monthFilter bool) ([]ClaimedRankResponse, error)

	// 获取总奖励排行榜（status = 0 + status = 1）
	GetTotalRewardRank(monthFilter bool) ([]TotalRewardRankResponse, error)

	GetUserStakingInfo(
		address string,
		status *int,
		page, size int,
	) ([]UserStakingInfoResponse, int64, error)
}

// stakeInfoService 结构体实现
type stakeInfoService struct {
	Db *database.DB
}

// StakeRankResponse — 接口层返回结构（带 rank、month）
type StakeRankResponse struct {
	Rank        int      `json:"rank"`
	UserAddress string   `json:"userAddress"`
	TotalStake  *big.Int `json:"totalStake"`
	Month       bool     `json:"month"`
}

// ClaimedRankResponse — 接口层返回结构（带 rank、month）
type ClaimedRankResponse struct {
	Rank        int      `json:"rank"`        // 排名序号
	UserAddress string   `json:"userAddress"` // 用户钱包地址
	Claimed     *big.Int `json:"claimed"`     // 已领取奖励
	Month       bool     `json:"month"`       // 是否为月度排行
}

// TotalRewardRankResponse — 接口层返回结构（含已领取与未领取）
type TotalRewardRankResponse struct {
	Rank        int      `json:"rank"`        // 排名序号
	UserAddress string   `json:"userAddress"` // 用户钱包地址
	TotalReward *big.Int `json:"totalReward"` // 总奖励 = 已领取 + 未领取
	Claimed     *big.Int `json:"claimed"`     // 已领取奖励
	Unclaimed   *big.Int `json:"unclaimed"`   // 预估未领取奖励（动态计算）
	Month       bool     `json:"month"`       // 是否为月度排行
}

// UserStakingInfoResponse 用户质押信息响应结构体
type UserStakingInfoResponse struct {
	TokenId       int64  `json:"tokenId"`
	Amount        string `json:"amount"`
	StakingType   int16  `json:"stakingType"`
	StartTime     int64  `json:"startTime"`
	EndTime       int64  `json:"endTime"`
	NftApr        int64  `json:"nftApr"`
	IsAutoRenew   bool   `json:"isAutoRenew"`
	MessageNonce  int64  `json:"messageNonce"`
	TxMessageHash string `json:"txMessageHash"`
	StakingReward string `json:"stakingReward"`
	StakingStatus int16  `json:"stakingStatus"`
	CreateTime    int64  `json:"createTime"`
}

// NewStakingInfoService 创建新的实例
func NewStakingInfoService(db *database.DB) StakingInfoService {
	return &stakeInfoService{Db: db}
}

// GetStakeRank 查询 stake 排名
func (s *stakeInfoService) GetStakeRank(monthFilter bool) ([]StakeRankResponse, error) {
	ranks, err := s.Db.StakingDB.GetStakeRank(monthFilter)
	if err != nil {
		return nil, err
	}

	// var res []StakeRankResponse
	res := make([]StakeRankResponse, 0)

	for i, item := range ranks {
		res = append(res, StakeRankResponse{
			Rank:        i + 1,
			UserAddress: item.UserAddress,
			TotalStake:  item.TotalStake,
			Month:       monthFilter,
		})
	}

	return res, nil
}

// GetUserStakingInfo 查询用户质押记录
func (s *stakeInfoService) GetUserStakingInfo(
	address string,
	status *int,
	page, size int,
) ([]UserStakingInfoResponse, int64, error) {

	// === 1. 调用 DAO 层 ===
	stakingList, total, err := s.Db.StakingDB.GetUserStakingInfo(address, status, page, size)
	if err != nil {
		return nil, 0, err
	}

	// === 2. 组装 service 层 Response ===
	var resp []UserStakingInfoResponse
	for _, item := range stakingList {
		resp = append(resp, UserStakingInfoResponse{
			TokenId:       item.TokenID,
			Amount:        item.Amount.String(),
			StakingType:   item.StakingType,
			StartTime:     item.StartTime,
			EndTime:       item.EndTime,
			NftApr:        item.NftApr,
			IsAutoRenew:   item.IsAutoRenew,
			MessageNonce:  item.MessageNonce,
			TxMessageHash: item.TxMessageHash,
			StakingReward: item.StakingReward.String(),
			StakingStatus: item.StakingStatus,
			CreateTime:    item.CreateTime.Unix(),
		})
	}

	return resp, total, nil
}

// GetClaimedRank 查询已领取奖励排名（status = 1）
func (s *stakeInfoService) GetClaimedRank(monthFilter bool) ([]ClaimedRankResponse, error) {
	ranks, err := s.Db.StakingDB.GetClaimedRank(monthFilter)
	if err != nil {
		return nil, err
	}
	var res []ClaimedRankResponse
	for i, item := range ranks {
		res = append(res, ClaimedRankResponse{
			Rank:        i + 1,
			UserAddress: item.UserAddress,
			Claimed:     item.Claimed,
			Month:       monthFilter,
		})
	}
	return res, nil
}

// GetTotalRewardRank 查询总奖励排名（包含 status = 0 和 status = 1）
func (s *stakeInfoService) GetTotalRewardRank(monthFilter bool) ([]TotalRewardRankResponse, error) {
	ranks, err := s.Db.StakingDB.GetTotalRewardRank(monthFilter)
	if err != nil {
		return nil, err
	}
	var res []TotalRewardRankResponse
	for i, item := range ranks {
		res = append(res, TotalRewardRankResponse{
			Rank:        i + 1,
			UserAddress: item.UserAddress,
			TotalReward: item.TotalReward,
			Claimed:     item.Claimed,
			Unclaimed:   item.Unclaimed,
			Month:       monthFilter,
		})
	}
	return res, nil
}
