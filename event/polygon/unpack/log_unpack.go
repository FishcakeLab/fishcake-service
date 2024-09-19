package unpack

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/account_nft_info"
	"github.com/FishcakeLab/fishcake-service/database/activity"
	"github.com/FishcakeLab/fishcake-service/database/drop"
	"github.com/FishcakeLab/fishcake-service/database/event"
	"github.com/FishcakeLab/fishcake-service/database/token_nft"
	"github.com/FishcakeLab/fishcake-service/event/polygon/abi"
)

var (
	NftTokenUnpack, _ = abi.NewNftManagerFilterer(common.Address{}, nil)
	MerchantUnpack, _ = abi.NewFishcakeEventManagerFilterer(common.Address{}, nil)
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
		ReturnAmount:       big.NewInt(0),
		MinedAmount:        big.NewInt(0),
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
	ReturnAmount := uEvent.ReturnAmount
	MinedAmount := uEvent.MinedAmount
	return db.ActivityInfoDB.ActivityFinish(ActivityId, ReturnAmount, MinedAmount)
}

func MintNft(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := NftTokenUnpack.ParseCreateNFT(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	token := token_nft.TokenNft{
		TokenId:         uEvent.TokenId.Int64(),
		Who:             uEvent.Creator.String(),
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
	accountNftInfo := account_nft_info.AccountNftInfo{
		Address: uEvent.Creator.String(),
	}
	if uEvent.Type == 1 {
		accountNftInfo.ProDeadline = uEvent.Deadline.Uint64()
	} else {
		accountNftInfo.BasicDeadline = uEvent.Deadline.Uint64()
	}
	if err := db.Transaction(func(tx *database.DB) error {
		if err := tx.TokenNftDB.StoreTokenNft(token); err != nil {
			return err
		}
		if err := tx.AccountNftInfoDB.StoreAccountNftInfo(accountNftInfo); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
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
		Address:         uEvent.Who.String(),
		DropAmount:      uEvent.DropAmt,
		ActivityId:      uEvent.ActivityId.Int64(),
		DropType:        1,
		Timestamp:       event.Timestamp,
		TransactionHash: event.TransactionHash.String(),
		EventSignature:  event.EventSignature.String(),
	}

	if err := db.Transaction(func(tx *database.DB) error {
		resultErr, exist := tx.DropInfoDB.IsExist(drop.TransactionHash, drop.EventSignature, drop.DropType)
		if !exist && resultErr == nil {
			if err := tx.DropInfoDB.StoreDropInfo(drop); err != nil {
				return err
			}
			if err := tx.ActivityInfoDB.UpdateActivityInfo(uEvent.ActivityId.String()); err != nil {
				return err
			}
		} else {
			return resultErr
		}
		// create merchant drop record
		activityInfo := tx.ActivityInfoDB.ActivityInfo(int(drop.ActivityId))
		drop.Address = activityInfo.BusinessAccount
		drop.DropType = 2
		if err := tx.DropInfoDB.StoreDropInfo(drop); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
