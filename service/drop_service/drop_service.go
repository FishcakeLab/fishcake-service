package drop_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/drop"
)

type DropService interface {
	DropInfoList(pageNum, pageSize int, address, dropType string) ([]drop.DropInfo, int)
}

type dropService struct {
	Db *database.DB
}

func NewDropService(db *database.DB) DropService {
	return &dropService{Db: db}
}

func (n *dropService) DropInfoList(pageNum, pageSize int, address, dropType string) ([]drop.DropInfo, int) {
	infos, count := n.Db.DropInfoDB.List(int(uint64(pageNum)), pageSize, address, dropType)
	for i, info := range infos {
		systemDropType := info.SystemDropType
		if systemDropType == "" {
			continue
		}
		infos[i].BusinessName = GetDropTypeDesc(systemDropType)
	}
	return infos, count
}

func GetDropTypeDesc(systemDropType string) string {
	switch systemDropType {
	case "1":
		return "Fishcake Wallet 🍥"
	case "2":
		return "Congrats First Claim 🍥"
	default:
		return "Unknown"
	}
}
