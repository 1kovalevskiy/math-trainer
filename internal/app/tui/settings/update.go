package settings

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch typedMsg := msg.(type) {
	case tea.MouseMsg:
		if !shared.IsLeftClick(typedMsg) {
			return m, nil
		}
		m.errText = ""
		return m.handleMouseClick(typedMsg)
	case tea.KeyMsg:
		switch typedMsg.String() {
		case "up", "k":
			m.errText = ""
			m = m.moveCursorUp()
		case "down", "j":
			m.errText = ""
			m = m.moveCursorDown()
		case "left", "h":
			m.errText = ""
			m = m.decrementCurrent()
		case "right", "l":
			m.errText = ""
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
	m.cursor = ui.MoveIndex(m.cursor, -1, rowAddDifficulty, rowExamplesCount)

	return m
}

func (m Model) moveCursorDown() Model {
	if m.cursor == rowExamplesCount {
		m.cursor = rowApply
		return m
	}
	if m.cursor == rowApply || m.cursor == rowBack {
		return m
	}
	m.cursor = ui.MoveIndex(m.cursor, 1, rowAddDifficulty, rowExamplesCount)

	return m
}

func (m Model) handleMouseClick(msg tea.MouseMsg) (Model, tea.Cmd) {
	switch {
	case shared.InZone(zoneAddPrev, msg):
		m.cursor = rowAddDifficulty
		m.settings.AddDifficulty = m.prevDifficulty(m.settings.AddDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneAddNext, msg):
		m.cursor = rowAddDifficulty
		m.settings.AddDifficulty = m.nextDifficulty(m.settings.AddDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneSubtractPrev, msg):
		m.cursor = rowSubtractDifficulty
		m.settings.SubtractDifficulty = m.prevDifficulty(m.settings.SubtractDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneSubtractNext, msg):
		m.cursor = rowSubtractDifficulty
		m.settings.SubtractDifficulty = m.nextDifficulty(m.settings.SubtractDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneMultiplyPrev, msg):
		m.cursor = rowMultiplyDifficulty
		m.settings.MultiplyDifficulty = m.prevDifficulty(m.settings.MultiplyDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneMultiplyNext, msg):
		m.cursor = rowMultiplyDifficulty
		m.settings.MultiplyDifficulty = m.nextDifficulty(m.settings.MultiplyDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneDividePrev, msg):
		m.cursor = rowDivideDifficulty
		m.settings.DivideDifficulty = m.prevDifficulty(m.settings.DivideDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneDivideNext, msg):
		m.cursor = rowDivideDifficulty
		m.settings.DivideDifficulty = m.nextDifficulty(m.settings.DivideDifficulty)
		m.settings = m.normalizeSettings()
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
	case rowAddDifficulty:
		m.settings.AddDifficulty = m.nextDifficulty(m.settings.AddDifficulty)
		m.settings = m.normalizeSettings()
	case rowSubtractDifficulty:
		m.settings.SubtractDifficulty = m.nextDifficulty(m.settings.SubtractDifficulty)
		m.settings = m.normalizeSettings()
	case rowMultiplyDifficulty:
		m.settings.MultiplyDifficulty = m.nextDifficulty(m.settings.MultiplyDifficulty)
		m.settings = m.normalizeSettings()
	case rowDivideDifficulty:
		m.settings.DivideDifficulty = m.nextDifficulty(m.settings.DivideDifficulty)
		m.settings = m.normalizeSettings()
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
	case rowAddDifficulty:
		m.settings.AddDifficulty = m.prevDifficulty(m.settings.AddDifficulty)
		m.settings = m.normalizeSettings()
	case rowSubtractDifficulty:
		m.settings.SubtractDifficulty = m.prevDifficulty(m.settings.SubtractDifficulty)
		m.settings = m.normalizeSettings()
	case rowMultiplyDifficulty:
		m.settings.MultiplyDifficulty = m.prevDifficulty(m.settings.MultiplyDifficulty)
		m.settings = m.normalizeSettings()
	case rowDivideDifficulty:
		m.settings.DivideDifficulty = m.prevDifficulty(m.settings.DivideDifficulty)
		m.settings = m.normalizeSettings()
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

func (m Model) nextDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty {
	if m.rules == nil {
		return current
	}

	return m.rules.GetNextDifficulty(current)
}

func (m Model) prevDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty {
	if m.rules == nil {
		return current
	}

	return m.rules.GetPreviousDifficulty(current)
}
