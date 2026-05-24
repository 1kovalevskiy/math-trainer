package tui

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	zone "github.com/lrstanley/bubblezone"
)

func (m Model) View() string {
	if m.width <= 0 || m.height <= 0 {
		content := m.viewCurrentScreen(ui.MinPanelContentWidth, ui.MinPanelContentWidth)
		return zone.Scan(ui.Panel.Render(centerBlock(content, ui.MinPanelContentWidth)))
	}

	panelWidth := m.width - ui.Panel.GetHorizontalFrameSize()
	panelHeight := m.height - ui.Panel.GetVerticalFrameSize()
	if panelWidth < 1 {
		panelWidth = 1
	}
	if panelHeight < 1 {
		panelHeight = 1
	}

	chrome := screenChrome(m.screen)
	contentHeight := contentHeightForScreen(chrome.hints, panelHeight)
	content := m.viewCurrentScreen(panelWidth, contentHeight)
	if chrome.fitContent {
		content = renderScreenContent(content, chrome.hints, panelWidth, panelHeight)
	} else {
		content = renderScreenContentNoFit(content, chrome.hints, panelWidth, panelHeight)
	}

	return zone.Scan(ui.Panel.Width(panelWidth).Height(panelHeight).Render(content))
}

func (m Model) viewCurrentScreen(width int, height int) string {
	switch m.screen {
	case ScreenStart:
		return m.startModel.ViewWithSize(width, height)
	case ScreenSettings:
		return m.settingsModel.ViewWithSize(width, height)
	case ScreenTask:
		return m.taskModel.ViewWithSize(width, height)
	case ScreenResult:
		return m.resultModel.ViewWithSize(width, height)
	default:
		return "Неизвестный экран"
	}
}
