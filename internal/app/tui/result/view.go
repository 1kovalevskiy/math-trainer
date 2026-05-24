package result

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
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
	b.WriteString(ui.Label.Render("Сложности: ") + ui.Value.Render(settingsDifficultySummary(m.summary.Settings)) + "\n\n")

	for _, entry := range m.summary.Results {
		b.WriteString(renderEntry(entry) + "\n")
	}

	b.WriteString("\n")
	for i, option := range m.options {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(zone.Mark(optionZoneID(i), ui.MenuItem(m.cursor == i, option)))
	}

	b.WriteString("\n" + ui.Accent.Render(fmt.Sprintf("Правильных: %d из %d", m.summary.Correct, m.summary.Total)))

	return b.String()
}

func optionZoneID(index int) string {
	return fmt.Sprintf("result:option:%d", index)
}

func settingsDifficultySummary(settings mathmodels.TrainingSettings) string {
	parts := []string{
		fmt.Sprintf("+ %s", shared.DifficultyLabel(settings.AddDifficulty)),
		fmt.Sprintf("- %s", shared.DifficultyLabel(settings.SubtractDifficulty)),
		fmt.Sprintf("* %s", shared.DifficultyLabel(settings.MultiplyDifficulty)),
		fmt.Sprintf("/ %s", shared.DifficultyLabel(settings.DivideDifficulty)),
	}
	return strings.Join(parts, ", ")
}

func renderEntry(entry mathmodels.ExampleResult) string {
	base := expressionStyle.Render(fmt.Sprintf("%d) %s = ", entry.Order, shared.ExerciseText(entry.Exercise)))

	switch entry.Status {
	case mathmodels.ResultStatusCorrect:
		answer := 0
		if entry.UserAnswer != nil {
			answer = *entry.UserAnswer
		}
		return base + correctStyle.Render(fmt.Sprintf("%d", answer))
	case mathmodels.ResultStatusIncorrect:
		userAnswer := "_"
		if entry.UserAnswer != nil {
			userAnswer = fmt.Sprintf("%d", *entry.UserAnswer)
		}
		return base + incorrectStyle.Render(userAnswer) +
			fmt.Sprintf(" (ответ: %d)", entry.CorrectAnswer)
	case mathmodels.ResultStatusSkipped:
		return base + skippedStyle.Render("____") +
			fmt.Sprintf(" (ответ: %d)", entry.CorrectAnswer)
	default:
		return fmt.Sprintf("%d) %s", entry.Order, shared.ExerciseText(entry.Exercise))
	}
}
