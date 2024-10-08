package activity

import (
	"errors"
	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	_ "github.com/FishcakeLab/fishcake-service/database/utils/serializers"
	"gorm.io/gorm"
	"math/big"
	"time"
)

type ActivityInfo struct {
	Id                 string   `gorm:"id" json:"id"`
	ActivityId         int64    `gorm:"activity_id" json:"activityId"`
	BusinessAccount    string   `gorm:"business_account" json:"businessAccount"`
	BusinessName       string   `gorm:"business_name" json:"businessName"`
	ActivityContent    string   `gorm:"activity_content" json:"activityContent"`
	LatitudeLongitude  string   `gorm:"latitude_longitude" json:"latitudeLongitude"`
	ActivityCreateTime int64    `gorm:"activity_create_time" json:"activityCreateTime"`
	ActivityDeadline   int64    `gorm:"activity_deadline" json:"activityDeadline"`
	DropType           int8     `gorm:"drop_type" json:"dropType"`
	DropNumber         int64    `gorm:"drop_number" json:"dropNumber"`
	MinDropAmt         *big.Int `gorm:"serializer:u256;column:min_drop_amt" json:"minDropAmt"`
	MaxDropAmt         *big.Int `gorm:"serializer:u256;column:max_drop_amt" json:"maxDropAmt"`
	TokenContractAddr  string   `gorm:"token_contract_addr" json:"tokenContractAddr"`
	ActivityStatus     int8     `gorm:"activity_status" json:"activityStatus"`
	AlreadyDropNumber  int64    `gorm:"already_drop_number" json:"alreadyDropNumber"`
	BasicDeadline      uint64   `gorm:"basic_deadline" json:"basicDeadline" `
	ProDeadline        uint64   `gorm:"pro_deadline" json:"proDeadline"`
	ReturnAmount       *big.Int `gorm:"serializer:u256;column:return_amount" json:"returnAmount"`
	MinedAmount        *big.Int `gorm:"serializer:u256;column:mined_amount" json:"minedAmount"`
}

func (ActivityInfo) TableName() string {
	return "activity_info"
}

type ActivityInfoView interface {
	ActivityInfoList(activityFilter, businessAccount, activityStatus, businessName, tokenContractAddr, latitude, longitude, scope string,
		activityId,
		pageNum, pageSize int) ([]ActivityInfo, int)
	ActivityInfo(activityId int) ActivityInfo
}

type ActivityInfoDB interface {
	ActivityInfoView
	StoreActivityInfo(activityInfo ActivityInfo) error
	ActivityFinish(activityId string, ReturnAmount, MinedAmount *big.Int) error
	UpdateActivityInfo(activityId string) error
}

type activityInfoDB struct {
	db *gorm.DB
}

func (a activityInfoDB) UpdateActivityInfo(activityId string) error {
	sql := `update activity_info set already_drop_number = already_drop_number + 1 where activity_id = ?`
	err := a.db.Exec(sql, activityId).Error
	return err
}

func (a activityInfoDB) ActivityFinish(activityId string, ReturnAmount, MinedAmount *big.Int) error {
	finishSql := `update activity_info set activity_status = 2, return_amount = ?, mined_amount = ? where activity_id = ?`
	err := a.db.Exec(finishSql, ReturnAmount, MinedAmount, activityId).Error
	return err
}

func (a activityInfoDB) StoreActivityInfo(activityInfo ActivityInfo) error {
	activityInfoRecord := new(ActivityInfo)
	var exist ActivityInfo
	err := a.db.Table(activityInfoRecord.TableName()).Where("activity_id = ?", activityInfo.ActivityId).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := a.db.Table(activityInfoRecord.TableName()).Omit("id, basic_deadline, pro_deadline").Create(&activityInfo)
			return result.Error
		}
	}
	return err
}

func (a activityInfoDB) ActivityInfo(activityId int) ActivityInfo {
	var activityInfo ActivityInfo
	this := a.db.Table(ActivityInfo{}.TableName())
	this = this.Joins("LEFT JOIN account_nft_info ON activity_info.business_account = account_nft_info.address")
	this = this.Select("activity_info.*,account_nft_info.basic_deadline,account_nft_info.pro_deadline")
	result := this.Where("activity_id = ?", activityId).Take(&activityInfo)
	if result.Error == nil {
		return activityInfo
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errors_h.NewErrorByEnum(enum.DataErr)
		return ActivityInfo{}
	} else {
		return ActivityInfo{}
	}
}

func (a activityInfoDB) ActivityInfoList(activityFilter, businessAccount, activityStatus, businessName, tokenContractAddr, latitude, longitude, scope string, activityId, pageNum, pageSize int) ([]ActivityInfo, int) {
	var activityInfo []ActivityInfo
	var count int64
	this := a.db.Table(ActivityInfo{}.TableName())
	if businessAccount != "" {
		this = this.Where("business_account ILIKE ?", "%"+businessAccount+"%")
	}
	if activityStatus != "" {
		this = this.Where("activity_status = ?", activityStatus)
		if activityStatus == "2" {
			this = this.Where("activity_deadline < ?", time.Now().Unix())
		} else {
			this = this.Where("activity_deadline > ?", time.Now().Unix())
		}
	}
	if businessName != "" {
		this = this.Where("business_name ILIKE ?", "%"+businessName+"%")
	}
	if tokenContractAddr != "" {
		this = this.Where("token_contract_addr ILIKE ?", "%"+tokenContractAddr+"%")
	}
	if activityId > 0 {
		this = this.Where("activity_id = ?", activityId)
	}
	if latitude != "" && longitude != "" {
		this = this.Where("ST_DWithin(ST_SetSRID(ST_MakePoint("+
			"CAST(SPLIT_PART(latitude_longitude, ',', 1) AS numeric), "+
			"CAST(SPLIT_PART(latitude_longitude, ',', 2) AS numeric)), 4326),ST_SetSRID(ST_MakePoint(?, ?), 4326),?)", latitude, longitude, scope)
	}
	if activityFilter == "1" {
		this = this.Where("account_nft_info.pro_deadline >= ?", time.Now().Unix())
	}
	if activityFilter == "2" {
		this = this.Where("account_nft_info.basic_deadline >= ?", time.Now().Unix())
	}
	this = this.Joins("LEFT JOIN account_nft_info ON activity_info.business_account = account_nft_info.address")
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	this = this.Count(&count)
	this = this.Order("activity_create_time DESC, activity_status ASC").Select("activity_info.*,account_nft_info.basic_deadline,account_nft_info.pro_deadline")
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
