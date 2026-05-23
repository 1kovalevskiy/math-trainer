package mathcontroller_test

import (
	"context"
	"errors"
	"testing"

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

	started, err := controller.StartTraining(ctx, mathmodels.TrainingSettings{
		Difficulty:    mathmodels.DifficultyMedium,
		ExamplesCount: 2,
	})
	if err != nil {
		t.Fatalf("StartTraining() unexpected error: %v", err)
	}
	assertCurrentExercise(t, started, 1, 2, mathmodels.Exercise{
		Left: 2, Right: 3, Operator: mathmodels.OperatorAdd,
	})

	afterSubmit, err := controller.SubmitAnswer(ctx, "5")
	if err != nil {
		t.Fatalf("SubmitAnswer() unexpected error: %v", err)
	}
	assertCurrentExercise(t, afterSubmit, 2, 2, mathmodels.Exercise{
		Left: 7, Right: 4, Operator: mathmodels.OperatorSubtract,
	})

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
	if got, want := finished.Summary.Total, 2; got != want {
		t.Fatalf("summary total mismatch: got %d, want %d", got, want)
	}
	if got, want := len(finished.Summary.Results), 2; got != want {
		t.Fatalf("summary results count mismatch: got %d, want %d", got, want)
	}
	if got, want := finished.Summary.Results[0].Status, mathmodels.ResultStatusCorrect; got != want {
		t.Fatalf("first result status mismatch: got %q, want %q", got, want)
	}
	if got, want := finished.Summary.Results[0].CorrectAnswer, 5; got != want {
		t.Fatalf("first result correct answer mismatch: got %d, want %d", got, want)
	}
	if got, want := finished.Summary.Results[1].Status, mathmodels.ResultStatusSkipped; got != want {
		t.Fatalf("second result status mismatch: got %q, want %q", got, want)
	}
	if finished.Summary.Results[1].UserAnswer != nil {
		t.Fatalf("second result user answer mismatch: got %v, want nil", *finished.Summary.Results[1].UserAnswer)
	}
	if got, want := finished.Summary.Results[1].CorrectAnswer, 3; got != want {
		t.Fatalf("second result correct answer mismatch: got %d, want %d", got, want)
	}
}

func TestController_SubmitAnswerRejectsInvalidInputAndKeepsState(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := mathmemory.New()
	controller := mathcontroller.New(store, mathcontroller.WithExerciseGenerator(sequenceGenerator(
		mathmodels.Exercise{Left: 8, Right: 2, Operator: mathmodels.OperatorSubtract},
	)))

	started, err := controller.StartTraining(ctx, mathmodels.TrainingSettings{
		Difficulty:    mathmodels.DifficultyEasy,
		ExamplesCount: 1,
	})
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
	if got := len(state.Results); got != 0 {
		t.Fatalf("results changed after invalid answer: got %d results", got)
	}
}

func TestController_GeneratesNextExerciseWithoutRepeatingSolvedOnes(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := mathmemory.New()
	first := mathmodels.Exercise{Left: 2, Right: 3, Operator: mathmodels.OperatorAdd}
	second := mathmodels.Exercise{Left: 9, Right: 4, Operator: mathmodels.OperatorSubtract}
	controller := mathcontroller.New(store, mathcontroller.WithExerciseGenerator(sequenceGenerator(
		first,
		first,
		second,
	)))

	started, err := controller.StartTraining(ctx, mathmodels.TrainingSettings{
		Difficulty:    mathmodels.DifficultyEasy,
		ExamplesCount: 2,
	})
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

	normalized := controller.NormalizeSettings(mathmodels.TrainingSettings{
		Difficulty:    mathmodels.Difficulty("unknown"),
		ExamplesCount: -10,
	})
	if got, want := normalized.Difficulty, mathmodels.DifficultyEasy; got != want {
		t.Fatalf("difficulty mismatch: got %q, want %q", got, want)
	}
	if got, want := normalized.ExamplesCount, mathmodels.MinExamplesCount; got != want {
		t.Fatalf("examples count mismatch: got %d, want %d", got, want)
	}

	normalized = controller.NormalizeSettings(mathmodels.TrainingSettings{
		Difficulty:    mathmodels.DifficultyHard,
		ExamplesCount: mathmodels.MaxExamplesCount + 1,
	})
	if got, want := normalized.ExamplesCount, mathmodels.MaxExamplesCount; got != want {
		t.Fatalf("examples count mismatch: got %d, want %d", got, want)
	}
}

func assertCurrentExercise(
	t *testing.T,
	snapshot mathmodels.TrainingSnapshot,
	order int,
	total int,
	exercise mathmodels.Exercise,
) {
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
	return func(mathmodels.Difficulty) (mathmodels.Exercise, error) {
		if next >= len(exercises) {
			return mathmodels.Exercise{}, errors.New("exercise generator exhausted")
		}

		exercise := exercises[next]
		next++
		return exercise, nil
	}
}
