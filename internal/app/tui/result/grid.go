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
	columnWidths := fullRowColumnWidths(entries, columns, entryWidth)
	lines := make([]string, 0, rows+(rows-1)*resultRowGap)
	for row := 0; row < rows; row++ {
		if row > 0 {
			for i := 0; i < resultRowGap; i++ {
				lines = append(lines, "")
			}
		}

		line := renderGridRow(entries, row, columns, columnWidths, entryWidth, contentWidth)
		lines = append(lines, line)
	}
	return lines
}

func renderGridRow(entries []string, row int, columns int, columnWidths []int, entryWidth int, contentWidth int) string {
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
		width := entryWidth
		if count == columns && col < len(columnWidths) {
			width = columnWidths[col]
		}
		cells = append(cells, resultPadRight(entries[start+col], width))
	}

	if count == 3 && columns == 3 && len(columnWidths) >= 3 {
		return renderFullThreeColumnRow(cells, columnWidths, contentWidth)
	}

	line := strings.Join(cells, resultPad(resultColumnGap))
	if count < columns {
		return resultPadCenter(line, contentWidth)
	}

	return resultPadRight(line, contentWidth)
}

func renderFullThreeColumnRow(cells []string, columnWidths []int, contentWidth int) string {
	if len(cells) != 3 || len(columnWidths) < 3 {
		return resultPadRight(strings.Join(cells, resultPad(resultColumnGap)), contentWidth)
	}

	middleStart := (contentWidth - columnWidths[1]) / 2
	leftStart := middleStart - resultColumnGap - columnWidths[0]
	rightStart := middleStart + columnWidths[1] + resultColumnGap
	if leftStart < 0 || rightStart+columnWidths[2] > contentWidth {
		return resultPadRight(strings.Join(cells, resultPad(resultColumnGap)), contentWidth)
	}

	return strings.Join([]string{
		resultPad(leftStart),
		cells[0],
		resultPad(middleStart - leftStart - columnWidths[0]),
		cells[1],
		resultPad(rightStart - middleStart - columnWidths[1]),
		cells[2],
		resultPad(contentWidth - rightStart - columnWidths[2]),
	}, "")
}

func fullRowColumnWidths(entries []string, columns int, fallbackWidth int) []int {
	if columns < 1 {
		return nil
	}
	if fallbackWidth < 1 {
		fallbackWidth = 1
	}

	widths := make([]int, columns)
	for rowStart := 0; rowStart+columns <= len(entries); rowStart += columns {
		for col := 0; col < columns; col++ {
			widths[col] = max(widths[col], ui.Width(entries[rowStart+col]))
		}
	}
	for col := range widths {
		if widths[col] < 1 {
			widths[col] = fallbackWidth
		}
	}

	return widths
}
