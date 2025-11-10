package token_transfer_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/token_transfer"
)

type TokenTransferService interface {
	ListSent(address, tokenType string, lastTimestamp uint64, limit int) ([]token_transfer.TokenSent, error)
	ListReceived(address, tokenType string, lastTimestamp uint64, limit int) ([]token_transfer.TokenReceived, error)
}

type tokenTransferService struct {
	Db *database.DB
}

func NewTokenTransferService(db *database.DB) TokenTransferService {
	return &tokenTransferService{Db: db}
}

func (n *tokenTransferService) ListSent(address, tokenType string, lastTimestamp uint64, limit int) ([]token_transfer.TokenSent, error) {
	infos, err := n.Db.TokenSentDB.List(address, tokenType, lastTimestamp, limit)
	return infos, err
}

func (n *tokenTransferService) ListReceived(address, tokenType string, lastTimestamp uint64, limit int) ([]token_transfer.TokenReceived, error) {
	infos, err := n.Db.TokenReceivedDB.List(address, tokenType, lastTimestamp, limit)
	return infos, err
}
