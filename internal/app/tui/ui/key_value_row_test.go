package ui

import (
	"strings"
	"testing"
)

func TestRenderKeyValueRowKeepsStableWidthAcrossActiveStates(t *testing.T) {
	t.Parallel()

	layout := KeyValueRowLayout{LabelWidth: 8, ValueWidth: 5, RowWidth: 24}
	inactive := RenderKeyValueRow(KeyValueRow{
		Label:  "Param",
		Value:  []Segment{{Text: "1", Style: SettingRowStyle(false)}},
		Layout: layout,
	})
	active := RenderKeyValueRow(KeyValueRow{
		Active: true,
		Label:  "Param",
		Value:  []Segment{{Text: "1", Style: SettingRowStyle(true)}},
		Layout: layout,
	})

	if Width(inactive) != Width(active) {
		t.Fatalf("row width mismatch: inactive %d, active %d", Width(inactive), Width(active))
	}
	if Width(active) != layout.RowWidth {
		t.Fatalf("active width mismatch: got %d, want %d", Width(active), layout.RowWidth)
	}
	if strings.Contains(active, "\n") {
		t.Fatalf("active row should be a single line: %q", active)
	}
}
