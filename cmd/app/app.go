package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1kovalevskiy/math-trainer/internal/configs"
	mathController "github.com/1kovalevskiy/math-trainer/internal/controllers/math"
	tea "github.com/charmbracelet/bubbletea"
)

type closer struct {
	name string
	fn   func(context.Context) error
}

type App struct {
	cfg            *config.Config
	logger         *slog.Logger
	mathController *mathController.Controller
	program        *tea.Program
	closers        []closer
}

func InitApp(ctx context.Context) (*App, error) {
	app := &App{}

	if err := app.initConfig(); err != nil {
		return nil, fmt.Errorf("init config: %w", err)
	}

	if err := app.initLogger(); err != nil {
		return nil, fmt.Errorf("init logger: %w", err)
	}

	if err := app.initProviders(ctx); err != nil {
		return nil, fmt.Errorf("init providers: %w", err)
	}

	if err := app.initStorages(ctx); err != nil {
		return nil, fmt.Errorf("init storages: %w", err)
	}

	if err := app.initControllers(ctx); err != nil {
		return nil, fmt.Errorf("init controllers: %w", err)
	}

	if err := app.initPrograms(); err != nil {
		return nil, fmt.Errorf("init programs: %w", err)
	}

	return app, nil
}

func (a *App) RunApp(ctx context.Context) error {
	if a.program == nil {
		return errors.New("bubble tea program is not initialized")
	}

	runCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	runErrCh := make(chan error, 1)
	go func() {
		_, err := a.program.Run()
		runErrCh <- err
	}()

	var runErr error
	var closeErr error
	isClosed := false
	select {
	case <-runCtx.Done():
		closeErr = a.closeAll(shutdownCtx)
		isClosed = true
		runErr = <-runErrCh
	case runErr = <-runErrCh:
	}

	if !isClosed {
		closeErr = a.closeAll(shutdownCtx)
	}

	return errors.Join(runErr, closeErr)
}

func (a *App) addCloser(name string, fn func(context.Context) error) {
	a.closers = append(a.closers, closer{name: name, fn: fn})
}

func (a *App) closeAll(ctx context.Context) error {
	var joined error

	for i := len(a.closers) - 1; i >= 0; i-- {
		current := a.closers[i]
		if current.fn == nil {
			continue
		}

		a.logger.Info("closing resource", "name", current.name)
		if err := current.fn(ctx); err != nil {
			a.logger.Error("failed to close resource", "name", current.name, "error", err)
			joined = errors.Join(joined, fmt.Errorf("close %s: %w", current.name, err))
			continue
		}

		a.logger.Info("resource closed", "name", current.name)
	}

	return joined
}
