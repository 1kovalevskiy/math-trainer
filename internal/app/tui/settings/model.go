package settings

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

type rules interface {
	NormalizeSettings(settings mathmodels.TrainingSettings) mathmodels.TrainingSettings
	GetNextDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty
	GetPreviousDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty
}

type Model struct {
	settings mathmodels.TrainingSettings
	rules    rules
	focus    settingsFocus
	errText  string
}

func NewModel(settings mathmodels.TrainingSettings, rules rules) Model {
	if rules != nil {
		settings = rules.NormalizeSettings(settings)
	}
	return Model{settings: settings, rules: rules, focus: newSettingsFocus()}
}

func (m Model) WithError(text string) Model {
	m.errText = text
	return m
}
