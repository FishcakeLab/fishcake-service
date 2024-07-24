package activity_info

import (
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/FishcakeLab/fishcake-service/service"
)

func ActivityInfoExtApi(rg *gin.Engine) {
	r := rg.Group("/v1/activityExt")
	r.GET("list", extList)
	r.GET("info", extInfo)
}

func extList(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageNumStr := c.Query("pageNum")
	pageSize := bigint.StringToInt(pageSizeStr)
	pageNum := bigint.StringToInt(pageNumStr)
	infos, count := service.BaseService.ActivityInfoExtService.ActivityInfoExtList(pageNum, pageSize)
	page := api_result.NewPage(infos, count, pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}

func extInfo(c *gin.Context) {
	activityIdStr := c.Query("activityId")
	activityId := bigint.StringToInt(activityIdStr)
	info := service.BaseService.ActivityInfoExtService.ActivityInfoExt(activityId)
	api_result.NewApiResult(c).Success(info)
}
