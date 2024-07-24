package main

import (
	"context"
	"github.com/urfave/cli/v2"
	"log"

	"github.com/ethereum/go-ethereum/params"

	fishcake_service "github.com/FishcakeLab/fishcake-service"
	"github.com/FishcakeLab/fishcake-service/common/cliapp"
	"github.com/FishcakeLab/fishcake-service/common/opio"
	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	flag2 "github.com/FishcakeLab/fishcake-service/flags"
)

func runIndexer(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Printf("run indexer start")
	cfg, err := config.New("./config.yaml")
	if err != nil {
		log.Printf("failed to load config", "err", err)
		return nil, err
	}
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database", "err", err)
		return nil, err
	}
	return fishcake_service.NewIndex(ctx, cfg, db, shutdown), nil
}

func runApi(ctx *cli.Context, _ context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Printf("run api start")
	cfg, err := config.New("./config.yaml")

	if err != nil {
		log.Printf("Failed to load config", "err", err)
		return nil, err
	}
	log.Printf("Run api start", "HttpHostHost", cfg.HttpHost, "HttpHostHost", cfg.HttpPort)
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database", "err", err)
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
	log.Println("Running migrations...")
	cfg, err := config.New("./config.yaml")
	if err != nil {
		log.Printf("Failed to load config", "err", err)
		return err
	}
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database", "err", err)
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

func newCli(GitCommit string, GitDate string) *cli.App {
	flags := flag2.Flags
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitDate),
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
