package task

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	zone "github.com/lrstanley/bubblezone"
)

const (
	zoneSubmit = "task:submit"
	zoneSkip   = "task:skip"
	zoneBack   = "task:back"
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

	b.WriteString("\n")
	b.WriteString(zone.Mark(zoneSubmit, ui.Button("Ответить", m.canSubmit())))
	b.WriteString(" ")
	b.WriteString(zone.Mark(zoneSkip, ui.Button("Пропустить", false)))
	b.WriteString(" ")
	b.WriteString(zone.Mark(zoneBack, ui.Button("В меню", false)))
	b.WriteString("\n\n" + ui.Hint.Render("Введите число и нажмите Enter"))
	b.WriteString("\n" + ui.Hint.Render("S - пропустить, Esc - в меню, Click - мышь"))

	return b.String()
}
