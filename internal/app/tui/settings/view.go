package settings

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString("Настройки тренировки\n\n")
	b.WriteString(withCursor(m.cursor == rowDifficulty, "Сложность: "+m.settings.Difficulty.String()) + "\n")
	b.WriteString(withCursor(m.cursor == rowExamplesCount, fmt.Sprintf("Количество примеров: %d", m.settings.ExamplesCount)) + "\n\n")
	b.WriteString(withCursor(m.cursor == rowApply, "[ Применить ]") + "\n")
	b.WriteString(withCursor(m.cursor == rowBack, "[ Назад ]") + "\n")
	b.WriteString("\n↑/↓ - выбор пункта")
	b.WriteString("\n←/→ - изменить значение")
	b.WriteString("\nEnter - подтвердить, Esc - назад")

	return b.String()
}

func withCursor(active bool, text string) string {
	if active {
		return "> " + text
	}

	return "  " + text
}
