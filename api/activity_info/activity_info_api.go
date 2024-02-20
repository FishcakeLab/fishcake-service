package activity_info

import (
	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/bigint"
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
	pageSize := bigint.StringToInt(pageSizeStr)
	pageNum := bigint.StringToInt(pageNumStr)
	infos, count := service.BaseService.ActivityInfoService.ActivityInfoList(pageNum, pageSize)
	page := api_result.NewPage(infos, count, pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}

func info(c *gin.Context) {
	activityIdStr := c.Query("activityId")
	activityId := bigint.StringToInt(activityIdStr)
	info := service.BaseService.ActivityInfoService.ActivityInfo(activityId)
	api_result.NewApiResult(c).Success(info)
}
