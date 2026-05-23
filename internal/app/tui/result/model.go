package result

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

type Option int

const (
	OptionRetry Option = iota
	OptionSettings
	OptionHome
)

type Model struct {
	summary mathmodels.TrainingSummary
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

func (m Model) WithSummary(summary *mathmodels.TrainingSummary) Model {
	if summary != nil {
		m.summary = *summary
	}
	m.cursor = 0
	return m
}

func (m Model) selectedOption() Option {
	if m.cursor < 0 || m.cursor >= len(m.options) {
		return OptionRetry
	}

	return Option(m.cursor)
}
