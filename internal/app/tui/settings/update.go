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
		m.errText = ""
		return m.handleMouseClick(typedMsg)
	case tea.KeyMsg:
		switch typedMsg.String() {
		case "up", "k":
			m.errText = ""
			m.focus = m.focus.moveUp()
		case "down", "j":
			m.errText = ""
			m.focus = m.focus.moveDown()
		case "left", "h":
			m.errText = ""
			m = m.decrementCurrent()
		case "right", "l":
			m.errText = ""
			m = m.incrementCurrent()
		case "enter":
			switch {
			case m.focus.isAction(actionApply):
				return m, emit(ApplySettingsMsg{Settings: m.settings})
			case m.focus.isAction(actionBack):
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
	case shared.InZone(zoneAddPrev, msg):
		m.focus = settingsFocus{section: focusSettings, setting: settingAddDifficulty}
		m.settings.AddDifficulty = m.prevDifficulty(m.settings.AddDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneAddNext, msg):
		m.focus = settingsFocus{section: focusSettings, setting: settingAddDifficulty}
		m.settings.AddDifficulty = m.nextDifficulty(m.settings.AddDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneSubtractPrev, msg):
		m.focus = settingsFocus{section: focusSettings, setting: settingSubtractDifficulty}
		m.settings.SubtractDifficulty = m.prevDifficulty(m.settings.SubtractDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneSubtractNext, msg):
		m.focus = settingsFocus{section: focusSettings, setting: settingSubtractDifficulty}
		m.settings.SubtractDifficulty = m.nextDifficulty(m.settings.SubtractDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneMultiplyPrev, msg):
		m.focus = settingsFocus{section: focusSettings, setting: settingMultiplyDifficulty}
		m.settings.MultiplyDifficulty = m.prevDifficulty(m.settings.MultiplyDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneMultiplyNext, msg):
		m.focus = settingsFocus{section: focusSettings, setting: settingMultiplyDifficulty}
		m.settings.MultiplyDifficulty = m.nextDifficulty(m.settings.MultiplyDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneDividePrev, msg):
		m.focus = settingsFocus{section: focusSettings, setting: settingDivideDifficulty}
		m.settings.DivideDifficulty = m.prevDifficulty(m.settings.DivideDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneDivideNext, msg):
		m.focus = settingsFocus{section: focusSettings, setting: settingDivideDifficulty}
		m.settings.DivideDifficulty = m.nextDifficulty(m.settings.DivideDifficulty)
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneCountDec, msg):
		m.focus = settingsFocus{section: focusSettings, setting: settingExamplesCount}
		m.settings.ExamplesCount--
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneCountInc, msg):
		m.focus = settingsFocus{section: focusSettings, setting: settingExamplesCount}
		m.settings.ExamplesCount++
		m.settings = m.normalizeSettings()
		return m, nil
	case shared.InZone(zoneApply, msg):
		m.focus = settingsFocus{section: focusActions, action: actionApply}
		return m, emit(ApplySettingsMsg{Settings: m.settings})
	case shared.InZone(zoneBack, msg):
		m.focus = settingsFocus{section: focusActions, action: actionBack}
		return m, emit(BackMsg{})
	default:
		return m, nil
	}
}

func (m Model) incrementCurrent() Model {
	switch {
	case m.focus.isSetting(settingAddDifficulty):
		m.settings.AddDifficulty = m.nextDifficulty(m.settings.AddDifficulty)
		m.settings = m.normalizeSettings()
	case m.focus.isSetting(settingSubtractDifficulty):
		m.settings.SubtractDifficulty = m.nextDifficulty(m.settings.SubtractDifficulty)
		m.settings = m.normalizeSettings()
	case m.focus.isSetting(settingMultiplyDifficulty):
		m.settings.MultiplyDifficulty = m.nextDifficulty(m.settings.MultiplyDifficulty)
		m.settings = m.normalizeSettings()
	case m.focus.isSetting(settingDivideDifficulty):
		m.settings.DivideDifficulty = m.nextDifficulty(m.settings.DivideDifficulty)
		m.settings = m.normalizeSettings()
	case m.focus.isSetting(settingExamplesCount):
		m.settings.ExamplesCount++
		m.settings = m.normalizeSettings()
	case m.focus.section == focusActions:
		m.focus = m.focus.moveRight()
	}

	return m
}

func (m Model) decrementCurrent() Model {
	switch {
	case m.focus.isSetting(settingAddDifficulty):
		m.settings.AddDifficulty = m.prevDifficulty(m.settings.AddDifficulty)
		m.settings = m.normalizeSettings()
	case m.focus.isSetting(settingSubtractDifficulty):
		m.settings.SubtractDifficulty = m.prevDifficulty(m.settings.SubtractDifficulty)
		m.settings = m.normalizeSettings()
	case m.focus.isSetting(settingMultiplyDifficulty):
		m.settings.MultiplyDifficulty = m.prevDifficulty(m.settings.MultiplyDifficulty)
		m.settings = m.normalizeSettings()
	case m.focus.isSetting(settingDivideDifficulty):
		m.settings.DivideDifficulty = m.prevDifficulty(m.settings.DivideDifficulty)
		m.settings = m.normalizeSettings()
	case m.focus.isSetting(settingExamplesCount):
		m.settings.ExamplesCount--
		m.settings = m.normalizeSettings()
	case m.focus.section == focusActions:
		m.focus = m.focus.moveLeft()
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
