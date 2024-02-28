package unpack

import (
	"github.com/FishcakeLab/fishcake-service/common/global_const"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/activity"
	"github.com/FishcakeLab/fishcake-service/database/drop"
	"github.com/FishcakeLab/fishcake-service/database/event"
	"github.com/FishcakeLab/fishcake-service/database/token_nft"
	"github.com/FishcakeLab/fishcake-service/event/polygon/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	NftTokenUnpack, _ = abi.NewNftTokenManagerFilterer(common.Address{}, nil)
	MerchantUnpack, _ = abi.NewMerchantMangerFilterer(common.Address{}, nil)
)

func ActivityAdd(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := MerchantUnpack.ParseActivityAdd(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	activityInfo := activity.ActivityInfo{
		ActivityId:         uEvent.ActivityId.Int64(),
		BusinessName:       uEvent.BusinessName,
		BusinessAccount:    uEvent.Who.String(),
		ActivityContent:    uEvent.ActivityContent,
		LatitudeLongitude:  uEvent.LatitudeLongitude,
		ActivityCreateTime: int64(event.Timestamp),
		ActivityDeadline:   uEvent.ActivityDeadLine.Int64(),
		DropType:           int8(uEvent.DropType),
		DropNumber:         uEvent.DropNumber.Int64(),
		MinDropAmt:         uEvent.MinDropAmt.Int64(),
		MaxDropAmt:         uEvent.MaxDropAmt.Int64(),
		TokenContractAddr:  uEvent.TokenContractAddr.String(),
		ActivityStatus:     1,
	}
	return db.ActivityInfoDB.StoreActivityInfo(activityInfo)
}

func ActivityFinish(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := MerchantUnpack.ParseActivityFinish(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	ActivityId := uEvent.ActivityId.String()
	return db.ActivityInfoDB.ActivityFinish(ActivityId)
}

func MintNft(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := NftTokenUnpack.ParseTransfer(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	if uEvent.From == common.HexToAddress(global_const.ZeroAddress) {
		token := token_nft.TokenNft{
			TokenId:         uEvent.TokenId.Int64(),
			Address:         uEvent.To.String(),
			ContractAddress: event.ContractAddress.String(),
			TokenAmount:     1,
			TokenUrl:        "baseUrl:" + uEvent.TokenId.String(),
			Timestamp:       event.Timestamp,
		}
		return db.TokenNftDB.StoreTokenNft(token)
	}
	return nil
}

func Drop(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := MerchantUnpack.ParseDrop(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	drop := drop.DropInfo{
		Address:    uEvent.Who.String(),
		DropAmount: uEvent.DropAmt.Int64(),
		ActivityId: uEvent.ActivityId.Int64(),
		Timestamp:  event.Timestamp,
	}
	return db.DropInfoDB.StoreDropInfo(drop)

	return nil
}
