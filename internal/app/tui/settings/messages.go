package settings

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

type ApplySettingsMsg struct {
	Settings mathmodels.TrainingSettings
}

type BackMsg struct{}
