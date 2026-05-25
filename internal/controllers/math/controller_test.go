package mathcontroller_test

import (
	"context"
	"errors"
	"testing"
	"time"

	mathcontroller "github.com/1kovalevskiy/math-trainer/internal/controllers/math"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	mathmemory "github.com/1kovalevskiy/math-trainer/internal/storages/memory/math"
)

func TestController_StartSubmitSkipFinishUsesStorage(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := mathmemory.New()
	controller := mathcontroller.New(store, mathcontroller.WithExerciseGenerator(sequenceGenerator(
		mathmodels.Exercise{Left: 2, Right: 3, Operator: mathmodels.OperatorAdd},
		mathmodels.Exercise{Left: 7, Right: 4, Operator: mathmodels.OperatorSubtract},
	)))

	started, err := controller.StartTraining(ctx, mathmodels.TrainingSettings{AddDifficulty: mathmodels.DifficultyMedium, SubtractDifficulty: mathmodels.DifficultyEasy, MultiplyDifficulty: mathmodels.DifficultyDisabled, DivideDifficulty: mathmodels.DifficultyDisabled, ExamplesCount: 2})
	if err != nil {
		t.Fatalf("StartTraining() unexpected error: %v", err)
	}
	assertCurrentExercise(t, started, 1, 2, mathmodels.Exercise{Left: 2, Right: 3, Operator: mathmodels.OperatorAdd})

	afterSubmit, err := controller.SubmitAnswer(ctx, "5")
	if err != nil {
		t.Fatalf("SubmitAnswer() unexpected error: %v", err)
	}
	assertCurrentExercise(t, afterSubmit, 2, 2, mathmodels.Exercise{Left: 7, Right: 4, Operator: mathmodels.OperatorSubtract})

	finished, err := controller.SkipCurrent(ctx)
	if err != nil {
		t.Fatalf("SkipCurrent() unexpected error: %v", err)
	}
	if finished.Phase != mathmodels.TrainingPhaseFinished {
		t.Fatalf("finished snapshot phase mismatch: got %q", finished.Phase)
	}
	if finished.Summary == nil {
		t.Fatal("finished snapshot expected summary")
	}
	if got, want := finished.Summary.Correct, 1; got != want {
		t.Fatalf("summary correct mismatch: got %d, want %d", got, want)
	}
}

func TestController_SubmitAnswerRejectsInvalidInputAndKeepsState(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := mathmemory.New()
	controller := mathcontroller.New(store, mathcontroller.WithExerciseGenerator(sequenceGenerator(
		mathmodels.Exercise{Left: 8, Right: 2, Operator: mathmodels.OperatorSubtract},
	)))

	started, err := controller.StartTraining(ctx, mathmodels.TrainingSettings{AddDifficulty: mathmodels.DifficultyDisabled, SubtractDifficulty: mathmodels.DifficultyEasy, MultiplyDifficulty: mathmodels.DifficultyDisabled, DivideDifficulty: mathmodels.DifficultyDisabled, ExamplesCount: 1})
	if err != nil {
		t.Fatalf("StartTraining() unexpected error: %v", err)
	}

	_, err = controller.SubmitAnswer(ctx, "not-a-number")
	if !errors.Is(err, mathmodels.ErrInvalidAnswer) {
		t.Fatalf("SubmitAnswer() error mismatch: got %v, want %v", err, mathmodels.ErrInvalidAnswer)
	}

	state, err := store.GetState(ctx)
	if err != nil {
		t.Fatalf("GetState() unexpected error: %v", err)
	}
	if got, want := state.CurrentExercise, started.Current.Exercise; got != want {
		t.Fatalf("current exercise changed after invalid answer: got %+v, want %+v", got, want)
	}
}

func TestController_GeneratesNextExerciseWithoutRepeatingSolvedOnes(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := mathmemory.New()
	first := mathmodels.Exercise{Left: 2, Right: 3, Operator: mathmodels.OperatorAdd}
	second := mathmodels.Exercise{Left: 9, Right: 4, Operator: mathmodels.OperatorSubtract}
	controller := mathcontroller.New(store, mathcontroller.WithExerciseGenerator(sequenceGenerator(first, first, second)))

	started, err := controller.StartTraining(ctx, mathmodels.TrainingSettings{AddDifficulty: mathmodels.DifficultyEasy, SubtractDifficulty: mathmodels.DifficultyEasy, MultiplyDifficulty: mathmodels.DifficultyDisabled, DivideDifficulty: mathmodels.DifficultyDisabled, ExamplesCount: 2})
	if err != nil {
		t.Fatalf("StartTraining() unexpected error: %v", err)
	}
	assertCurrentExercise(t, started, 1, 2, first)

	afterSubmit, err := controller.SubmitAnswer(ctx, "5")
	if err != nil {
		t.Fatalf("SubmitAnswer() unexpected error: %v", err)
	}
	assertCurrentExercise(t, afterSubmit, 2, 2, second)
}

