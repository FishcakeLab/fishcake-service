package drop_info

import (
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/bigint"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/service"
)

func DropInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/drop")
	r.GET("list", list)
}

// list godoc
// @Summary Drop/Receive records
// @Description Query drop/receive records, corresponds to contract Drop events
// @Tags History
// @Accept json
// @Produce json
// @Param pageNum query int true "Page number, starts from 1"
// @Param pageSize query int true "Page size"
// @Param address query string false "Wallet address (case insensitive)"
// @Param dropType query string false "1=user received, 2=merchant dropped"
// @Success 200 {object} api_result.ApiResult
// @Router /v1/drop/list [get]
func list(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageNumStr := c.Query("pageNum")
	if pageSizeStr == "" || pageNumStr == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
		return
	}
	address := c.Query("address")
	dropType := c.Query("dropType")
	pageSize := bigint.StringToInt(pageSizeStr)
	pageNum := bigint.StringToInt(pageNumStr)
	infos, count := service.BaseService.DropService.DropInfoList(pageNum, pageSize, address, dropType)
	page := api_result.NewPage(infos, count, pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}
