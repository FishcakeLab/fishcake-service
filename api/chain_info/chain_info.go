package chain_info

import (
	"context"
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/global_const"
	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"github.com/FishcakeLab/fishcake-service/service"
)

type SignReturnValue struct {
	Nonce                string `json:"nonce"`
	GasLimit             string `json:"gas_limit"`
	MaxFeePerGas         string `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas"`
	GasPrice             string `json:"gas_price"`
}

func ChainInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/chain_info")
	r.GET("balance", balance)
	r.GET("sign_info", signInfo)
	r.GET("send_tx", sentRawTransaction)
	r.GET("transactions", transactions)
}

func balance(c *gin.Context) {
	address := c.Query("address")
	contractAddress := c.Query("contractAddress")
	if address == "" || contractAddress == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
	}
	req := &account.AccountRequest{
		Chain:           "Polygon",
		Network:         "mainnet",
		Address:         address,
		ContractAddress: contractAddress,
	}
	response, _ := service.BaseService.RpcService.GetAccount(context.Background(), req)
	if response.Code == global_const.RpcReturnCodeSuccess {
		api_result.NewApiResult(c).Success(response.Balance)
		return
	}
	api_result.NewApiResult(c).Error(enum.GrpcErr.Code, response.Msg)
}

func signInfo(c *gin.Context) {
	address := c.Query("address")
	reqAccount := &account.AccountRequest{
		Chain:   "Polygon",
		Network: "mainnet",
		Address: address,
	}
	responseAccount, _ := service.BaseService.RpcService.GetAccount(context.Background(), reqAccount)
	if responseAccount.Code == global_const.RpcReturnCodeError {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, responseAccount.Msg)
		return
	}

	reqFee := &account.FeeRequest{
		Chain:   "Polygon",
		Network: "mainnet",
	}
	responseFee, _ := service.BaseService.RpcService.GetFee(context.Background(), reqFee)
	if responseFee.Code == global_const.RpcReturnCodeError {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, responseFee.Msg)
		return
	}
	retValue := SignReturnValue{
		Nonce:                responseAccount.Sequence,
		GasLimit:             "150000",
		MaxFeePerGas:         responseFee.FastFee,
		MaxPriorityFeePerGas: "30000000000",
		GasPrice:             responseFee.FastFee,
	}
	api_result.NewApiResult(c).Success(retValue)
	return
}

func sentRawTransaction(c *gin.Context) {
	rawTx := c.Query("rawTx")
	req := &account.SendTxRequest{
		Chain:   "Polygon",
		Network: "mainnet",
		RawTx:   rawTx,
	}
	response, _ := service.BaseService.RpcService.SendTx(context.Background(), req)
	if response.Code == global_const.RpcReturnCodeSuccess {
		api_result.NewApiResult(c).Success(response.TxHash)
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
	req := &account.TxAddressRequest{
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
