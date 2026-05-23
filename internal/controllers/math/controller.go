package mathcontroller

import (
	"context"
	"fmt"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

type trainingStorage interface {
	GetState(ctx context.Context) (mathmodels.TrainingState, error)
	SaveState(ctx context.Context, state mathmodels.TrainingState) error
	ClearState(ctx context.Context) error
}

type ExerciseGenerator func(difficulty mathmodels.Difficulty) (mathmodels.Exercise, error)

type Option func(*Controller)

type Controller struct {
	storage  trainingStorage
	generate ExerciseGenerator
}

func New(storage trainingStorage, opts ...Option) *Controller {
	controller := &Controller{
		storage:  storage,
		generate: generateExercise,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(controller)
		}
	}

	return controller
}

func (c *Controller) validate() error {
	if c == nil {
		return fmt.Errorf("math controller is nil")
	}
	if c.storage == nil {
		return fmt.Errorf("math training storage is nil")
	}
	if c.generate == nil {
		return fmt.Errorf("math exercise generator is nil")
	}

	return nil
}
