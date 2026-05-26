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
	zone "github.com/lrstanley/bubblezone"
)

const (
	repositoryFooterText   = "github.com/1kovalevskiy/math-trainer"
	repositoryFooterURL    = "https://github.com/1kovalevskiy/math-trainer"
	repositoryFooterZoneID = "footer:repository"
	hintAreaHeight         = 2
	hintBottomPadding      = 1
	footerAreaHeight       = 1
)

var panelSurface = lipgloss.NewStyle().Background(ui.PanelBackgroundColor)

type screenChromeConfig struct {
	hints      []string
	fitContent bool
	footer     bool
}

func renderScreenContent(content string, hints []string, width int, height int, footer bool, footerHovered bool) string {
	if width < 1 {
		width = 1
	}
	if height < 1 {
		height = 1
	}

	hintArea := renderHintArea(hints, width)
	hintHeight := hintAreaHeight
	if strings.TrimSpace(hintArea) == "" {
		hintHeight = 0
	}

	footerArea := ""
	footerHeight := 0
	if footer {
		footerArea = renderFooterArea(width, footerHovered)
		footerHeight = footerAreaHeight
	} else if hintHeight > 0 {
		footerArea = panelPad(width)
		footerHeight = hintBottomPadding
	}

	contentHeight := height - hintHeight - footerHeight
	if contentHeight < 1 {
		contentHeight = 1
	}

	body := renderCenteredBlock(fitBlock(centerBlock(content, width), width, contentHeight), width, contentHeight)
	parts := []string{body}
	if hintHeight > 0 {
		parts = append(parts, hintArea)
	}
	if footerArea != "" {
		parts = append(parts, footerArea)
	}
	return strings.Join(parts, "\n")
}

func renderScreenContentNoFit(content string, hints []string, width int, height int, footer bool, footerHovered bool) string {
	if width < 1 {
		width = 1
	}
	if height < 1 {
		height = 1
	}

	hintArea := renderHintArea(hints, width)
	hintHeight := hintAreaHeight
	if strings.TrimSpace(hintArea) == "" {
		hintHeight = 0
	}

	footerArea := ""
	footerHeight := 0
	if footer {
		footerArea = renderFooterArea(width, footerHovered)
		footerHeight = footerAreaHeight
	} else if hintHeight > 0 {
		footerArea = panelPad(width)
		footerHeight = hintBottomPadding
	}

	contentHeight := height - hintHeight - footerHeight
	if contentHeight < 1 {
		contentHeight = 1
	}

	body := renderFixedBlock(content, width, contentHeight)
	parts := []string{body}
	if hintHeight > 0 {
		parts = append(parts, hintArea)
	}
	if footerArea != "" {
		parts = append(parts, footerArea)
	}
	return strings.Join(parts, "\n")
}

func contentHeightForScreen(hints []string, panelHeight int, footer bool) int {
	hintHeight := 0
	if len(hints) > 0 {
		hintHeight = hintAreaHeight
	}
	footerHeight := 0
	if footer {
		footerHeight = footerAreaHeight
	} else if hintHeight > 0 {
		footerHeight = hintBottomPadding
	}
	contentHeight := panelHeight - hintHeight - footerHeight
	if contentHeight < 1 {
		contentHeight = 1
	}
	return contentHeight
}

func renderHintArea(hints []string, width int) string {
	return renderBottomBlock(renderHintBlock(hints, width), width, hintAreaHeight)
}

func renderFooterArea(width int, hovered bool) string {
	style := ui.Footer
	if hovered {
		style = ui.FooterHover
	}
	return panelPadCenter(zone.Mark(repositoryFooterZoneID, style.Render(repositoryFooterText)), width)
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
		lines = append(lines, panelPadCenter(ui.Hint.Render(hint), width))
	}
	if len(lines) == 0 {
		return ""
	}

	return strings.Join(lines, "\n")
}

func renderFixedBlock(content string, width int, height int) string {
	if height < 1 {
		height = 1
	}

	lines := strings.Split(content, "\n")
	if len(lines) > height {
		lines = lines[:height]
	}
	for len(lines) < height {
		lines = append(lines, "")
	}
	for i, line := range lines {
		lines[i] = panelPadRight(line, width)
	}

	return strings.Join(lines, "\n")
}

func renderCenteredBlock(content string, width int, height int) string {
	if height < 1 {
		height = 1
	}

	lines := strings.Split(content, "\n")
	if len(lines) > height {
		lines = lines[:height]
	}
	centered := make([]string, 0, height)
	top := (height - len(lines)) / 2
	for i := 0; i < top; i++ {
		centered = append(centered, panelPad(width))
	}
	centered = append(centered, lines...)
	for len(centered) < height {
		centered = append(centered, panelPad(width))
	}
	for i, line := range centered {
		centered[i] = panelPadCenter(line, width)
	}

	return strings.Join(centered, "\n")
}

func renderBottomBlock(content string, width int, height int) string {
	if height < 1 {
		height = 1
	}

	lines := strings.Split(content, "\n")
	if len(lines) > height {
		lines = lines[len(lines)-height:]
	}
	block := make([]string, 0, height)
	for len(block)+len(lines) < height {
		block = append(block, panelPad(width))
	}
	block = append(block, lines...)
	for i, line := range block {
		block[i] = panelPadCenter(line, width)
	}

	return strings.Join(block, "\n")
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
		return panelPadCenter("…", width)
	}

	visible := height - 1
	top := visible / 2
	bottom := visible - top

	fitted := make([]string, 0, height)
	fitted = append(fitted, lines[:top]...)
	fitted = append(fitted, panelPadCenter(fmt.Sprintf("… еще %d строк …", len(lines)-visible), width))
	fitted = append(fitted, lines[len(lines)-bottom:]...)

	return strings.Join(fitted, "\n")
}

func centerBlock(content string, width int) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = panelPadCenter(line, width)
	}

	return strings.Join(lines, "\n")
}

func panelPad(width int) string {
	return ui.StyledPad(panelSurface, width)
}

func panelPadRight(content string, width int) string {
	return ui.StyledPadRight(panelSurface, content, width)
}

func panelPadCenter(content string, width int) string {
	return ui.StyledPadCenter(panelSurface, content, width)
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
			footer:     true,
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
