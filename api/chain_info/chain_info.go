package chain_info

import (
	"context"
	"github.com/gin-gonic/gin"
	"math/big"
	"strings"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/global_const"
	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"github.com/FishcakeLab/fishcake-service/service"
)

type SignReturnValue struct {
	Nonce                string `json:"nonce"`
	NativeTokenGasLimit  string `json:"native_token_gas_limit"`
	Erc20TokenGasLimit   string `json:"erc20_token_gas_limit"`
	MaxFeePerGas         string `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas"`
	GasPrice             string `json:"gas_price"`
}

type BalanceResultValue struct {
	PolBalance  string `json:"pol_balance"`
	UsdtBalance string `json:"usdt_balance"`
	FccBalance  string `json:"fcc_balance"`
}

func ChainInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/chain_info")
	r.GET("balance", balance)
	r.GET("sign_info", signInfo)
	r.GET("send_tx", sentRawTransaction)
	r.GET("transactions", transactions)
	r.POST("send_sign", sendSign)
}

func balance(c *gin.Context) {
	address := c.Query("address")

	if address == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
	}
	reqFcc := &account.AccountRequest{
		Chain:           "Polygon",
		Network:         "mainnet",
		Address:         address,
		ContractAddress: service.BaseService.RewardService.FccAddress(),
	}
	responseFcc, _ := service.BaseService.RpcService.GetAccount(context.Background(), reqFcc)
	if responseFcc.Code == global_const.RpcReturnCodeError {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, responseFcc.Msg)
		return
	}

	reqPol := &account.AccountRequest{
		Chain:           "Polygon",
		Network:         "mainnet",
		Address:         address,
		ContractAddress: "0x00",
	}
	responsePol, _ := service.BaseService.RpcService.GetAccount(context.Background(), reqPol)
	if responsePol.Code == global_const.RpcReturnCodeError {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, responsePol.Msg)
		return
	}

	reqUsdt := &account.AccountRequest{
		Chain:           "Polygon",
		Network:         "mainnet",
		Address:         address,
		ContractAddress: service.BaseService.RewardService.UsdtAddress(),
	}
	responseUsdt, _ := service.BaseService.RpcService.GetAccount(context.Background(), reqUsdt)
	if responseUsdt.Code == global_const.RpcReturnCodeError {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, responseUsdt.Msg)
		return
	}

	balanceRet := BalanceResultValue{
		PolBalance:  responsePol.Balance,
		UsdtBalance: responseUsdt.Balance,
		FccBalance:  responseFcc.Balance,
	}
	api_result.NewApiResult(c).Success(balanceRet)
	return
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

	parts := strings.Split(responseFee.FastFee, "|")
	// Extract first part and convert to big.Int
	firstNumberStr := parts[0]
	bigIntValue := new(big.Int)
	_, _ = bigIntValue.SetString(firstNumberStr, 10)

	retValue := SignReturnValue{
		Nonce:                responseAccount.Sequence,
		NativeTokenGasLimit:  "21000",
		Erc20TokenGasLimit:   "150000",
		MaxFeePerGas:         bigIntValue.String(),
		MaxPriorityFeePerGas: "30000000000",
		GasPrice:             bigIntValue.String(),
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
	if response == nil {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, "send tx fail, please try again later")
		return
	}
	if response.Code == global_const.RpcReturnCodeError {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, response.Msg)
		return
	}
	api_result.NewApiResult(c).Success(response.TxHash)
	return
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

type Offline struct {
	RawTx string
}

func sendSign(c *gin.Context) {
	var offline Offline
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&offline); err != nil {
		api_result.NewApiResult(c).Error("400", "param parse fail")
		return
	}
	exist := service.BaseService.SignService.IsExistSign(offline.RawTx)
	if exist {
		api_result.NewApiResult(c).Error("400", "sign is already exist")
		return
	}
	err := service.BaseService.SignService.StoreSign(offline.RawTx)
	if err != nil {
		api_result.NewApiResult(c).Error("400", "store sign fail")
		return
	}
	api_result.NewApiResult(c).Success("ok")
}
