package result

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

var (
	expressionStyle = lipgloss.NewStyle().Bold(true)
	correctStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	incorrectStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Strikethrough(true).Bold(true)
	skippedStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Bold(true)
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(ui.Title.Render("Результаты тренировки") + "\n")
	b.WriteString(ui.Subtitle.Render("Сводка по всем примерам") + "\n\n")
	b.WriteString(ui.Label.Render("Сложность: ") + ui.Value.Render(m.summary.Difficulty.String()) + "\n\n")

	for _, entry := range m.summary.Results {
		b.WriteString(renderEntry(entry) + "\n")
	}

	b.WriteString("\n")
	for i, option := range m.options {
		b.WriteString(zone.Mark(optionZoneID(i), ui.MenuItem(m.cursor == i, option)) + "\n")
	}

	b.WriteString("\n" + ui.Accent.Render(fmt.Sprintf("Правильных: %d из %d", m.summary.Correct, len(m.summary.Results))))
	b.WriteString("\n" + ui.Hint.Render("↑/↓ - выбор, Enter - подтвердить"))

	return b.String()
}

func optionZoneID(index int) string {
	return fmt.Sprintf("result:option:%d", index)
}

func renderEntry(entry shared.ExampleResult) string {
	base := expressionStyle.Render(fmt.Sprintf("%d) %s = ", entry.Order, entry.Expression))

	switch entry.Status {
	case shared.ResultStatusCorrect:
		answer := 0
		if entry.UserAnswer != nil {
			answer = *entry.UserAnswer
		}
		return base + correctStyle.Render(fmt.Sprintf("%d", answer))
	case shared.ResultStatusIncorrect:
		userAnswer := "_"
		if entry.UserAnswer != nil {
			userAnswer = fmt.Sprintf("%d", *entry.UserAnswer)
		}
		return base + incorrectStyle.Render(userAnswer) +
			fmt.Sprintf(" (ответ: %d)", entry.CorrectAnswer)
	case shared.ResultStatusSkipped:
		return base + skippedStyle.Render("____") +
			fmt.Sprintf(" (ответ: %d)", entry.CorrectAnswer)
	default:
		return fmt.Sprintf("%d) %s", entry.Order, entry.Expression)
	}
}
