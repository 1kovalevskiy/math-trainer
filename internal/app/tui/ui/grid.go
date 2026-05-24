package ui

import "math"

type GridOptions struct {
	MaxColumns       int
	PreferredColumns int
	ColumnGap        int
	ScrollbarWidth   int
}

type GridLayout struct {
	Columns       int
	Rows          int
	ContentRows   int
	ContentWidth  int
	ShowScrollbar bool
}

func BuildGridLayout(width int, viewportRows int, totalItems int, itemWidth int, options GridOptions) GridLayout {
	if viewportRows < 1 {
		viewportRows = 1
	}
	if totalItems < 0 {
		totalItems = 0
	}
	if itemWidth < 1 {
		itemWidth = 1
	}
	if options.ColumnGap < 0 {
		options.ColumnGap = 0
	}
	if options.ScrollbarWidth < 0 {
		options.ScrollbarWidth = 0
	}
	if options.MaxColumns < 1 {
		options.MaxColumns = 1
	}
	if options.PreferredColumns < 1 {
		options.PreferredColumns = options.MaxColumns
	}

	columns := maxGridColumns(width, itemWidth, options.ColumnGap)
	columns = min(columns, options.MaxColumns, options.PreferredColumns)
	if columns < 1 {
		columns = 1
	}

	rows := ceilDiv(totalItems, columns)
	if rows <= viewportRows {
		return GridLayout{
			Columns:      columns,
			Rows:         rows,
			ContentRows:  rows,
			ContentWidth: gridContentWidth(columns, itemWidth, options.ColumnGap),
		}
	}

	if options.ScrollbarWidth > 0 {
		available := width - options.ScrollbarWidth
		adjustedColumns := maxGridColumns(available, itemWidth, options.ColumnGap)
		adjustedColumns = min(adjustedColumns, options.MaxColumns, options.PreferredColumns)
		if adjustedColumns < 1 {
			adjustedColumns = 1
		}
		if adjustedColumns < columns {
			columns = adjustedColumns
			rows = ceilDiv(totalItems, columns)
		}
	}

	return GridLayout{
		Columns:       columns,
		Rows:          rows,
		ContentRows:   rows,
		ContentWidth:  gridContentWidth(columns, itemWidth, options.ColumnGap),
		ShowScrollbar: true,
	}
}

func maxGridColumns(width int, itemWidth int, gap int) int {
	if width < itemWidth {
		return 1
	}

	return (width + gap) / (itemWidth + gap)
}

func gridContentWidth(columns int, itemWidth int, gap int) int {
	if columns < 1 {
		return 0
	}

	return columns*itemWidth + (columns-1)*gap
}

func ceilDiv(a int, b int) int {
	if b < 1 {
		return 0
	}

	return int(math.Ceil(float64(a) / float64(b)))
}

func min(values ...int) int {
	if len(values) == 0 {
		return 0
	}
	current := values[0]
	for _, value := range values[1:] {
		if value < current {
			current = value
		}
	}

	return current
}
