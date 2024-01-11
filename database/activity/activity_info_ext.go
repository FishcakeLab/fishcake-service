package activity

import "gorm.io/gorm"

type ActivityInfoExt struct {
	Id                         string `gorm:"id" json:"id"`
	ActivityId                 int64  `gorm:"activity_id" json:"activityId"`
	AlreadyDropAmts            int64  `gorm:"already_drop_amts" json:"alreadyDropAmts"`
	AlreadyDropNumber          int64  `gorm:"already_drop_number" json:"alreadyDropNumber"`
	BusinessMinedAmt           int64  `gorm:"business_mined_amt" json:"businessMinedAmt"`
	BusinessMinedWithdrawedAmt int64  `gorm:"business_mined_withdrawed_amt" json:"businessMinedWithdrawedAmt"`
	ActivityStatus             int8   `gorm:"activity_status" json:"activityStatus"`
}

type ActivityInfoExtView interface {
}

type ActivityInfoExtDB interface {
	ActivityInfoExtView
}

type activityInfoExtDB struct {
	db *gorm.DB
}

func NewActivityInfoExtDB(db *gorm.DB) ActivityInfoExtDB {
	return &activityInfoExtDB{db: db}
}
