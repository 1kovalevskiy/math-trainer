package result

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"

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
