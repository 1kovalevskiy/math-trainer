package app

import (
	"fmt"
	"os"

	config "github.com/1kovalevskiy/math-trainer/internal/configs"
)

func (a *App) initConfig() error {
	configPath, err := config.ResolvePath(os.Getenv("APP_CONFIG_PATH"))
	if err != nil {
		return fmt.Errorf("resolve config path: %w", err)
	}

	cfg, err := config.LoadOrCreate(configPath)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	a.configPath = configPath
	a.cfg = cfg

	return nil
}
