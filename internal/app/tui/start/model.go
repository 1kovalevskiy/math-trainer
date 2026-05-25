package start

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"

type Option int

const (
	OptionPractice Option = iota
	OptionSettings
	OptionQuit
)

type Model struct {
	cursor  int
	focus   ui.LinearFocus
	options []string
}

func NewModel() Model {
	options := []string{
		"Начать тренировку",
		"Настройки сложности",
		"Выход",
	}
	return Model{
		cursor:  0,
		focus:   ui.NewLinearFocus(0, len(options)-1),
		options: options,
	}
}

func (m Model) withCursor(cursor int) Model {
	m.focus = ui.NewLinearFocus(cursor, len(m.options)-1)
	m.cursor = m.focus.Index()
	return m
}

func (m Model) moveCursor(delta int) Model {
	m.focus = m.focus.Move(delta)
	m.cursor = m.focus.Index()
	return m
}

func (m Model) selectedOption() Option {
	if m.cursor < 0 || m.cursor >= len(m.options) {
		return OptionPractice
	}

	return Option(m.cursor)
}
