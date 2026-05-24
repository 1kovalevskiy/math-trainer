package ui

import "strings"

type KeyValueRowLayout struct {
	LabelWidth int
	ValueWidth int
	RowWidth   int
}

type KeyValueRow struct {
	Active bool
	Label  string
	Value  []Segment
	Layout KeyValueRowLayout
}

func RenderKeyValueRow(row KeyValueRow) string {
	rowStyle := SettingRowStyle(row.Active)
	leftMarker := rowStyle.Render(" ")
	rightMarker := rowStyle.Render(" ")
	if row.Active {
		leftMarker = SettingRowMarkerStyle().Render("▸")
		rightMarker = SettingRowMarkerStyle().Render("◂")
	}

	label := row.Label + ":"
	labelPadding := maxInt(0, row.Layout.LabelWidth-Width(label))
	value := RenderSegments(row.Value...)
	valuePadding := maxInt(0, row.Layout.ValueWidth-Width(value))
	valuePadLeft := valuePadding / 2
	valuePadRight := valuePadding - valuePadLeft

	content := RenderSegments(
		Segment{Text: "  ", Style: rowStyle},
		Segment{Text: leftMarker},
		Segment{Text: label, Style: rowStyle},
		Segment{Text: strings.Repeat(" ", labelPadding), Style: rowStyle},
		Segment{Text: " ", Style: rowStyle},
		Segment{Text: strings.Repeat(" ", valuePadLeft), Style: rowStyle},
		Segment{Text: value},
		Segment{Text: strings.Repeat(" ", valuePadRight), Style: rowStyle},
		Segment{Text: rightMarker},
		Segment{Text: "  ", Style: rowStyle},
	)

	return StyledPadRight(rowStyle, content, row.Layout.RowWidth)
}

func KeyValueBaseRowWidth(labelWidth int, valueWidth int) int {
	return 2 + 1 + labelWidth + 1 + valueWidth + 1 + 2
}

func maxInt(values ...int) int {
	current := values[0]
	for _, value := range values[1:] {
		if value > current {
			current = value
		}
	}

	return current
}
