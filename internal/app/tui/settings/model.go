package settings

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

type rules interface {
	NormalizeSettings(settings mathmodels.TrainingSettings) mathmodels.TrainingSettings
	GetNextDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty
	GetPreviousDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty
}

const (
	rowAddDifficulty = iota
	rowSubtractDifficulty
	rowMultiplyDifficulty
	rowDivideDifficulty
	rowExamplesCount
	rowApply
	rowBack
	lastRow = rowBack
)

type Model struct {
	cursor   int
	settings mathmodels.TrainingSettings
	rules    rules
	errText  string
}

func NewModel(settings mathmodels.TrainingSettings, rules rules) Model {
	if rules != nil {
		settings = rules.NormalizeSettings(settings)
	}
	return Model{cursor: rowAddDifficulty, settings: settings, rules: rules}
}

func (m Model) WithError(text string) Model {
	m.errText = text
	return m
}
