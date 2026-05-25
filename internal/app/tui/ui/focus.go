package ui

type LinearFocus struct {
	index int
	min   int
	max   int
}

func NewLinearFocus(index int, max int) LinearFocus {
	return LinearFocus{index: MoveIndex(index, 0, 0, max), min: 0, max: max}
}

func (f LinearFocus) Index() int {
	return f.index
}

func (f LinearFocus) Move(delta int) LinearFocus {
	f.index = MoveIndex(f.index, delta, f.min, f.max)
	return f
}
