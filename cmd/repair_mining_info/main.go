package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/FishcakeLab/fishcake-service/config"
)

type repairRow struct {
	NormAddress    string  `gorm:"column:norm_address"`
	DisplayAddress string  `gorm:"column:display_address"`
	CanonicalMined string  `gorm:"column:canonical_mined"`
	KeeperID       *string `gorm:"column:keeper_id"`
	CurrentMined   *string `gorm:"column:current_mined"`
	CurrentPower   *string `gorm:"column:current_power"`
	LastMintTime   *string `gorm:"column:last_mint_time"`
	MiningRowCount int64   `gorm:"column:mining_row_count"`
}

type repairPlan struct {
	Address        string
	KeeperID       *string
	CanonicalMined string
	NewMinedAmount string
	NewPower       string
	LastMintTime   string
	ConsumedPower  string
	CurrentMined   string
	CurrentPower   string
	NeedsWrite     bool
	NeedsInsert    bool
	MiningRowCount int64
}

const repairQuery = `
WITH activity_totals AS (
    SELECT
        lower(ai.business_account) AS norm_address,
        MIN(ai.business_account) AS display_address,
        COALESCE(SUM(ai.mined_amount::numeric), 0)::text AS canonical_mined
    FROM activity_info ai
    WHERE ai.activity_status = 2
      AND ai.business_account IS NOT NULL
      AND length(ai.business_account) > 0
    GROUP BY lower(ai.business_account)
),
mining_ranked AS (
    SELECT
        mi.id,
        mi.address,
        lower(mi.address) AS norm_address,
        mi.mined_amount::text AS current_mined,
        mi.mined_fishcakepower::text AS current_power,
        mi.last_mint_time::text AS last_mint_time,
        ROW_NUMBER() OVER (PARTITION BY lower(mi.address) ORDER BY mi.id) AS rn,
        COUNT(*) OVER (PARTITION BY lower(mi.address)) AS row_count
    FROM mining_info mi
),
mining_primary AS (
    SELECT *
    FROM mining_ranked
    WHERE rn = 1
)
SELECT
    COALESCE(a.norm_address, m.norm_address) AS norm_address,
    COALESCE(m.address, a.display_address) AS display_address,
    COALESCE(a.canonical_mined, '0') AS canonical_mined,
    m.id AS keeper_id,
    m.current_mined,
    m.current_power,
    m.last_mint_time,
    COALESCE(m.row_count, 0) AS mining_row_count
FROM activity_totals a
FULL OUTER JOIN mining_primary m ON a.norm_address = m.norm_address
WHERE (? = '' OR COALESCE(a.norm_address, m.norm_address) = lower(?))
ORDER BY COALESCE(a.norm_address, m.norm_address)
`

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))

	configPath := flag.String("config", "./config.yaml", "Path to config file")
	address := flag.String("address", "", "Only repair one address")
	apply := flag.Bool("apply", false, "Write changes to the database")
	flag.Parse()

	cfg, err := config.New(*configPath)
	if err != nil {
		log.Error("failed to load config", "err", err)
		os.Exit(1)
	}

	db, err := openDB(cfg)
	if err != nil {
		log.Error("failed to connect database", "err", err)
		os.Exit(1)
	}

	rows, err := loadRepairRows(db, strings.TrimSpace(*address))
	if err != nil {
		log.Error("failed to load repair rows", "err", err)
		os.Exit(1)
	}

	plans, err := buildPlans(rows)
	if err != nil {
		log.Error("failed to build repair plan", "err", err)
		os.Exit(1)
	}

	printPlan(plans, *apply)

	if !*apply {
		log.Info("dry run complete", "rows", len(plans))
		return
	}

	if err := applyPlans(db, plans); err != nil {
		log.Error("failed to apply repair plan", "err", err)
		os.Exit(1)
	}

	log.Info("repair complete", "rows", len(plans))
}

func openDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", cfg.DbHost, cfg.DbName)
	if cfg.DbPort != 0 {
		dsn += fmt.Sprintf(" port=%d", cfg.DbPort)
	}
	if cfg.DbUser != "" {
		dsn += fmt.Sprintf(" user=%s", cfg.DbUser)
	}
	if cfg.DbPassword != "" {
		dsn += fmt.Sprintf(" password=%s", cfg.DbPassword)
	}

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func loadRepairRows(db *gorm.DB, address string) ([]repairRow, error) {
	var rows []repairRow
	err := db.Raw(repairQuery, address, address).Scan(&rows).Error
	return rows, err
}

