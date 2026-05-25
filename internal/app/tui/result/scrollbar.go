package result

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
