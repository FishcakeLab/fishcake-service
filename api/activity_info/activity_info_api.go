package activity_info

import (
	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/service"
)

func ActivityInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/activity_info")
	r.GET("list", list)
}

func list(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageNumStr := c.Query("pageNum")
	pageSize := bigint.StringToInt(pageSizeStr)
	pageNum := bigint.StringToInt(pageNumStr)
	infos, count := service.BaseService.ActivityInfoService.List(pageNum, pageSize)
	page := api_result.NewPage(infos, count, pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}

//func
