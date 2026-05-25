package config

import (
	"os"
	"path/filepath"
	"testing"

	configmodel "github.com/1kovalevskiy/math-trainer/internal/models/config"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestResolvePathUsesEnvOverride(t *testing.T) {
	t.Parallel()

	path, err := ResolvePath("/tmp/custom.yaml")
	if err != nil {
		t.Fatalf("ResolvePath() unexpected error: %v", err)
	}
	if path != "/tmp/custom.yaml" {
		t.Fatalf("ResolvePath() path mismatch: got %q", path)
	}
}

func TestLoadOrCreateCreatesMissingFile(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "math-trainer.yaml")

	cfg, err := LoadOrCreate(path)
	if err != nil {
		t.Fatalf("LoadOrCreate() unexpected error: %v", err)
	}

	if cfg.Training.ToMath() != mathmodels.DefaultTrainingSettings() {
		t.Fatalf("created default settings mismatch: got %+v", cfg.Training.ToMath())
	}
	if got, want := cfg.Training.AddDifficulty, mathmodels.DifficultyStarter; got != want {
		t.Fatalf("default add difficulty mismatch: got %q, want %q", got, want)
	}
	if got, want := cfg.Training.SubtractDifficulty, mathmodels.DifficultyStarter; got != want {
		t.Fatalf("default subtract difficulty mismatch: got %q, want %q", got, want)
	}
	if got, want := cfg.Training.MultiplyDifficulty, mathmodels.DifficultyDisabled; got != want {
		t.Fatalf("default multiply difficulty mismatch: got %q, want %q", got, want)
	}
	if got, want := cfg.Training.DivideDifficulty, mathmodels.DifficultyDisabled; got != want {
		t.Fatalf("default divide difficulty mismatch: got %q, want %q", got, want)
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("created file missing: %v", err)
	}
}

func TestSaveTrainingSettingsUpdatesOnlyTrainingSection(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "math-trainer.yaml")

	initial := configmodel.Config{
		App:      configmodel.App{LogLevel: "ERROR"},
		Training: configmodel.TrainingFromMath(mathmodels.DefaultTrainingSettings()),
	}
	if err := writeConfig(path, initial); err != nil {
		t.Fatalf("writeConfig() unexpected error: %v", err)
	}

	updated := mathmodels.TrainingSettings{
		AddDifficulty:      mathmodels.DifficultyMedium,
		SubtractDifficulty: mathmodels.DifficultyHard,
		MultiplyDifficulty: mathmodels.DifficultyEasy,
		DivideDifficulty:   mathmodels.DifficultyDisabled,
		ExamplesCount:      17,
	}
	if err := SaveTrainingSettings(path, updated); err != nil {
		t.Fatalf("SaveTrainingSettings() unexpected error: %v", err)
	}

	rawCfg, err := loadFileOnly(path)
	if err != nil {
		t.Fatalf("loadFileOnly() unexpected error: %v", err)
	}

	if got, want := rawCfg.App.LogLevel, "ERROR"; got != want {
		t.Fatalf("app.log_level mismatch: got %q, want %q", got, want)
	}
	if got := rawCfg.Training.ToMath(); got != updated {
		t.Fatalf("training settings mismatch: got %+v, want %+v", got, updated)
	}
}
