package nft_info

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/service"
)

func BoosterInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/booster")
	r.GET("list", listBooster)
}

func listBooster(c *gin.Context) {
	address := strings.ToLower(c.Query("address"))
	if address == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
		return
	}

	pageNum := 1
	pageSize := 50

	if v := c.Query("pageNum"); v != "" {
		pageNum = bigint.StringToInt(v)
	}
	if v := c.Query("pageSize"); v != "" {
		pageSize = bigint.StringToInt(v)
	}

	infos, count := service.BaseService.BoosterService.
		ListByAddress(address, pageNum, pageSize)

	page := api_result.NewPage(infos, count, pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}
