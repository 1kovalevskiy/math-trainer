package tui

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/settings"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) handleCommandResultMsg(msg tea.Msg) (Model, tea.Cmd, bool) {
	switch typedMsg := msg.(type) {
	case trainingSnapshotMsg:
		if typedMsg.err != nil {
			m.taskModel = m.taskModel.WithError(errorText(typedMsg.err))
			return m, nil, true
		}
		return m.applySnapshot(typedMsg.snapshot), nil, true
	case persistSettingsMsg:
		if typedMsg.err != nil {
			m.settingsModel = m.settingsModel.WithError(errorText(typedMsg.err))
			m.screen = ScreenSettings
			return m, nil, true
		}
		m.settings = typedMsg.settings
		m.settingsModel = settings.NewModel(m.settings, m.mathController)
		m.screen = ScreenStart
		return m, nil, true
	case cancelTrainingMsg:
		return m, nil, true
	default:
		return m, nil, false
	}
}
