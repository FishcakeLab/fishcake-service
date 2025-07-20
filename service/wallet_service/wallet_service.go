package wallet_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/wallet"
)

type WalletService interface {
	WalletExist(address, device string) bool
	StoreWallet(address, device string) error
	IsExistRawTx(rawTx string) bool
	StoreRawTx(rawTx string) error
	QueryTxInfoByHash(txHash string) (*wallet.QueueTx, error)
}

type walletService struct {
	Db *database.DB
}

func NewWalletService(db *database.DB) WalletService {
	return &walletService{Db: db}
}

func (w walletService) WalletExist(address, device string) bool {
	return w.Db.WalletInfoDB.ExistWalletAddr(address, device)
}

func (w walletService) StoreWallet(address, device string) error {
	return w.Db.WalletInfoDB.StoreWalletAddr(address, device)
}

func (w walletService) IsExistRawTx(rawTx string) bool {
	return w.Db.QueueTxDB.ExistQueueTx(rawTx)
}

func (w walletService) StoreRawTx(rawTx string) error {
	return w.Db.QueueTxDB.StoreRawTx(rawTx)
}

func (w walletService) QueryTxInfoByHash(txHash string) (*wallet.QueueTx, error) {
	return w.Db.QueueTxDB.QueryTxInfoByHash(txHash)
}
