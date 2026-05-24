package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Width(text string) int {
	return lipgloss.Width(text)
}

func PadRight(text string, width int) string {
	padding := width - Width(text)
	if padding <= 0 {
		return text
	}

	return text + strings.Repeat(" ", padding)
}

func PadCenter(text string, width int) string {
	padding := width - Width(text)
	if padding <= 0 {
		return text
	}

	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
}

func MaxWidth(values ...string) int {
	width := 0
	for _, value := range values {
		if current := Width(value); current > width {
			width = current
		}
	}

	return width
}
