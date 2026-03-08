package start

type Option int

const (
	OptionPractice Option = iota
	OptionSettings
	OptionQuit
)

type Model struct {
	cursor  int
	options []string
}

func NewModel() Model {
	return Model{
		cursor: 0,
		options: []string{
			"Начать тренировку",
			"Настройки сложности",
			"Выход",
		},
	}
}

func (m Model) selectedOption() Option {
	if m.cursor < 0 || m.cursor >= len(m.options) {
		return OptionPractice
	}

	return Option(m.cursor)
}
