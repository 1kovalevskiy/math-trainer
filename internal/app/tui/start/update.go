package start

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch typedMsg := msg.(type) {
	case tea.MouseMsg:
		if !shared.IsLeftClick(typedMsg) {
			return m, nil
		}

		for i := range m.options {
			if shared.InZone(optionZoneID(i), typedMsg) {
				m.cursor = i
				return m, m.selectCurrent()
			}
		}

		return m, nil
	case tea.KeyMsg:
		switch typedMsg.String() {
		case "up", "k":
			m.cursor = ui.MoveIndex(m.cursor, -1, 0, len(m.options)-1)
		case "down", "j":
			m.cursor = ui.MoveIndex(m.cursor, 1, 0, len(m.options)-1)
		case "enter":
			return m, m.selectCurrent()
		}

		return m, nil
	}

	return m, nil
}

func (m Model) selectCurrent() tea.Cmd {
	switch m.selectedOption() {
	case OptionPractice:
		return emit(OpenTaskMsg{})
	case OptionSettings:
		return emit(OpenSettingsMsg{})
	case OptionQuit:
		return emit(QuitMsg{})
	default:
		return nil
	}
}
