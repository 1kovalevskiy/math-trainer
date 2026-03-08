package result

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch key.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.options)-1 {
			m.cursor++
		}
	case "enter":
		switch m.selectedOption() {
		case OptionRetry:
			return m, emit(RetryTaskMsg{})
		case OptionSettings:
			return m, emit(OpenSettingsMsg{})
		case OptionHome:
			return m, emit(BackToStartMsg{})
		}
	}

	return m, nil
}
