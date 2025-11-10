package activity

import (
	"errors"
	"math/big"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	_ "github.com/FishcakeLab/fishcake-service/database/utils/serializers"
)

// ActivityParticipantAddress represents the participant's address information for an activity
type ActivityParticipantAddress struct {
	ActivityId int64  `gorm:"activity_id" json:"activityId"`
	Address    string `gorm:"address" json:"address"`
	JoinTime   int64  `gorm:"join_time" json:"joinTime"`
	Status     int8   `gorm:"column:status;default:0"`
}

// ActivityInfo represents the main activity information
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
	IsExpired          bool     `gorm:"-" json:"isExpired"`
}
type WalletAddress struct {
	ID        string `gorm:"column:id;primaryKey;default:replace((uuid_generate_v4())::text, '-'::text, ''::text)"`
	Address   string `gorm:"column:address;unique;not null"`
	Device    string `gorm:"column:device;unique;not null"`
	CreatedAt int64  `gorm:"column:created_at;not null"`
	Status    int8   `gorm:"column:status;default:1"`
	Remark    string `gorm:"column:remark"`
}

type BusinessRank struct {
	BusinessAccount string `json:"businessAccount"`
	TotalMined      string `json:"totalMined"`
}

func (ActivityInfo) TableName() string {
	return "activity_info"
}

// ActivityInfoView defines the interface for viewing activity information
type ActivityInfoView interface {
	ActivityInfoList(activityFilter, businessAccount, activityStatus, businessName, tokenContractAddr, latitude, longitude, scope string,
		activityId,
		pageNum, pageSize int) ([]ActivityInfo, int)
	ActivityInfo(activityId int) ActivityInfo
	UnDropActivityParticipantAddresses() []ActivityParticipantAddress
}

// ActivityInfoDB defines the interface for activity database operations
type ActivityInfoDB interface {
	ActivityInfoView
	StoreActivityInfo(activityInfo ActivityInfo) error
	ActivityFinish(activityId string, ReturnAmount, MinedAmount *big.Int) error
	UpdateActivityInfo(activityId string) error
	MarkActivityParticipantAddressDropped(string) error
	GetActivityRank(monthFilter bool) ([]BusinessRank, error)
}

type activityInfoDB struct {
	db *gorm.DB
}

func (d activityInfoDB) GetActivityRank(monthFilter bool) ([]BusinessRank, error) {
	db := d.db.Table("activity_info").
		Select("business_account, SUM(mined_amount::numeric) as total_mined").
		Where("activity_status = ?", 2).
		Group("business_account").
		Order("total_mined DESC")

	if monthFilter {
		// 当前时间戳减去 30 天
		now := time.Now().Unix()
		oneMonthAgo := now - 30*24*3600
		db = db.Where("activity_deadline >= ?", oneMonthAgo)
	}

	var ranks []BusinessRank
	if err := db.Scan(&ranks).Error; err != nil {
		return nil, err
	}
	return ranks, nil
}

func (a activityInfoDB) UnDropActivityParticipantAddresses() []ActivityParticipantAddress {
	var dropAddressList []ActivityParticipantAddress
	err := a.db.Table("activity_participants_addresses").Where("status = ?", 0).Find(&dropAddressList).Error
	if err != nil {
		log.Error("Query undrop address list fail", "err", err)
		return nil
	}
	return dropAddressList
}

func (a activityInfoDB) MarkActivityParticipantAddressDropped(address string) error {
	err := a.db.Table("activity_participants_addresses").Where("address = ?", address).Update("status", 1).Error
	if err != nil {
		log.Error("mark drop address status fail", "err", err)
		return err
	}
	return nil
}

// UpdateActivityInfo increments the already_drop_number for an activity
func (a activityInfoDB) UpdateActivityInfo(activityId string) error {
	sql := `update activity_info set already_drop_number = already_drop_number + 1 where activity_id = ? AND already_drop_number < drop_number;`
	err := a.db.Exec(sql, activityId).Error
	return err
}

// optional: use Model and Updates replace SQL
// func (a activityInfoDB) UpdateActivityInfo(activityId string) error {
// 	result := a.db.Table("activity_info").
// 		Where("activity_id = ?", activityId).
// 		UpdateColumn("already_drop_number", gorm.Expr("already_drop_number + 1"))
// 	return result.Error
// }

// ActivityFinish updates activity status and amounts when activity is finished
func (a activityInfoDB) ActivityFinish(activityId string, ReturnAmount, MinedAmount *big.Int) error {
	finishSql := `update activity_info set activity_status = 2, return_amount = ?, mined_amount = ? where activity_id = ?`
	err := a.db.Exec(finishSql, ReturnAmount, MinedAmount, activityId).Error
	return err
}

// optional: use Model and Updates replace SQL
// func (a activityInfoDB) ActivityFinish(activityId string, ReturnAmount, MinedAmount *big.Int) error {
// 	// use Model and Updates replace SQL
// 	return a.db.Model(&ActivityInfo{}).
// 		Where("activity_id = ?", activityId).
// 		Updates(map[string]interface{}{
// 			"activity_status": 2,
// 			"return_amount":   ReturnAmount,
// 			"mined_amount":    MinedAmount,
// 		}).Error
// }

