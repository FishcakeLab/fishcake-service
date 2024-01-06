package activity

import (
	"errors"
	"gorm.io/gorm"

	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
)

type ActivityInfo struct {
	Id                 string `gorm:"id" `
	ActivityId         int64  `gorm:"activity_id" `
	BusinessAccount    string `gorm:"business_account"`
	BusinessName       string `gorm:"business_name"`
	ActivityContent    string `gorm:"activity_content"`
	LatitudeLongitude  string `gorm:"latitude_longitude"`
	ActivityCreateTime int64  `gorm:"activity_create_time"`
	ActivityDeadline   int64  `gorm:"activity_deadline"`
	DropType           int8   `gorm:"drop_type"`
	DropNumber         int64  `gorm:"drop_number"`
	MinDropAmt         int64  `gorm:"min_drop_amt"`
	MaxDropAmt         int64  `gorm:"max_drop_amt"`
	TokenContractAddr  string `gorm:"token_contract_addr"`
}

type ActivityInfoView interface {
	ListActivityInfo(pageNum, pageSize int) ([]ActivityInfo, int64, error)
}

type ActivityInfoDB interface {
	ActivityInfoView
}

type activityInfoDB struct {
	db *gorm.DB
}

func (a activityInfoDB) ListActivityInfo(pageNum, pageSize int) ([]ActivityInfo, int64, error) {
	var activityInfo []ActivityInfo
	var count int64
	this := a.db
	this = this.Count(&count)
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	result := this.Find(&activityInfo)
	if result.Error == nil {
		return activityInfo, count, nil
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errors_h.NewErrorByEnum(enum.DataErr)
		return nil, count, nil
	} else {
		return nil, count, nil
	}
}

func NewActivityDB(db *gorm.DB) ActivityInfoDB {
	return &activityInfoDB{db: db}
}
