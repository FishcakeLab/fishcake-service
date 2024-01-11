package nft_service

import "github.com/FishcakeLab/fishcake-service/database"

type NftService interface {
}

type nftService struct {
	db *database.DB
}

func NewNftService(db *database.DB) NftService {
	return &nftService{db: db}
}
