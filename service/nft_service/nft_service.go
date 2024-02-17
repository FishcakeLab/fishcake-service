package nft_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/token_nft"
)

type NftService interface {
	NftInfoList(pageNum, pageSize int, contractAddress string) ([]token_nft.TokenNft, int)
	NftInfo(activityId int) token_nft.TokenNft
}

type nftService struct {
	Db *database.DB
}

func NewNftService(db *database.DB) NftService {
	return &nftService{Db: db}
}

func (n *nftService) NftInfoList(pageNum, pageSize int, contractAddress string) ([]token_nft.TokenNft, int) {
	infos, count := n.Db.TokenNftDb.List(int(uint64(pageNum)), pageSize, contractAddress)
	return infos, count
}

func (n *nftService) NftInfo(tokenId int) token_nft.TokenNft {
	infos := n.Db.TokenNftDb.NftInfo(tokenId)
	return infos
}
