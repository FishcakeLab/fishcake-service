package main

import (
	"context"
	"github.com/FishcakeLab/fishcake-service/config"
	"log"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/params"

	"github.com/FishcakeLab/fishcake-service/common/cliapp"
	flag2 "github.com/FishcakeLab/fishcake-service/flags"
)

func runIndexer(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Printf("run indexer start")
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.Printf("failed to load config", "err", err)
		return nil, err
	}
	log.Printf("run api start", "HttpHostHost", cfg.HttpHost, "HttpHostHost", cfg.HttpPort)
	return nil, nil
}

func runApi(ctx *cli.Context, _ context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Printf("run api start")
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.Printf("failed to load config", "err", err)
		return nil, err
	}
	log.Printf("run api start", "HttpHostHost", cfg.HttpHost, "HttpHostHost", cfg.HttpPort)
	return nil, nil
}

func runMigrations(ctx *cli.Context) error {
	log.Printf("run migration start")
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.Printf("failed to load config", "err", err)
		return err
	}
	log.Printf("run api start", "HttpHostHost", cfg.HttpHost, "HttpHostHost", cfg.HttpPort)
	return nil
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
