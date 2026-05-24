package tui

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/result"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/settings"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/start"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/task"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	"github.com/charmbracelet/lipgloss"
)

const (
	hintAreaHeight    = 2
	hintBottomPadding = 1
)

type screenChromeConfig struct {
	hints      []string
	fitContent bool
}

func renderScreenContent(content string, hints []string, width int, height int) string {
	if width < 1 {
		width = 1
	}
	if height < 1 {
		height = 1
	}

	hintArea := renderHintArea(hints, width)
	hintHeight := hintAreaHeight
	bottomPadding := hintBottomPadding
	if strings.TrimSpace(hintArea) == "" {
		hintHeight = 0
		bottomPadding = 0
	}

	contentHeight := height - hintHeight - bottomPadding
	if contentHeight < 1 {
		contentHeight = 1
		bottomPadding = 0
	}

	body := lipgloss.Place(
		width,
		contentHeight,
		lipgloss.Center,
		lipgloss.Center,
		fitBlock(centerBlock(content, width), width, contentHeight),
	)
	if hintHeight == 0 {
		return body
	}

	return strings.Join([]string{
		body,
		hintArea,
		lipgloss.NewStyle().Width(width).Render(""),
	}, "\n")
}

func renderScreenContentNoFit(content string, hints []string, width int, height int) string {
	if width < 1 {
		width = 1
	}
	if height < 1 {
		height = 1
	}

	hintArea := renderHintArea(hints, width)
	hintHeight := hintAreaHeight
	bottomPadding := hintBottomPadding
	if strings.TrimSpace(hintArea) == "" {
		hintHeight = 0
		bottomPadding = 0
	}

	contentHeight := height - hintHeight - bottomPadding
	if contentHeight < 1 {
		contentHeight = 1
		bottomPadding = 0
	}

	body := lipgloss.Place(
		width,
		contentHeight,
		lipgloss.Left,
		lipgloss.Top,
		content,
	)
	if hintHeight == 0 {
		return body
	}

	return strings.Join([]string{
		body,
		hintArea,
		lipgloss.NewStyle().Width(width).Render(""),
	}, "\n")
}

func contentHeightForScreen(hints []string, panelHeight int) int {
	hintHeight := 0
	bottomPadding := 0
	if len(hints) > 0 {
		hintHeight = hintAreaHeight
		bottomPadding = hintBottomPadding
	}
	contentHeight := panelHeight - hintHeight - bottomPadding
	if contentHeight < 1 {
		contentHeight = 1
	}
	return contentHeight
}

func renderHintArea(hints []string, width int) string {
	return lipgloss.Place(
		width,
		hintAreaHeight,
		lipgloss.Center,
		lipgloss.Bottom,
		renderHintBlock(hints, width),
	)
}

func renderHintBlock(hints []string, width int) string {
	if len(hints) == 0 {
		return ""
	}

	lines := make([]string, 0, len(hints))
	for _, hint := range hints {
		if hint == "" {
			continue
		}
		lines = append(lines, ui.Hint.Width(width).Align(lipgloss.Center).Render(hint))
	}
	if len(lines) == 0 {
		return ""
	}

	return strings.Join(lines, "\n")
}

func fitBlock(content string, width int, height int) string {
	if height < 1 {
		height = 1
	}

	lines := strings.Split(content, "\n")
	if len(lines) <= height {
		return content
	}
	if height == 1 {
		return lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Render("…")
	}

	visible := height - 1
	top := visible / 2
	bottom := visible - top

	fitted := make([]string, 0, height)
	fitted = append(fitted, lines[:top]...)
	fitted = append(fitted, lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Render(
		fmt.Sprintf("… еще %d строк …", len(lines)-visible),
	))
	fitted = append(fitted, lines[len(lines)-bottom:]...)

	return strings.Join(fitted, "\n")
}

func centerBlock(content string, width int) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Render(line)
	}

	return strings.Join(lines, "\n")
}

func screenHints(screen Screen) []string {
	return screenChrome(screen).hints
}

func screenChrome(screen Screen) screenChromeConfig {
	switch screen {
	case ScreenStart:
		return screenChromeConfig{
			hints:      start.Hints(),
			fitContent: true,
		}
	case ScreenSettings:
		return screenChromeConfig{
			hints:      settings.Hints(),
			fitContent: true,
		}
	case ScreenTask:
		return screenChromeConfig{
			hints:      task.Hints(),
			fitContent: true,
		}
	case ScreenResult:
		return screenChromeConfig{
			hints:      result.Hints(),
			fitContent: false,
		}
	default:
		return screenChromeConfig{fitContent: true}
	}
}
