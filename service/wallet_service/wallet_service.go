package wallet_service

import "github.com/FishcakeLab/fishcake-service/database"

type WalletService interface {
	WalletExist(address, device string) bool
	StoreWallet(address, device string) error
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
