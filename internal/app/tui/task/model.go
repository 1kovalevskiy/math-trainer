package task

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

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
	buttonFocus  ui.LinearFocus
}

func NewModel(current *mathmodels.CurrentExercise, settings mathmodels.TrainingSettings) Model {
	model := Model{
		settings:     settings,
		buttonCursor: buttonSubmit,
		buttonFocus:  ui.NewLinearFocus(buttonSubmit, lastButton),
	}
	if current != nil {
		model.current = *current
	}
	return model
}

func (m Model) withButtonCursor(cursor int) Model {
	m.buttonFocus = ui.NewLinearFocus(cursor, lastButton)
	m.buttonCursor = m.buttonFocus.Index()
	return m
}

func (m Model) moveButtonCursor(delta int) Model {
	m.buttonFocus = m.buttonFocus.Move(delta)
	m.buttonCursor = m.buttonFocus.Index()
	return m
}

func (m Model) WithError(errText string) Model {
	m.errText = errText
	return m
}

func (m Model) canSubmit() bool {
	return m.input != ""
}
