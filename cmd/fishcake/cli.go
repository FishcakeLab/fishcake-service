package main

import (
	"context"
	"github.com/urfave/cli/v2"

	fishcake_service "github.com/FishcakeLab/fishcake-service"
	"github.com/FishcakeLab/fishcake-service/common/cliapp"
	"github.com/FishcakeLab/fishcake-service/common/opio"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	flag2 "github.com/FishcakeLab/fishcake-service/flags"
	"github.com/ethereum/go-ethereum/log"
)

func runIndexer(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("run indexer start")
	cfg, err := config.New("./config.yaml")
	if err != nil {
		log.Info("failed to load config", "err", err)
		return nil, err
	}
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return nil, err
	}
	return fishcake_service.NewIndex(ctx, cfg, db, shutdown), nil
}

func runApi(ctx *cli.Context, _ context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("run api start")
	cfg, err := config.New("./config.yaml")

	if err != nil {
		log.Info("Failed to load config", "err", err)
		return nil, err
	}
	log.Info("Run api start", "HttpHostHost", cfg.HttpHost, "HttpHostHost", cfg.HttpPort)
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return nil, err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	return fishcake_service.NewFishCake(cfg, db), nil
}

func runMigrations(ctx *cli.Context) error {
	log.Info("Running migrations...")
	cfg, err := config.New("./config.yaml")
	if err != nil {
		log.Info("Failed to load config", "err", err)
		return err
	}
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	return db.ExecuteSQLMigration(cfg.Migrations)
}

func newCli() *cli.App {
	flags := flag2.Flags
	return &cli.App{
		Version:              "0.0.1",
		Description:          "An indexer of all optimism events with a serving api layer",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "api",
				Flags:       flags,
				Description: "Runs the api service",
				Action:      cliapp.LifecycleCmd(runApi),
			},
			{
				Name:        "index",
				Flags:       flags,
				Description: "Runs the indexing service",
				Action:      cliapp.LifecycleCmd(runIndexer),
			},
			{
				Name:        "migrate",
				Flags:       flags,
				Description: "Runs the database migrations",
				Action:      runMigrations,
			},
			{
				Name:        "version",
				Description: "print version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
