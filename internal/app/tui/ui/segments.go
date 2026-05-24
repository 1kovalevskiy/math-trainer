package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Segment struct {
	Text  string
	Style lipgloss.Style
}

func RenderSegments(segments ...Segment) string {
	var b strings.Builder
	for _, segment := range segments {
		if segment.Text == "" {
			continue
		}
		b.WriteString(segment.Style.Render(segment.Text))
	}

	return b.String()
}

func StyledPad(style lipgloss.Style, width int) string {
	if width <= 0 {
		return ""
	}

	return style.Render(strings.Repeat(" ", width))
}

func StyledPadRight(style lipgloss.Style, content string, width int) string {
	return content + StyledPad(style, width-Width(content))
}
