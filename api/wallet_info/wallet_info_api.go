package wallet_info

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"github.com/FishcakeLab/fishcake-service/service/reward_service"
	"github.com/FishcakeLab/fishcake-service/service/rpc_service"
)

func WalletInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/wallet")
	r.GET("walletaddradd", walletaddradd)
}

func walletaddradd(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Info("walletaddradd panic: %v", err)
			panic(err) // Re-throw to be caught by Recover middleware
		}
	}()

	// Add logging for key steps
	log.Info("Processing wallet address add request")

	addr := c.Query("addr")

	log.Info("Starting indexer")
	cfg, err := config.New("./config.yaml")

	db, err := database.NewDB(cfg)
	if err != nil {
		api_result.NewApiResult(c).Success("Failed to connect to database")
		return
	}
	isAdd := db.WalletInfoDB.SelectWalletAddr(addr) // Check if address exists
	if isAdd == false {
		api_result.NewApiResult(c).Success("Already exists, no reward")
		return
	} else {
		// Get decryption key
		privateKey, err := reward_service.NewRewardService("").DecryptPrivateKey()
		log.Info("Got decryption key:", privateKey)
		if err != nil {
			api_result.NewApiResult(c).Success("decrypt private key error")
			return
		}

		// Calculate reward amount (50 * 10^6)
		amount := new(big.Int).Mul(
			big.NewInt(50),
			new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil),
		)

		// Create and send transaction
		txHex, _, err := reward_service.NewRewardService("").CreateOfflineTransaction(
			reward_service.FCC,
			privateKey,
			addr,
			amount,
		)
		if err != nil {
			api_result.NewApiResult(c).Success("create offline transaction error")
			return
		}
		req := &account.SendTxRequest{
			Chain:   "Polygon",
			Network: "mainnet",
			RawTx:   txHex,
		}
		sendtx, err := rpc_service.NewRpcService(cfg.RpcUrl).SendTx(context.Background(), req)
		if err != nil {
			log.Info("RPC send tx error: %v", err)
			api_result.NewApiResult(c).Error("Failed to send reward:", err.Error())
			return
		} else {
			err = db.WalletInfoDB.AddWalletAddr(addr)
			if err != nil {
				api_result.NewApiResult(c).Error("Failed to insert into database:", err.Error())
			}
			api_result.NewApiResult(c).Success("Successfully inserted into database")
			return
		}

		log.Info("sendtx.Msg:", sendtx.Msg)
	}
}
