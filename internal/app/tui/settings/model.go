package settings

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"

type Model struct {
	cursor       int
	current      shared.Difficulty
	difficulties []shared.Difficulty
}

func NewModel(current shared.Difficulty) Model {
	options := shared.AllDifficulties()
	cursor := 0
	for i, difficulty := range options {
		if difficulty == current {
			cursor = i
			break
		}
	}

	return Model{
		cursor:       cursor,
		current:      current,
		difficulties: options,
	}
}

func (m Model) selectedDifficulty() shared.Difficulty {
	if m.cursor < 0 || m.cursor >= len(m.difficulties) {
		return m.current
	}

	return m.difficulties[m.cursor]
}
