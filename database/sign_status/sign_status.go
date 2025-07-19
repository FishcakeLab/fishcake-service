package sign_status

import (
	"errors"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	"time"
)

type SignStatus struct {
	Id              int
	Signature       string
	Status          int8
	TransactionHash string
	ResultMsg       string
	CreateTime      time.Time
}

func (SignStatus) TableName() string {
	return "sign_status"
}

type SignStatusDBView interface {
	IsExistSign(sign string) bool
	GetSignsByStatus(status int) []SignStatus
}
type SignStatusDB interface {
	SignStatusDBView
	StoreSignStatus(sign string) error
	UpdateSignStatus(id int, status int, tx string, msg string) error
}

type signStatusDB struct {
	gorm *gorm.DB
}

func (db signStatusDB) GetSignsByStatus(status int) []SignStatus {
	var notConfirmSign []SignStatus
	var confirm SignStatus
	err := db.gorm.Table(confirm.TableName()).Where("status = ?", status).Find(&notConfirmSign).Error
	if err != nil {
		log.Error("Query notConfirm sign list fail", "err", err)
		return nil
	}
	return notConfirmSign
}

func (db signStatusDB) UpdateSignStatus(id int, status int, tx string, msg string) error {
	sql := `update sign_status set status = ?,transaction_hash = ?,result_msg = ? where id = ?`
	err := db.gorm.Exec(sql, status, tx, msg, id).Error
	return err
}

func (db signStatusDB) StoreSignStatus(sign string) error {
	existAddress := SignStatus{
		Signature:  sign,
		Status:     0,
		CreateTime: time.Now(),
	}
	if err := db.gorm.Table("sign_status").Create(&existAddress).Error; err != nil {
		return err
	}
	return nil
}

func (db signStatusDB) IsExistSign(sign string) bool {
	var signStatus SignStatus
	result := db.gorm.Table("sign_status").Where("signature = ?", sign).First(&signStatus)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true //存在返回true
}

func NewSignStatusDB(db *gorm.DB) SignStatusDB {
	return &signStatusDB{gorm: db}
}
