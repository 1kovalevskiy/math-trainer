package task

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString("Математическое задание\n\n")
	b.WriteString(fmt.Sprintf("Пример %d из %d\n", m.index, m.total))
	b.WriteString("Сложность: " + m.difficulty.String() + "\n")
	b.WriteString("Пример: " + m.exercise.Expression() + " = ?\n\n")
	b.WriteString("Ваш ответ: " + m.input + "\n")

	if m.errText != "" {
		b.WriteString("\nОшибка: " + m.errText + "\n")
	}

	b.WriteString("\nВведите число и нажмите Enter")
	b.WriteString("\nS - пропустить пример, Esc - в главное меню")

	return b.String()
}
