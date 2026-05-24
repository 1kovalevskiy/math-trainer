package mathcontroller

import (
	"context"
	"errors"
	"fmt"
	"strings"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

type trainingStorage interface {
	GetState(ctx context.Context) (mathmodels.TrainingState, error)
	SaveState(ctx context.Context, state mathmodels.TrainingState) error
	ClearState(ctx context.Context) error
}

type ExerciseGenerator func(settings mathmodels.TrainingSettings) (mathmodels.Exercise, error)

type Option func(*Controller)

type Controller struct {
	storage         trainingStorage
	generate        ExerciseGenerator
	defaultSettings mathmodels.TrainingSettings
}

func New(storage trainingStorage, opts ...Option) *Controller {
	controller := &Controller{
		storage:         storage,
		generate:        generateExercise,
		defaultSettings: mathmodels.DefaultTrainingSettings(),
	}

	for _, opt := range opts {
		if opt != nil {
			opt(controller)
		}
	}

	controller.defaultSettings = controller.NormalizeSettings(controller.defaultSettings)

	return controller
}

func (c *Controller) validate() error {
	if c == nil {
		return errors.New("controller is nil")
	}
	if c.storage == nil {
		return errors.New("training storage is nil")
	}
	if c.generate == nil {
		return errors.New("exercise generator is nil")
	}

	return nil
}

func enabledOperatorDifficulties(settings mathmodels.TrainingSettings) []operatorDifficulty {
	pairs := []operatorDifficulty{
		{operator: mathmodels.OperatorAdd, difficulty: settings.AddDifficulty},
		{operator: mathmodels.OperatorSubtract, difficulty: settings.SubtractDifficulty},
		{operator: mathmodels.OperatorMultiply, difficulty: settings.MultiplyDifficulty},
		{operator: mathmodels.OperatorDivide, difficulty: settings.DivideDifficulty},
	}

	enabled := make([]operatorDifficulty, 0, len(pairs))
	for _, pair := range pairs {
		if pair.difficulty != mathmodels.DifficultyDisabled {
			enabled = append(enabled, pair)
		}
	}

	return enabled
}

func formatSettingsDifficulties(settings mathmodels.TrainingSettings) string {
	parts := []string{
		fmt.Sprintf("+: %s", settings.AddDifficulty),
		fmt.Sprintf("-: %s", settings.SubtractDifficulty),
		fmt.Sprintf("*: %s", settings.MultiplyDifficulty),
		fmt.Sprintf("/: %s", settings.DivideDifficulty),
	}

	return strings.Join(parts, ", ")
}

type operatorDifficulty struct {
	operator   mathmodels.Operator
	difficulty mathmodels.Difficulty
}
