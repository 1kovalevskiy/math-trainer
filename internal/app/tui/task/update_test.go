package task

import (
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	tea "github.com/charmbracelet/bubbletea"
)

func TestUpdateSelectsTaskButtonsHorizontally(t *testing.T) {
	model := NewModel(&mathmodels.CurrentExercise{}, mathmodels.DifficultyEasy)

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})

	_, cmd := model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected command")
	}
	if _, ok := cmd().(SkipMsg); !ok {
		t.Fatalf("expected SkipMsg, got %T", cmd())
	}
}

func TestUpdateSelectsBackButtonHorizontally(t *testing.T) {
	model := NewModel(&mathmodels.CurrentExercise{}, mathmodels.DifficultyEasy)

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})

	_, cmd := model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected command")
	}
	if _, ok := cmd().(BackMsg); !ok {
		t.Fatalf("expected BackMsg, got %T", cmd())
	}
}

func TestUpdateIgnoresVerticalArrowsForTaskButtons(t *testing.T) {
	model := NewModel(&mathmodels.CurrentExercise{}, mathmodels.DifficultyEasy)

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})

	_, cmd := model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd != nil {
		t.Fatalf("expected no command because submit is still selected without input, got %T", cmd())
	}
}
