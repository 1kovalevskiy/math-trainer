package task

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	zone "github.com/lrstanley/bubblezone"
)

var (
	subtitleStyle = ui.Subtitle.Copy().Background(ui.PanelBackgroundColor)
	labelStyle    = ui.Label.Copy().Background(ui.PanelBackgroundColor)
	valueStyle    = ui.Value.Copy().Background(ui.PanelBackgroundColor)
	accentStyle   = ui.Accent.Copy().Background(ui.PanelBackgroundColor)
	errorStyle    = ui.Error.Copy().Background(ui.PanelBackgroundColor)
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
	b.WriteString(subtitleStyle.Render(fmt.Sprintf("Пример %d из %d", m.current.Order, m.current.Total)) + "\n\n")
	b.WriteString(labelStyle.Render("Сложность: ") + valueStyle.Render(shared.DifficultyLabel(difficulty)) + "\n")
	b.WriteString(labelStyle.Render("Пример: ") + accentStyle.Render(shared.ExerciseText(m.current.Exercise)+" = ?") + "\n\n")
	b.WriteString(labelStyle.Render("Ваш ответ: ") + valueStyle.Render(m.input) + "\n")

	if m.errText != "" {
		b.WriteString("\n" + errorStyle.Render("Ошибка: "+m.errText) + "\n")
	}

	b.WriteString("\n")
	b.WriteString(ui.JoinInline([]string{
		zone.Mark(zoneSubmit, ui.Button("Ответить", m.buttonCursor == buttonSubmit)),
		zone.Mark(zoneSkip, ui.Button("Пропустить", m.buttonCursor == buttonSkip)),
		zone.Mark(zoneBack, ui.Button("В меню", m.buttonCursor == buttonBack)),
	}, 1))

	return b.String()
}

func (m Model) ViewWithSize(_, _ int) string {
	return m.View()
}
