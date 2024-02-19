package activity_info

import (
	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/FishcakeLab/fishcake-service/database/activity"
	"github.com/FishcakeLab/fishcake-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ActivityInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/activity")
	r.GET("list", list)
	r.GET("info", info)
	r.POST("activityAdd", activityAdd)
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

func activityAdd(c *gin.Context) {
	var activityInfo activity.ActivityInfo
	if err := c.ShouldBindJSON(&activityInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	info := service.BaseService.ActivityInfoService.AddActivityInfo(activityInfo)
	api_result.NewApiResult(c).Success(info)
}
