package tui

import (
	"context"
	"testing"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/start"
)

func TestHandleScreenMsgOpensSettingsFromStart(t *testing.T) {
	t.Parallel()

	model := NewModel(context.Background(), persistTestController{}, nil)
	updated, cmd, handled := model.handleScreenMsg(start.OpenSettingsMsg{})

	if !handled {
		t.Fatal("expected start settings message to be handled")
	}
	if cmd != nil {
		t.Fatalf("expected no command, got %T", cmd())
	}
	if updated.screen != ScreenSettings {
		t.Fatalf("screen mismatch: got %v, want %v", updated.screen, ScreenSettings)
	}
}
