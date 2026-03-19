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

// listBooster godoc
// @Summary User Booster NFT list
// @Description Query Booster NFTs owned by a user. nftType >= 10 means used by staking (original type + 10)
// @Tags Earnings
// @Accept json
// @Produce json
// @Param address query string true "Wallet address"
// @Param pageNum query int false "Page number, default 1"
// @Param pageSize query int false "Page size, default 50, max 100"
// @Success 200 {object} api_result.ApiResult
// @Router /v1/booster/list [get]
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
