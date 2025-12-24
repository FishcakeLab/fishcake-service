package activity

import (
	"errors"
	"fmt"
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

type MiningInfo struct {
	Id                 string   `gorm:"column:id" json:"id"`
	Address            string   `gorm:"column:address" json:"address"`
	MinedAmount        *big.Int `gorm:"serializer:u256;column:mined_amount" json:"minedAmount"`
	MinedFishCakePower *big.Int `gorm:"serializer:u256;column:mined_fishcake_power" json:"minedFishCakePower"`
	LastMintTime       *big.Int `gorm:"serializer:u256;column:last_mint_time" json:"lastMintTime"`
}

func (ActivityInfo) TableName() string {
	return "activity_info"
}

func (MiningInfo) TableName() string {
	return "mining_info"
}

// ActivityInfoView defines the interface for viewing activity information
type ActivityInfoView interface {
	ActivityInfoList(activityFilter, businessAccount, activityStatus, businessName, tokenContractAddr, latitude, longitude, scope string,
		activityId,
		pageNum, pageSize int) ([]ActivityInfo, int)
	ActivityInfo(activityId int) ActivityInfo
	UnDropActivityParticipantAddresses() []ActivityParticipantAddress
	GetUserMinedAmount(address string, monthFilter bool) (*big.Int, error)
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

type MiningInfoDB interface {
	Create(info *MiningInfo) error
	GetByAddress(address string) (*MiningInfo, error)
	GetByAddressForUpdate(address string) (*MiningInfo, error)
	Update(info *MiningInfo) error
	Upsert(info *MiningInfo) error
	ListAll() ([]MiningInfo, error)
}

type activityInfoDB struct {
	db *gorm.DB
}

type miningInfoDB struct {
	db *gorm.DB
}

// GetUserMinedAmount 查询某个地址的累计挖矿数量（可选按月）
func (d activityInfoDB) GetUserMinedAmount(address string, monthFilter bool) (*big.Int, error) {
	db := d.db.Table("activity_info").
		Select("SUM(mined_amount::numeric) as total_mined").
		Where("business_account = ? AND activity_status = 2", address)

	if monthFilter {
		now := time.Now().Unix()
		oneMonthAgo := now - 30*24*3600
		db = db.Where("activity_deadline >= ?", oneMonthAgo)
	}

	var result struct {
		TotalMined *big.Int `gorm:"column:total_mined"`
	}

	if err := db.Scan(&result).Error; err != nil {
		return nil, err
	}

	// 没数据时为 nil → 返回 0
	if result.TotalMined == nil {
		return big.NewInt(0), nil
	}

	return result.TotalMined, nil
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
func (a activityInfoDB) ActivityFinish(activityId string, returnAmount, minedAmount *big.Int) error {

	tx := a.db.Begin()
	if tx.Error != nil {
		log.Warn("begin transaction failed", "err", tx.Error)
		return tx.Error
	}

	// 1. 更新 activity_info
	finishSql := `UPDATE activity_info 
				  SET activity_status = 2, return_amount = ?, mined_amount = ? 
				  WHERE activity_id = ?`
	if err := tx.Exec(finishSql, returnAmount, minedAmount, activityId).Error; err != nil {
		tx.Rollback()
		log.Error("update activity_info fail", "err", err)
		return err
	}

	// 2. 读取 business_account（因为 mining_info 需要 address）
	var address string
	if err := tx.Table("activity_info").
		Select("business_account").
		Where("activity_id = ?", activityId).
		Scan(&address).Error; err != nil {
		tx.Rollback()
		log.Error("query business_account fail", "err", err)
		return err
	}

	// address 非空检查
	if len(address) == 0 {
		tx.Rollback()
		return fmt.Errorf("empty business_account for activity_id %s", activityId)
	}

	// 3. 查询 mining_info 里是否已有记录
	var mi MiningInfo
	err := tx.Where(`address = ?`, address).First(&mi).Error

	// 3.1 如果不存在，插入一条新的
	if errors.Is(err, gorm.ErrRecordNotFound) {

		newInfo := MiningInfo{
			Address:            address,
			MinedAmount:        new(big.Int).Set(minedAmount), // 初始挖矿量
			MinedFishCakePower: new(big.Int).Set(minedAmount), // 初始power = 挖矿量
		}

		if err := tx.Create(&newInfo).Error; err != nil {
			tx.Rollback()
			log.Error("insert mining_info fail", "err", err)
			return err
		}

		return tx.Commit().Error
	}

	// 3.2 如果存在，则累加
	if err != nil {
		tx.Rollback()
		log.Error("query mining_info fail", "err", err)
		return err
	}

	mi.MinedAmount = new(big.Int).Add(mi.MinedAmount, minedAmount)
	mi.MinedFishCakePower = new(big.Int).Add(mi.MinedFishCakePower, minedAmount)

	if err := tx.Model(&MiningInfo{}).
		Where(`address = ?`, address).
		Updates(map[string]interface{}{
			"mined_amount":         mi.MinedAmount,
			"mined_fishcake_power": mi.MinedFishCakePower,
		}).Error; err != nil {

		tx.Rollback()
		log.Error("update mining_info fail", "err", err)
		return err
	}

	// return tx.Commit().Error
	if err := tx.Commit().Error; err != nil {
		log.Error("commit transaction fail", "err", err)
		return err
	}
	return nil
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
		switch activityStatus {
		case "2":
			this = this.Where("activity_status = ?", 2) // 如果是 2，就是finish了，不考虑过期
		case "1":
			this = this.Where("activity_status = ?", 1)
			this = this.Where("activity_deadline > ?", time.Now().Unix()) // 如果是1，就是还没finish，也没过期
		case "3":
			this = this.Where("activity_status = ?", 1)
			this = this.Where("activity_deadline < ?", time.Now().Unix()) // 如果是3，就是还没finish，但过期了
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

	//【修复点 1】必须先 JOIN，再加 account_nft_info 的筛选条件
	// 否则 WHERE 右表字段会让 LEFT JOIN 退化成 INNER JOIN
	this = this.Joins("LEFT JOIN account_nft_info ON activity_info.business_account = account_nft_info.address")

	if activityFilter == "1" {
		this = this.Where("account_nft_info.pro_deadline >= ?", time.Now().Unix())
	}
	if activityFilter == "2" {
		this = this.Where("account_nft_info.basic_deadline >= ?", time.Now().Unix())
	}

	//【修复点 2】Count 需要单独 query，不能和 Limit/Offset/Order 混在一起
	// 且分页逻辑要放在 count 之后，不影响总数计算
	countQuery := this.Session(&gorm.Session{})
	if err := countQuery.Debug().Count(&count).Error; err != nil {
		log.Error("count query failed", "err", err)
	}

	// Apply pagination
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}

	// Get total count and results
	// this = this.Count(&count)
	this = this.Order("activity_create_time DESC, activity_status ASC").Select("activity_info.*,account_nft_info.basic_deadline,account_nft_info.pro_deadline")
	result := this.Debug().Find(&activityInfo)

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

func (m *miningInfoDB) Create(info *MiningInfo) error {
	return m.db.Create(info).Error
}

func (m *miningInfoDB) GetByAddress(address string) (*MiningInfo, error) {
	var info MiningInfo
	err := m.db.Where("address = ?", address).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}
func (m *miningInfoDB) GetByAddressForUpdate(address string) (*MiningInfo, error) {
	var info MiningInfo
	err := m.db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("address = ?", address).
		Take(&info).Error
	return &info, err
}

func (m *miningInfoDB) Update(info *MiningInfo) error {
	return m.db.Model(&MiningInfo{}).
		Where("address = ?", info.Address).
		Updates(map[string]interface{}{
			"mined_amount":         info.MinedAmount,
			"mined_fishcake_power": info.MinedFishCakePower,
			"last_mint_time":       info.LastMintTime,
		}).Error
}

func (m *miningInfoDB) Upsert(info *MiningInfo) error {
	var existing MiningInfo
	err := m.db.Where("address = ?", info.Address).First(&existing).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 插入
			return m.Create(info)
		}
		return err
	}

	// 更新
	info.Id = existing.Id // 保留原 id
	return m.Update(info)
}

func (m *miningInfoDB) ListAll() ([]MiningInfo, error) {
	var list []MiningInfo
	err := m.db.Find(&list).Error
	return list, err
}

// NewActivityDB creates a new instance of ActivityInfoDB
func NewActivityDB(db *gorm.DB) ActivityInfoDB {
	return &activityInfoDB{db: db}
}

func NewMiningInfoDB(db *gorm.DB) MiningInfoDB {
	return &miningInfoDB{db: db}
}
