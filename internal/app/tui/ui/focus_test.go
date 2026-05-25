package ui

import "testing"

func TestLinearFocusClampsMovement(t *testing.T) {
	t.Parallel()

	focus := NewLinearFocus(0, 2)

	if got := focus.Move(-1).Index(); got != 0 {
		t.Fatalf("expected clamp at first index, got %d", got)
	}
	if got := focus.Move(3).Index(); got != 2 {
		t.Fatalf("expected clamp at last index, got %d", got)
	}
}
