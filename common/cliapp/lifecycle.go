package cliapp

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/FishcakeLab/fishcake-service/common/opio"
)

var interruptErr = errors.New("interrupt signal")

type Lifecycle interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Stopped() bool
}

type LifecycleAction func(ctx *cli.Context, close context.CancelCauseFunc) (Lifecycle, error)

func LifecycleCmd(fn LifecycleAction) cli.ActionFunc {
	return lifecycleCmd(fn, opio.BlockOnInterruptsContext)
}

type waitSignalFn func(ctx context.Context, signals ...os.Signal)

func lifecycleCmd(fn LifecycleAction, blockOnInterrupt waitSignalFn) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		hostCtx := ctx.Context
		appCtx, appCancel := context.WithCancelCause(hostCtx)
		ctx.Context = appCtx

		go func() {
			blockOnInterrupt(appCtx)
			appCancel(interruptErr)
		}()

		appLifecycle, err := fn(ctx, appCancel)
		if err != nil {
			return errors.Join(
				fmt.Errorf("failed to setup: %w", err),
				context.Cause(appCtx),
			)
		}

		if err := appLifecycle.Start(appCtx); err != nil {
			return errors.Join(
				fmt.Errorf("failed to start: %w", err),
				context.Cause(appCtx),
			)
		}

		<-appCtx.Done()

		stopCtx, stopCancel := context.WithCancelCause(hostCtx)
		go func() {
			blockOnInterrupt(stopCtx)
			stopCancel(interruptErr)
		}()

		stopErr := appLifecycle.Stop(stopCtx)
		stopCancel(nil)
		if stopErr != nil {
			return errors.Join(
				fmt.Errorf("failed to stop: %w", stopErr),
				context.Cause(stopCtx),
			)
		}
		return nil
	}
}
