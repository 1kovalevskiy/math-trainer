package tui

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	"github.com/charmbracelet/lipgloss"
)

const (
	hintAreaHeight    = 2
	hintBottomPadding = 1
)

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
	switch screen {
	case ScreenStart:
		return []string{"↑/↓ - выбор, Enter - подтвердить, Ctrl+C - выход"}
	case ScreenSettings:
		return []string{
			"↑/↓ - строки, ←/→ - изменить значение или выбрать кнопку",
			"Enter - подтвердить, Esc - назад, Click - мышь",
		}
	case ScreenTask:
		return []string{
			"←/→ - выбор кнопки, Enter - подтвердить",
			"S - пропустить, Esc - в меню, Click - мышь",
		}
	case ScreenResult:
		return []string{"←/→ - выбор, Enter - подтвердить"}
	default:
		return nil
	}
}
