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

const maxResultColumns = 3

var (
	expressionStyle = lipgloss.NewStyle().Bold(true)
	correctStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	incorrectStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Strikethrough(true).Bold(true)
	skippedStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Bold(true)
)

func (m Model) View() string {
	return m.ViewWithSize(ui.MinPanelContentWidth, 24)
}

func (m Model) ViewWithSize(width int, height int) string {
	if width < 1 {
		width = 1
	}
	if height < 1 {
		height = 1
	}

	headerLines := []string{
		ui.Title.Render("Результаты тренировки"),
		ui.Subtitle.Render("Сводка по всем примерам"),
		"",
		ui.Label.Render("Сложности: ") + ui.Value.Render(settingsDifficultySummary(m.summary.Settings)),
		ui.Accent.Render(fmt.Sprintf("Правильных: %d из %d", m.summary.Correct, m.summary.Total)),
		"",
	}
	actionsLine := m.renderActionButtons()

	headerHeight := len(headerLines)
	actionsHeight := 1
	resultsHeight := height - headerHeight - actionsHeight
	if resultsHeight < 1 {
		resultsHeight = 1
	}

	viewModel := m.WithViewport(width, resultsHeight)
	viewModel.refreshScrollBounds()
	resultsBlock := viewModel.renderResultsBlock(width, resultsHeight)

	lines := make([]string, 0, headerHeight+resultsHeight+actionsHeight)
	lines = append(lines, headerLines...)
	lines = append(lines, strings.Split(resultsBlock, "\n")...)
	lines = append(lines, actionsLine)

	if len(lines) > height {
		lines = lines[:height]
	}
	for len(lines) < height {
		lines = append(lines, "")
	}

	for i := range lines {
		lines[i] = lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Render(lines[i])
	}

	return strings.Join(lines, "\n")
}

func (m Model) renderResultsBlock(width int, viewportHeight int) string {
	entries := make([]string, 0, len(m.summary.Results))
	for _, entry := range m.summary.Results {
		entries = append(entries, renderEntry(entry))
	}

	if len(entries) == 0 {
		line := lipgloss.NewStyle().Width(width).Render(ui.Hint.Render("Нет ответов"))
		return strings.Join(fillLines([]string{line}, viewportHeight), "\n")
	}

	entryWidth := 1
	for _, entry := range entries {
		entryWidth = max(entryWidth, lipgloss.Width(entry))
	}

	layout := m.gridLayout(width, viewportHeight, len(entries), entryWidth)
	offset := ui.ScrollState{
		Offset:       m.scrollOffset,
		ViewportRows: viewportHeight,
		ContentRows:  layout.ContentRows,
	}.ClampOffset(m.scrollOffset)

	rowLines := m.renderGridRows(entries, layout.Columns, layout.Rows, entryWidth, layout.ContentWidth)
	visible := cropRows(rowLines, offset, viewportHeight)

	if layout.ShowScrollbar {
		visible = addVerticalScrollbar(visible, layout.ContentWidth, viewportHeight, offset, layout.ContentRows)
	}

	return strings.Join(fillLines(visible, viewportHeight), "\n")
}

func (m Model) gridLayout(width int, viewportHeight int, total int, entryWidth int) ui.GridLayout {
	return ui.BuildGridLayout(width, viewportHeight, total, entryWidth, ui.GridOptions{
		MaxColumns:       maxResultColumns,
		PreferredColumns: maxResultColumns,
		ColumnGap:        3,
		ScrollbarWidth:   2,
	})
}

func (m Model) renderGridRows(entries []string, columns int, rows int, entryWidth int, contentWidth int) []string {
	columnGap := "   "
	lines := make([]string, 0, rows)
	for row := 0; row < rows; row++ {
		cells := make([]string, 0, columns)
		for col := 0; col < columns; col++ {
			idx := row*columns + col
			if idx >= len(entries) {
				cells = append(cells, lipgloss.NewStyle().Width(entryWidth).Render(""))
				continue
			}
			cells = append(cells, lipgloss.NewStyle().Width(entryWidth).Render(entries[idx]))
		}
		line := strings.Join(cells, columnGap)
		lines = append(lines, lipgloss.NewStyle().Width(contentWidth).Render(line))
	}
	return lines
}

func (m Model) renderActionButtons() string {
	buttons := make([]string, 0, len(m.options))
	for i, option := range m.options {
		buttons = append(buttons, zone.Mark(optionZoneID(i), ui.MenuItem(m.cursor == i, option)))
	}
	return ui.JoinInline(buttons, 1)
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

func cropRows(rows []string, offset int, viewportHeight int) []string {
	if offset < 0 {
		offset = 0
	}
	if offset > len(rows) {
		offset = len(rows)
	}
	end := offset + viewportHeight
	if end > len(rows) {
		end = len(rows)
	}
	return append([]string(nil), rows[offset:end]...)
}

func fillLines(lines []string, height int) []string {
	for len(lines) < height {
		lines = append(lines, "")
	}
	return lines
}

func addVerticalScrollbar(lines []string, contentWidth int, viewportHeight int, offset int, totalRows int) []string {
	if viewportHeight < 1 {
		return lines
	}
	if totalRows <= viewportHeight {
		return lines
	}

	trackHeight := viewportHeight
	thumbHeight := max(1, int(float64(trackHeight*viewportHeight)/float64(totalRows)+0.5))
	maxThumbTop := trackHeight - thumbHeight
	maxOffset := totalRows - viewportHeight
	thumbTop := 0
	if maxOffset > 0 && maxThumbTop > 0 {
		thumbTop = int(float64(offset*maxThumbTop)/float64(maxOffset) + 0.5)
	}

	res := make([]string, 0, len(lines))
	for i, line := range lines {
		glyph := "│"
		if i >= thumbTop && i < thumbTop+thumbHeight {
			glyph = "█"
		}
		res = append(res, lipgloss.NewStyle().Width(contentWidth).Render(line)+" "+glyph)
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
