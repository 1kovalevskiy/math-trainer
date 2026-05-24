package mathcontroller_test

import (
	"testing"

	mathcontroller "github.com/1kovalevskiy/math-trainer/internal/controllers/math"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	mathmemory "github.com/1kovalevskiy/math-trainer/internal/storages/memory/math"
)

func TestController_GetDefaultSettingsFromOptionAndNormalize(t *testing.T) {
	t.Parallel()

	controller := mathcontroller.New(
		mathmemory.New(),
		mathcontroller.WithDefaultSettings(mathmodels.TrainingSettings{
			AddDifficulty:      mathmodels.DifficultyDisabled,
			SubtractDifficulty: mathmodels.DifficultyDisabled,
			MultiplyDifficulty: mathmodels.DifficultyDisabled,
			DivideDifficulty:   mathmodels.DifficultyDisabled,
			ExamplesCount:      mathmodels.MaxExamplesCount + 100,
		}),
	)

	defaults := controller.GetDefaultSettings()
	if defaults.ExamplesCount != mathmodels.MaxExamplesCount {
		t.Fatalf("examples_count mismatch: got %d", defaults.ExamplesCount)
	}
	if defaults.AddDifficulty == mathmodels.DifficultyDisabled &&
		defaults.SubtractDifficulty == mathmodels.DifficultyDisabled &&
		defaults.MultiplyDifficulty == mathmodels.DifficultyDisabled &&
		defaults.DivideDifficulty == mathmodels.DifficultyDisabled {
		t.Fatal("expected at least one enabled operator")
	}
}
