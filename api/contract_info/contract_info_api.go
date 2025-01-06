package contract_info

import (
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/service"
)

func ContractInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/contract")
	r.GET("info", info)
}

func info(c *gin.Context) {
	infos := service.BaseService.ContractInfoService.ContractInfo()
	api_result.NewApiResult(c).Success(infos)
}
