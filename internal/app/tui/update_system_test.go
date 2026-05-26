package tui

import (
	"context"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
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

func TestHandleSystemMsgOpensRepositoryFooterOnStartScreenClick(t *testing.T) {
	oldOpenRepositoryURLCmd := openRepositoryURLCmd
	defer func() {
		openRepositoryURLCmd = oldOpenRepositoryURLCmd
	}()

	called := false
	openRepositoryURLCmd = func() tea.Cmd {
		return func() tea.Msg {
			called = true
			return repositoryOpenMsg{}
		}
	}

	zone.Scan(renderScreenContent("Main", []string{"Hint"}, 42, 9, true, false))
	footerZone := waitForZone(t, repositoryFooterZoneID)
	if footerZone == nil {
		t.Fatal("expected repository footer zone")
	}

	model := NewModel(context.Background(), persistTestController{}, nil)
	_, cmd, handled := model.handleSystemMsg(tea.MouseMsg{
		X:    footerZone.StartX,
		Y:    footerZone.StartY,
		Type: tea.MouseLeft,
	})

	if !handled {
		t.Fatal("expected footer click to be handled")
	}
	if cmd == nil {
		t.Fatal("expected footer click to return command")
	}
	cmd()
	if !called {
		t.Fatal("expected repository open command to be called")
	}
}

func TestHandleSystemMsgIgnoresRepositoryFooterClickOutsideStartScreen(t *testing.T) {
	zone.Scan(renderScreenContent("Main", []string{"Hint"}, 42, 9, true, false))
	footerZone := waitForZone(t, repositoryFooterZoneID)
	if footerZone == nil {
		t.Fatal("expected repository footer zone")
	}

	model := NewModel(context.Background(), persistTestController{}, nil)
	model.screen = ScreenSettings
	_, cmd, handled := model.handleSystemMsg(tea.MouseMsg{
		X:    footerZone.StartX,
		Y:    footerZone.StartY,
		Type: tea.MouseLeft,
	})

	if handled {
		t.Fatal("expected footer click outside start screen to be ignored")
	}
	if cmd != nil {
		t.Fatal("expected no command outside start screen")
	}
}

func TestHandleSystemMsgTracksRepositoryFooterHover(t *testing.T) {
	zone.Scan(renderScreenContent("Main", []string{"Hint"}, 42, 9, true, false))
	footerZone := waitForZone(t, repositoryFooterZoneID)
	if footerZone == nil {
		t.Fatal("expected repository footer zone")
	}

	model := NewModel(context.Background(), persistTestController{}, nil)
	updated, _, handled := model.handleSystemMsg(tea.MouseMsg{
		X:      footerZone.StartX,
		Y:      footerZone.StartY,
		Type:   tea.MouseMotion,
		Action: tea.MouseActionMotion,
	})

	if !handled {
		t.Fatal("expected footer hover to be handled")
	}
	if !updated.repositoryFooterHovered {
		t.Fatal("expected repository footer hover state")
	}

	updated, _, handled = updated.handleSystemMsg(tea.MouseMsg{
		X:      footerZone.EndX + 1,
		Y:      footerZone.StartY,
		Type:   tea.MouseMotion,
		Action: tea.MouseActionMotion,
	})

	if !handled {
		t.Fatal("expected footer hover exit to be handled")
	}
	if updated.repositoryFooterHovered {
		t.Fatal("expected repository footer hover state to be cleared")
	}
}
