package activity

import (
	"errors"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	"gorm.io/gorm"
)

type ActivityInfo struct {
	Id                 string `gorm:"id" json:"id"`
	ActivityId         int64  `gorm:"activity_id" json:"activityId"`
	BusinessAccount    string `gorm:"business_account" json:"businessAccount"`
	BusinessName       string `gorm:"business_name" json:"businessName"`
	ActivityContent    string `gorm:"activity_content" json:"activityContent"`
	LatitudeLongitude  string `gorm:"latitude_longitude" json:"latitudeLongitude"`
	ActivityCreateTime int64  `gorm:"activity_create_time" json:"activityCreateTime"`
	ActivityDeadline   int64  `gorm:"activity_deadline" json:"activityDeadline"`
	DropType           int8   `gorm:"drop_type" json:"dropType"`
	DropNumber         int64  `gorm:"drop_number" json:"dropNumber"`
	MinDropAmt         int64  `gorm:"min_drop_amt" json:"minDropAmt"`
	MaxDropAmt         int64  `gorm:"max_drop_amt" json:"maxDropAmt"`
	TokenContractAddr  string `gorm:"token_contract_addr" json:"tokenContractAddr"`
}

func (ActivityInfo) TableName() string {
	return "activity_info"
}

type ActivityInfoView interface {
	ListActivityInfo(pageNum, pageSize int) ([]ActivityInfo, int)
}

type ActivityInfoDB interface {
	ActivityInfoView
}

type activityInfoDB struct {
	db *gorm.DB
}

func (a activityInfoDB) ListActivityInfo(pageNum, pageSize int) ([]ActivityInfo, int) {
	var activityInfo []ActivityInfo
	var count int64
	this := a.db.Table(ActivityInfo{}.TableName())
	this = this.Count(&count)
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	result := this.Find(&activityInfo)
	if result.Error == nil {
		return activityInfo, int(count)
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errors_h.NewErrorByEnum(enum.DataErr)
		return nil, int(count)
	} else {
		return nil, int(count)
	}
}

func NewActivityDB(db *gorm.DB) ActivityInfoDB {
	return &activityInfoDB{db: db}
}
