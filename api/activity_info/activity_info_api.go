package activity_info

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/service"
)

func ActivityInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/activity")
	r.GET("list", list)
	r.GET("info", info)
	r.GET("miningRank", rank) // router for mining rank
	r.GET("minedAmount", getUserMinedAmount)
	r.GET("miningInfo", getUserMiningInfo)
}

// getUserMinedAmount godoc
// @Summary Get user mined amount
// @Description Get cumulative mining amount for a user (business account)
// @Tags Activity
// @Accept json
// @Produce json
// @Param address query string true "Wallet address"
// @Param month query int false "1=last 30 days only, omit for all-time"
// @Success 200 {object} api_result.ApiResult{obj=object{userAddress=string,totalMined=string,month=bool}}
// @Router /v1/activity/minedAmount [get]
func getUserMinedAmount(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, "address is required")
		return
	}

	monthStr := c.Query("month")
	var monthFilter bool
	if monthStr != "" && bigint.StringToInt(monthStr) == 1 {
		monthFilter = true
	}

	log.Info(">>> getUserMinedAmount", "address", address, "monthFilter", monthFilter)

	res, err := service.BaseService.ActivityInfoService.GetUserMinedAmount(address, monthFilter)
	if err != nil {
		api_result.NewApiResult(c).Error(enum.DataErr.Code, err.Error())
		return
	}

	api_result.NewApiResult(c).Success(res)
}

// list godoc
// @Summary Activity list
// @Description Query activity list with filters, corresponds to contract ActivityAdd events
// @Tags Activity
// @Accept json
// @Produce json
// @Param pageNum query int true "Page number, starts from 1"
// @Param pageSize query int true "Page size"
// @Param activityId query int false "Activity ID"
// @Param activityStatus query string false "Status: 1=active, 2=finished, 3=expired"
// @Param businessName query string false "Business name (fuzzy search)"
// @Param businessAccount query string false "Business wallet address (fuzzy search)"
// @Param tokenContractAddr query string false "Token contract address (fuzzy search)"
// @Param latitude query string false "Latitude (requires longitude and scope)"
// @Param longitude query string false "Longitude"
// @Param scope query string false "Search radius in meters"
// @Param activityFilter query string false "NFT filter: 1=Pro NFT valid, 2=Basic NFT valid"
// @Success 200 {object} api_result.ApiResult
// @Router /v1/activity/list [get]
func list(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageNumStr := c.Query("pageNum")
	if pageSizeStr == "" || pageNumStr == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
		return
	}
	pageSize := bigint.StringToInt(pageSizeStr)
	pageNum := bigint.StringToInt(pageNumStr)
	activityIdStr := c.Query("activityId")
	activityId := bigint.StringToInt(activityIdStr)
	activityStatus := c.Query("activityStatus")
	businessName := c.Query("businessName")
	tokenContractAddr := c.Query("tokenContractAddr")
	businessAccount := c.Query("businessAccount")
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")
	scope := c.Query("scope")
	activityFilter := c.Query("activityFilter")
	infos, count := service.BaseService.ActivityInfoService.ActivityInfoList(activityFilter, businessAccount, activityStatus, businessName, tokenContractAddr, latitude, longitude, scope, activityId, pageNum, pageSize)
	page := api_result.NewPage(infos, count, pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}

// info godoc
// @Summary Activity detail
// @Description Get single activity info by ID
// @Tags Activity
// @Accept json
// @Produce json
// @Param activityId query int true "Activity ID"
// @Success 200 {object} api_result.ApiResult
// @Router /v1/activity/info [get]
func info(c *gin.Context) {
	activityIdStr := c.Query("activityId")
	activityId := bigint.StringToInt(activityIdStr)
	info := service.BaseService.ActivityInfoService.ActivityInfo(activityId)
	api_result.NewApiResult(c).Success(info)
}

// rank godoc
// @Summary Mining leaderboard
// @Description Get mining rank by business account mined amount
// @Tags Leaderboard
// @Accept json
// @Produce json
// @Param month query int false "1=last 30 days, omit for all-time"
// @Success 200 {object} api_result.ApiResult
// @Router /v1/activity/miningRank [get]
func rank(c *gin.Context) {
	monthStr := c.Query("month") // 可选参数
	var monthFilter bool
	if monthStr != "" && bigint.StringToInt(monthStr) == 1 {
		monthFilter = true
	}

	log.Info(">>> rank API called", "monthFilter", monthFilter)
	ranks, err := service.BaseService.ActivityInfoService.GetActivityRank(monthFilter)
	if err != nil {
		api_result.NewApiResult(c).Error(enum.DataErr.Code, err.Error())
		return
	}

	api_result.NewApiResult(c).Success(ranks)
}

// getUserMiningInfo godoc
// @Summary User mining info & Fishcake Power
// @Description Get real-time mining info including Fishcake Power for a user
// @Tags Earnings
// @Accept json
// @Produce json
// @Param address query string true "Wallet address"
// @Success 200 {object} api_result.ApiResult
// @Router /v1/activity/miningInfo [get]
func getUserMiningInfo(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, "address is required")
		return
	}

	log.Info(">>> getUserMiningInfo", "address", address)
	res, err := service.BaseService.ActivityInfoService.GetUserMiningInfo(address)
	if err != nil {
		api_result.NewApiResult(c).Error(enum.DataErr.Code, err.Error())
		return
	}

	api_result.NewApiResult(c).Success(res)
}
