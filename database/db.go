package database

import (
	"context"
	"fmt"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database/account_nft_info"
	"github.com/FishcakeLab/fishcake-service/database/activity"
	"github.com/FishcakeLab/fishcake-service/database/block_listener"
	"github.com/FishcakeLab/fishcake-service/database/common"
	"github.com/FishcakeLab/fishcake-service/database/drop"
	"github.com/FishcakeLab/fishcake-service/database/event"
	"github.com/FishcakeLab/fishcake-service/database/token_nft"
	"github.com/FishcakeLab/fishcake-service/database/wallet"
	"github.com/FishcakeLab/fishcake-service/synchronizer/retry"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type DB struct {
	gorm              *gorm.DB
	Blocks            common.BlocksDB
	ContractEvent     event.ContractEventDB
	ActivityInfoDB    activity.ActivityInfoDB
	ActivityInfoExtDB activity.ActivityInfoExtDB
	TokenNftDB        token_nft.TokenNftDB
	DropInfoDB        drop.DropInfoDB
	BlockListener     block_listener.BlockListenerDB
	AccountNftInfoDB  account_nft_info.AccountNftInfoDB
	WalletInfoDB      wallet.WalletInfoDB
}

func NewDB(cfg *config.Config) (*DB, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}
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
	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		CreateBatchSize:        3_000,
	}
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	gorm, err := retry.Do[*gorm.DB](context.Background(), 10, retryStrategy, func() (*gorm.DB, error) {
		gorm, err := gorm.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return gorm, nil
	})
	if err != nil {
		return nil, err
	}
	db := &DB{
		gorm:              gorm,
		Blocks:            common.NewBlocksDB(gorm),
		ContractEvent:     event.NewContractEventsDB(gorm),
		ActivityInfoDB:    activity.NewActivityDB(gorm),
		ActivityInfoExtDB: activity.NewActivityInfoExtDB(gorm),
		TokenNftDB:        token_nft.NewTokenNftDB(gorm),
		AccountNftInfoDB:  account_nft_info.NewAccountNftInfoDB(gorm),
		DropInfoDB:        drop.NewDropInfoDB(gorm),
		WalletInfoDB:      wallet.NewWalletInfoDB(gorm),
		BlockListener:     block_listener.NewBlockListenerDB(gorm),
	}
	return db, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(tx *gorm.DB) error {
		txDB := &DB{
			gorm:              tx,
			Blocks:            common.NewBlocksDB(tx),
			ContractEvent:     event.NewContractEventsDB(tx),
			ActivityInfoDB:    activity.NewActivityDB(tx),
			ActivityInfoExtDB: activity.NewActivityInfoExtDB(tx),
			TokenNftDB:        token_nft.NewTokenNftDB(tx),
			DropInfoDB:        drop.NewDropInfoDB(tx),
			BlockListener:     block_listener.NewBlockListenerDB(tx),
			AccountNftInfoDB:  account_nft_info.NewAccountNftInfoDB(tx),
			WalletInfoDB:      wallet.NewWalletInfoDB(tx),
		}
		return fn(txDB)
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to process migration file: %s", path))
		}
		if info.IsDir() {
			return nil
		}
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("Error reading SQL file: %s", path))
		}
		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("Error executing SQL script: %s", path))
		}
		return nil
	})
	return err
}
