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

const (
	maxResultColumns = 3
	resultColumnGap  = 3
	resultRowGap     = 1
)

var (
	resultSurfaceStyle  = lipgloss.NewStyle().Background(ui.PanelBackgroundColor)
	resultSubtitleStyle = ui.Subtitle.Copy().Background(ui.PanelBackgroundColor)
	resultLabelStyle    = ui.Label.Copy().Background(ui.PanelBackgroundColor)
	resultValueStyle    = ui.Value.Copy().Background(ui.PanelBackgroundColor)
	resultAccentStyle   = ui.Accent.Copy().Background(ui.PanelBackgroundColor)
	correctStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("235")).Background(lipgloss.Color("151")).Bold(true)
	incorrectStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("235")).Background(lipgloss.Color("217")).Bold(true)
	skippedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("230")).Background(lipgloss.Color("240")).Bold(true)
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

	headerLines := m.renderHeaderLines()
	actionsLine := m.renderActionButtons()

	headerHeight := len(headerLines)
	actionsHeight := 1
	resultsHeight := m.resultsViewportHeight(height)

	viewModel := m.WithContentSize(width, height)
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
		lines[i] = resultPadCenter(lines[i], width)
	}

	return strings.Join(lines, "\n")
}

func (m Model) renderHeaderLines() []string {
	return []string{
		ui.Title.Render("Результаты тренировки"),
		resultSubtitleStyle.Render("Сводка по всем примерам"),
		"",
		resultLabelStyle.Render("Сложности: ") + resultValueStyle.Render(settingsDifficultySummary(m.summary.Settings)),
		resultAccentStyle.Render(fmt.Sprintf("Правильных: %d из %d", m.summary.Correct, m.summary.Total)),
		"",
	}
}

func (m Model) resultsViewportHeight(contentHeight int) int {
	resultsHeight := contentHeight - len(m.renderHeaderLines()) - 1
	if resultsHeight < 1 {
		return 1
	}

	return resultsHeight
}

func (m Model) renderResultsBlock(width int, viewportHeight int) string {
	entries := make([]string, 0, len(m.summary.Results))
	for _, entry := range m.summary.Results {
		entries = append(entries, renderEntry(entry))
	}

	if len(entries) == 0 {
		line := ui.Hint.Render("Нет ответов")
		return strings.Join(centerLines([]string{line}, viewportHeight), "\n")
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
		visible = addVerticalScrollbar(visible, width, viewportHeight, offset, layout.ContentRows)
		return strings.Join(fillLines(visible, viewportHeight), "\n")
	}

	return strings.Join(centerLines(visible, viewportHeight), "\n")
}

func (m Model) gridLayout(width int, viewportHeight int, total int, entryWidth int) ui.GridLayout {
	return ui.BuildGridLayout(width, viewportHeight, total, entryWidth, ui.GridOptions{
		MaxColumns:       maxResultColumns,
		PreferredColumns: maxResultColumns,
		ColumnGap:        resultColumnGap,
		RowGap:           resultRowGap,
		ScrollbarWidth:   2,
	})
}

func (m Model) renderGridRows(entries []string, columns int, rows int, entryWidth int, contentWidth int) []string {
	lines := make([]string, 0, rows+(rows-1)*resultRowGap)
	for row := 0; row < rows; row++ {
		if row > 0 {
			for i := 0; i < resultRowGap; i++ {
				lines = append(lines, "")
			}
		}

		line := renderGridRow(entries, row, columns, entryWidth, contentWidth)
		lines = append(lines, line)
	}
	return lines
}

func renderGridRow(entries []string, row int, columns int, entryWidth int, contentWidth int) string {
	start := row * columns
	if start >= len(entries) {
		return ""
	}

	count := columns
	if remaining := len(entries) - start; remaining < count {
		count = remaining
	}

	cells := make([]string, 0, count)
	for col := 0; col < count; col++ {
		cells = append(cells, resultPadRight(entries[start+col], entryWidth))
	}

	line := strings.Join(cells, resultPad(resultColumnGap))
	if count < columns {
		return resultPadCenter(line, contentWidth)
	}

	return resultPadRight(line, contentWidth)
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
	base := fmt.Sprintf("%d) %s = ", entry.Order, shared.ExerciseText(entry.Exercise))

	switch entry.Status {
	case mathmodels.ResultStatusCorrect:
		answer := 0
		if entry.UserAnswer != nil {
			answer = *entry.UserAnswer
		}
		return correctStyle.Render(base + fmt.Sprintf("%d", answer))
	case mathmodels.ResultStatusIncorrect:
		userAnswer := "_"
		if entry.UserAnswer != nil {
			userAnswer = fmt.Sprintf("%d", *entry.UserAnswer)
		}
		return incorrectStyle.Render(base + userAnswer + fmt.Sprintf(" (ответ: %d)", entry.CorrectAnswer))
	case mathmodels.ResultStatusSkipped:
		return skippedStyle.Render(base + fmt.Sprintf("____ (ответ: %d)", entry.CorrectAnswer))
	default:
		return resultSurfaceStyle.Render(fmt.Sprintf("%d) %s", entry.Order, shared.ExerciseText(entry.Exercise)))
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

func centerLines(lines []string, height int) []string {
	if len(lines) >= height {
		return lines
	}

	top := (height - len(lines)) / 2
	centered := make([]string, 0, height)
	for i := 0; i < top; i++ {
		centered = append(centered, "")
	}
	centered = append(centered, lines...)
	return fillLines(centered, height)
}

func addVerticalScrollbar(lines []string, viewportWidth int, viewportHeight int, offset int, totalRows int) []string {
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
		res = append(res, resultPadCenterWithRightGlyph(line, viewportWidth, glyph))
	}

	return res
}

func resultPad(width int) string {
	return ui.StyledPad(resultSurfaceStyle, width)
}

func resultPadRight(content string, width int) string {
	return ui.StyledPadRight(resultSurfaceStyle, content, width)
}

func resultPadCenter(content string, width int) string {
	return ui.StyledPadCenter(resultSurfaceStyle, content, width)
}

func resultPadCenterWithRightGlyph(content string, width int, glyph string) string {
	contentWidth := ui.Width(content)
	padding := width - contentWidth
	if padding <= 0 {
		return content
	}

	left := padding / 2
	right := padding - left - ui.Width(glyph)
	if right < 0 {
		right = 0
	}

	return resultPad(left) + content + resultPad(right) + glyph
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
