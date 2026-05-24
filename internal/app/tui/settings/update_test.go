package settings

import (
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	tea "github.com/charmbracelet/bubbletea"
)

type testRules struct{}

func (testRules) NormalizeSettings(settings mathmodels.TrainingSettings) mathmodels.TrainingSettings {
	return settings
}

func (testRules) GetNextDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty {
	switch current {
	case mathmodels.DifficultyDisabled:
		return mathmodels.DifficultyEasy
	case mathmodels.DifficultyEasy:
		return mathmodels.DifficultyMedium
	case mathmodels.DifficultyMedium:
		return mathmodels.DifficultyHard
	default:
		return mathmodels.DifficultyDisabled
	}
}

func (testRules) GetPreviousDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty {
	return current
}

func TestUpdateSelectsActionButtonsHorizontally(t *testing.T) {
	model := NewModel(mathmodels.TrainingSettings{}, testRules{})

	for i := 0; i < 5; i++ {
		model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	}
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
	model := NewModel(mathmodels.TrainingSettings{ExamplesCount: 10}, testRules{})

	for i := 0; i < 5; i++ {
		model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	}
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyUp})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})

	if model.settings.ExamplesCount != 11 {
		t.Fatalf("expected examples count to increment, got %d", model.settings.ExamplesCount)
	}
}

func TestUpdateDoesNotMoveActionRowDownToSettingsRows(t *testing.T) {
	model := NewModel(mathmodels.TrainingSettings{ExamplesCount: 10}, testRules{})

	for i := 0; i < 5; i++ {
		model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	}
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})

	if model.cursor != rowApply {
		t.Fatalf("cursor mismatch: got %d, want %d", model.cursor, rowApply)
	}
}

func TestUpdateChangesEachOperationDifficulty(t *testing.T) {
	model := NewModel(mathmodels.TrainingSettings{
		AddDifficulty:      mathmodels.DifficultyEasy,
		SubtractDifficulty: mathmodels.DifficultyEasy,
		MultiplyDifficulty: mathmodels.DifficultyDisabled,
		DivideDifficulty:   mathmodels.DifficultyDisabled,
		ExamplesCount:      10,
	}, testRules{})

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyRight})

	if got, want := model.settings.AddDifficulty, mathmodels.DifficultyMedium; got != want {
		t.Fatalf("add difficulty mismatch: got %q, want %q", got, want)
	}
	if got, want := model.settings.SubtractDifficulty, mathmodels.DifficultyMedium; got != want {
		t.Fatalf("subtract difficulty mismatch: got %q, want %q", got, want)
	}
	if got, want := model.settings.MultiplyDifficulty, mathmodels.DifficultyEasy; got != want {
		t.Fatalf("multiply difficulty mismatch: got %q, want %q", got, want)
	}
	if got, want := model.settings.DivideDifficulty, mathmodels.DifficultyEasy; got != want {
		t.Fatalf("divide difficulty mismatch: got %q, want %q", got, want)
	}
}
