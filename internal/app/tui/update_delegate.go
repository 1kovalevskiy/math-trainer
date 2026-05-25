package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) updateCurrentScreen(msg tea.Msg) (Model, tea.Cmd, bool) {
	var cmd tea.Cmd

	switch m.screen {
	case ScreenStart:
		m.startModel, cmd = m.startModel.Update(msg)
	case ScreenSettings:
		m.settingsModel, cmd = m.settingsModel.Update(msg)
	case ScreenTask:
		m.taskModel, cmd = m.taskModel.Update(msg)
	case ScreenResult:
		m.resultModel = m.resultModelWithCurrentViewport(m.resultModel)
		m.resultModel, cmd = m.resultModel.Update(msg)
	default:
		return m, nil, false
	}

	return m, cmd, true
}
