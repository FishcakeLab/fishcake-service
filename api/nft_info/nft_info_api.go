package nft_info

import (
	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/service"
)

func NftInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/nft")
	r.GET("list", list)
	r.GET("info", info)
}

func list(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageNumStr := c.Query("pageNum")
	contractAddress := c.Query("contractAddress")
	pageSize := bigint.StringToInt(pageSizeStr)
	pageNum := bigint.StringToInt(pageNumStr)
	infos, count := service.BaseService.NftService.NftInfoList(pageNum, pageSize, contractAddress)
	page := api_result.NewPage(infos, count, pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}

func info(c *gin.Context) {
	activityIdStr := c.Query("activityId")
	activityId := bigint.StringToInt(activityIdStr)
	info := service.BaseService.NftService.NftInfo(activityId)
	api_result.NewApiResult(c).Success(info)
}
