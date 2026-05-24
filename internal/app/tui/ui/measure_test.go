package ui

import "testing"

func TestPadCenterKeepsRequestedVisibleWidth(t *testing.T) {
	t.Parallel()

	got := PadCenter("abc", 7)
	if Width(got) != 7 {
		t.Fatalf("width mismatch: got %d, want 7", Width(got))
	}
}

func TestJoinInlineUsesRequestedGap(t *testing.T) {
	t.Parallel()

	got := JoinInline([]string{"a", "b", "c"}, 2)
	if got != "a  b  c" {
		t.Fatalf("join mismatch: got %q", got)
	}
}
