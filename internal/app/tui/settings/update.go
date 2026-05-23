package settings

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
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
			m = m.moveCursorUp()
		case "down", "j":
			m = m.moveCursorDown()
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

func (m Model) moveCursorUp() Model {
	if m.cursor == rowApply || m.cursor == rowBack {
		m.cursor = rowExamplesCount
		return m
	}
	if m.cursor > rowDifficulty {
		m.cursor--
	}

	return m
}

func (m Model) moveCursorDown() Model {
	if m.cursor == rowExamplesCount {
		m.cursor = rowApply
		return m
	}
	if m.cursor < rowExamplesCount {
		m.cursor++
	}

	return m
}

func (m Model) handleMouseClick(msg tea.MouseMsg) (Model, tea.Cmd) {
	switch {
	case shared.InZone(zoneDifficultyPrev, msg):
		m.cursor = rowDifficulty
		m.settings.Difficulty = m.prevDifficulty()
		return m, nil
	case shared.InZone(zoneDifficultyNext, msg):
		m.cursor = rowDifficulty
		m.settings.Difficulty = m.nextDifficulty()
		return m, nil
	case shared.InZone(zoneCountDec, msg):
		m.cursor = rowExamplesCount
		m.settings.ExamplesCount--
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneCountInc, msg):
		m.cursor = rowExamplesCount
		m.settings.ExamplesCount++
		m.settings = m.normalizeSettings()
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
		m.settings.Difficulty = m.nextDifficulty()
	case rowExamplesCount:
		m.settings.ExamplesCount++
		m.settings = m.normalizeSettings()
	case rowApply:
		m.cursor = rowBack
	}

	return m
}

func (m Model) decrementCurrent() Model {
	switch m.cursor {
	case rowDifficulty:
		m.settings.Difficulty = m.prevDifficulty()
	case rowExamplesCount:
		m.settings.ExamplesCount--
		m.settings = m.normalizeSettings()
	case rowBack:
		m.cursor = rowApply
	}

	return m
}

func (m Model) normalizeSettings() mathmodels.TrainingSettings {
	if m.rules == nil {
		return m.settings
	}

	return m.rules.NormalizeSettings(m.settings)
}

func (m Model) nextDifficulty() mathmodels.Difficulty {
	if m.rules == nil {
		return m.settings.Difficulty
	}

	return m.rules.GetNextDifficulty(m.settings.Difficulty)
}

func (m Model) prevDifficulty() mathmodels.Difficulty {
	if m.rules == nil {
		return m.settings.Difficulty
	}

	return m.rules.GetPreviousDifficulty(m.settings.Difficulty)
}
