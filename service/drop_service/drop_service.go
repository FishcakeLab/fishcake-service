package drop_service

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/drop"
	"math/big"
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
		if info.ActivityId != 0 {
			continue
		}
		amount5 := new(big.Int).Mul(
			big.NewInt(5),
			new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil),
		)
		amount10 := new(big.Int).Mul(
			big.NewInt(10),
			new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil),
		)

		if infos[i].DropAmount.Cmp(amount5) == 0 {
			infos[i].BusinessName = "Fishcake Wallet 🍥"
		} else if infos[i].DropAmount.Cmp(amount10) == 0 {
			infos[i].BusinessName = "Congrats First Claim 🍥"
		}
	}
	return infos, count
}
