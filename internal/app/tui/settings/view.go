package settings

import "strings"

func (m Model) View() string {
	var b strings.Builder

	b.WriteString("Настройки сложности\n\n")
	b.WriteString("Текущая сложность: " + m.current.String() + "\n\n")
	for i, difficulty := range m.difficulties {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		b.WriteString(cursor + " " + difficulty.String() + "\n")
	}

	b.WriteString("\n↑/↓ - выбор, Enter - применить, Esc - назад")

	return b.String()
}
