package settings

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch key.String() {
	case "up", "k":
		if m.cursor > rowDifficulty {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < lastRow {
			m.cursor++
		}
	case "left", "h":
		switch m.cursor {
		case rowDifficulty:
			m.settings.Difficulty = m.settings.Difficulty.Prev()
		case rowExamplesCount:
			m.settings.ExamplesCount = shared.NormalizeExamplesCount(m.settings.ExamplesCount - 1)
		}
	case "right", "l":
		switch m.cursor {
		case rowDifficulty:
			m.settings.Difficulty = m.settings.Difficulty.Next()
		case rowExamplesCount:
			m.settings.ExamplesCount = shared.NormalizeExamplesCount(m.settings.ExamplesCount + 1)
		}
	case "enter":
		switch m.cursor {
		case rowApply:
			return m, emit(ApplySettingsMsg{Settings: m.settings})
		case rowBack:
			return m, emit(BackMsg{})
		}
	case "esc":
		return m, emit(BackMsg{})
	}

	return m, nil
}
