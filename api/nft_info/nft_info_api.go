package nft_info

import (
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/service"
)

func NftInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/nft")
	r.GET("list", list)
	r.GET("info", info)
	r.GET("detail", detail)
	r.GET("nft_count", nftInfo)
}

func list(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageNumStr := c.Query("pageNum")
	if pageSizeStr == "" || pageNumStr == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
		return
	}
	contractAddress := c.Query("contractAddress")
	address := c.Query("address")
	pageSize := bigint.StringToInt(pageSizeStr)
	pageNum := bigint.StringToInt(pageNumStr)
	infos, count := service.BaseService.NftService.NftInfoList(pageNum, pageSize, contractAddress, address)
	page := api_result.NewPage(infos, count, pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}

func nftInfo(c *gin.Context) {
	contractAddress := c.Query("contractAddress")
	count := service.BaseService.NftService.NftCount(contractAddress)
	api_result.NewApiResult(c).Success(count)
}

func info(c *gin.Context) {
	tokenIdStr := c.Query("tokenId")
	tokenId := bigint.StringToInt(tokenIdStr)
	info := service.BaseService.NftService.NftInfo(tokenId)
	api_result.NewApiResult(c).Success(info)
}

func detail(c *gin.Context) {
	businessAccount := c.Query("businessAccount")
	deadline := c.Query("deadline")
	info := service.BaseService.NftService.NftDetail(businessAccount, deadline)
	api_result.NewApiResult(c).Success(info)
}
