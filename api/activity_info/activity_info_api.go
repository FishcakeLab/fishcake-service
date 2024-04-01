package activity_info

import (
	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/service"
	"github.com/gin-gonic/gin"
)

func ActivityInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/activity")
	r.GET("list", list)
	r.GET("info", info)
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
	activityStatus := c.Query("activityStatus")
	businessName := c.Query("businessName")
	tokenContractAddr := c.Query("tokenContractAddr")
	businessAccount := c.Query("businessAccount")
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")
	scope := c.Query("scope")
	infos, count := service.BaseService.ActivityInfoService.ActivityInfoList(businessAccount, activityStatus, businessName, tokenContractAddr, latitude, longitude, scope, pageNum, pageSize)
	page := api_result.NewPage(infos, count, pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}

func info(c *gin.Context) {
	activityIdStr := c.Query("activityId")
	activityId := bigint.StringToInt(activityIdStr)
	info := service.BaseService.ActivityInfoService.ActivityInfo(activityId)
	api_result.NewApiResult(c).Success(info)
}
