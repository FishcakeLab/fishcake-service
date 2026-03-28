package mining_record_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/mining_record"
)

type MiningRecordService interface {
	List(address, recordType string, lastTimestamp uint64, lastId string, limit int) ([]mining_record.MiningRecord, error)
}

type miningRecordService struct {
	Db *database.DB
}

func NewMiningRecordService(db *database.DB) MiningRecordService {
	return &miningRecordService{Db: db}
}

func (n *miningRecordService) List(address, recordType string, lastTimestamp uint64, lastId string, limit int) ([]mining_record.MiningRecord, error) {
	infos, err := n.Db.MiningRecordDB.List(address, recordType, lastTimestamp, lastId, limit)
	return infos, err
}
