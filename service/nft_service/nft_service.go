package nft_service

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/event"
	"github.com/FishcakeLab/fishcake-service/database/token_nft"
)

type NftService interface {
	NftInfoList(pageNum, pageSize int, contractAddress, address string) ([]token_nft.TokenNft, int)
	NftInfo(tokenId int) token_nft.TokenNft
	NftDetail(businessAccount, deadline string) token_nft.TokenNft
	NftCount(contractAddress string) int64
	TransactionCount(address string) int64
}

type nftService struct {
	Db *database.DB
}

func (n *nftService) TransactionCount(address string) int64 {
	contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(address)}
	count, err := n.Db.ContractEvent.ContractEventCount(contractEventFilter)
	if err != nil {
		return 13756 // old system have 1276 events interactions, migrate to aws delete them
	}
	return count + 11401
}

func NewNftService(db *database.DB) NftService {
	return &nftService{Db: db}
}

func (n *nftService) NftCount(contractAddress string) int64 {
	count := n.Db.TokenNftDB.NftCount(contractAddress)
	return count
}

func (n *nftService) NftInfoList(pageNum, pageSize int, contractAddress, address string) ([]token_nft.TokenNft, int) {
	infos, count := n.Db.TokenNftDB.List(int(uint64(pageNum)), pageSize, contractAddress, address)
	return infos, count
}

func (n *nftService) NftInfo(tokenId int) token_nft.TokenNft {
	infos := n.Db.TokenNftDB.NftInfo(tokenId)
	return infos
}

func (n *nftService) NftDetail(businessAccount, deadline string) token_nft.TokenNft {
	infos := n.Db.TokenNftDB.NftDetail(businessAccount, deadline)
	return infos
}
