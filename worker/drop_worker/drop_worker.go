package drop_worker

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/tasks"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/drop"
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
					if value.Address == "" {
						continue
					}
					privateKey, address, err := dp.baseService.RewardService.DecryptPrivateKey()
					log.Info("Got decryption key:", privateKey)
					if err != nil {
						return err
					}

					amount := new(big.Int).Mul(
						big.NewInt(10),
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
					log.Info("RpcService GetFee", "fee", fee)

					if fee.Code == account.ReturnCode_ERROR {
						log.Error("get fee fail code is", account.ReturnCode_ERROR)
						continue
					}

					log.Info("RpcService GetFee Eip1559Wallet MaxFeePerGas", "MaxFeePerGas", fee.Eip1559Wallet.MaxFeePerGas)

					bigIntValue := new(big.Int)
					_, _ = bigIntValue.SetString(fee.Eip1559Wallet.MaxFeePerGas, 10)

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
					if sendTx != nil {
						log.Info("send tx success", "txHash", sendTx.TxHash)
					} else {
						continue
					}
					errMark := dp.db.ActivityInfoDB.MarkActivityParticipantAddressDropped(value.Address)
					if errMark != nil {
						log.Error("Marked activity participant address dropped fail", "err", err)
						return errMark
					}

					//store system drop info
					systemDropInfo := drop.SystemDropInfo{
						Address:         value.Address,
						DropAmount:      amount,
						DropType:        2,
						Timestamp:       uint64(time.Now().Unix()),
						TransactionHash: sendTx.TxHash,
					}
					resultErr, exist := dp.db.SystemDropInfoDB.IsExist(systemDropInfo.TransactionHash, systemDropInfo.DropType)
					if !exist && resultErr == nil {
						if err := dp.db.SystemDropInfoDB.StoreSystemDropInfo(systemDropInfo); err != nil {
							return err
						}
					} else {
						return resultErr
					}
					return nil

				}
			}
		}
		return nil
	})
	return nil
}
