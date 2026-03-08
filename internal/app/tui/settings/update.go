package settings

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
		if m.cursor < len(m.difficulties)-1 {
			m.cursor++
		}
	case "enter":
		selected := m.selectedDifficulty()
		m.current = selected
		return m, emit(ApplyDifficultyMsg{Difficulty: selected})
	case "esc":
		return m, emit(BackMsg{})
	}

	return m, nil
}
