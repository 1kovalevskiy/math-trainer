package ui

func MoveIndex(index int, delta int, minValue int, maxValue int) int {
	next := index + delta
	if next < minValue {
		return minValue
	}
	if next > maxValue {
		return maxValue
	}

	return next
}
