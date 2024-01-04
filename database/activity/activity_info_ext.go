package activity

import "gorm.io/gorm"

type ActivityInfoExt struct {
	Id                         string `gorm:"id"`
	ActivityId                 int64  `gorm:"activity_id"`
	AlreadyDropAmts            int64  `gorm:"already_drop_amts"`
	AlreadyDropNumber          int64  `gorm:"already_drop_number"`
	BusinessMinedAmt           int64  `gorm:"business_mined_amt"`
	BusinessMinedWithdrawedAmt int64  `gorm:"business_mined_withdrawed_amt"`
	ActivityStatus             int8   `gorm:"activity_status"`
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
