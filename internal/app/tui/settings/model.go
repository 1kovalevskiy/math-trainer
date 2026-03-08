package settings

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"

const (
	rowDifficulty = iota
	rowExamplesCount
	rowApply
	rowBack
	lastRow = rowBack
)

type Model struct {
	cursor   int
	settings shared.TrainingSettings
}

func NewModel(settings shared.TrainingSettings) Model {
	settings.ExamplesCount = shared.NormalizeExamplesCount(settings.ExamplesCount)
	return Model{
		cursor:   rowDifficulty,
		settings: settings,
	}
}
