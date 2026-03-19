package drop

import (
	"errors"
	"math/big"
	"time"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	"github.com/FishcakeLab/fishcake-service/database/activity"
	_ "github.com/FishcakeLab/fishcake-service/database/utils/serializers"
)

type DropInfo struct {
	Id                string   `json:"id" gorm:"id"`
	ActivityId        int64    `gorm:"activity_id" json:"activityId"`
	Address           string   `json:"address" gorm:"address"`
	DropAmount        *big.Int `json:"dropAmount" gorm:"serializer:u256;column:drop_amount"`
	DropType          int8     `gorm:"drop_type" json:"dropType"`
	SystemDropType    string   `gorm:"system_drop_type" json:"-"`
	Timestamp         uint64   `json:"timestamp" gorm:"timestamp"`
	TokenContractAddr string   `gorm:"token_contract_addr" json:"tokenContractAddr"`
	BusinessName      string   `gorm:"business_name" json:"businessName"`
	TransactionHash   string   `gorm:"transaction_hash" json:"transactionHash"`
	EventSignature    string   `gorm:"event_signature" json:"eventSignature"`
	ReturnAmount      *big.Int `gorm:"serializer:u256;column:return_amount" json:"returnAmount"`
	MinedAmount       *big.Int `gorm:"serializer:u256;column:mined_amount" json:"minedAmount"`

	BlockNumber uint64 `gorm:"column:block_number" json:"blockNumber"`
	LogIndex    uint   `gorm:"column:log_index" json:"logIndex"`
}
type WalletAddress struct {
	ID        string `gorm:"column:id;primaryKey;default:replace((uuid_generate_v4())::text, '-'::text, ''::text)"`
	Address   string `gorm:"column:address;unique;not null"`
	Device    string `gorm:"column:device;unique;not null"`
	CreatedAt int64  `gorm:"column:created_at;not null"`
	Status    int8   `gorm:"column:status;default:1"`
	Remark    string `gorm:"column:remark"`
}

func (DropInfo) TableName() string {
	return "drop_info"
}

type DropInfoView interface {
	List(pageNum, pageSize int, address, dropType string) ([]DropInfo, int)
	IsExist(transactionHash string, logIndex uint, dropType int8, address string, blockNumber uint64) (error, bool)
}

type DropInfoDB interface {
	DropInfoView
	StoreDropInfo(drop DropInfo) error
}

type dropInfoDB struct {
	db *gorm.DB
}

func (d dropInfoDB) List(pageNum, pageSize int, address, dropType string) ([]DropInfo, int) {
	var tokenNft []DropInfo
	var count int64
	var this = d.db.Table(DropInfo{}.TableName())
	// dropType=1 is Token Received need inject system drop info
	if dropType == "1" {
		this = d.db.Table("(SELECT id,activity_id,address,drop_amount,drop_type,timestamp,transaction_hash,event_signature,null as system_drop_type,block_number,log_index FROM drop_info" +
			" UNION ALL SELECT id,null as activity_id,address,drop_amount,1 as drop_type,timestamp,transaction_hash,null as event_signature,drop_type as system_drop_type,0 as block_number,0 as log_index FROM system_drop_info where status = 1) AS drop_info")
	}
	if address != "" {
		this = this.Where("address ILIKE ?", address)
	}
	if dropType != "" {
		this = this.Where("drop_info.drop_type = ?", dropType)
	}
	this = this.Joins("LEFT JOIN activity_info ON drop_info.activity_id = activity_info.activity_id")

	// 用独立 session 计算总数，避免 Count 污染后续 Select
	countQuery := this.Session(&gorm.Session{})
	if err := countQuery.Count(&count).Error; err != nil {
		log.Error("count query failed", "err", err)
	}

	this = this.Order("drop_info.timestamp DESC")
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	result := this.Select("drop_info.*, activity_info.token_contract_addr, activity_info.business_name, activity_info.return_amount, activity_info.mined_amount").Scan(&tokenNft)
	if result.Error == nil {
		return tokenNft, int(count)
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errors_h.NewErrorByEnum(enum.DataErr)
		return nil, int(count)
	} else {
		return nil, int(count)
	}
}

func (d dropInfoDB) IsExist(transactionHash string, logIndex uint, dropType int8, address string, blockNumber uint64) (error, bool) {
	var drop DropInfo
	// 1. 尝试物理匹配 (未来标准格式)
	err := d.db.Table(DropInfo{}.TableName()).
		Where("transaction_hash = ? AND log_index = ? AND drop_type = ?", transactionHash, logIndex, dropType).
		Take(&drop).Error
	if err == nil {
		return nil, true
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err, false
	}

	// 2. 物理匹配失败，尝试业务匹配 (针对 index=0 的历史数据)
	var exist DropInfo
	errLocate := d.db.Table(DropInfo{}.TableName()).
		Where("transaction_hash = ? AND address ILIKE ? AND drop_type = ?", transactionHash, address, dropType).
		Take(&exist).Error

	if errLocate == nil {
		// 命中了历史数据，如果它没有 index，顺手帮它补上，实现数据“原地升级”
		if exist.LogIndex == 0 {
			d.db.Table(DropInfo{}.TableName()).Where("id = ?", exist.Id).
				Updates(map[string]interface{}{
					"log_index":    logIndex,
					"block_number": blockNumber,
				})
		}
		// 告诉调用者：这活儿已经干过了，别再加人数了
		return nil, true
	}

	return nil, false
}

func (d dropInfoDB) StoreDropInfo(drop DropInfo) error {

	var exist DropInfo
	err := d.db.Table(DropInfo{}.TableName()).Where("transaction_hash = ? AND address ILIKE ? AND log_index = ? AND drop_type = ?", drop.TransactionHash, drop.Address, drop.LogIndex, drop.DropType).Take(&exist).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		errCreate := d.db.Table(DropInfo{}.TableName()).
			Omit("id", "token_contract_addr", "business_name", "return_amount", "mined_amount").
			Create(&drop).Error
		if errCreate != nil {
			log.Error("create drop fail", "err", errCreate)
			return errCreate
		}
	}

	var walletAddress WalletAddress
	this := d.db.Table("wallet_addresses")

	if drop.Address == "" {
		return nil
	}
	result := this.Where("address = ?", drop.Address).Take(&walletAddress)
	if result.Error == nil {

		var existAddress activity.ActivityParticipantAddress
		errDrop := d.db.Table("activity_participants_addresses").Where("address = ?", drop.Address).Take(&existAddress).Error
		if errDrop != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				existAddress = activity.ActivityParticipantAddress{
					ActivityId: drop.ActivityId,
					Address:    drop.Address,
					Status:     0,
					JoinTime:   time.Now().Unix(),
				}
				if errCreate := d.db.Table("activity_participants_addresses").Create(&existAddress).Error; errCreate != nil {
					log.Error("Create activity_participants_addresses fail", "err", errCreate)
					return errCreate
				}
			}
		}
	}
	return nil
}

func NewDropInfoDB(db *gorm.DB) DropInfoDB {
	return &dropInfoDB{db: db}
}
