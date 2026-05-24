package tui

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	zone "github.com/lrstanley/bubblezone"
)

func (m Model) View() string {
	var content string

	switch m.screen {
	case ScreenStart:
		content = m.startModel.View()
	case ScreenSettings:
		content = m.settingsModel.View()
	case ScreenTask:
		content = m.taskModel.View()
	case ScreenResult:
		content = m.resultModel.View()
	default:
		content = "Неизвестный экран"
	}

	if m.width <= 0 || m.height <= 0 {
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

	if m.screen == ScreenResult {
		content = m.resultModel.ViewWithSize(panelWidth, contentHeightForScreen(screenHints(m.screen), panelHeight))
		content = renderScreenContentNoFit(content, screenHints(m.screen), panelWidth, panelHeight)
		return zone.Scan(ui.Panel.Width(panelWidth).Height(panelHeight).Render(content))
	}

	content = renderScreenContent(content, screenHints(m.screen), panelWidth, panelHeight)

	return zone.Scan(ui.Panel.Width(panelWidth).Height(panelHeight).Render(content))
}
