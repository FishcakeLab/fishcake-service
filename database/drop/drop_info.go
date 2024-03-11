package drop

import (
	"errors"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	_ "github.com/FishcakeLab/fishcake-service/database/utils/serializers"
	"gorm.io/gorm"
	"math/big"
)

type DropInfo struct {
	Id         string   `json:"id" gorm:"id"`
	ActivityId int64    `gorm:"activity_id" json:"activityId"`
	Address    string   `json:"address" gorm:"address"`
	DropAmount *big.Int `json:"dropAmount" gorm:"serializer:u256;column:drop_amount"`
	Timestamp  uint64   `json:"timestamp" gorm:"timestamp"`
}

func (DropInfo) TableName() string {
	return "drop_info"
}

type DropInfoView interface {
	List(pageNum, pageSize int, Address string) ([]DropInfo, int)
}

type DropInfoDB interface {
	DropInfoView
	StoreDropInfo(drop DropInfo) error
}

type dropInfoDB struct {
	db *gorm.DB
}

func (d dropInfoDB) StoreDropInfo(drop DropInfo) error {
	drpoInfo := new(DropInfo)
	result := d.db.Table(drpoInfo.TableName()).Omit("id").Create(&drop)
	return result.Error
}

func (d dropInfoDB) List(pageNum, pageSize int, address string) ([]DropInfo, int) {
	var tokenNft []DropInfo
	var count int64
	this := d.db.Table(DropInfo{}.TableName())
	if address != "" {
		this = this.Where("address = ?", address)
	}
	this = this.Count(&count)
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	result := this.Find(&tokenNft)
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
