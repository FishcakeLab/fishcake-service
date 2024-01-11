package token_nft

import (
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
}

type TokenNftDB interface {
	TokenNftView
}

type tokenNftDB struct {
	db *gorm.DB
}

func NewTokenNftDB(db *gorm.DB) TokenNftDB {
	return &tokenNftDB{db: db}
}
