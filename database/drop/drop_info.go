package drop

import (
	"context"
	"errors"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database/activity"
	_ "github.com/FishcakeLab/fishcake-service/database/utils/serializers"
	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"github.com/FishcakeLab/fishcake-service/service/reward_service"
	"github.com/FishcakeLab/fishcake-service/service/rpc_service"
	"gorm.io/gorm"
	"log"
	"math/big"
	"time"
)

type DropInfo struct {
	Id                string   `json:"id" gorm:"id"`
	ActivityId        int64    `gorm:"activity_id" json:"activityId"`
	Address           string   `json:"address" gorm:"address"`
	DropAmount        *big.Int `json:"dropAmount" gorm:"serializer:u256;column:drop_amount"`
	DropType          int8     `gorm:"drop_type" json:"dropType"`
	Timestamp         uint64   `json:"timestamp" gorm:"timestamp"`
	TokenContractAddr string   `gorm:"token_contract_addr" json:"tokenContractAddr"`
	BusinessName      string   `gorm:"business_name" json:"businessName"`
	TransactionHash   string   `gorm:"transaction_hash" json:"transactionHash"`
	EventSignature    string   `gorm:"event_signature" json:"eventSignature"`
	ReturnAmount      *big.Int `gorm:"serializer:u256;column:return_amount" json:"returnAmount"`
	MinedAmount       *big.Int `gorm:"serializer:u256;column:mined_amount" json:"minedAmount"`
}

func (DropInfo) TableName() string {
	return "drop_info"
}

type DropInfoView interface {
	List(pageNum, pageSize int, address, dropType string) ([]DropInfo, int)
	IsExist(transactionHash, eventSignature string, dropType int8) (error, bool)
}

type DropInfoDB interface {
	DropInfoView
	StoreDropInfo(drop DropInfo) error
}

type dropInfoDB struct {
	db *gorm.DB
}

func (d dropInfoDB) IsExist(transactionHash, eventSignature string, dropType int8) (error, bool) {
	var drop DropInfo
	err := d.db.Table(drop.TableName()).Where("transaction_hash = ? and event_signature = ? and drop_type = ?", transactionHash, eventSignature, dropType).Take(&drop).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false
		}
		return err, false
	}
	return nil, true
}

func (d dropInfoDB) StoreDropInfo(drop DropInfo) error {
	drpoInfo := new(DropInfo)
	var exist DropInfo
	err := d.db.Table(drpoInfo.TableName()).Where("transaction_hash = ? and event_signature = ? and drop_type = ?", drop.TransactionHash, drop.EventSignature, drop.DropType).Take(&exist).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var existAddress activity.ActivityParticipantAddress

			err := d.db.Table("activity_participants_addresses").
				Where("address = ?", drop.Address).
				Take(&existAddress).Error

			if err != nil && drop.DropType == 1 {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					existAddress = activity.ActivityParticipantAddress{
						ActivityId: drop.ActivityId,
						Address:    drop.Address,
						JoinTime:   time.Now().Unix(),
					}
					if err := d.db.Table("activity_participants_addresses").Create(&existAddress).Error; err != nil {
						return err
					}
					privateKey, err := reward_service.NewRewardService("").DecryptPrivateKey() //获取解密密钥
					if err != nil {
						log.Printf("decrypt private key error: %v", err)

					}

					amount := new(big.Int).Mul(
						big.NewInt(50),
						new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil),
					)

					txHex, txHash, err := reward_service.NewRewardService("").CreateOfflineTransaction(
						reward_service.FCC,
						privateKey,
						drop.Address,
						amount,
					)
					log.Println(txHex, txHash)
					req :=
						&account.SendTxRequest{
							Chain:   "Polygon",
							Network: "mainnet",
							RawTx:   txHex,
						} // 目标链
					cfg, err := config.New("")

					sendtx, _ := rpc_service.NewRpcService(cfg.RpcUrl).SendTx(context.Background(), req)
					log.Println(sendtx.TxHash)

				}
			}
			result := d.db.Table(drpoInfo.TableName()).Omit("id, token_contract_addr, business_name, return_amount, mined_amount").Create(&drop)
			return result.Error
		}
	}
	return err
}

func (d dropInfoDB) List(pageNum, pageSize int, address, dropType string) ([]DropInfo, int) {
	var tokenNft []DropInfo
	var count int64
	this := d.db.Table(DropInfo{}.TableName())
	if address != "" {
		this = this.Where("address ILIKE ?", address)
	}
	if dropType != "" {
		this = this.Where("drop_info.drop_type = ?", dropType)
	}
	this = this.Joins("LEFT JOIN activity_info ON drop_info.activity_id = activity_info.activity_id")
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	this = this.Count(&count)
	result := this.Select("drop_info.*, activity_info.token_contract_addr, activity_info.business_name, activity_info.return_amount, activity_info.mined_amount").Scan(&tokenNft)
	if result.Error == nil {
		return tokenNft, int(count)
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errors_h.NewErrorByEnum(enum.DataErr)
		return nil, int(count)
	} else {
		return nil, int(count)
	}
}

func NewDropInfoDB(db *gorm.DB) DropInfoDB {
	return &dropInfoDB{db: db}
}
