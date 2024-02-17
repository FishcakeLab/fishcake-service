package activity

import (
	"errors"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	"gorm.io/gorm"
)

type ActivityInfoExt struct {
	Id                         string `gorm:"id" json:"id"`
	ActivityId                 int64  `gorm:"activity_id" json:"activityId"`
	AlreadyDropAmts            int64  `gorm:"already_drop_amts" json:"alreadyDropAmts"`
	AlreadyDropNumber          int64  `gorm:"already_drop_number" json:"alreadyDropNumber"`
	BusinessMinedAmt           int64  `gorm:"business_mined_amt" json:"businessMinedAmt"`
	BusinessMinedWithdrawedAmt int64  `gorm:"business_mined_withdrawed_amt" json:"businessMinedWithdrawedAmt"`
	ActivityStatus             int8   `gorm:"activity_status" json:"activityStatus"`
}

func (ActivityInfoExt) TableName() string {
	return "activity_info_ext"
}

type ActivityInfoExtView interface {
	ActivityInfoExtList(pageNum, pageSize int) ([]ActivityInfoExt, int)
	ActivityInfoExt(activityId int) ActivityInfoExt
}

type ActivityInfoExtDB interface {
	ActivityInfoExtView
}

type activityInfoExtDB struct {
	db *gorm.DB
}

func (a activityInfoExtDB) ActivityInfoExtList(pageNum, pageSize int) ([]ActivityInfoExt, int) {
	var activityInfoExt []ActivityInfoExt
	var count int64
	this := a.db.Table(ActivityInfo{}.TableName())
	this = this.Count(&count)
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	result := this.Find(&activityInfoExt)
	if result.Error == nil {
		return activityInfoExt, int(count)
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errors_h.NewErrorByEnum(enum.DataErr)
		return nil, int(count)
	} else {
		return nil, int(count)
	}
}

func (a activityInfoExtDB) ActivityInfoExt(activityId int) ActivityInfoExt {
	var activityInfoExt ActivityInfoExt
	this := a.db.Table(ActivityInfoExt{}.TableName())
	result := this.Where("activity_id = ?", activityId).Take(&activityInfoExt)
	if result.Error == nil {
		return activityInfoExt
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errors_h.NewErrorByEnum(enum.DataErr)
		return ActivityInfoExt{}
	} else {
		return ActivityInfoExt{}
	}
}

func NewActivityInfoExtDB(db *gorm.DB) ActivityInfoExtDB {
	return &activityInfoExtDB{db: db}
}
