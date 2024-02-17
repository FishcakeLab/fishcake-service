package token_nft

import (
	"errors"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	"gorm.io/gorm"
)

type TokenNft struct {
	Id              string `json:"id" gorm:"id"`
	TokenId         int64  `json:"tokenId" gorm:"token_id"`
	Address         string `json:"address" gorm:"address"`
	ContractAddress string `json:"contractAddress" gorm:"contract_address"`
	TokenUrl        string `json:"tokenUrl" gorm:"token_url"`
	TokenAmount     int64  `json:"tokenAmount" gorm:"token_amount"`
	Timestamp       int64  `json:"timestamp" gorm:"timestamp"`
}

func (TokenNft) TableName() string {
	return "token_nft"
}

type TokenNftView interface {
	List(pageNum, pageSize int, contractAddress string) ([]TokenNft, int)
	NftInfo(activityId int) TokenNft
}

type TokenNftDB interface {
	TokenNftView
}

type tokenNftDB struct {
	db *gorm.DB
}

func (t tokenNftDB) List(pageNum, pageSize int, contractAddress string) ([]TokenNft, int) {
	var tokenNft []TokenNft
	var count int64
	this := t.db.Table(TokenNft{}.TableName())
	if contractAddress != "" {
		this = this.Where("contract_address = ?", contractAddress)
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

func (t tokenNftDB) NftInfo(tokenId int) TokenNft {
	var tokenNft TokenNft
	this := t.db.Table(TokenNft{}.TableName())
	result := this.Where("token_id = ?", tokenId).Take(&tokenNft)
	if result.Error == nil {
		return tokenNft
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errors_h.NewErrorByEnum(enum.DataErr)
		return TokenNft{}
	} else {
		return TokenNft{}
	}
}

func NewTokenNftDB(db *gorm.DB) TokenNftDB {
	return &tokenNftDB{db: db}
}
