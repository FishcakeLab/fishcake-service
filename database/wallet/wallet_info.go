package wallet

import (
	"errors"
	"time"

	_ "github.com/FishcakeLab/fishcake-service/database/utils/serializers"
	"gorm.io/gorm"
)

type WalletAddress struct {
	ID        string `gorm:"column:id;primaryKey;default:replace((uuid_generate_v4())::text, '-'::text, ''::text)"`
	Address   string `gorm:"column:address;unique;not null"`
	Device    string `gorm:"column:device;unique;not null"`
	CreatedAt int64  `gorm:"column:created_at;not null"`
	Status    int8   `gorm:"column:status;default:1"`
	Remark    string `gorm:"column:remark"`
}

type WalletInfoDB interface {
	StoreWalletAddr(string, string) error
	ExistWalletAddr(string, string) bool
}

type walletInfoDB struct {
	db *gorm.DB
}

func NewWalletInfoDB(db *gorm.DB) WalletInfoDB {
	return &walletInfoDB{db: db}
}

func (w walletInfoDB) StoreWalletAddr(addr string, device string) error {
	existAddress := WalletAddress{
		Address:   addr,
		Device:    device,
		CreatedAt: time.Now().Unix(),
	}
	if err := w.db.Table("wallet_addresses").Create(&existAddress).Error; err != nil {
		return err
	}
	return nil
}

func (w walletInfoDB) ExistWalletAddr(addr string, device string) bool {
	var walletAddress WalletAddress
	this := w.db.Table("wallet_addresses")
	result := this.Where("address = ? and device = ?", addr, device).Take(&walletAddress)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}
