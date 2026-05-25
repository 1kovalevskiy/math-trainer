package result

import (
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	"github.com/charmbracelet/lipgloss"
)

const (
	maxResultColumns = 3
	resultColumnGap  = 3
	resultRowGap     = 1
)

var (
	resultSurfaceStyle  = lipgloss.NewStyle().Background(ui.PanelBackgroundColor)
	resultTitleStyle    = ui.Title.Copy()
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
	return renderSummaryHeaderLines(m.summary)
}

func (m Model) resultsViewportHeight(contentHeight int) int {
	resultsHeight := contentHeight - len(m.renderHeaderLines()) - 1
	if resultsHeight < 1 {
		return 1
	}

	return resultsHeight
}

func (m Model) renderResultsBlock(width int, viewportHeight int) string {
	entries := renderResultEntries(m.summary.Results)
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
