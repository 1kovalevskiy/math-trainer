package result

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
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
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter":
			return m, m.selectCurrent()
		}
	}

	return m, nil
}

func (m Model) selectCurrent() tea.Cmd {
	switch m.selectedOption() {
	case OptionRetry:
		return emit(RetryTaskMsg{})
	case OptionSettings:
		return emit(OpenSettingsMsg{})
	case OptionHome:
		return emit(BackToStartMsg{})
	default:
		return nil
	}
}
