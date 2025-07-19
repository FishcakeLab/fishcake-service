package sign_service

import "github.com/FishcakeLab/fishcake-service/database"

type SignService interface {
	IsExistSign(sign string) bool
	StoreSign(sign string) error
	UpdateSignStatus(id int, status int, tx string, msg string) error
}

type signService struct {
	Db *database.DB
}

func NewSignService(db *database.DB) SignService {
	return &signService{Db: db}
}

func (w signService) IsExistSign(sign string) bool {
	return w.Db.SignStatusDB.IsExistSign(sign)
}

func (w signService) StoreSign(sign string) error {
	return w.Db.SignStatusDB.StoreSignStatus(sign)
}

func (w signService) UpdateSignStatus(id int, status int, tx string, msg string) error {
	return w.Db.SignStatusDB.UpdateSignStatus(id, status, tx, msg)
}
