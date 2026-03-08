package config

import (
	"fmt"
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	App App `json:"app" yaml:"app"`
}

type App struct {
	LogLevel string `json:"log_level" yaml:"log_level" env:"APP_LOG_LEVEL" env-default:"INFO"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		slog.Debug("no .env file found")
	}
}

func NewConfig(path string) (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("config env error: %w", err)
	}

	return cfg, nil
}
