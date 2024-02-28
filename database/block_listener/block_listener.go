package block_listener

import (
	"errors"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	"math/big"
	"time"
)

type BlockListener struct {
	GUID        string   `json:"guid"`
	BlockNumber *big.Int `json:"block_number" gorm:"serializer:u256"`
	Created     uint64   `json:"created"`
	Updated     uint64   `json:"updated"`
}

func (BlockListener) TableName() string {
	return "block_listener"
}

type blockListenerDB struct {
	gorm *gorm.DB
}

type BlockListenerDB interface {
	BlockListenerDBView
	SaveOrUpdateLastBlockNumber(lastBlock BlockListener) error
}

type BlockListenerDBView interface {
	GetLastBlockNumber() (lastBlock *BlockListener, err error)
}

func NewBlockListenerDB(db *gorm.DB) BlockListenerDB {
	return &blockListenerDB{gorm: db}
}

func (db blockListenerDB) SaveOrUpdateLastBlockNumber(lastBlock BlockListener) error {
	bl := new(BlockListener)
	var exitsLastBlock BlockListener

	err := db.gorm.Table(bl.TableName()).Take(&exitsLastBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info("No last block number found, will insert it")
			lastBlock.Created = uint64(time.Now().Unix())
			lastBlock.Updated = uint64(time.Now().Unix())
			result := db.gorm.Table(bl.TableName()).Omit("guid").Create(&lastBlock)
			if result.Error != nil {
				return result.Error
			}
			return nil
		}
		return err
	} else {
		lastBlock.Updated = uint64(time.Now().Unix())
		updateResult := db.gorm.Table(bl.TableName()).Where("1 = ?", 1).Omit("guid", "created").Updates(&lastBlock)
		if updateResult.Error != nil {
			return updateResult.Error
		}
	}
	return nil
}

func (db blockListenerDB) GetLastBlockNumber() (lastBlock *BlockListener, err error) {
	bl := new(BlockListener)
	err = db.gorm.Table(bl.TableName()).Take(&lastBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return lastBlock, nil
}
