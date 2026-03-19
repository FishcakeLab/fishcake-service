package token_nft

import (
	"errors"
	"math/big"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
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

	BlockNumber uint64 `gorm:"column:block_number" json:"blockNumber"`
	LogIndex    uint   `gorm:"column:log_index" json:"logIndex"`
	TxHash      string `gorm:"column:tx_hash" json:"txHash"`
}

func (TokenNft) TableName() string {
	return "token_nft"
}

type TokenNftView interface {
	List(pageNum, pageSize int, contractAddress, address string) ([]TokenNft, int)
	NftInfo(tokenId int) TokenNft
	NftDetail(businessAccount, deadline string) TokenNft
}

type TokenNftDB interface {
	TokenNftView
	StoreTokenNft(token TokenNft) error
	NftCount(contractAddress string) int64
	UpsertBoosterNft(token TokenNft) error
	ListBoosterByAddress(
		address string,
		pageNum,
		pageSize int,
	) ([]TokenNft, int)
	MarkBoosterUsed(tokenId int64) error
}

type tokenNftDB struct {
	db *gorm.DB
}

func (t tokenNftDB) MarkBoosterUsed(tokenId int64) error {
	// 只允许未 used 的 booster 被更新（幂等）
	result := t.db.Model(&TokenNft{}).
		Where("token_id = ? AND nft_type < 10", tokenId).
		Update("nft_type", gorm.Expr("nft_type + 10"))

	if result.Error != nil {
		return result.Error
	}

	// 已经 used 的情况：不报错，直接吞掉
	return nil
}

func (t tokenNftDB) ListBoosterByAddress(
	address string,
	pageNum,
	pageSize int,
) ([]TokenNft, int) {

	var list []TokenNft
	var count int64

	db := t.db.
		Select("token_id, nft_type").
		Where("who ILIKE ? AND business_name = 'BoosterNFT'", address)

	db.Count(&count)

	db.
		Order("token_id desc").
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&list)

	return list, int(count)
}

func (t tokenNftDB) UpsertBoosterNft(token TokenNft) error {
	var exist TokenNft

	err := t.db.
		Where("token_id = ?", token.TokenId).
		Take(&exist).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 首次 mint：加入物理幂等保护
		return t.db.Table(TokenNft{}.TableName()).
			Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "tx_hash"}, {Name: "log_index"}},
				DoNothing: true,
			}).
			Omit("id").
			Create(&token).Error
	}

	if err != nil {
		return err
	}

	// 已存在：允许 nft_type 演进（3 → 13）
	update := map[string]interface{}{
		"nft_type": token.NftType,
		"who":      token.Who,
	}

	return t.db.Model(&TokenNft{}).
		Where("id = ?", exist.Id).
		Updates(update).Error
}

func (t tokenNftDB) NftCount(contractAddress string) int64 {
	var count int64
	this := t.db.Table(TokenNft{}.TableName())
	result := this.Where("LOWER(contract_address) = LOWER(?)", contractAddress).Count(&count)
	if result.Error == nil {
		return count
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errors_h.NewErrorByEnum(enum.DataErr)
		return count
	} else {
		return count
	}
}

func (t tokenNftDB) StoreTokenNft(token TokenNft) error {
	var exist TokenNft
	err := t.db.Table(TokenNft{}.TableName()).Where("token_id = ? AND contract_address ILIKE ?", token.TokenId, token.ContractAddress).Take(&exist).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 彻底不存在，直接插入，并带上物理唯一保护
	return t.db.Table(TokenNft{}.TableName()).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "tx_hash"}, {Name: "log_index"}},
			DoNothing: true,
		}).
		Omit("id").
		Create(&token).Error
}

	if err != nil {
		return err
	}

	// 存在旧数据 (可能缺失指纹)：执行“指纹补全计划”
	if exist.TxHash == "" || exist.LogIndex == 0 {
		return t.db.Table(TokenNft{}.TableName()).
			Where("id = ?", exist.Id).
			Updates(map[string]interface{}{
				"tx_hash":      token.TxHash,
				"log_index":    token.LogIndex,
				"block_number": token.BlockNumber,
			}).Error
	}

	return nil
}

func (t tokenNftDB) List(pageNum, pageSize int, contractAddress, address string) ([]TokenNft, int) {
	var tokenNft []TokenNft
	var count int64
	this := t.db.Table(TokenNft{}.TableName())
	if contractAddress != "" {
		this = this.Where("contract_address ILIKE ?", contractAddress)
	}
	if address != "" {
		this = this.Where("who ILIKE ?", address)
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

func (t tokenNftDB) NftDetail(businessAccount, deadline string) TokenNft {
	var tokenNft TokenNft
	this := t.db.Table(TokenNft{}.TableName())
	result := this.Where("who ILIKE ? and deadline = ?", businessAccount, deadline).Take(&tokenNft)
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
