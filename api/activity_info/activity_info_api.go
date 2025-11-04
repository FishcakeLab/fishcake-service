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
	r.GET("rank", rank) // router for mining rank
}

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

func info(c *gin.Context) {
	activityIdStr := c.Query("activityId")
	activityId := bigint.StringToInt(activityIdStr)
	info := service.BaseService.ActivityInfoService.ActivityInfo(activityId)
	api_result.NewApiResult(c).Success(info)
}

// /v1/activity/rank?month=1
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
