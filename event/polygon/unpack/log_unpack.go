package unpack

import (
	"github.com/FishcakeLab/fishcake-service/common/global_const"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/event"
	"github.com/FishcakeLab/fishcake-service/database/token_nft"
	"github.com/FishcakeLab/fishcake-service/event/polygon/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	NftTokenUnpack, _ = abi.NewNftTokenManagerFilterer(common.Address{}, nil)
	MerchantUnpack, _ = abi.NewMerchantMangerFilterer(common.Address{}, nil)
)

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
			TokenUrl:        "1111",
			Timestamp:       event.Timestamp,
		}
		return db.TokenNftDb.StoreTokenNft(token)
	}
	return nil
}
