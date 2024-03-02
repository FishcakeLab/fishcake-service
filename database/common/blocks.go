package common

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"

	"github.com/FishcakeLab/fishcake-service/database/utils"
	_ "github.com/FishcakeLab/fishcake-service/database/utils/serializers"
	"github.com/ethereum/go-ethereum/common"
)

type BlockHeader struct {
	GUID       uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	Hash       common.Hash `gorm:"serializer:bytes"`
	ParentHash common.Hash `gorm:"serializer:bytes"`
	Number     *big.Int    `gorm:"serializer:u256"`
	Timestamp  uint64
	RLPHeader  *utils.RLPHeader `gorm:"serializer:rlp;column:rlp_bytes"`
}

func (BlockHeader) TableName() string {
	return "block_header"
}

type BlocksView interface {
	BlockHeader(common.Hash) (*BlockHeader, error)
	BlockHeaderWithFilter(BlockHeader) (*BlockHeader, error)
	BlockHeaderWithScope(func(db *gorm.DB) *gorm.DB) (*BlockHeader, error)
	LatestBlockHeader() (*BlockHeader, error)
}

type BlocksDB interface {
	BlocksView
	StoreBlockHeaders([]BlockHeader) error
}

type blocksDB struct {
	gorm *gorm.DB
}

func (b blocksDB) BlockHeader(hash common.Hash) (*BlockHeader, error) {
	return b.BlockHeaderWithFilter(BlockHeader{Hash: hash})
}

func (b blocksDB) BlockHeaderWithFilter(header BlockHeader) (*BlockHeader, error) {
	return b.BlockHeaderWithScope(func(gorm *gorm.DB) *gorm.DB { return gorm.Where(&header) })
}

func (b blocksDB) BlockHeaderWithScope(f func(db *gorm.DB) *gorm.DB) (*BlockHeader, error) {
	var header BlockHeader
	result := b.gorm.Table("block_headers").Scopes(f).Take(&header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &header, nil
}

func (b blocksDB) LatestBlockHeader() (*BlockHeader, error) {
	var header BlockHeader
	result := b.gorm.Table("block_headers").Order("number DESC").Take(&header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &header, nil
}

func (b blocksDB) StoreBlockHeaders(headers []BlockHeader) error {
	result := b.gorm.Table("block_headers").Omit("guid").Create(&headers)
	return result.Error
}

func NewBlocksDB(db *gorm.DB) BlocksDB {
	return &blocksDB{gorm: db}
}
