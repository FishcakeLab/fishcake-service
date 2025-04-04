package drop

import (
	"errors"
	"gorm.io/gorm"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"

	_ "github.com/FishcakeLab/fishcake-service/database/utils/serializers"
)

type SystemDropInfo struct {
	Id              string   `json:"id" gorm:"id"`
	Address         string   `json:"address" gorm:"address"`
	DropAmount      *big.Int `json:"dropAmount" gorm:"serializer:u256;column:drop_amount"`
	DropType        int8     `gorm:"drop_type" json:"dropType"`
	Timestamp       uint64   `json:"timestamp" gorm:"timestamp"`
	TransactionHash string   `gorm:"transaction_hash" json:"transactionHash"`
	Status          int8     `gorm:"column:status;default:0"`
}

func (SystemDropInfo) TableName() string {
	return "system_drop_info"
}

type SystemDropInfoView interface {
	IsExist(transactionHash string, dropType int8) (error, bool)
}

type SystemDropInfoDB interface {
	SystemDropInfoView
	StoreSystemDropInfo(drop SystemDropInfo) error
	UnDropConfirmAddresses() []SystemDropInfo
	UpdateStatus(id string) error
}

type systemDropInfoDB struct {
	db *gorm.DB
}

func (d systemDropInfoDB) IsExist(transactionHash string, dropType int8) (error, bool) {
	var drop SystemDropInfo
	err := d.db.Table(drop.TableName()).Where("transaction_hash = ? and drop_type = ?", transactionHash, dropType).Take(&drop).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false
		}
		return err, false
	}
	return nil, true
}

func (d systemDropInfoDB) UnDropConfirmAddresses() []SystemDropInfo {
	var unDropConfirmAddresses []SystemDropInfo
	var drop SystemDropInfo
	// past 2 hour
	yesterday := time.Now().Add(-2 * time.Hour)
	timestamp := yesterday.Unix()
	err := d.db.Table(drop.TableName()).Where("status = ? and timestamp > ?", 0, timestamp).Find(&unDropConfirmAddresses).Error
	if err != nil {
		log.Error("Query UnDrop Confirm Addresses list fail", "err", err)
		return nil
	}
	return unDropConfirmAddresses
}

func (d systemDropInfoDB) UpdateStatus(id string) error {
	sql := `update system_drop_info set status = 1 where id = ?`
	err := d.db.Exec(sql, id).Error
	return err
}

func (d systemDropInfoDB) StoreSystemDropInfo(drop SystemDropInfo) error {
	systemDropInfo := new(SystemDropInfo)
	var exist SystemDropInfo
	err := d.db.Table(systemDropInfo.TableName()).Where("transaction_hash = ?", drop.TransactionHash).Take(&exist).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		errCreate := d.db.Table(systemDropInfo.TableName()).Omit("id").Create(&drop).Error
		if errCreate != nil {
			log.Error("create system drop fail", "err", errCreate)
			return errCreate
		}
	}
	return nil
}

func NewSystemDropInfoDB(db *gorm.DB) SystemDropInfoDB {
	return &systemDropInfoDB{db: db}
}
