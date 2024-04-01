package token_nft

import (
	"errors"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	"gorm.io/gorm"
	"math/big"
)

type TokenNft struct {
	Id              string   `json:"id" gorm:"id"`
	Who             string   `json:"who" gorm:"who"`
	TokenId         int64    `json:"tokenId" gorm:"token_id"`
	BusinessName    string   `json:"businessName" gorm:"business_name"`
	Description     string   `json:"description" gorm:"description"`
	ImgUrl          string   `json:"imgUrl" gorm:"img_url"`
	BusinessAddress string   `json:"businessAddress" gorm:"business_address"`
	WebSite         string   `json:"webSite" gorm:"web_site"`
	Social          string   `json:"social" gorm:"social"`
	ContractAddress string   `json:"contractAddress" gorm:"contract_address"`
	CostValue       *big.Int `json:"costValue" gorm:"serializer:u256;column:cost_value"`
	Deadline        uint64   `json:"deadline" gorm:"deadline"`
	NftType         int8     `json:"nftType" gorm:"nft_type"`
}

func (TokenNft) TableName() string {
	return "token_nft"
}

type TokenNftView interface {
	List(pageNum, pageSize int, contractAddress string) ([]TokenNft, int)
	NftInfo(tokenId int) TokenNft
}

type TokenNftDB interface {
	TokenNftView
	StoreTokenNft(token TokenNft) error
}

type tokenNftDB struct {
	db *gorm.DB
}

func (t tokenNftDB) StoreTokenNft(token TokenNft) error {
	tokenNft := new(TokenNft)
	var exist TokenNft
	err := t.db.Table(tokenNft.TableName()).Where("token_id = ?", token.TokenId).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := t.db.Table(tokenNft.TableName()).Omit("id").Create(&token)
			return result.Error
		}
	}
	return err
}

func (t tokenNftDB) List(pageNum, pageSize int, contractAddress string) ([]TokenNft, int) {
	var tokenNft []TokenNft
	var count int64
	this := t.db.Table(TokenNft{}.TableName())
	if contractAddress != "" {
		this = this.Where("contract_address = ?", contractAddress)
	}
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	this = this.Count(&count)
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
