package main

import (
	"context"
	"log"
	"os"

	"github.com/FishcakeLab/fishcake-service/common/opio"
)

var (
	GitCommit = ""
	GitDate   = ""
)

func main() {
	app := newCli(GitCommit, GitDate)
	ctx := opio.WithInterruptBlocker(context.Background())
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Println("application failed", "err", err)
		os.Exit(1)
	}
}
