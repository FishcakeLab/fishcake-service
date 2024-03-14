package unpack

import (
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/activity"
	"github.com/FishcakeLab/fishcake-service/database/drop"
	"github.com/FishcakeLab/fishcake-service/database/event"
	"github.com/FishcakeLab/fishcake-service/database/token_nft"
	"github.com/FishcakeLab/fishcake-service/event/polygon/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	NftTokenUnpack, _ = abi.NewNFTManagerFilterer(common.Address{}, nil)
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
		MinDropAmt:         uEvent.MinDropAmt,
		MaxDropAmt:         uEvent.MaxDropAmt,
		TokenContractAddr:  uEvent.TokenContractAddr.String(),
		ActivityStatus:     1,
		AlreadyDropNumber:  0,
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
	uEvent, unpackErr := NftTokenUnpack.ParseMintNewEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	token := token_nft.TokenNft{
		TokenId:         uEvent.TokenId.Int64(),
		BusinessName:    uEvent.BusinessName,
		Description:     uEvent.Description,
		ImgUrl:          uEvent.ImgUrl,
		BusinessAddress: uEvent.BusinessAddress,
		WebSite:         uEvent.WebSite,
		Social:          uEvent.Social,
		ContractAddress: event.ContractAddress.String(),
		CostValue:       uEvent.Value,
		Deadline:        uEvent.Deadline.Uint64(),
		NftType:         int8(uEvent.Type),
	}
	return db.TokenNftDB.StoreTokenNft(token)
}

func Drop(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := MerchantUnpack.ParseDrop(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	drop := drop.DropInfo{
		Address:    uEvent.Who.String(),
		DropAmount: uEvent.DropAmt,
		ActivityId: uEvent.ActivityId.Int64(),
		Timestamp:  event.Timestamp,
	}

	if err := db.Transaction(func(tx *database.DB) error {
		if err := tx.DropInfoDB.StoreDropInfo(drop); err != nil {
			return err
		}
		if err := tx.ActivityInfoDB.UpdateActivityInfo(uEvent.ActivityId.String()); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