func buildPlans(rows []repairRow) ([]repairPlan, error) {
	plans := make([]repairPlan, 0, len(rows))
	for _, row := range rows {
		if row.MiningRowCount > 1 {
			return nil, fmt.Errorf("address %s has %d mining_info rows; aborting automatic repair", row.DisplayAddress, row.MiningRowCount)
		}

		canonical, err := parseBigInt(row.CanonicalMined)
		if err != nil {
			return nil, fmt.Errorf("address %s invalid canonical mined amount: %w", row.DisplayAddress, err)
		}
		currentMined, err := parseBigIntPtr(row.CurrentMined)
		if err != nil {
			return nil, fmt.Errorf("address %s invalid current mined amount: %w", row.DisplayAddress, err)
		}
		currentPower, err := parseBigIntPtr(row.CurrentPower)
		if err != nil {
			return nil, fmt.Errorf("address %s invalid current power amount: %w", row.DisplayAddress, err)
		}

		consumedPower := new(big.Int).Sub(currentMined, currentPower)
		if consumedPower.Sign() < 0 {
			consumedPower = big.NewInt(0)
		}

		newPower := new(big.Int).Sub(new(big.Int).Set(canonical), consumedPower)
		if newPower.Sign() < 0 {
			newPower = big.NewInt(0)
		}

		lastMintTime := "0"
		if row.LastMintTime != nil && strings.TrimSpace(*row.LastMintTime) != "" {
			lastMintTime = strings.TrimSpace(*row.LastMintTime)
		}

		currentMinedStr := currentMined.String()
		currentPowerStr := currentPower.String()
		newMinedStr := canonical.String()
		newPowerStr := newPower.String()
		needsInsert := row.KeeperID == nil || strings.TrimSpace(*row.KeeperID) == ""
		needsWrite := needsInsert || currentMinedStr != newMinedStr || currentPowerStr != newPowerStr

		plans = append(plans, repairPlan{
			Address:        row.DisplayAddress,
			KeeperID:       row.KeeperID,
			CanonicalMined: newMinedStr,
			NewMinedAmount: newMinedStr,
			NewPower:       newPowerStr,
			LastMintTime:   lastMintTime,
			ConsumedPower:  consumedPower.String(),
			CurrentMined:   currentMinedStr,
			CurrentPower:   currentPowerStr,
			NeedsWrite:     needsWrite,
			NeedsInsert:    needsInsert,
			MiningRowCount: row.MiningRowCount,
		})
	}
	return plans, nil
}

func applyPlans(db *gorm.DB, plans []repairPlan) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, plan := range plans {
			if !plan.NeedsWrite {
				continue
			}

			if plan.NeedsInsert {
				if err := tx.Exec(
					`INSERT INTO mining_info (address, mined_amount, mined_fishcakepower, last_mint_time) VALUES (?, ?, ?, ?)`,
					plan.Address,
					plan.NewMinedAmount,
					plan.NewPower,
					plan.LastMintTime,
				).Error; err != nil {
					return fmt.Errorf("insert mining_info for %s: %w", plan.Address, err)
				}
				continue
			}

			if err := tx.Exec(
				`UPDATE mining_info SET address = ?, mined_amount = ?, mined_fishcakepower = ?, last_mint_time = ? WHERE id = ?`,
				plan.Address,
				plan.NewMinedAmount,
				plan.NewPower,
				plan.LastMintTime,
				*plan.KeeperID,
			).Error; err != nil {
				return fmt.Errorf("update mining_info for %s: %w", plan.Address, err)
			}
		}
		return nil
	})
}

func printPlan(plans []repairPlan, apply bool) {
	mode := "DRY-RUN"
	if apply {
		mode = "APPLY"
	}

	var toWrite int
	for _, plan := range plans {
		if plan.NeedsWrite {
			toWrite++
		}
	}

	fmt.Printf("[%s] scanned=%d changed=%d\n", mode, len(plans), toWrite)
	for _, plan := range plans {
		if !plan.NeedsWrite {
			continue
		}

		action := "update"
		if plan.NeedsInsert {
			action = "insert"
		}

		fmt.Printf(
			"%s\t%s\tcurrent=%s/%s\tnew=%s/%s\tconsumed=%s\tlastMint=%s\n",
			action,
			plan.Address,
			plan.CurrentMined,
			plan.CurrentPower,
			plan.NewMinedAmount,
			plan.NewPower,
			plan.ConsumedPower,
			plan.LastMintTime,
		)
	}
}

func parseBigIntPtr(value *string) (*big.Int, error) {
	if value == nil {
		return big.NewInt(0), nil
	}
	return parseBigInt(*value)
}

func parseBigInt(value string) (*big.Int, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return big.NewInt(0), nil
	}

	result, ok := new(big.Int).SetString(trimmed, 10)
	if !ok {
		return nil, fmt.Errorf("invalid integer %q", value)
	}
	return result, nil
}
