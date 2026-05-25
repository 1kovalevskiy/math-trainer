package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) handleSystemMsg(msg tea.Msg) (Model, tea.Cmd, bool) {
	switch typedMsg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = typedMsg.Width
		m.height = typedMsg.Height
		if m.screen == ScreenResult {
			m.resultModel = m.resultModelWithCurrentViewport(m.resultModel)
		}
		return m, nil, true
	case tea.KeyMsg:
		if typedMsg.String() == "ctrl+c" {
			return m, tea.Quit, true
		}
	}

	return m, nil, false
}
