package result

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch typedMsg := msg.(type) {
	case tea.MouseMsg:
		m.refreshScrollBounds()
		if shared.IsWheelUp(typedMsg) {
			m.scrollOffset--
			m.clampScroll()
			return m, nil
		}
		if shared.IsWheelDown(typedMsg) {
			m.scrollOffset++
			m.clampScroll()
			return m, nil
		}
		if !shared.IsLeftClick(typedMsg) {
			return m, nil
		}

		for i := range m.options {
			if shared.InZone(optionZoneID(i), typedMsg) {
				m = m.withCursor(i)
				return m, m.selectCurrent()
			}
		}
		return m, nil
	case tea.KeyMsg:
		m.refreshScrollBounds()
		switch typedMsg.String() {
		case "left", "h":
			m = m.moveCursor(-1)
		case "right", "l":
			m = m.moveCursor(1)
		case "up", "k":
			m.scrollOffset--
			m.clampScroll()
		case "down", "j":
			m.scrollOffset++
			m.clampScroll()
		case "pgup":
			m.scrollOffset -= m.pageSize()
			m.clampScroll()
		case "pgdown":
			m.scrollOffset += m.pageSize()
			m.clampScroll()
		case "home":
			m.scrollOffset = 0
		case "end":
			m.scrollOffset = m.maxScrollOffset()
		case "enter":
			return m, m.selectCurrent()
		}
	}

	return m, nil
}

func (m Model) pageSize() int {
	return ui.ScrollState{ViewportRows: m.lastViewportHeight}.PageSize()
}

func (m Model) selectCurrent() tea.Cmd {
	switch m.selectedOption() {
	case OptionRetry:
		return emit(RetryTaskMsg{})
	case OptionSettings:
		return emit(OpenSettingsMsg{})
	case OptionHome:
		return emit(BackToStartMsg{})
	default:
		return nil
	}
}
