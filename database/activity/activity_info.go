package activity

import "gorm.io/gorm"

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
	ListActivityInfo() ([]ActivityInfo, error)
}

type ActivityInfoDB interface {
	ActivityInfoView
}

type activityInfoDB struct {
	db *gorm.DB
}

func (a activityInfoDB) ListActivityInfo() ([]ActivityInfo, error) {
	var activityInfo []ActivityInfo
	err := a.db.Find(&activityInfo).Error
	if err != nil {
		return nil, err
	}
	return activityInfo, nil
}

func NewActivityInfoDb(db *gorm.DB) ActivityInfoDB {
	return &activityInfoDB{db: db}
}
