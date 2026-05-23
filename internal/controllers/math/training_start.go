package mathcontroller

import (
	"context"
	"fmt"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func (c *Controller) StartTraining(
	ctx context.Context,
	settings mathmodels.TrainingSettings,
) (mathmodels.TrainingSnapshot, error) {
	if err := c.validate(); err != nil {
		return mathmodels.TrainingSnapshot{}, err
	}

	settings = c.NormalizeSettings(settings)
	exercise, err := c.generateUnused(settings.Difficulty, nil)
	if err != nil {
		return mathmodels.TrainingSnapshot{}, fmt.Errorf("generate first exercise: %w", err)
	}

	state := mathmodels.TrainingState{
		Settings:        settings,
		CurrentOrder:    1,
		CurrentExercise: exercise,
		Results:         make([]mathmodels.ExampleResult, 0, settings.ExamplesCount),
	}
	if err := c.storage.SaveState(ctx, state); err != nil {
		return mathmodels.TrainingSnapshot{}, fmt.Errorf("save training state: %w", err)
	}

	return makeSnapshot(state), nil
}
