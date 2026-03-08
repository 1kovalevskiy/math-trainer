package result

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	status := "Неверно"
	if m.outcome.Correct {
		status = "Верно"
	}

	b.WriteString("Результат\n\n")
	b.WriteString("Сложность: " + m.outcome.Difficulty.String() + "\n")
	b.WriteString("Пример: " + m.outcome.Expression + "\n")
	b.WriteString("Ваш ответ: " + fmt.Sprintf("%d", m.outcome.Answer) + "\n")
	b.WriteString("Правильный ответ: " + fmt.Sprintf("%d", m.outcome.Expected) + "\n")
	b.WriteString("Статус: " + status + "\n\n")

	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		b.WriteString(cursor + " " + option + "\n")
	}

	b.WriteString("\n↑/↓ - выбор, Enter - подтвердить")

	return b.String()
}
