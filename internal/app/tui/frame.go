package tui

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"

type panelFrame struct {
	panelWidth         int
	panelHeight        int
	contentWidth       int
	contentPanelHeight int
	hasWindowSize      bool
}

func newPanelFrame(width int, height int) panelFrame {
	if width <= 0 || height <= 0 {
		return panelFrame{
			contentWidth:       ui.MinPanelContentWidth,
			contentPanelHeight: ui.MinPanelContentWidth,
		}
	}

	panelWidth := width - ui.Panel.GetHorizontalBorderSize()
	panelHeight := height - ui.Panel.GetVerticalBorderSize()
	if panelWidth < 1 {
		panelWidth = 1
	}
	if panelHeight < 1 {
		panelHeight = 1
	}

	contentWidth := panelWidth - ui.Panel.GetHorizontalPadding()
	contentPanelHeight := panelHeight - ui.Panel.GetVerticalPadding()
	if contentWidth < 1 {
		contentWidth = 1
	}
	if contentPanelHeight < 1 {
		contentPanelHeight = 1
	}

	return panelFrame{
		panelWidth:         panelWidth,
		panelHeight:        panelHeight,
		contentWidth:       contentWidth,
		contentPanelHeight: contentPanelHeight,
		hasWindowSize:      true,
	}
}
