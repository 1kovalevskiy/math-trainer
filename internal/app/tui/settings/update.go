package settings

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch typedMsg := msg.(type) {
	case tea.MouseMsg:
		if !shared.IsLeftClick(typedMsg) {
			return m, nil
		}
		return m.handleMouseClick(typedMsg)
	case tea.KeyMsg:
		switch typedMsg.String() {
		case "up", "k":
			if m.cursor > rowDifficulty {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < lastRow {
				m.cursor++
			}
		case "left", "h":
			m = m.decrementCurrent()
		case "right", "l":
			m = m.incrementCurrent()
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
	}

	return m, nil
}

func (m Model) handleMouseClick(msg tea.MouseMsg) (Model, tea.Cmd) {
	switch {
	case shared.InZone(zoneDifficultyPrev, msg):
		m.cursor = rowDifficulty
		m.settings.Difficulty = m.settings.Difficulty.Prev()
		return m, nil
	case shared.InZone(zoneDifficultyNext, msg):
		m.cursor = rowDifficulty
		m.settings.Difficulty = m.settings.Difficulty.Next()
		return m, nil
	case shared.InZone(zoneCountDec, msg):
		m.cursor = rowExamplesCount
		m.settings.ExamplesCount = shared.NormalizeExamplesCount(m.settings.ExamplesCount - 1)
		return m, nil
	case shared.InZone(zoneCountInc, msg):
		m.cursor = rowExamplesCount
		m.settings.ExamplesCount = shared.NormalizeExamplesCount(m.settings.ExamplesCount + 1)
		return m, nil
	case shared.InZone(zoneApply, msg):
		m.cursor = rowApply
		return m, emit(ApplySettingsMsg{Settings: m.settings})
	case shared.InZone(zoneBack, msg):
		m.cursor = rowBack
		return m, emit(BackMsg{})
	default:
		return m, nil
	}
}

func (m Model) incrementCurrent() Model {
	switch m.cursor {
	case rowDifficulty:
		m.settings.Difficulty = m.settings.Difficulty.Next()
	case rowExamplesCount:
		m.settings.ExamplesCount = shared.NormalizeExamplesCount(m.settings.ExamplesCount + 1)
	}

	return m
}

func (m Model) decrementCurrent() Model {
	switch m.cursor {
	case rowDifficulty:
		m.settings.Difficulty = m.settings.Difficulty.Prev()
	case rowExamplesCount:
		m.settings.ExamplesCount = shared.NormalizeExamplesCount(m.settings.ExamplesCount - 1)
	}

	return m
}
