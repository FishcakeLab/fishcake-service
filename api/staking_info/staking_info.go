package staking_info

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/service"
)

// StakingInfoApi 注册 staking 信息相关路由
func StakingInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/staking")

	// stake总量及月度排行
	r.GET("stakingRank", stakingRank)

	// 已领取奖励的排行（status = 1）
	r.GET("claimedRank", claimedRank)

	// 已领取 + 未领取（status = 0）的总奖励排行
	r.GET("totalRewardRank", totalRewardRank)
}

// /v1/staking/stakeRank?month=1
func stakingRank(c *gin.Context) {
	monthStr := c.Query("month")
	var monthFilter bool
	if monthStr != "" && bigint.StringToInt(monthStr) == 1 {
		monthFilter = true
	}

	log.Info(">>> stakeRank API called", "monthFilter", monthFilter)
	res, err := service.BaseService.StakingInfoService.GetStakeRank(monthFilter)
	if err != nil {
		api_result.NewApiResult(c).Error(enum.DataErr.Code, err.Error())
		return
	}

	api_result.NewApiResult(c).Success(res)
}

// /v1/staking/claimedRank?month=1
func claimedRank(c *gin.Context) {
	monthStr := c.Query("month")
	var monthFilter bool
	if monthStr != "" && bigint.StringToInt(monthStr) == 1 {
		monthFilter = true
	}

	log.Info(">>> claimedRank API called", "monthFilter", monthFilter)
	ranks, err := service.BaseService.StakingInfoService.GetClaimedRank(monthFilter)
	if err != nil {
		api_result.NewApiResult(c).Error(enum.DataErr.Code, err.Error())
		return
	}

	api_result.NewApiResult(c).Success(ranks)
}

// /v1/staking/totalRewardRank?month=1
func totalRewardRank(c *gin.Context) {
	monthStr := c.Query("month")
	var monthFilter bool
	if monthStr != "" && bigint.StringToInt(monthStr) == 1 {
		monthFilter = true
	}

	log.Info(">>> totalRewardRank API called", "monthFilter", monthFilter)
	ranks, err := service.BaseService.StakingInfoService.GetTotalRewardRank(monthFilter)
	if err != nil {
		api_result.NewApiResult(c).Error(enum.DataErr.Code, err.Error())
		return
	}

	api_result.NewApiResult(c).Success(ranks)
}
