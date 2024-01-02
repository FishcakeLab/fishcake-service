package fishcake

import (
	"context"
	"log"
	"os"
)

var (
	GitCommit = ""
	GitDate   = ""
)

func main() {
	app := newCli(GitCommit, GitDate)
	// sub-commands set up their individual interrupt lifecycles, which can block on the given interrupt as needed.
	ctx := opio.WithInterruptBlocker(context.Background())
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Println("application failed", "err", err)
		os.Exit(1)
	}
}
