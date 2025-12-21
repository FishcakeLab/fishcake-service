package unpack

import (
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/database/account_nft_info"
	"github.com/FishcakeLab/fishcake-service/database/activity"
	"github.com/FishcakeLab/fishcake-service/database/drop"
	"github.com/FishcakeLab/fishcake-service/database/event"
	"github.com/FishcakeLab/fishcake-service/database/stake"
	"github.com/FishcakeLab/fishcake-service/database/token_nft"
	"github.com/FishcakeLab/fishcake-service/database/token_transfer"
	"github.com/FishcakeLab/fishcake-service/event/polygon/abi"
)

var (
	NftTokenUnpack, _ = abi.NewNftManagerFilterer(common.Address{}, nil)
	MerchantUnpack, _ = abi.NewFishcakeEventManagerFilterer(common.Address{}, nil)
	StakingUnpack, _  = abi.NewStakingManagerFilterer(common.Address{}, nil)
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

	tokenSent := token_transfer.TokenSent{
		Address:      activityInfo.BusinessAccount,
		TokenAddress: activityInfo.TokenContractAddr,
		Amount: new(big.Int).Mul(
			activityInfo.MaxDropAmt,
			uEvent.DropNumber,
		),
		Description: uEvent.ActivityContent,
		Timestamp:   uint64(activityInfo.ActivityCreateTime),
	}

	if err := db.Transaction(func(tx *database.DB) error {

		if err := tx.ActivityInfoDB.StoreActivityInfo(activityInfo); err != nil {
			return err
		}

		if err := tx.TokenSentDB.StoreTokenSent(tokenSent); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil
	}
	return nil

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

	address := db.ActivityInfoDB.ActivityInfo(int(uEvent.ActivityId.Int64())).BusinessAccount
	content := db.ActivityInfoDB.ActivityInfo(int(uEvent.ActivityId.Int64())).ActivityContent

	tokenReceived := token_transfer.TokenReceived{
		Address:      address,
		TokenAddress: db.ActivityInfoDB.ActivityInfo(int(uEvent.ActivityId.Int64())).TokenContractAddr,
		Amount:       ReturnAmount,
		Description:  content,
		Timestamp:    event.Timestamp,
	}

	if err := db.Transaction(func(tx *database.DB) error {
		if err := tx.ActivityInfoDB.ActivityFinish(ActivityId, ReturnAmount, MinedAmount); err != nil {
			return err
		}

		if err := tx.TokenReceivedDB.StoreTokenReceived(tokenReceived); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil

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

func MintBoosterNft(event event.ContractEvent, db *database.DB) error {
	return db.Transaction(func(tx *database.DB) error {
		rlpLog := event.RLPLog
		uEvent, err := NftTokenUnpack.ParseMintBoosterNFT(*rlpLog)
		if err != nil {
			return err
		}

		address := strings.ToLower(uEvent.Miner.Hex())
		tokenId := uEvent.TokenId.Int64()
		nftType := int8(uEvent.NftType)
		mintTime := uEvent.MintTime.Int64()
		usedPower := uEvent.UsedFishCakePower

		// ---------- 1. 锁行读取 mining_info ----------
		miningInfo, err := tx.MiningInfoDB.GetByAddressForUpdate(address)
		if err != nil {
			return err
		}

		// ---------- 2. 幂等判断（必须在事务里） ----------
		if miningInfo.LastMintTime.Int64() >= mintTime {
			return nil
		}

		// ---------- 3. 更新 mining_info ----------
		newMiningInfo := &activity.MiningInfo{
			Id:                 miningInfo.Id,
			Address:            miningInfo.Address,
			MinedAmount:        miningInfo.MinedAmount,
			MinedFishCakePower: new(big.Int).Sub(miningInfo.MinedFishCakePower, usedPower),
			LastMintTime:       big.NewInt(mintTime),
		}

		if err := tx.MiningInfoDB.Update(newMiningInfo); err != nil {
			return err
		}

		// ---------- 4. Upsert Booster NFT ----------
		token := token_nft.TokenNft{
			TokenId:         tokenId,
			Who:             address,
			NftType:         nftType,
			ContractAddress: strings.ToLower(event.ContractAddress.Hex()),
			BusinessName:    "BoosterNFT",
			Description:     "Miner Booster NFT",
		}

		if err := tx.TokenNftDB.UpsertBoosterNft(token); err != nil {
			return err
		}

		return nil
	})
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

	tokenReceived := token_transfer.TokenReceived{
		Address:      drop.Address,
		TokenAddress: db.ActivityInfoDB.ActivityInfo(int(uEvent.ActivityId.Int64())).TokenContractAddr,
		Amount:       uEvent.DropAmt,
		Description:  db.ActivityInfoDB.ActivityInfo(int(uEvent.ActivityId.Int64())).ActivityContent,
		Timestamp:    uint64(event.Timestamp),
	}

	if err := db.Transaction(func(tx *database.DB) error {
		resultErr, exist := tx.DropInfoDB.IsExist(drop.TransactionHash, drop.EventSignature, drop.DropType)
		if !exist && resultErr == nil {
			if err := tx.DropInfoDB.StoreDropInfo(drop); err != nil {
				return err
			}
			if err := tx.TokenReceivedDB.StoreTokenReceived(tokenReceived); err != nil {
				return err
			}
			// update activity info already drop number
			if err := tx.ActivityInfoDB.UpdateActivityInfo(uEvent.ActivityId.String()); err != nil {
				return err
			} // 重复加了
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

// StakeHolderDepositStaking 事件解析与入库
func StakeHolderDepositStaking(event event.ContractEvent, db *database.DB) error {
	return db.Transaction(func(tx *database.DB) error {
		rlpLog := event.RLPLog
		uEvent, err := StakingUnpack.ParseStakeHolderDepositStaking(*rlpLog)
		if err != nil {
			log.Error("unpack StakeHolderDepositStaking fail", "err", err)
			return err
		}

		record := stake.StakeHolderStaking{
			UserAddress:   uEvent.Staker.Hex(),
			Amount:        uEvent.Amount,
			StakingType:   int16(uEvent.StakingType),
			StartTime:     uEvent.StartStakingTime.Int64(),
			EndTime:       uEvent.EndStakingTime.Int64(),
			TokenID:       uEvent.BindingNft.Int64(), // Booster tokenId
			NftApr:        uEvent.NftApr.Int64(),
			IsAutoRenew:   uEvent.IsAutoRenew,
			MessageNonce:  uEvent.MessageNonce.Int64(),
			TxMessageHash: "",
			StakingReward: big.NewInt(0),
			StakingStatus: 0,
			CreateTime:    time.Now(),
		}

		// 1. 插入 staking 记录
		if err := tx.StakingDB.InsertDepositRecord(record); err != nil {
			log.Error("insert StakeHolderDepositStaking fail",
				"user", uEvent.Staker.String(),
				"nonce", uEvent.MessageNonce.Int64(),
				"block", event.BlockNumber,
				"err", err)
			return err
		}

		// 2. 标记 Booster NFT 为 used（nft_type + 10）
		tokenId := uEvent.BindingNft.Int64()
		if tokenId > 0 {
			if err := tx.TokenNftDB.MarkBoosterUsed(tokenId); err != nil {
				log.Error("mark booster used fail",
					"tokenId", tokenId,
					"err", err)
				return err
			}
		}

		log.Info("StakeHolderDepositStaking success",
			"user", uEvent.Staker.String(),
			"nonce", uEvent.MessageNonce.Int64(),
			"amount", uEvent.Amount.String(),
			"boosterTokenId", tokenId,
			"block", event.BlockNumber)

		return nil
	})
}

// StakeHolderWithdrawStaking 事件解析与更新
func StakeHolderWithdrawStaking(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := StakingUnpack.ParseStakeHolderWithdrawStaking(*rlpLog)
	if unpackErr != nil {
		log.Error("unpack StakeHolderWithdrawStaking fail", "err", unpackErr)
		return unpackErr
	}

	userAddr := uEvent.Recipant.String()
	messageNonce := uEvent.MessageNonce.Int64()
	txHash := common.BytesToHash(uEvent.MessageHash[:]).Hex()
	reward := uEvent.RewardAprFunding

	if err := db.StakingDB.UpdateWithdrawRecord(userAddr, messageNonce, txHash, reward); err != nil {
		log.Error("update StakeHolderWithdrawStaking fail",
			"user", userAddr,
			"nonce", messageNonce,
			"block", event.BlockNumber,
			"err", err)
		return err
	}

	// 成功日志
	log.Info("✅ StakeHolderWithdrawStaking success",
		"user", userAddr,
		"nonce", messageNonce,
		"reward", reward.String(),
		"block", event.BlockNumber)

	return nil
}

// Transfer 事件解析与入库
func Transfer(event event.ContractEvent, db *database.DB, address string) error {
	rlpLog := event.RLPLog
	from := rlpLog.Topics[1].Hex()
	to := rlpLog.Topics[2].Hex()
	value := new(big.Int).SetBytes(rlpLog.Data)

	// 忽略给事件合约地址转账的记录，以及从事件合约地址转出的记录
	if from == "0x2CAf752814f244b3778e30c27051cc6B45CB1fc9" || to == "0x2CAf752814f244b3778e30c27051cc6B45CB1fc9" {
		return nil
	}

	tokenSent := token_transfer.TokenSent{
		Address:      from,
		TokenAddress: address,
		Amount:       value,
		Description:  "ERC20 Token Transfer Sent",
		Timestamp:    uint64(event.Timestamp),

		TxHash:   event.RLPLog.TxHash.Hex(),
		LogIndex: uint(event.RLPLog.Index),
	}

	tokenReceived := token_transfer.TokenReceived{
		Address:      to,
		TokenAddress: address,
		Amount:       value,
		Description:  "ERC20 Token Transfer Received",
		Timestamp:    uint64(event.Timestamp),

		TxHash:   event.RLPLog.TxHash.Hex(),
		LogIndex: uint(event.RLPLog.Index),
	}

	if err := db.Transaction(func(tx *database.DB) error {
		if err := tx.TokenSentDB.StoreTokenSent(tokenSent); err != nil {
			return err
		}
		if err := tx.TokenReceivedDB.StoreTokenReceived(tokenReceived); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil

}
