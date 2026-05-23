package mathcontroller

import (
	"context"
	"fmt"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func (c *Controller) completeCurrent(
	ctx context.Context,
	answer *int,
	defaultStatus mathmodels.ResultStatus,
) (mathmodels.TrainingSnapshot, error) {
	if err := c.validate(); err != nil {
		return mathmodels.TrainingSnapshot{}, err
	}

	state, err := c.storage.GetState(ctx)
	if err != nil {
		return mathmodels.TrainingSnapshot{}, fmt.Errorf("get training state: %w", err)
	}
	if state.Finished {
		return mathmodels.TrainingSnapshot{}, mathmodels.ErrTrainingFinished
	}

	correctAnswer := correctAnswer(state.CurrentExercise)
	status := defaultStatus
	if answer != nil && *answer == correctAnswer {
		status = mathmodels.ResultStatusCorrect
	}

	state.Results = append(state.Results, mathmodels.ExampleResult{
		Order:         state.CurrentOrder,
		Exercise:      state.CurrentExercise,
		CorrectAnswer: correctAnswer,
		UserAnswer:    answer,
		Status:        status,
	})

	if len(state.Results) >= state.Settings.ExamplesCount {
		state.Finished = true
		state.CurrentOrder = 0
		state.CurrentExercise = mathmodels.Exercise{}
	} else {
		state.CurrentOrder++
		nextExercise, err := c.generateUnused(state.Settings.Difficulty, usedExercises(state))
		if err != nil {
			return mathmodels.TrainingSnapshot{}, fmt.Errorf("generate next exercise: %w", err)
		}
		state.CurrentExercise = nextExercise
	}

	if err := c.storage.SaveState(ctx, state); err != nil {
		return mathmodels.TrainingSnapshot{}, fmt.Errorf("save training state: %w", err)
	}

	return makeSnapshot(state), nil
}
