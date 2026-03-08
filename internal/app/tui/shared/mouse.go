package shared

import (
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func IsLeftClick(msg tea.MouseMsg) bool {
	if msg.Type == tea.MouseLeft {
		return true
	}

	switch msg.Action {
	case tea.MouseActionPress:
		return msg.Button == tea.MouseButtonLeft
	default:
		return false
	}
}

func InZone(id string, msg tea.MouseMsg) bool {
	zoneInfo := zone.Get(id)
	return zoneInfo != nil && zoneInfo.InBounds(msg)
}
