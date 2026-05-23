package result

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestUpdateSelectsResultButtonsHorizontally(t *testing.T) {
	model := NewModel()

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})

	_, cmd := model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected command")
	}
	if _, ok := cmd().(OpenSettingsMsg); !ok {
		t.Fatalf("expected OpenSettingsMsg, got %T", cmd())
	}
}

func TestUpdateIgnoresVerticalArrowsForResultButtons(t *testing.T) {
	model := NewModel()

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})

	_, cmd := model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected command")
	}
	if _, ok := cmd().(RetryTaskMsg); !ok {
		t.Fatalf("expected RetryTaskMsg, got %T", cmd())
	}
}
