package wallet_info

import (
	"context"
	"math/big"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"github.com/FishcakeLab/fishcake-service/service"
	"github.com/FishcakeLab/fishcake-service/service/reward_service"
)

func WalletInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/wallet")
	r.GET("createWallet", CreateWallet)
}

func CreateWallet(c *gin.Context) {
	addr := c.Query("address")
	device := c.Query("device")
	if addr == "" || device == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
		return
	}

	isExist := service.BaseService.WalletService.WalletExist(addr, device)
	if isExist == false {
		api_result.NewApiResult(c).Success("This address and device has already exist, don't submit again")
		return
	}

	privateKey, address, err := service.BaseService.RewardService.DecryptPrivateKey()
	log.Info("Got decryption key:", privateKey)
	if err != nil {
		api_result.NewApiResult(c).Success("decrypt private key error")
		return
	}

	amount := new(big.Int).Mul(
		big.NewInt(50),
		new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil),
	)
	log.Info("Drop amount", "amount", amount)

	fccAddress := service.BaseService.RewardService.FccAddress()
	log.Info("Drop fccAddress", "fccAddress", fccAddress)
	reqAccount := &account.AccountRequest{
		Chain:           "Polygon",
		Network:         "mainnet",
		Address:         address.String(),
		ContractAddress: fccAddress,
	}

	accountInfo, err := service.BaseService.RpcService.GetAccount(context.Background(), reqAccount)
	if err != nil {
		api_result.NewApiResult(c).Success("get account info fail")
		return
	}
	log.Info("Drop sign nonce", "nonce", accountInfo.Sequence)
	nonce, _ := strconv.ParseUint(accountInfo.Sequence, 10, 64)

	reqFee := &account.FeeRequest{
		Chain:   "Polygon",
		Network: "mainnet",
	}
	fee, err := service.BaseService.RpcService.GetFee(context.Background(), reqFee)
	if err != nil {
		api_result.NewApiResult(c).Success("get fee fail")
		return
	}
	log.Info("Drop fee", "fee", fee.FastFee)

	parts := strings.Split(fee.FastFee, "|")
	firstNumberStr := parts[0]
	bigIntValue := new(big.Int)
	_, _ = bigIntValue.SetString(firstNumberStr, 10)

	rawTx, _, err := service.BaseService.RewardService.CreateOfflineTransaction(
		big.NewInt(137),
		reward_service.ERC20,
		privateKey,
		addr,
		nonce,
		bigIntValue,
		amount,
	)
	if err != nil {
		api_result.NewApiResult(c).Success("create offline transaction error")
		return
	}
	log.Info("sign transaction tx", "rawTx", rawTx)

	reqTx := &account.SendTxRequest{
		Chain:   "Polygon",
		Network: "mainnet",
		RawTx:   rawTx,
	}

	sendTx, err := service.BaseService.RpcService.SendTx(context.Background(), reqTx)
	if err != nil {
		log.Error("RPC send tx error: %v", err)
		api_result.NewApiResult(c).Error("4000", "Failed to send reward, Please resubmit again")
		return
	} else {
		err = service.BaseService.WalletService.StoreWallet(addr, device)
		if err != nil {
			api_result.NewApiResult(c).Error("4000", "store wallet address fail")
		}
		api_result.NewApiResult(c).Success(sendTx.TxHash)
		return
	}
}
