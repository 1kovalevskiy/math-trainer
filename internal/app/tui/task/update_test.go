package task

import (
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	tea "github.com/charmbracelet/bubbletea"
)

func TestUpdateMovesButtonSelectionHorizontally(t *testing.T) {
	model := NewModel(&mathmodels.CurrentExercise{}, mathmodels.TrainingSettings{})

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})

	if model.buttonCursor != buttonBack {
		t.Fatalf("button cursor mismatch: got %d, want %d", model.buttonCursor, buttonBack)
	}
}

func TestUpdateBackspaceRemovesInputSymbol(t *testing.T) {
	model := NewModel(&mathmodels.CurrentExercise{}, mathmodels.TrainingSettings{})

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'2'}})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyBackspace})

	if got, want := model.input, "1"; got != want {
		t.Fatalf("input mismatch: got %q, want %q", got, want)
	}
}

func TestUpdateEnterOnSubmitWithoutInputDoesNothing(t *testing.T) {
	model := NewModel(&mathmodels.CurrentExercise{}, mathmodels.TrainingSettings{})

	_, cmd := model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd != nil {
		t.Fatalf("unexpected cmd for empty submit: %v", cmd)
	}
}
