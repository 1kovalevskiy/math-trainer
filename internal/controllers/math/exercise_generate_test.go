package mathcontroller

import (
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestGenerateExercise_UsesOnlyEnabledOperators(t *testing.T) {
	t.Parallel()

	settings := mathmodels.TrainingSettings{
		AddDifficulty:      mathmodels.DifficultyDisabled,
		SubtractDifficulty: mathmodels.DifficultyDisabled,
		MultiplyDifficulty: mathmodels.DifficultyDisabled,
		DivideDifficulty:   mathmodels.DifficultyEasy,
		ExamplesCount:      5,
	}

	for i := 0; i < 20; i++ {
		exercise, err := generateExercise(settings)
		if err != nil {
			t.Fatalf("generateExercise() unexpected error: %v", err)
		}
		if got, want := exercise.Operator, mathmodels.OperatorDivide; got != want {
			t.Fatalf("operator mismatch: got %q, want %q", got, want)
		}
	}
}

func TestGenerateExercise_DivisionAlwaysWhole(t *testing.T) {
	t.Parallel()

	settings := mathmodels.TrainingSettings{
		AddDifficulty:      mathmodels.DifficultyDisabled,
		SubtractDifficulty: mathmodels.DifficultyDisabled,
		MultiplyDifficulty: mathmodels.DifficultyDisabled,
		DivideDifficulty:   mathmodels.DifficultyMedium,
		ExamplesCount:      5,
	}

	for i := 0; i < 100; i++ {
		exercise, err := generateExercise(settings)
		if err != nil {
			t.Fatalf("generateExercise() unexpected error: %v", err)
		}
		if exercise.Right == 0 {
			t.Fatal("division right operand must not be zero")
		}
		if exercise.Left%exercise.Right != 0 {
			t.Fatalf("expected whole division, got %d / %d", exercise.Left, exercise.Right)
		}
	}
}
