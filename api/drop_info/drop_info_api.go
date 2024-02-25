package drop_info

import (
	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/service"
)

func DropInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/drop")
	r.GET("list", list)
}

func list(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageNumStr := c.Query("pageNum")
	address := c.Query("address")
	pageSize := bigint.StringToInt(pageSizeStr)
	pageNum := bigint.StringToInt(pageNumStr)
	infos, count := service.BaseService.DropService.DropInfoList(pageNum, pageSize, address)
	page := api_result.NewPage(infos, count, pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}
