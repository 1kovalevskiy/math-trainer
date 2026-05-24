package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	configmodel "github.com/1kovalevskiy/math-trainer/internal/models/config"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

const defaultConfigFileName = "math-trainer.yaml"

type Config = configmodel.Config

type App = configmodel.App

func init() {
	if err := godotenv.Load(); err != nil {
		slog.Debug("no .env file found")
	}
}

func ResolvePath(configPathEnv string) (string, error) {
	if configPathEnv != "" {
		return configPathEnv, nil
	}

	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("resolve executable path: %w", err)
	}

	return filepath.Join(filepath.Dir(exePath), defaultConfigFileName), nil
}

func LoadOrCreate(path string) (*Config, error) {
	cfg, err := load(path)
	if err == nil {
		return cfg, nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	if err := writeConfig(path, configmodel.Default()); err != nil {
		return nil, fmt.Errorf("create config file: %w", err)
	}

	cfg, err = load(path)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func SaveTrainingSettings(path string, settings mathmodels.TrainingSettings) error {
	cfg, err := loadFileOnly(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		cfg = configmodel.Default()
	}

	cfg.Training = configmodel.TrainingFromMath(settings)
	if err := writeConfig(path, cfg); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}

func load(path string) (*Config, error) {
	cfg := &configmodel.Config{}

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("config env error: %w", err)
	}

	return cfg, nil
}

func loadFileOnly(path string) (configmodel.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return configmodel.Config{}, fmt.Errorf("read config file: %w", err)
	}

	cfg := configmodel.Default()
	if len(data) == 0 {
		return cfg, nil
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return configmodel.Config{}, fmt.Errorf("parse config yaml: %w", err)
	}

	return cfg, nil
}

func writeConfig(path string, cfg configmodel.Config) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create config directory: %w", err)
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshal config yaml: %w", err)
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write config file: %w", err)
	}

	return nil
}
