package task

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

const (
	buttonSubmit = iota
	buttonSkip
	buttonBack
	lastButton = buttonBack
)

type Model struct {
	settings     mathmodels.TrainingSettings
	current      mathmodels.CurrentExercise
	input        string
	errText      string
	buttonCursor int
}

func NewModel(current *mathmodels.CurrentExercise, settings mathmodels.TrainingSettings) Model {
	model := Model{settings: settings, buttonCursor: buttonSubmit}
	if current != nil {
		model.current = *current
	}
	return model
}

func (m Model) WithError(errText string) Model {
	m.errText = errText
	return m
}

func (m Model) canSubmit() bool {
	return m.input != ""
}
