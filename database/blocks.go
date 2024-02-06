package database

import (
	"math/big"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"

	"github.com/FishcakeLab/fishcake-service/database/utils"
)

type BlockHeader struct {
	Hash       common.Hash `gorm:"primaryKey;serializer:bytes"`
	ParentHash common.Hash `gorm:"serializer:bytes"`
	Number     *big.Int    `gorm:"serializer:u256"`
	Timestamp  uint64
	RLPHeader  *utils.RLPHeader `gorm:"serializer:rlp;column:rlp_bytes"`
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
	StoreL2BlockHeaders([]BlockHeader) error
}

type blocksDB struct {
	gorm *gorm.DB
}

func (b blocksDB) BlockHeader(hash common.Hash) (*BlockHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (b blocksDB) BlockHeaderWithFilter(header BlockHeader) (*BlockHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (b blocksDB) BlockHeaderWithScope(f func(db *gorm.DB) *gorm.DB) (*BlockHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (b blocksDB) LatestBlockHeader() (*BlockHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (b blocksDB) StoreBlockHeaders(headers []BlockHeader) error {
	//TODO implement me
	panic("implement me")
}

func (b blocksDB) StoreL2BlockHeaders(headers []BlockHeader) error {
	//TODO implement me
	panic("implement me")
}

func NewBlocksDB(db *gorm.DB) BlocksDB {
	return &blocksDB{gorm: db}
}
