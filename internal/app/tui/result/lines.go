package result

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