func TestController_NormalizeSettings(t *testing.T) {
	t.Parallel()

	controller := mathcontroller.New(mathmemory.New())

	normalized := controller.NormalizeSettings(mathmodels.TrainingSettings{AddDifficulty: mathmodels.Difficulty("unknown"), SubtractDifficulty: mathmodels.Difficulty("unknown"), MultiplyDifficulty: mathmodels.Difficulty("unknown"), DivideDifficulty: mathmodels.Difficulty("unknown"), ExamplesCount: -10})
	if got, want := normalized.AddDifficulty, mathmodels.DifficultyEasy; got != want {
		t.Fatalf("add difficulty mismatch: got %q, want %q", got, want)
	}
	if got, want := normalized.SubtractDifficulty, mathmodels.DifficultyEasy; got != want {
		t.Fatalf("subtract difficulty mismatch: got %q, want %q", got, want)
	}
	if got, want := normalized.MultiplyDifficulty, mathmodels.DifficultyDisabled; got != want {
		t.Fatalf("multiply difficulty mismatch: got %q, want %q", got, want)
	}
	if got, want := normalized.DivideDifficulty, mathmodels.DifficultyDisabled; got != want {
		t.Fatalf("divide difficulty mismatch: got %q, want %q", got, want)
	}

	normalized = controller.NormalizeSettings(mathmodels.TrainingSettings{AddDifficulty: mathmodels.DifficultyDisabled, SubtractDifficulty: mathmodels.DifficultyDisabled, MultiplyDifficulty: mathmodels.DifficultyDisabled, DivideDifficulty: mathmodels.DifficultyDisabled, ExamplesCount: mathmodels.MaxExamplesCount + 1})
	if got, want := normalized.ExamplesCount, mathmodels.MaxExamplesCount; got != want {
		t.Fatalf("examples count mismatch: got %d, want %d", got, want)
	}
	if normalized.AddDifficulty == mathmodels.DifficultyDisabled && normalized.SubtractDifficulty == mathmodels.DifficultyDisabled && normalized.MultiplyDifficulty == mathmodels.DifficultyDisabled && normalized.DivideDifficulty == mathmodels.DifficultyDisabled {
		t.Fatal("expected at least one enabled operator")
	}
}

func TestController_DivisionAnswer(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := mathmemory.New()
	controller := mathcontroller.New(store, mathcontroller.WithExerciseGenerator(sequenceGenerator(
		mathmodels.Exercise{Left: 12, Right: 3, Operator: mathmodels.OperatorDivide},
	)))

	_, err := controller.StartTraining(ctx, mathmodels.TrainingSettings{AddDifficulty: mathmodels.DifficultyDisabled, SubtractDifficulty: mathmodels.DifficultyDisabled, MultiplyDifficulty: mathmodels.DifficultyDisabled, DivideDifficulty: mathmodels.DifficultyEasy, ExamplesCount: 1})
	if err != nil {
		t.Fatalf("StartTraining() unexpected error: %v", err)
	}

	finished, err := controller.SubmitAnswer(ctx, "4")
	if err != nil {
		t.Fatalf("SubmitAnswer() unexpected error: %v", err)
	}
	if got, want := finished.Summary.Correct, 1; got != want {
		t.Fatalf("summary correct mismatch: got %d, want %d", got, want)
	}
}

func TestController_FinishedSummaryIncludesElapsedTrainingTime(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := mathmemory.New()
	startedAt := time.Date(2026, 5, 25, 10, 0, 0, 0, time.UTC)
	finishedAt := startedAt.Add(2*time.Minute + 15*time.Second)
	controller := mathcontroller.New(
		store,
		mathcontroller.WithExerciseGenerator(sequenceGenerator(
			mathmodels.Exercise{Left: 2, Right: 3, Operator: mathmodels.OperatorAdd},
		)),
		mathcontroller.WithClock(sequenceClock(startedAt, finishedAt)),
	)

	_, err := controller.StartTraining(ctx, mathmodels.TrainingSettings{AddDifficulty: mathmodels.DifficultyEasy, SubtractDifficulty: mathmodels.DifficultyDisabled, MultiplyDifficulty: mathmodels.DifficultyDisabled, DivideDifficulty: mathmodels.DifficultyDisabled, ExamplesCount: 1})
	if err != nil {
		t.Fatalf("StartTraining() unexpected error: %v", err)
	}

	finished, err := controller.SubmitAnswer(ctx, "5")
	if err != nil {
		t.Fatalf("SubmitAnswer() unexpected error: %v", err)
	}
	if finished.Summary == nil {
		t.Fatal("finished snapshot expected summary")
	}
	if got, want := finished.Summary.Elapsed, 2*time.Minute+15*time.Second; got != want {
		t.Fatalf("summary elapsed mismatch: got %v, want %v", got, want)
	}
}

func assertCurrentExercise(t *testing.T, snapshot mathmodels.TrainingSnapshot, order int, total int, exercise mathmodels.Exercise) {
	t.Helper()

	if snapshot.Phase != mathmodels.TrainingPhaseInProgress {
		t.Fatalf("snapshot phase mismatch: got %q, want %q", snapshot.Phase, mathmodels.TrainingPhaseInProgress)
	}
	if snapshot.Current == nil {
		t.Fatal("snapshot expected current exercise")
	}
	if got, want := snapshot.Current.Order, order; got != want {
		t.Fatalf("current order mismatch: got %d, want %d", got, want)
	}
	if got, want := snapshot.Current.Total, total; got != want {
		t.Fatalf("current total mismatch: got %d, want %d", got, want)
	}
	if got, want := snapshot.Current.Exercise, exercise; got != want {
		t.Fatalf("current exercise mismatch: got %+v, want %+v", got, want)
	}
}

func sequenceGenerator(exercises ...mathmodels.Exercise) mathcontroller.ExerciseGenerator {
	next := 0
	return func(mathmodels.TrainingSettings) (mathmodels.Exercise, error) {
		if next >= len(exercises) {
			return mathmodels.Exercise{}, errors.New("exercise generator exhausted")
		}

		exercise := exercises[next]
		next++
		return exercise, nil
	}
}

func sequenceClock(times ...time.Time) mathcontroller.Clock {
	next := 0
	return func() time.Time {
		if next >= len(times) {
			return times[len(times)-1]
		}

		current := times[next]
		next++
		return current
	}
}
