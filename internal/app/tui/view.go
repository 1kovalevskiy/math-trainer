package tui

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	zone "github.com/lrstanley/bubblezone"
)

func (m Model) View() string {
	frame := newPanelFrame(m.width, m.height)
	if !frame.hasWindowSize {
		content := m.viewCurrentScreen(frame.contentWidth, frame.contentPanelHeight)
		return zone.Scan(ui.Panel.Render(centerBlock(content, frame.contentWidth)))
	}

	chrome := screenChrome(m.screen)
	contentHeight := contentHeightForScreen(chrome.hints, frame.contentPanelHeight)
	content := m.viewCurrentScreen(frame.contentWidth, contentHeight)
	if chrome.fitContent {
		content = renderScreenContent(content, chrome.hints, frame.contentWidth, frame.contentPanelHeight)
	} else {
		content = renderScreenContentNoFit(content, chrome.hints, frame.contentWidth, frame.contentPanelHeight)
	}

	return zone.Scan(ui.Panel.Width(frame.panelWidth).Height(frame.panelHeight).Render(content))
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
