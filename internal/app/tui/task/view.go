package task

import "strings"

func (m Model) View() string {
	var b strings.Builder

	b.WriteString("Математическое задание\n\n")
	b.WriteString("Сложность: " + m.difficulty.String() + "\n")
	b.WriteString("Пример: " + m.exercise.Expression() + " = ?\n\n")
	b.WriteString("Ваш ответ: " + m.input + "\n")

	if m.errText != "" {
		b.WriteString("\nОшибка: " + m.errText + "\n")
	}

	b.WriteString("\nВведите число и нажмите Enter")
	b.WriteString("\nR - новый пример, Esc - в главное меню")

	return b.String()
}
