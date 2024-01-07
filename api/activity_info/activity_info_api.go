package activity_info

import (
	"github.com/FishcakeLab/fishcake-service/api/service"
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
)

func ActivityInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/activity_info")
	r.GET("list", list)
}

func list(c *gin.Context) {
	pageSize := c.Query("")
	pageNo := c.Query("")

	infos, count := service.BaseService.ActivityInfoService.List(pageSize, pageNo)
	api_result.NewApiResult(c).Success(gin.H{
		"result": infos,
		"total":  count,
	})

}
