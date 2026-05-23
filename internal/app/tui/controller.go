package tui

import (
	"context"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

type mathController interface {
	GetDefaultSettings() mathmodels.TrainingSettings
	NormalizeSettings(settings mathmodels.TrainingSettings) mathmodels.TrainingSettings
	GetNextDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty
	GetPreviousDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty
	StartTraining(ctx context.Context, settings mathmodels.TrainingSettings) (mathmodels.TrainingSnapshot, error)
	SubmitAnswer(ctx context.Context, rawAnswer string) (mathmodels.TrainingSnapshot, error)
	SkipCurrent(ctx context.Context) (mathmodels.TrainingSnapshot, error)
	CancelTraining(ctx context.Context) error
}
