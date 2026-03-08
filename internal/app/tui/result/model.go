package result

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"

type Option int

const (
	OptionRetry Option = iota
	OptionSettings
	OptionHome
)

type Outcome struct {
	Difficulty shared.Difficulty
	Expression string
	Expected   int
	Answer     int
	Correct    bool
}

type Model struct {
	outcome Outcome
	cursor  int
	options []string
}

func NewModel() Model {
	return Model{
		cursor: 0,
		options: []string{
			"Решить еще один пример",
			"Настройки сложности",
			"В главное меню",
		},
	}
}

func (m Model) WithOutcome(outcome Outcome) Model {
	m.outcome = outcome
	m.cursor = 0
	return m
}

func (m Model) selectedOption() Option {
	if m.cursor < 0 || m.cursor >= len(m.options) {
		return OptionRetry
	}

	return Option(m.cursor)
}
