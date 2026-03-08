package app

import (
	"fmt"
	"os"

	"github.com/1kovalevskiy/math-trainer/internal/configs"
)

func (a *App) initConfig() error {
	configPath := os.Getenv("APP_CONFIG_PATH")
	if configPath != "" {
		cfg, err := config.NewConfig(configPath)
		if err != nil {
			return fmt.Errorf("load config from APP_CONFIG_PATH: %w", err)
		}
		a.cfg = cfg
		return nil
	}

	logLevel := os.Getenv("APP_LOG_LEVEL")
	if logLevel == "" {
		logLevel = "INFO"
	}

	a.cfg = &config.Config{
		App: config.App{LogLevel: logLevel},
	}

	return nil
}
