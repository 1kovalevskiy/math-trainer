package ui

import "strings"

func JoinInline(items []string, gap int) string {
	if len(items) == 0 {
		return ""
	}
	if gap < 0 {
		gap = 0
	}

	return strings.Join(items, strings.Repeat(" ", gap))
}

func MaxButtonWidth(labels []string) int {
	width := 1
	for _, label := range labels {
		if current := Width(Button(label, false)); current > width {
			width = current
		}
	}

	return width
}

func StretchTwoButtonWidths(totalLineWidth int, leftText string, leftActive bool, rightText string, rightActive bool) (int, int) {
	leftWidth := Width(leftText)
	rightWidth := Width(rightText)

	for {
		line := ButtonFixed(leftText, leftActive, leftWidth) + " " + ButtonFixed(rightText, rightActive, rightWidth)
		if Width(line) >= totalLineWidth {
			break
		}
		if leftWidth <= rightWidth {
			leftWidth++
		} else {
			rightWidth++
		}
	}

	return leftWidth, rightWidth
}
