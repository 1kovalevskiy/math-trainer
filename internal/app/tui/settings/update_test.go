package settings

import (
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	tea "github.com/charmbracelet/bubbletea"
)

func TestUpdateSelectsActionButtonsHorizontally(t *testing.T) {
	model := NewModel(mathmodels.TrainingSettings{}, nil)

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})

	_, cmd := model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected command")
	}
	if _, ok := cmd().(BackMsg); !ok {
		t.Fatalf("expected BackMsg, got %T", cmd())
	}
}

func TestUpdateReturnsFromActionRowToSettingsRow(t *testing.T) {
	model := NewModel(mathmodels.TrainingSettings{}, nil)

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyUp})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})

	if model.settings.ExamplesCount != 1 {
		t.Fatalf("expected examples count to increment, got %d", model.settings.ExamplesCount)
	}
}
