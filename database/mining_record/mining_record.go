package mining_record

import (
	"math/big"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	RecordTypeActivityFinish = "activity_finish"
	RecordTypeMintBoosterNFT = "mint_booster_nft"
)

type MiningRecord struct {
	ID               string   `gorm:"column:id" json:"id"`
	Address          string   `gorm:"column:address" json:"address"`
	RecordType       string   `gorm:"column:record_type" json:"recordType"`
	MinedAmountDelta *big.Int `gorm:"serializer:u256;column:mined_amount_delta" json:"minedAmountDelta"`
	PowerIncrease    *big.Int `gorm:"serializer:u256;column:power_increase" json:"powerIncrease"`
	PowerDecrease    *big.Int `gorm:"serializer:u256;column:power_decrease" json:"powerDecrease"`
	ActivityID       *int64   `gorm:"column:activity_id" json:"activityId"`
	TokenID          *int64   `gorm:"column:token_id" json:"tokenId"`
	Description      string   `gorm:"column:description" json:"description"`
	Timestamp        uint64   `gorm:"column:timestamp" json:"timestamp"`
	TxHash           string   `gorm:"column:tx_hash" json:"txHash"`
	LogIndex         uint     `gorm:"column:log_index" json:"logIndex"`
	BlockNumber      uint64   `gorm:"column:block_number" json:"blockNumber"`
}

func (MiningRecord) TableName() string {
	return "mining_record"
}

type MiningRecordDB interface {
	Store(record MiningRecord) error
	List(address, recordType string, lastTimestamp uint64, lastId string, limit int) ([]MiningRecord, error)
}

type miningRecordDB struct {
	db *gorm.DB
}

func (m miningRecordDB) Store(record MiningRecord) error {
	return m.db.
		Table(record.TableName()).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "tx_hash"},
				{Name: "log_index"},
				{Name: "record_type"},
			},
			DoNothing: true,
		}).
		Omit("id").
		Create(&record).
		Error
}

func (m miningRecordDB) List(address, recordType string, lastTimestamp uint64, lastId string, limit int) ([]MiningRecord, error) {
	var records []MiningRecord

	query := m.db.Table(MiningRecord{}.TableName()).
		Where("address ILIKE ?", address)

	if recordType != "" {
		query = query.Where("record_type = ?", recordType)
	}

	if lastTimestamp > 0 {
		if lastId != "" {
			query = query.Where("timestamp < ? OR (timestamp = ? AND id < ?)", lastTimestamp, lastTimestamp, lastId)
		} else {
			query = query.Where("timestamp < ?", lastTimestamp)
		}
	}

	err := query.
		Order("timestamp DESC, id DESC").
		Limit(limit).
		Find(&records).Error

	return records, err
}

func NewMiningRecordDB(db *gorm.DB) MiningRecordDB {
	return &miningRecordDB{db: db}
}
