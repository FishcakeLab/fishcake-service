package chain_info

import (
	"context"
	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/global_const"
	"github.com/FishcakeLab/fishcake-service/service"
)

func ChainInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/chain_info")
	r.GET("balance", balance)
	r.GET("transactions", transactions)
}

func balance(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
	}

	req :=
		&account.AccountRequest{
			Chain:   "Polygon",
			Network: "mainnet",
			Address: address,
		} // 目标链

	response, _ := service.BaseService.RpcService.GetAccount(context.Background(), req)
	if response.Code == global_const.RpcReturnCodeSuccess {
		api_result.NewApiResult(c).Success(response.Balance)
		return
	}
	api_result.NewApiResult(c).Error(enum.GrpcErr.Code, response.Msg)
}

func transactions(c *gin.Context) {
	address := c.Query("address")
	contractAddress := c.Query("contractAddress")
	if address == "" || contractAddress == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
	}
	req :=
		&account.TxAddressRequest{
			Chain:   "Polygon",
			Network: "mainnet",
			Address: address,
		}

	response, _ := service.BaseService.RpcService.GetTxByAddress(context.Background(), req)
	if response.Code == global_const.RpcReturnCodeSuccess {
		api_result.NewApiResult(c).Success(response)
		return
	}
	api_result.NewApiResult(c).Error(enum.GrpcErr.Code, response.Msg)
}
