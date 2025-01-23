package wallet_info

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/FishcakeLab/fishcake-service/common/api_result"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"github.com/FishcakeLab/fishcake-service/service/reward_service"
	"github.com/FishcakeLab/fishcake-service/service/rpc_service"
	"github.com/gin-gonic/gin"
)

func WalletInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/wallet")
	r.GET("walletaddradd", walletaddradd)
}

func walletaddradd(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("walletaddradd panic: %v", err)
			panic(err) // Re-throw to be caught by Recover middleware
		}
	}()

	// Add logging for key steps
	log.Printf("Processing wallet address add request")

	addr := c.Query("addr")

	log.Printf("Starting indexer")
	cfg, err := config.New("./config.yaml")

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Printf("Failed to connect to database")
	}
	isAdd := db.WalletInfoDB.SelectWalletAddr(addr) // Check if address exists
	if isAdd == false {
		api_result.NewApiResult(c).Success("Already exists, no reward")

	} else {
		// Get decryption key
		privateKey, err := reward_service.NewRewardService("").DecryptPrivateKey()
		log.Println("Got decryption key:", privateKey)
		if err != nil {
			log.Printf("decrypt private key error: %v", err)
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
		println("txHex", txHex)

		fmt.Println(cfg.RpcUrl)

		req := &account.SendTxRequest{
			Chain:   "Polygon",
			Network: "mainnet",
			RawTx:   txHex,
		}
		fmt.Println(req.String())

		sendtx, err := rpc_service.NewRpcService(cfg.RpcUrl).SendTx(context.Background(), req)
		if err != nil {
			fmt.Printf("RPC send tx error: %v", err)
			api_result.NewApiResult(c).Error("Failed to send reward:", err.Error())
		} else {
			err = db.WalletInfoDB.AddWalletAddr(addr)
			if err != nil {
				api_result.NewApiResult(c).Error("Failed to insert into database:", err.Error())
			}
			api_result.NewApiResult(c).Success("Successfully inserted into database")
		}

		fmt.Println("sendtx.Msg:", sendtx.Msg)
	}
}
