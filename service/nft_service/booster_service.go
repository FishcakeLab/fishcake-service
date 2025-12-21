package nft_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
)

type BoosterInfo struct {
	TokenId int64 `json:"tokenId"`
	NftType int8  `json:"nftType"`
}

type BoosterService interface {
	ListByAddress(
		address string,
		pageNum,
		pageSize int,
	) ([]BoosterInfo, int)
}

type boosterService struct {
	Db *database.DB
}

func NewBoosterService(db *database.DB) BoosterService {
	return &boosterService{Db: db}

}

func (b *boosterService) ListByAddress(
	address string,
	pageNum,
	pageSize int,
) ([]BoosterInfo, int) {

	if pageSize <= 0 || pageSize > 100 {
		pageSize = 50
	}
	if pageNum <= 0 {
		pageNum = 1
	}

	nfts, count := b.Db.TokenNftDB.ListBoosterByAddress(
		address,
		pageNum,
		pageSize,
	)

	infos := make([]BoosterInfo, 0, len(nfts))
	for _, nft := range nfts {
		infos = append(infos, BoosterInfo{
			TokenId: nft.TokenId,
			NftType: nft.NftType,
		})
	}

	return infos, count
}
