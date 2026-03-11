package account_nft_info

import (
	"errors"

	"gorm.io/gorm"
)

type AccountNftInfo struct {
	Id            string `json:"id" gorm:"id"`
	Address       string `json:"address" gorm:"address"`
	BasicDeadline uint64 `json:"basicDeadline" gorm:"basic_deadline"`
	ProDeadline   uint64 `json:"proDeadline" gorm:"pro_deadline"`
}

func (AccountNftInfo) TableName() string {
	return "account_nft_info"
}

type AccountNftInfoDBView interface {
}

type AccountNftInfoDB interface {
	AccountNftInfoDBView
	StoreAccountNftInfo(accountNftInfo AccountNftInfo) error
}

type accountNftInfoDB struct {
	db *gorm.DB
}

func (a accountNftInfoDB) StoreAccountNftInfo(accountNftInfo AccountNftInfo) error {
	nftInfo := new(AccountNftInfo)
	var exist AccountNftInfo
	err := a.db.Table(nftInfo.TableName()).Where("address ILIKE ?", accountNftInfo.Address).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := a.db.Table(nftInfo.TableName()).Omit("id").Create(&accountNftInfo)
			return result.Error
		}
	} else {
		needsUpdate := false
		if accountNftInfo.BasicDeadline > exist.BasicDeadline {
			exist.BasicDeadline = accountNftInfo.BasicDeadline
			needsUpdate = true
		}
		if accountNftInfo.ProDeadline > exist.ProDeadline {
			exist.ProDeadline = accountNftInfo.ProDeadline
			needsUpdate = true
		}

		if needsUpdate {
			updateResult := a.db.Table(AccountNftInfo{}.TableName()).Save(&exist)
			if updateResult.Error != nil {
				return updateResult.Error
			}
		}
	}
	return err
}

func NewAccountNftInfoDB(db *gorm.DB) AccountNftInfoDB {
	return &accountNftInfoDB{db: db}
}
