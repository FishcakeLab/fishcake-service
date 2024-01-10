package chain_info

import (
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/global_const"
	"github.com/FishcakeLab/fishcake-service/service"
)

func ChainInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/chain_info")
	r.GET("balance", balance)
}

func balance(c *gin.Context) {
	address := c.Query("address")
	//contractAddress := c.Query("contractAddress")
	response := service.BaseService.RpcService.GetBalance(address)
	if response.Code == global_const.RpcReturnCodeSUCCESS {
		api_result.NewApiResult(c).Success(response.Balance)
		return
	}
	api_result.NewApiResult(c).Error(enum.GrpcErr.Code, response.Msg)
}
