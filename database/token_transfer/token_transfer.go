package token_transfer

import (
	"math/big"

	"gorm.io/gorm"
)

type TokenSent struct {
	Id           string   `json:"id" gorm:"id"`
	Address      string   `json:"address" gorm:"address"`
	TokenAddress string   `json:"token_address" gorm:"token_address"`
	Description  string   `json:"description" gorm:"description"`
	Amount       *big.Int `json:"amount" gorm:"serializer:u256;column:amount"`
	Timestamp    uint64   `json:"timestamp" gorm:"timestamp"`
}

type TokenReceived struct {
	Id           string   `json:"id" gorm:"id"`
	Address      string   `json:"address" gorm:"address"`
	TokenAddress string   `json:"token_address" gorm:"token_address"`
	Description  string   `json:"description" gorm:"description"`
	Amount       *big.Int `json:"amount" gorm:"serializer:u256;column:amount"`
	Timestamp    uint64   `json:"timestamp" gorm:"timestamp"`
}

func (TokenSent) TableName() string {
	return "token_sent"
}

func (TokenReceived) TableName() string {
	return "token_received"
}

type TokenSentView interface {
	List(address, tokenType string, lastTimestamp uint64, limit int) ([]TokenSent, error)
}

type TokenSentDB interface {
	TokenSentView
	StoreTokenSent(tokenSent TokenSent) error
}

type tokenSentDB struct {
	db *gorm.DB
}

type TokenReceivedView interface {
	List(address, tokenType string, lastTimestamp uint64, limit int) ([]TokenReceived, error)
}

type TokenReceivedDB interface {
	TokenReceivedView
	StoreTokenReceived(tokenReceived TokenReceived) error
}

type tokenReceivedDB struct {
	db *gorm.DB
}

func (ts tokenSentDB) List(address, tokenType string, lastTimestamp uint64, limit int) ([]TokenSent, error) {
	var records []TokenSent

	query := ts.db.Table(TokenSent{}.TableName()).
		Where("address = ?", address)

	if tokenType != "" {
		query = query.Where("token_address = ?", tokenType)
	}

	// 查询比上次游标新的数据
	if lastTimestamp > 0 {
		query = query.Where("timestamp > ?", lastTimestamp)
	}

	// 返回最新数据，时间升序保证前端拼接顺序稳定
	err := query.
		Order("timestamp asc").
		Limit(limit).
		Find(&records).Error

	return records, err
}

func (ts tokenSentDB) StoreTokenSent(tokenSent TokenSent) error {
	err := ts.db.Table(tokenSent.TableName()).Create(&tokenSent).Error
	if err != nil {
		return err
	}
	return nil
}

func (ts tokenReceivedDB) List(address, tokenType string, lastTimestamp uint64, limit int) ([]TokenReceived, error) {
	var records []TokenReceived

	query := ts.db.Table(TokenReceived{}.TableName()).
		Where("address = ?", address)

	if tokenType != "" {
		query = query.Where("token_address = ?", tokenType)
	}

	// 查询比上次游标新的数据
	if lastTimestamp > 0 {
		query = query.Where("timestamp > ?", lastTimestamp)
	}

	// 返回最新数据，时间升序保证前端拼接顺序稳定
	err := query.
		Order("timestamp asc").
		Limit(limit).
		Find(&records).Error

	return records, err
}

func (tr tokenReceivedDB) StoreTokenReceived(tokenReceived TokenReceived) error {
	err := tr.db.Table(tokenReceived.TableName()).Create(&tokenReceived).Error
	if err != nil {
		return err
	}
	return nil
}

func NewTokenSentDB(db *gorm.DB) TokenSentDB {
	return &tokenSentDB{db: db}
}
func NewTokenReceivedDB(db *gorm.DB) TokenReceivedDB {
	return &tokenReceivedDB{db: db}
}
