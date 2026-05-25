package tui

import (
	"context"
	"testing"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/start"
	tea "github.com/charmbracelet/bubbletea"
)

func TestUpdateCurrentScreenDelegatesToStartModel(t *testing.T) {
	t.Parallel()

	model := NewModel(context.Background(), persistTestController{}, nil)
	_, cmd, handled := model.updateCurrentScreen(tea.KeyMsg{Type: tea.KeyEnter})

	if !handled {
		t.Fatal("expected start screen update to be handled")
	}
	if cmd == nil {
		t.Fatal("expected command from start submodel")
	}
	if _, ok := cmd().(start.OpenTaskMsg); !ok {
		t.Fatalf("expected OpenTaskMsg, got %T", cmd())
	}
}
