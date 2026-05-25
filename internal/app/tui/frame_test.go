package tui

import (
	"testing"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
)

func TestNewPanelFrameNormalizesTerminalAndContentSizes(t *testing.T) {
	t.Parallel()

	frame := newPanelFrame(90, 20)

	if got, want := frame.panelWidth, 88; got != want {
		t.Fatalf("panel width mismatch: got %d, want %d", got, want)
	}
	if got, want := frame.panelHeight, 18; got != want {
		t.Fatalf("panel height mismatch: got %d, want %d", got, want)
	}
	if got, want := frame.contentWidth, 84; got != want {
		t.Fatalf("content width mismatch: got %d, want %d", got, want)
	}
	if got, want := frame.contentPanelHeight, 16; got != want {
		t.Fatalf("content panel height mismatch: got %d, want %d", got, want)
	}
}

func TestNewPanelFrameFallsBackBeforeWindowSize(t *testing.T) {
	t.Parallel()

	frame := newPanelFrame(0, 0)

	if got, want := frame.contentWidth, ui.MinPanelContentWidth; got != want {
		t.Fatalf("content width mismatch: got %d, want %d", got, want)
	}
	if got, want := frame.contentPanelHeight, ui.MinPanelContentWidth; got != want {
		t.Fatalf("content panel height mismatch: got %d, want %d", got, want)
	}
}
