package event

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/FishcakeLab/fishcake-service/database/utils"
)

type ContractEvent struct {
	GUID            uuid.UUID      `gorm:"primaryKey"`
	BlockHash       common.Hash    `gorm:"serializer:bytes"`
	ContractAddress common.Address `gorm:"serializer:bytes"`
	TransactionHash common.Hash    `gorm:"serializer:bytes"`
	LogIndex        uint64
	BlockNumber     *big.Int    `gorm:"serializer:u256"`
	EventSignature  common.Hash `gorm:"serializer:bytes"`
	Timestamp       uint64
	RLPLog          *types.Log `gorm:"serializer:rlp;column:rlp_bytes"`
}

func ContractEventFromLog(log *types.Log, timestamp uint64, blockNumber *big.Int) ContractEvent {
	eventSig := common.Hash{}
	if len(log.Topics) > 0 {
		eventSig = log.Topics[0]
	}
	return ContractEvent{
		GUID:            uuid.New(),
		BlockHash:       log.BlockHash,
		TransactionHash: log.TxHash,
		ContractAddress: log.Address,
		EventSignature:  eventSig,
		BlockNumber:     blockNumber,
		LogIndex:        uint64(log.Index),
		Timestamp:       timestamp,
		RLPLog:          log,
	}
}

func (c *ContractEvent) AfterFind(tx *gorm.DB) error {
	c.RLPLog.BlockHash = c.BlockHash
	c.RLPLog.TxHash = c.TransactionHash
	c.RLPLog.Index = uint(c.LogIndex)
	return nil
}

type ContractEventsView interface {
	ContractEvent(uuid.UUID) (*ContractEvent, error)
	ContractEventCount(ContractEvent) (int64, error)
	ContractEventWithFilter(ContractEvent) (*ContractEvent, error)
	ContractEventsWithFilter(ContractEvent, *big.Int, *big.Int) ([]ContractEvent, error)
	LatestContractEventWithFilter(ContractEvent) (*ContractEvent, error)
}

type ContractEventDB interface {
	ContractEventsView
	StoreContractEvents([]ContractEvent) error
}

type contractEventDB struct {
	gorm *gorm.DB
}

func (db *contractEventDB) ContractEvent(uuid uuid.UUID) (*ContractEvent, error) {
	return db.ContractEventWithFilter(ContractEvent{GUID: uuid})
}

func (db *contractEventDB) ContractEventCount(filter ContractEvent) (int64, error) {
	var count int64
	result := db.gorm.Table("contract_events").Where(&filter).Count(&count)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, result.Error
	}
	return count, nil
}

func (db *contractEventDB) ContractEventWithFilter(filter ContractEvent) (*ContractEvent, error) {
	var l2ContractEvent ContractEvent
	result := db.gorm.Table("contract_events").Where(&filter).Take(&l2ContractEvent)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l2ContractEvent, nil
}

func (db *contractEventDB) ContractEventsWithFilter(filter ContractEvent, fromHeight, toHeight *big.Int) ([]ContractEvent, error) {
	if fromHeight == nil {
		fromHeight = big.NewInt(0)
	}
	if toHeight == nil {
		return nil, errors.New("end height unspecified")
	}
	if fromHeight.Cmp(toHeight) > 0 {
		return nil, fmt.Errorf("fromHeight %d is greater than toHeight %d", fromHeight, toHeight)
	}
	query := db.gorm.Table("contract_events").Where(&filter)
	query = query.Joins("INNER JOIN block_headers ON contract_events.block_hash = block_headers.hash")
	query = query.Where("block_headers.number >= ? AND block_headers.number <= ?", fromHeight, toHeight)
	query = query.Order("block_headers.number ASC").Select("contract_events.*")
	var events []ContractEvent
	result := query.Find(&events)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return events, nil
}

func (db *contractEventDB) LatestContractEventWithFilter(filter ContractEvent) (*ContractEvent, error) {
	var l1ContractEvent ContractEvent
	result := db.gorm.Table("contract_events").Where(&filter).Order("timestamp DESC").Take(&l1ContractEvent)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l1ContractEvent, nil
}

func (db *contractEventDB) StoreContractEvents(events []ContractEvent) error {
	result := db.gorm.Table("contract_events").CreateInBatches(&events, utils.BatchInsertSize)
	return result.Error
}

func NewContractEventsDB(db *gorm.DB) ContractEventDB {
	return &contractEventDB{gorm: db}
}
