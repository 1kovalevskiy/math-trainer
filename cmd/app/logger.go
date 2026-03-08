package app

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func (a *App) initLogger() error {
	if a.cfg == nil {
		return fmt.Errorf("config is not initialized")
	}

	level := new(slog.LevelVar)
	switch strings.ToUpper(a.cfg.App.LogLevel) {
	case "DEBUG":
		level.Set(slog.LevelDebug)
	case "INFO":
		level.Set(slog.LevelInfo)
	case "WARN", "WARNING":
		level.Set(slog.LevelWarn)
	case "ERROR":
		level.Set(slog.LevelError)
	default:
		return fmt.Errorf("unsupported log level: %s", a.cfg.App.LogLevel)
	}

	a.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(a.logger)

	return nil
}
