package task

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(ui.Title.Render("Математическое задание") + "\n")
	b.WriteString(ui.Subtitle.Render(fmt.Sprintf("Пример %d из %d", m.index, m.total)) + "\n\n")
	b.WriteString(ui.Label.Render("Сложность: ") + ui.Value.Render(m.difficulty.String()) + "\n")
	b.WriteString(ui.Label.Render("Пример: ") + ui.Accent.Render(m.exercise.Expression()+" = ?") + "\n\n")
	b.WriteString(ui.Label.Render("Ваш ответ: ") + ui.Value.Render(m.input) + "\n")

	if m.errText != "" {
		b.WriteString("\n" + ui.Error.Render("Ошибка: "+m.errText) + "\n")
	}

	b.WriteString("\n" + ui.Hint.Render("Введите число и нажмите Enter"))
	b.WriteString("\n" + ui.Hint.Render("S - пропустить пример, Esc - в главное меню"))

	return b.String()
}
