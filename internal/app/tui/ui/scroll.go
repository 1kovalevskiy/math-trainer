package ui

type ScrollState struct {
	Offset       int
	ViewportRows int
	ContentRows  int
}

func (s ScrollState) MaxOffset() int {
	if s.ViewportRows < 1 || s.ContentRows <= s.ViewportRows {
		return 0
	}

	return s.ContentRows - s.ViewportRows
}

func (s ScrollState) ClampOffset(offset int) int {
	if offset < 0 {
		return 0
	}
	maxOffset := s.MaxOffset()
	if offset > maxOffset {
		return maxOffset
	}

	return offset
}

func (s ScrollState) PageSize() int {
	if s.ViewportRows <= 1 {
		return 1
	}

	return s.ViewportRows - 1
}
