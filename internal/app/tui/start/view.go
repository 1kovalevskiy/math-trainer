package start

import "strings"

func (m Model) View() string {
	var b strings.Builder

	b.WriteString("Математический тренажер\n\n")
	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		b.WriteString(cursor + " " + option + "\n")
	}

	b.WriteString("\n↑/↓ - выбор, Enter - подтвердить, Ctrl+C - выход")

	return b.String()
}
