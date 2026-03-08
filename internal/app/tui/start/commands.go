package start

import tea "github.com/charmbracelet/bubbletea"

func emit(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}
