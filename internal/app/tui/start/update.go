package start

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
		case OptionPractice:
			return m, emit(OpenTaskMsg{})
		case OptionSettings:
			return m, emit(OpenSettingsMsg{})
		case OptionQuit:
			return m, emit(QuitMsg{})
		}
	}

	return m, nil
}
