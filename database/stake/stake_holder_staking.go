package stake

import (
	"errors"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/FishcakeLab/fishcake-service/common/enum"
	"github.com/FishcakeLab/fishcake-service/common/errors_h"
	_ "github.com/FishcakeLab/fishcake-service/database/utils/serializers"
)

// StakeHolderStaking represents the on-chain staking record for each holder
type StakeHolderStaking struct {
	ID            string    `gorm:"column:id;primaryKey;default:replace((uuid_generate_v4())::text, '-'::text, ''::text)" json:"id"`
	UserAddress   string    `gorm:"column:user_address" json:"userAddress"`
	Amount        *big.Int  `gorm:"serializer:u256;column:amount" json:"amount"`
	StakingType   int16     `gorm:"column:staking_type" json:"stakingType"`
	StartTime     int64     `gorm:"column:start_time" json:"startTime"`
	EndTime       int64     `gorm:"column:end_time" json:"endTime"`
	TokenID       int64     `gorm:"column:token_id" json:"tokenId"`
	NftApr        int64     `gorm:"column:nft_apr" json:"nftApr"`
	IsAutoRenew   bool      `gorm:"column:is_auto_renew" json:"isAutoRenew"`
	MessageNonce  int64     `gorm:"column:message_nonce" json:"messageNonce"`
	TxMessageHash string    `gorm:"column:tx_message_hash" json:"txMessageHash"`
	StakingReward *big.Int  `gorm:"serializer:u256;column:staking_reward" json:"stakingReward"`
	StakingStatus int16     `gorm:"column:staking_status;default:0" json:"stakingStatus"` // 0=staking, 1=ended
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"createTime"`
}

func (StakeHolderStaking) TableName() string {
	return "stake_holder_staking"
}

// StakeHolderStakingView defines read-only queries
type StakeHolderStakingView interface {
	GetByAddressAndNonce(address string, nonce int64) (*StakeHolderStaking, error)
	ListByAddress(address string, pageNum, pageSize int) ([]StakeHolderStaking, int)
	ListAllActive() ([]StakeHolderStaking, error)
	GetStakeRank(monthFilter bool) ([]StakeRank, error)
	GetClaimedRank(monthFilter bool) ([]ClaimedRank, error)
	GetTotalRewardRank(monthFilter bool) ([]TotalRewardRank, error)
	GetUserStakingInfo(address string, statusFilter *int, pageNum, pageSize int) ([]StakeHolderStaking, int64, error)
}

// StakeHolderStakingDB defines full DB operations
type StakeHolderStakingDB interface {
	StakeHolderStakingView
	InsertDepositRecord(record StakeHolderStaking) error
	UpdateWithdrawRecord(address string, nonce int64, txHash string, reward *big.Int) error
	UpdateStatus(address string, nonce int64, status int16) error
}

// 数据结构：用于接口返回
type StakeRank struct {
	UserAddress string   `json:"userAddress"`
	TotalStake  *big.Int `json:"totalStake"`
}

type ClaimedRank struct {
	UserAddress string   `json:"userAddress"`
	Claimed     *big.Int `json:"claimed"`
}

type TotalRewardRank struct {
	UserAddress string   `json:"userAddress"`
	TotalReward *big.Int `json:"totalReward"`
	Claimed     *big.Int `json:"claimed"`
	Unclaimed   *big.Int `json:"unclaimed"`
}

// Implementation
type stakeHolderStakingDB struct {
	db *gorm.DB
}

// GetUserStakingInfo — paginated list of user staking records
func (d stakeHolderStakingDB) GetUserStakingInfo(
	address string,
	status *int,
	page, size int,
) ([]StakeHolderStaking, int64, error) {

	table := StakeHolderStaking{}.TableName()

	query := d.db.Table(table).
		Where("user_address = ?", address)

	if status != nil {
		query = query.Where("staking_status = ?", *status)
	}

	// Count total
	var total int64
	if err := query.Count(&total).Error; err != nil {
		log.Error("Count user staking records fail", "err", err)
		return nil, 0, err
	}

	// Pagination
	offset := (page - 1) * size

	var records []StakeHolderStaking
	err := query.
		Order("start_time DESC").
		Limit(size).
		Offset(offset).
		Find(&records).Error

	if err != nil {
		log.Error("Query user staking records fail", "err", err)
		return nil, 0, err
	}

	return records, total, nil
}

// InsertDepositRecord — triggered by StakeHolderDepositStaking event
func (d stakeHolderStakingDB) InsertDepositRecord(record StakeHolderStaking) error {
	err := d.db.Table(record.TableName()).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_address"}, {Name: "message_nonce"}},
			DoNothing: true,
		}).
		Create(&record).Error

	if err != nil {
		log.Error("Insert deposit staking record fail", "err", err)
		return err
	}
	return nil
}

// UpdateWithdrawRecord — triggered by StakeHolderWithdrawStaking event
func (d stakeHolderStakingDB) UpdateWithdrawRecord(address string, nonce int64, txHash string, reward *big.Int) error {
	result := d.db.Table(StakeHolderStaking{}.TableName()).
		Where("user_address = ? AND message_nonce = ?", address, nonce).
		Updates(map[string]interface{}{
			"staking_status":  1,
			"staking_reward":  reward,
			"tx_message_hash": txHash,
		})

	if result.Error != nil {
		log.Error("Update withdraw staking record fail", "err", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no matching staking record found")
	}
	return nil
}

// UpdateStatus — manually update staking status (e.g. job timeout cleanup)
func (d stakeHolderStakingDB) UpdateStatus(address string, nonce int64, status int16) error {
	err := d.db.Table(StakeHolderStaking{}.TableName()).
		Where("user_address = ? AND message_nonce = ?", address, nonce).
		Update("staking_status", status).Error

	if err != nil {
		log.Error("Update staking status fail", "err", err)
		return err
	}
	return nil
}

// GetByAddressAndNonce — get a single record by address + nonce
func (d stakeHolderStakingDB) GetByAddressAndNonce(address string, nonce int64) (*StakeHolderStaking, error) {
	var record StakeHolderStaking
	result := d.db.Table(record.TableName()).
		Where("user_address = ? AND message_nonce = ?", address, nonce).
		Take(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Error("Get staking record fail", "err", result.Error)
		return nil, result.Error
	}
	return &record, nil
}

// ListByAddress — paginated query by user address
func (d stakeHolderStakingDB) ListByAddress(address string, pageNum, pageSize int) ([]StakeHolderStaking, int) {
	var records []StakeHolderStaking
	var count int64

	this := d.db.Table(StakeHolderStaking{}.TableName()).
		Where("user_address = ?", address).
		Order("create_time DESC")

	this.Count(&count)

	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}

	result := this.Find(&records)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		errors_h.NewErrorByEnum(enum.DataErr)
	}
	return records, int(count)
}

// ListAllActive — list all staking in progress (status=0)
func (d stakeHolderStakingDB) ListAllActive() ([]StakeHolderStaking, error) {
	var records []StakeHolderStaking
	err := d.db.Table(StakeHolderStaking{}.TableName()).
		Where("staking_status = ?", 0).
		Find(&records).Error

	if err != nil {
		log.Error("Query active staking fail", "err", err)
		return nil, err
	}
	return records, nil
}

// NewStakeHolderStakingDB — constructor
func NewStakeHolderStakingDB(db *gorm.DB) StakeHolderStakingDB {
	return &stakeHolderStakingDB{db: db}
}

// GetStakeRank — 排名用户的质押总额
func (d *stakeHolderStakingDB) GetStakeRank(monthFilter bool) ([]StakeRank, error) {
	var ranks []StakeRank
	query := d.db.Table(StakeHolderStaking{}.TableName()).
		Select("user_address, SUM(amount) as total_stake").
		Group("user_address").
		Order("SUM(amount) DESC")

	// 如果 monthFilter=true，过滤当月数据
	if monthFilter {
		now := time.Now()
		start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		query = query.Where("create_time >= ?", start)
	}

	rows, err := query.Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var addr string
		var amountStr string
		if err := rows.Scan(&addr, &amountStr); err != nil {
			return nil, err
		}
		val, _ := new(big.Int).SetString(amountStr, 10)
		ranks = append(ranks, StakeRank{
			UserAddress: addr,
			TotalStake:  val,
		})
	}
	return ranks, nil
}

// GetClaimedRank — 排名已领取奖励的用户（staking_status=1）
func (d *stakeHolderStakingDB) GetClaimedRank(monthFilter bool) ([]ClaimedRank, error) {
	var ranks []ClaimedRank
	query := d.db.Table(StakeHolderStaking{}.TableName()).
		Select("user_address, SUM(staking_reward) as claimed").
		Where("staking_status = ?", 1).
		Group("user_address").
		Order("SUM(staking_reward) DESC")

	if monthFilter {
		now := time.Now()
		start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		query = query.Where("create_time >= ?", start)
	}

	rows, err := query.Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var addr string
		var amountStr string
		if err := rows.Scan(&addr, &amountStr); err != nil {
			return nil, err
		}
		val, _ := new(big.Int).SetString(amountStr, 10)
		ranks = append(ranks, ClaimedRank{
			UserAddress: addr,
			Claimed:     val,
		})
	}
	return ranks, nil
}

//	GetTotalRewardRank — 计算总奖励排行榜
//
// totalReward = SUM(staking_reward[status=1]) + SUM(预计未领取奖励)
func (d *stakeHolderStakingDB) GetTotalRewardRank(monthFilter bool) ([]TotalRewardRank, error) {
	var result []TotalRewardRank
	now := time.Now()

	//  1. 获取所有地址列表（已结束 + 未结束）
	var addresses []string
	addrQuery := d.db.Table(StakeHolderStaking{}.TableName()).Distinct("user_address")
	if monthFilter {
		start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		addrQuery = addrQuery.Where("create_time >= ?", start)
	}
	if err := addrQuery.Find(&addresses).Error; err != nil {
		return nil, err
	}

	// 2. 遍历每个地址计算 totalReward
	for _, addr := range addresses {
		var claimedStr string
		var unclaimedRecords []StakeHolderStaking

		// 2.1 计算已领取的 staking_reward（status = 1）
		claimedQuery := d.db.Table(StakeHolderStaking{}.TableName()).
			Select("COALESCE(SUM(staking_reward), 0)").
			Where("staking_status = 1 AND user_address = ?", addr)
		if monthFilter {
			start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			claimedQuery = claimedQuery.Where("create_time >= ?", start)
		}
		if err := claimedQuery.Scan(&claimedStr).Error; err != nil {
			return nil, err
		}
		claimed, _ := new(big.Int).SetString(claimedStr, 10)

		// 2.2 查询未结束记录（status = 0）
		unclaimedQuery := d.db.Table(StakeHolderStaking{}.TableName()).
			Where("staking_status = 0 AND user_address = ?", addr)
		if monthFilter {
			start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			unclaimedQuery = unclaimedQuery.Where("create_time >= ?", start)
		}
		if err := unclaimedQuery.Find(&unclaimedRecords).Error; err != nil {
			return nil, err
		}

		// 2.3 遍历未结束记录计算 unclaimedReward
		unclaimed := big.NewInt(0)
		for _, rec := range unclaimedRecords {
			reward := calculateAprFundingGo(rec, now)
			unclaimed.Add(unclaimed, reward)
		}

		total := new(big.Int).Add(claimed, unclaimed)
		result = append(result, TotalRewardRank{
			UserAddress: addr,
			Claimed:     claimed,
			Unclaimed:   unclaimed,
			TotalReward: total,
		})
	}

	// 3. 按 totalReward 降序排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].TotalReward.Cmp(result[j].TotalReward) > 0
	})

	return result, nil
}

// calculateAprFundingGo —— Golang复刻版 Solidity calculateAprFunding()
// 从数据库记录计算出未领取奖励 unclaimedReward
func calculateAprFundingGo(rec StakeHolderStaking, now time.Time) *big.Int {
	stakingAmount := new(big.Int).Set(rec.Amount)
	stakingType := rec.StakingType
	stakingTime := rec.StartTime
	isAutoRenew := rec.IsAutoRenew
	nftApr := rec.NftApr

	// ========== 获取锁定期和质押APR ==========
	lockTime, stakingApr := getStakingPeriodAndAprGo(stakingType)
	totalRewardApr := nftApr + stakingApr

	// ========== 实际质押时长 ==========
	actualDuration := now.Unix() - stakingTime
	if actualDuration <= lockTime || isAutoRenew {
		// 原样
	} else {
		actualDuration = lockTime
	}

	// ========== 计算奖励 ==========
	reward := big.NewInt(0)

	if actualDuration < lockTime {
		// reward = (amount * totalRewardApr * actualDuration) / (100 * 365 days)
		reward.Mul(stakingAmount, big.NewInt(int64(totalRewardApr)))
		reward.Mul(reward, big.NewInt(actualDuration))
		denominator := big.NewInt(100 * 365 * 24 * 3600)
		reward.Div(reward, denominator)
	} else {
		// baseReward = (amount * totalRewardApr * lockTime) / (100 * 365 days)
		base := new(big.Int).Mul(stakingAmount, big.NewInt(int64(totalRewardApr)))
		base.Mul(base, big.NewInt(lockTime))
		base.Div(base, big.NewInt(100*365*24*3600))

		// extraReward = (amount * stakingApr * (actual - lock)) / (100 * 365 days)
		extra := new(big.Int).Mul(stakingAmount, big.NewInt(int64(stakingApr)))
		extra.Mul(extra, big.NewInt(actualDuration-lockTime))
		extra.Div(extra, big.NewInt(100*365*24*3600))

		reward.Add(base, extra)
	}

	// ========== Halve after 2026-01-01 ==========
	halfAprTime := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	if now.After(halfAprTime) {
		reward.Div(reward, big.NewInt(2))
	}

	return reward
}

// getStakingPeriodAndAprGo —— 对应 Solidity getStakingPeriodAndApr()
// 返回 (lockTime, stakingApr)
func getStakingPeriodAndAprGo(stakingType int16) (lockTime int64, stakingApr int64) {
	switch stakingType {
	case 1:
		// 30天锁定期，APR=3%
		return 30 * 24 * 3600, 3
	case 2:
		// 60天锁定期，APR=6%
		return 60 * 24 * 3600, 6
	case 3:
		// 90天锁定期，APR=9%
		return 90 * 24 * 3600, 9
	case 4:
		// 半年锁定期，APR=15%
		return 180 * 24 * 3600, 15
	default:
		// Solidity 里 require 限定了 stakingType 必须在 1~4
		// 为保持一致，这里返回0并由上层逻辑判断跳过
		return 0, 0
	}
}
