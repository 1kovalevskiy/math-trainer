package settings

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"

type ApplySettingsMsg struct {
	Settings shared.TrainingSettings
}

type BackMsg struct{}
