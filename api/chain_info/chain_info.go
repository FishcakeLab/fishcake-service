package chain_info

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"math/big"
	"strconv"
	"strings"
	"sync"

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

type TransactionInfo struct {
	RawTx string `json:"raw_tx"`
}

func ChainInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/chain_info")
	r.GET("balance", balance)
	r.GET("balance_sync", balanceSync)
	r.GET("sign_info", signInfo)
	r.GET("send_tx", sentRawTransaction)
	r.GET("transactions", transactions)
	r.POST("submit_tx", submitTx)
	r.GET("txn_status", TxnStatus)
}

func balanceSync(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
		return
	}

	type balResult struct {
		balance string
		errCode int
		errMsg  string
	}

	var (
		wg      sync.WaitGroup
		resPol  balResult
		resUsdt balResult
		resFcc  balResult
	)
	wg.Add(3)

	// Polygon 主币余额
	go func() {
		defer wg.Done()
		balanceStr, err := service.BaseService.DappLinkService.GetPolBalanceByAddress(address)
		if err != nil {
			resPol.errCode = 500
			resPol.errMsg = err.Error()
			return
		}
		resPol.balance = balanceStr
	}()

	// USDT 余额
	go func() {
		defer wg.Done()
		balanceStr, err := service.BaseService.DappLinkService.GetErc20BalanceByAddress(service.BaseService.RewardService.UsdtAddress(), address)
		if err != nil {
			resUsdt.errCode = 500
			resUsdt.errMsg = err.Error()
			return
		}
		resUsdt.balance = balanceStr
	}()

	// FCC 余额
	go func() {
		defer wg.Done()
		responseFcc, err := service.BaseService.DappLinkService.GetErc20BalanceByAddress(service.BaseService.RewardService.FccAddress(), address)
		if err != nil {
			resFcc.errCode = 500
			resFcc.errMsg = err.Error()
			return
		}
		resFcc.balance = responseFcc
	}()

	wg.Wait()

	// 任意一个失败就直接返回
	if resPol.errCode != 0 {
		api_result.NewApiResult(c).Error(strconv.Itoa(resPol.errCode), resPol.errMsg)
		return
	}
	if resUsdt.errCode != 0 {
		api_result.NewApiResult(c).Error(strconv.Itoa(resUsdt.errCode), resUsdt.errMsg)
		return
	}
	if resFcc.errCode != 0 {
		api_result.NewApiResult(c).Error(strconv.Itoa(resFcc.errCode), resFcc.errMsg)
		return
	}

	balanceRet := BalanceResultValue{
		PolBalance:  resPol.balance,
		UsdtBalance: resUsdt.balance,
		FccBalance:  resFcc.balance,
	}
	api_result.NewApiResult(c).Success(balanceRet)
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

func submitTx(c *gin.Context) {
	var txInfo TransactionInfo
	if err := c.ShouldBindJSON(&txInfo); err != nil {
		api_result.NewApiResult(c).Error("400", "param parse fail")
		return
	}
	exist := service.BaseService.WalletService.IsExistRawTx(txInfo.RawTx)
	if !exist {
		api_result.NewApiResult(c).Error("400", "raw tx is already exist queue tx")
		return
	}
	err := service.BaseService.WalletService.StoreRawTx(txInfo.RawTx)
	if err != nil {
		api_result.NewApiResult(c).Error("400", "store raw tx to queue tx fail")
		return
	}
	api_result.NewApiResult(c).Success("ok")
}

func TxnStatus(c *gin.Context) {
	txHash := c.Query("hash")
	if txHash == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
	}
	queueTx, err := service.BaseService.WalletService.QueryTxInfoByHash(txHash)
	if err != nil {
		log.Error("query transaction info fail", "err", err)
		api_result.NewApiResult(c).Error("400", "query tx fail")
		return
	}
	api_result.NewApiResult(c).Success(queueTx)
}
