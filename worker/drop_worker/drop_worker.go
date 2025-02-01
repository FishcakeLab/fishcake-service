package drop_worker

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/tasks"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"github.com/FishcakeLab/fishcake-service/service"
	"github.com/FishcakeLab/fishcake-service/service/reward_service"
)

type DropWorkerProcessor struct {
	baseService *service.Service
	db          *database.DB
	tasks       tasks.Group
}

func NewDropWorkerProcessor(db *database.DB, cfg *config.Config, shutdown context.CancelCauseFunc) (*DropWorkerProcessor, error) {
	baseService := service.NewBaseService(db, cfg)
	dropWorkerProcessor := DropWorkerProcessor{
		baseService: baseService,
		db:          db,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker processor: %w", err))
		}},
	}
	return &dropWorkerProcessor, nil
}

func (dp *DropWorkerProcessor) DropWorkerStart() error {
	tickerRun := time.NewTicker(time.Second * 5)
	dp.tasks.Go(func() error {
		for range tickerRun.C {
			log.Info("drop service start.....")
			unDropAddress := dp.db.ActivityInfoDB.UnDropActivityParticipantAddresses()
			if unDropAddress != nil {
				for _, value := range unDropAddress {
					privateKey, address, err := dp.baseService.RewardService.DecryptPrivateKey()
					log.Info("Got decryption key:", privateKey)
					if err != nil {
						return err
					}

					amount := new(big.Int).Mul(
						big.NewInt(50),
						new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil),
					)

					fccAddress := dp.baseService.RewardService.FccAddress()

					reqAccount := &account.AccountRequest{
						Chain:           "Polygon",
						Network:         "mainnet",
						Address:         address.String(),
						ContractAddress: fccAddress,
					}

					accountInfo, err := dp.baseService.RpcService.GetAccount(context.Background(), reqAccount)
					if err != nil {
						log.Error("get address nonce fail", "err", err)
						return err
					}

					nonce, _ := strconv.ParseUint(accountInfo.Sequence, 10, 64)

					reqFee := &account.FeeRequest{
						Chain:   "Polygon",
						Network: "mainnet",
					}
					fee, err := dp.baseService.RpcService.GetFee(context.Background(), reqFee)
					if err != nil {
						log.Error("get fee fail", "err", err)
						return err
					}

					parts := strings.Split(fee.FastFee, "|")
					firstNumberStr := parts[0]
					bigIntValue := new(big.Int)
					_, _ = bigIntValue.SetString(firstNumberStr, 10)

					rawTx, _, err := dp.baseService.RewardService.CreateOfflineTransaction(
						big.NewInt(137),
						reward_service.ERC20,
						privateKey,
						value.Address,
						nonce,
						bigIntValue,
						amount,
					)
					if err != nil {
						return err
					}
					reqTx := &account.SendTxRequest{
						Chain:   "Polygon",
						Network: "mainnet",
						RawTx:   rawTx,
					}

					sendTx, err := dp.baseService.RpcService.SendTx(context.Background(), reqTx)
					if err != nil {
						log.Error("RPC send tx error: %v", err)
						return err
					}
					log.Info("send tx success", "txHash", sendTx.TxHash)

					errMark := dp.db.ActivityInfoDB.MarkActivityParticipantAddressDropped(value.Address)
					if errMark != nil {
						log.Error("Marked activity participant address dropped fail", "err", err)
						return errMark
					}
				}
			}
		}
		return nil
	})
	return nil
}