func (a activityInfoDB) StoreActivityInfo(activityInfo ActivityInfo) error {
	activityInfoRecord := new(ActivityInfo)

	err := a.db.Table(activityInfoRecord.TableName()).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "activity_id"}},
			DoNothing: true,
		}).Omit("id, basic_deadline, pro_deadline").Create(&activityInfo).Error
	if err != nil {
		log.Error("create activityInfo error", "err", err)
		return err
	}

	// var exist ActivityInfo

	// err := a.db.Table(activityInfoRecord.TableName()).Where("activity_id = ?", activityInfo.ActivityId).Take(&exist).Error
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		errCreate := a.db.Table(activityInfoRecord.TableName()).Omit("id, basic_deadline, pro_deadline").Create(&activityInfo).Error
	// 		if errCreate != nil {
	// 			log.Error("create activityInfo error", "err", errCreate)
	// 			return errCreate
	// 		}
	// 	}
	// }
	var walletAddress WalletAddress
	this := a.db.Table("wallet_addresses")
	result := this.Where("address = ?", activityInfo.BusinessAccount).Take(&walletAddress)
	if result.Error == nil {

		var existAddress ActivityParticipantAddress
		errApAddress := a.db.Table("activity_participants_addresses").Where("address = ?", activityInfo.BusinessAccount).Take(&existAddress).Error
		if errApAddress != nil && errors.Is(errApAddress, gorm.ErrRecordNotFound) {
			existAddress = ActivityParticipantAddress{
				ActivityId: activityInfo.ActivityId,
				Address:    activityInfo.BusinessAccount,
				Status:     0,
				JoinTime:   time.Now().Unix(),
			}
			if errCreate := a.db.Table("activity_participants_addresses").Create(&existAddress).Error; errCreate != nil {
				log.Error("Create activity_participants_addresses error", "err", errCreate)
				return errCreate
			}
		}

	}
	return nil
}

// ActivityInfo retrieves activity information by ID
func (a activityInfoDB) ActivityInfo(activityId int) ActivityInfo {
	var activityInfo ActivityInfo
	this := a.db.Table(ActivityInfo{}.TableName())
	this = this.Joins("LEFT JOIN account_nft_info ON activity_info.business_account = account_nft_info.address")
	this = this.Select("activity_info.*,account_nft_info.basic_deadline,account_nft_info.pro_deadline")

	result := this.Where("activity_id = ?", activityId).Take(&activityInfo)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			errors_h.NewErrorByEnum(enum.DataErr)
		}
		return ActivityInfo{}
	}

	// ====== 计算 isExpired 字段 ======
	activityInfo.IsExpired = activityInfo.ActivityDeadline <= time.Now().Unix()

	return activityInfo
}

/**
func (a activityInfoDB) ActivityInfo(activityId int) ActivityInfo {
	var activityInfo ActivityInfo

	query := a.db.Table(ActivityInfo{}.TableName()).
		Joins("LEFT JOIN account_nft_info ON activity_info.business_account = account_nft_info.address").
		Select("activity_info.*,account_nft_info.basic_deadline,account_nft_info.pro_deadline").
		Where("activity_id = ?", activityId)

	result := query.Take(&activityInfo)
	if result.Error == nil {
		return activityInfo
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errors_h.NewErrorByEnum(enum.DataErr)
		return ActivityInfo{}
	}
	return ActivityInfo{}
}
*/

// ActivityInfoList retrieves a list of activities based on various filters
func (a activityInfoDB) ActivityInfoList(activityFilter, businessAccount, activityStatus, businessName, tokenContractAddr, latitude, longitude, scope string, activityId, pageNum, pageSize int) ([]ActivityInfo, int) {
	var activityInfo []ActivityInfo
	var count int64
	this := a.db.Table(ActivityInfo{}.TableName())

	// Apply filters
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

	// Join with account_nft_info table
	this = this.Joins("LEFT JOIN account_nft_info ON activity_info.business_account = account_nft_info.address")

	// Apply pagination
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}

	// Get total count and results
	this = this.Count(&count)
	this = this.Order("activity_create_time DESC, activity_status ASC").Select("activity_info.*,account_nft_info.basic_deadline,account_nft_info.pro_deadline")
	result := this.Find(&activityInfo)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			errors_h.NewErrorByEnum(enum.DataErr)
		}
		return nil, int(count)
	}

	// ====== 计算 isExpired 字段 ======
	now := time.Now().Unix()
	for i := range activityInfo {
		activityInfo[i].IsExpired = activityInfo[i].ActivityDeadline <= now
	}

	return activityInfo, int(count)
}

// NewActivityDB creates a new instance of ActivityInfoDB
func NewActivityDB(db *gorm.DB) ActivityInfoDB {
	return &activityInfoDB{db: db}
}
