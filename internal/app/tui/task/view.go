package task

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
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

	difficulty := shared.DifficultyForOperator(m.settings, m.current.Exercise.Operator)

	b.WriteString(ui.Title.Render("Математическое задание") + "\n")
	b.WriteString(ui.Subtitle.Render(fmt.Sprintf("Пример %d из %d", m.current.Order, m.current.Total)) + "\n\n")
	b.WriteString(ui.Label.Render("Сложность: ") + ui.Value.Render(shared.DifficultyLabel(difficulty)) + "\n")
	b.WriteString(ui.Label.Render("Пример: ") + ui.Accent.Render(shared.ExerciseText(m.current.Exercise)+" = ?") + "\n\n")
	b.WriteString(ui.Label.Render("Ваш ответ: ") + ui.Value.Render(m.input) + "\n")

	if m.errText != "" {
		b.WriteString("\n" + ui.Error.Render("Ошибка: "+m.errText) + "\n")
	}

	b.WriteString("\n")
	b.WriteString(zone.Mark(zoneSubmit, ui.Button("Ответить", m.buttonCursor == buttonSubmit)))
	b.WriteString(" ")
	b.WriteString(zone.Mark(zoneSkip, ui.Button("Пропустить", m.buttonCursor == buttonSkip)))
	b.WriteString(" ")
	b.WriteString(zone.Mark(zoneBack, ui.Button("В меню", m.buttonCursor == buttonBack)))

	return b.String()
}
