package mathmemory_test

import (
	"context"
	"errors"
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	mathmemory "github.com/1kovalevskiy/math-trainer/internal/storages/memory/math"
)

func TestStorage_GetStateEmptyReturnsNoActiveTraining(t *testing.T) {
	t.Parallel()

	_, err := mathmemory.New().GetState(context.Background())
	if !errors.Is(err, mathmodels.ErrNoActiveTraining) {
		t.Fatalf("GetState() error mismatch: got %v, want %v", err, mathmodels.ErrNoActiveTraining)
	}
}

func TestStorage_SaveGetCopiesState(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := mathmemory.New()
	answer := 5
	state := mathmodels.TrainingState{
		Settings: mathmodels.TrainingSettings{
			AddDifficulty:      mathmodels.DifficultyMedium,
			SubtractDifficulty: mathmodels.DifficultyEasy,
			MultiplyDifficulty: mathmodels.DifficultyDisabled,
			DivideDifficulty:   mathmodels.DifficultyDisabled,
			ExamplesCount:      1,
		},
		CurrentOrder: 1,
		CurrentExercise: mathmodels.Exercise{
			Left:     2,
			Right:    3,
			Operator: mathmodels.OperatorAdd,
		},
		Results: []mathmodels.ExampleResult{{
			Order:         1,
			CorrectAnswer: 5,
			UserAnswer:    &answer,
			Status:        mathmodels.ResultStatusCorrect,
		}},
	}

	if err := store.SaveState(ctx, state); err != nil {
		t.Fatalf("SaveState() unexpected error: %v", err)
	}

	answer = 99
	state.Results[0].Status = mathmodels.ResultStatusIncorrect

	firstRead, err := store.GetState(ctx)
	if err != nil {
		t.Fatalf("GetState() unexpected error: %v", err)
	}
	if got, want := *firstRead.Results[0].UserAnswer, 5; got != want {
		t.Fatalf("stored user answer mismatch: got %d, want %d", got, want)
	}
	if got, want := firstRead.Results[0].Status, mathmodels.ResultStatusCorrect; got != want {
		t.Fatalf("stored status mismatch: got %q, want %q", got, want)
	}

	*firstRead.Results[0].UserAnswer = 42
	firstRead.Results[0].Status = mathmodels.ResultStatusSkipped

	secondRead, err := store.GetState(ctx)
	if err != nil {
		t.Fatalf("GetState() unexpected error: %v", err)
	}
	if got, want := *secondRead.Results[0].UserAnswer, 5; got != want {
		t.Fatalf("stored user answer changed through read alias: got %d, want %d", got, want)
	}
	if got, want := secondRead.Results[0].Status, mathmodels.ResultStatusCorrect; got != want {
		t.Fatalf("stored status changed through read alias: got %q, want %q", got, want)
	}
}

func TestStorage_ClearStateRemovesTraining(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := mathmemory.New()
	if err := store.SaveState(ctx, mathmodels.TrainingState{}); err != nil {
		t.Fatalf("SaveState() unexpected error: %v", err)
	}
	if err := store.ClearState(ctx); err != nil {
		t.Fatalf("ClearState() unexpected error: %v", err)
	}

	_, err := store.GetState(ctx)
	if !errors.Is(err, mathmodels.ErrNoActiveTraining) {
		t.Fatalf("GetState() error mismatch: got %v, want %v", err, mathmodels.ErrNoActiveTraining)
	}
}
