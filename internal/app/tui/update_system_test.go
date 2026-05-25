package tui

import (
	"context"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestHandleSystemMsgQuitsOnCtrlC(t *testing.T) {
	t.Parallel()

	model := NewModel(context.Background(), persistTestController{}, nil)
	_, cmd, handled := model.handleSystemMsg(tea.KeyMsg{Type: tea.KeyCtrlC})

	if !handled {
		t.Fatal("expected ctrl+c to be handled")
	}
	if cmd == nil {
		t.Fatal("expected quit command")
	}
}

func TestHandleSystemMsgStoresWindowSize(t *testing.T) {
	t.Parallel()

	model := NewModel(context.Background(), persistTestController{}, nil)
	updated, _, handled := model.handleSystemMsg(tea.WindowSizeMsg{Width: 100, Height: 30})

	if !handled {
		t.Fatal("expected window size to be handled")
	}
	if updated.width != 100 || updated.height != 30 {
		t.Fatalf("window size mismatch: got %dx%d, want 100x30", updated.width, updated.height)
	}
}
