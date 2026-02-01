package chain_info

import (
	"context"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/enum"
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
	TxHash string `json:"tx_hash"`
	RawTx  string `json:"raw_tx"`
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
	if responseFcc.Code == account.ReturnCode_ERROR {
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
	if responsePol.Code == account.ReturnCode_ERROR {
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
	if responseUsdt.Code == account.ReturnCode_ERROR {
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
	if address == "" {
		api_result.NewApiResult(c).Error("400", "missing address")
		return
	}

	// 1) 获取 nonce
	reqAccount := &account.AccountRequest{
		Chain:   "Polygon",
		Network: "mainnet",
		Address: address,
	}
	responseAccount, _ := service.BaseService.RpcService.GetAccount(context.Background(), reqAccount)
	if responseAccount.Code == account.ReturnCode_ERROR {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, responseAccount.Msg)
		return
	}

	// 2) feeHistory：用更多区块 + 适中分位数（别用 95）
	// blockCount 太小 + 取最后一块 = 极易被尖峰污染
	const blockCount = 20
	percentiles := []int{50, 75} // 更稳（50/75 通常比 95 稳定很多）

	var feeHist struct {
		OldestBlock   string     `json:"oldestBlock"`
		BaseFeePerGas []string   `json:"baseFeePerGas"`
		GasUsedRatio  []float64  `json:"gasUsedRatio"`
		Reward        [][]string `json:"reward"`
	}

	err := service.BaseService.Client.CallContext(
		context.Background(),
		&feeHist,
		"eth_feeHistory",
		blockCount,
		"latest",
		percentiles,
	)
	if err != nil {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, err.Error())
		return
	}

	parseHexBig := func(hexStr string) (*big.Int, bool) {
		s := strings.TrimPrefix(hexStr, "0x")
		if s == "" {
			return nil, false
		}
		v, ok := new(big.Int).SetString(s, 16)
		return v, ok
	}

	// 3) baseFee：取最新 baseFee
	if len(feeHist.BaseFeePerGas) == 0 {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, "invalid baseFee data")
		return
	}
	latestBaseFeeHex := feeHist.BaseFeePerGas[len(feeHist.BaseFeePerGas)-1]
	baseFee, ok := parseHexBig(latestBaseFeeHex)
	if !ok || baseFee.Sign() <= 0 {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, "invalid baseFee hex")
		return
	}

	// 4) tip：从多个区块里取“中位数”，而不是只取最后一块
	// feeHist.Reward 的长度通常是 blockCount（每块一个 reward 数组）
	// 每个 reward 数组的长度等于 len(percentiles)
	getMedian := func(vals []*big.Int) *big.Int {
		sort.Slice(vals, func(i, j int) bool { return vals[i].Cmp(vals[j]) < 0 })
		return new(big.Int).Set(vals[len(vals)/2])
	}

	// 从 reward 里拿 percentile[75] 对应的 tip（索引 1）
	var tipSamples []*big.Int
	for _, r := range feeHist.Reward {
		if len(r) < 2 { // 需要 50/75 两个分位数
			continue
		}
		v, ok := parseHexBig(r[1]) // 75th
		if ok && v.Sign() > 0 {
			tipSamples = append(tipSamples, v)
		}
	}

	// fallback tip：2 gwei
	gwei := big.NewInt(1_000_000_000)
	tip := new(big.Int).Mul(big.NewInt(2), gwei)
	if len(tipSamples) > 0 {
		tip = getMedian(tipSamples)
	}

	// 5) 强制 clamp：Polygon 上 priority fee 正常 1~5 gwei，给你到 10 gwei 已经很宽
	maxTip := new(big.Int).Mul(big.NewInt(10), gwei)
	if tip.Cmp(maxTip) > 0 {
		tip = maxTip
	}

	// 6) maxFee：2*baseFee + tip，并做 cap（防止 baseFee/节点异常）
	maxFee := new(big.Int).Mul(baseFee, big.NewInt(2))
	maxFee.Add(maxFee, tip)

	// maxFeeCap：给到 200 gwei 已经非常宽（按你业务也可调 100/150/200）
	maxFeeCap := new(big.Int).Mul(big.NewInt(200), gwei)
	if maxFee.Cmp(maxFeeCap) > 0 {
		maxFee = maxFeeCap
	}

	// 7) gasPrice：如果你还要给 legacy 用，给一个“baseFee + tip”即可（但不要叫 slow/normal 去误导前端）
	legacyGasPrice := new(big.Int).Add(baseFee, tip)
	legacyCap := new(big.Int).Mul(big.NewInt(200), gwei)
	if legacyGasPrice.Cmp(legacyCap) > 0 {
		legacyGasPrice = legacyCap
	}

	// 8) gasLimit 建议：approve/transfer 不要拍太大
	// approve 通常 70k～120k；你先给 100k 更合理
	retValue := SignReturnValue{
		Nonce:                responseAccount.Sequence, // 你之前返回是 string，这里保持原类型
		NativeTokenGasLimit:  "21000",
		Erc20TokenGasLimit:   "100000",
		MaxFeePerGas:         maxFee.String(),
		MaxPriorityFeePerGas: tip.String(),
		GasPrice:             legacyGasPrice.String(),
	}

	api_result.NewApiResult(c).Success(retValue)
}

func sentRawTransaction(c *gin.Context) {
	rawTx := c.Query("rawTx")
	req := &account.SendTxRequest{
		Chain:   "Polygon",
		Network: "mainnet",
		RawTx:   rawTx,
	}
	exist := service.BaseService.WalletService.IsExistRawTx(rawTx)
	if exist {
		api_result.NewApiResult(c).Error("400", "raw tx is already exist queue tx")
		return
	}
	response, _ := service.BaseService.RpcService.SendTx(context.Background(), req)
	if response == nil {
		api_result.NewApiResult(c).Error(enum.GrpcErr.Code, "send tx fail, please try again later")
		return
	}
	if response.Code == account.ReturnCode_ERROR {
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
	if response.Code == account.ReturnCode_ERROR {
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
	if exist {
		api_result.NewApiResult(c).Error("400", "raw tx is already exist queue tx")
		return
	}
	err := service.BaseService.WalletService.StoreRawTx(txInfo.RawTx, txInfo.TxHash)
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
