package result

import (
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
)

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
