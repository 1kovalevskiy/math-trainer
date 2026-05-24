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
	model := NewModel().WithViewport(80, 5)
	model.lastContentRows = 10

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})

	_, cmd := model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected command")
	}
	if _, ok := cmd().(RetryTaskMsg); !ok {
		t.Fatalf("expected RetryTaskMsg, got %T", cmd())
	}
}

func TestUpdateScrollsWithArrowKeysAndClamp(t *testing.T) {
	model := NewModel().WithSummary(testSummary(30)).WithViewport(80, 5)
	model.refreshScrollBounds()

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	if model.scrollOffset != 1 {
		t.Fatalf("expected scroll offset 1, got %d", model.scrollOffset)
	}

	for i := 0; i < 50; i++ {
		model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	}
	if model.scrollOffset != model.maxScrollOffset() {
		t.Fatalf("expected clamp to max offset %d, got %d", model.maxScrollOffset(), model.scrollOffset)
	}

	for i := 0; i < 50; i++ {
		model, _ = model.Update(tea.KeyMsg{Type: tea.KeyUp})
	}
	if model.scrollOffset != 0 {
		t.Fatalf("expected clamp to zero, got %d", model.scrollOffset)
	}
}

func TestUpdatePageAndHomeEndScrolling(t *testing.T) {
	model := NewModel().WithSummary(testSummary(30)).WithViewport(80, 5)
	model.refreshScrollBounds()

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyPgDown})
	if model.scrollOffset < 1 {
		t.Fatalf("expected pgdown to increase offset, got %d", model.scrollOffset)
	}

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyEnd})
	if model.scrollOffset != model.maxScrollOffset() {
		t.Fatalf("expected end to go max offset, got %d", model.scrollOffset)
	}

	model, _ = model.Update(tea.KeyMsg{Type: tea.KeyHome})
	if model.scrollOffset != 0 {
		t.Fatalf("expected home to go zero offset, got %d", model.scrollOffset)
	}
}

func TestUpdateScrollsWithMouseWheel(t *testing.T) {
	model := NewModel().WithSummary(testSummary(30)).WithViewport(80, 5)
	model.refreshScrollBounds()

	model, _ = model.Update(tea.MouseMsg{Button: tea.MouseButtonWheelDown})
	if model.scrollOffset != 1 {
		t.Fatalf("expected wheel down to increment offset, got %d", model.scrollOffset)
	}

	model, _ = model.Update(tea.MouseMsg{Button: tea.MouseButtonWheelUp})
	if model.scrollOffset != 0 {
		t.Fatalf("expected wheel up to decrement offset, got %d", model.scrollOffset)
	}
}
