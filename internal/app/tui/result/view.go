package result

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	"github.com/charmbracelet/lipgloss"
)

var (
	correctStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	incorrectStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Strikethrough(true)
	skippedStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	correctAnswerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
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
		b.WriteString(ui.MenuItem(m.cursor == i, option) + "\n")
	}

	b.WriteString("\n" + ui.Accent.Render(fmt.Sprintf("Правильных: %d из %d", m.summary.Correct, len(m.summary.Results))))
	b.WriteString("\n" + ui.Hint.Render("↑/↓ - выбор, Enter - подтвердить"))

	return b.String()
}

func renderEntry(entry shared.ExampleResult) string {
	switch entry.Status {
	case shared.ResultStatusCorrect:
		answer := 0
		if entry.UserAnswer != nil {
			answer = *entry.UserAnswer
		}
		line := fmt.Sprintf("%d) %s = %d", entry.Order, entry.Expression, answer)
		return correctStyle.Render(line)
	case shared.ResultStatusIncorrect:
		userAnswer := "?"
		if entry.UserAnswer != nil {
			userAnswer = fmt.Sprintf("%d", *entry.UserAnswer)
		}
		return fmt.Sprintf(
			"%d) %s = %s -> %s",
			entry.Order,
			entry.Expression,
			incorrectStyle.Render(userAnswer),
			correctAnswerStyle.Render(fmt.Sprintf("%d", entry.CorrectAnswer)),
		)
	case shared.ResultStatusSkipped:
		return fmt.Sprintf(
			"%d) %s = %s -> %d",
			entry.Order,
			entry.Expression,
			skippedStyle.Render("пропущено"),
			entry.CorrectAnswer,
		)
	default:
		return fmt.Sprintf("%d) %s", entry.Order, entry.Expression)
	}
}
