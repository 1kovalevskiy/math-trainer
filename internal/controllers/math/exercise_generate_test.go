package mathcontroller

import (
	"math/rand"
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

func TestGenerateExerciseByOperator_UsesDifficultyRules(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		operator   mathmodels.Operator
		difficulty mathmodels.Difficulty
		assert     func(t *testing.T, exercise mathmodels.Exercise)
	}{
		{
			name:       "starter addition keeps old easy range",
			operator:   mathmodels.OperatorAdd,
			difficulty: mathmodels.DifficultyStarter,
			assert: func(t *testing.T, exercise mathmodels.Exercise) {
				assertRange(t, "left", exercise.Left, 1, 10)
				assertRange(t, "right", exercise.Right, 1, 10)
			},
		},
		{
			name:       "expert addition reaches approved upper range",
			operator:   mathmodels.OperatorAdd,
			difficulty: mathmodels.DifficultyExpert,
			assert: func(t *testing.T, exercise mathmodels.Exercise) {
				assertRange(t, "left", exercise.Left, 1000, 9999)
				assertRange(t, "right", exercise.Right, 1000, 9999)
			},
		},
		{
			name:       "medium subtraction stays non-negative",
			operator:   mathmodels.OperatorSubtract,
			difficulty: mathmodels.DifficultyMedium,
			assert: func(t *testing.T, exercise mathmodels.Exercise) {
				assertRange(t, "left", exercise.Left, 50, 999)
				assertRange(t, "right", exercise.Right, 50, 999)
				if exercise.Left < exercise.Right {
					t.Fatalf("medium subtraction must be non-negative: got %d - %d", exercise.Left, exercise.Right)
				}
			},
		},
		{
			name:       "expert multiplication keeps old hard range",
			operator:   mathmodels.OperatorMultiply,
			difficulty: mathmodels.DifficultyExpert,
			assert: func(t *testing.T, exercise mathmodels.Exercise) {
				assertRange(t, "left", exercise.Left, 50, 100)
				assertRange(t, "right", exercise.Right, 50, 100)
			},
		},
		{
			name:       "expert division keeps old hard range",
			operator:   mathmodels.OperatorDivide,
			difficulty: mathmodels.DifficultyExpert,
			assert: func(t *testing.T, exercise mathmodels.Exercise) {
				if exercise.Right == 0 {
					t.Fatal("division right operand must not be zero")
				}
				if exercise.Left%exercise.Right != 0 {
					t.Fatalf("expected whole division, got %d / %d", exercise.Left, exercise.Right)
				}
				assertRange(t, "right", exercise.Right, 50, 100)
				assertRange(t, "quotient", exercise.Left/exercise.Right, 50, 100)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			random := rand.New(rand.NewSource(1))
			for i := 0; i < 100; i++ {
				exercise := generateExerciseByOperator(random, tt.operator, tt.difficulty)
				if got := exercise.Operator; got != tt.operator {
					t.Fatalf("operator mismatch: got %q, want %q", got, tt.operator)
				}
				tt.assert(t, exercise)
			}
		})
	}
}

func TestGenerateExerciseByOperator_HardAndExpertSubtractionCanBeNegative(t *testing.T) {
	t.Parallel()

	tests := []mathmodels.Difficulty{
		mathmodels.DifficultyHard,
		mathmodels.DifficultyExpert,
	}

	for _, difficulty := range tests {
		difficulty := difficulty
		t.Run(difficulty.String(), func(t *testing.T) {
			t.Parallel()

			random := rand.New(rand.NewSource(1))
			hasNegative := false
			for i := 0; i < 200; i++ {
				exercise := generateExerciseByOperator(random, mathmodels.OperatorSubtract, difficulty)
				if exercise.Left < exercise.Right {
					hasNegative = true
					break
				}
			}
			if !hasNegative {
				t.Fatalf("%s subtraction must sometimes produce negative results", difficulty)
			}
		})
	}
}

func assertRange(t *testing.T, name string, value int, min int, max int) {
	t.Helper()

	if value < min || value > max {
		t.Fatalf("%s out of range: got %d, want %d..%d", name, value, min, max)
	}
}
