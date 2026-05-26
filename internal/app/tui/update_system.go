package tui

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) handleSystemMsg(msg tea.Msg) (Model, tea.Cmd, bool) {
	switch typedMsg := msg.(type) {
	case repositoryOpenMsg:
		return m, nil, true
	case tea.WindowSizeMsg:
		m.width = typedMsg.Width
		m.height = typedMsg.Height
		if m.screen == ScreenResult {
			m.resultModel = m.resultModelWithCurrentViewport(m.resultModel)
		}
		return m, nil, true
	case tea.MouseMsg:
		if m.screen != ScreenStart {
			return m, nil, false
		}

		inRepositoryFooter := shared.InZone(repositoryFooterZoneID, typedMsg)
		if inRepositoryFooter != m.repositoryFooterHovered {
			m.repositoryFooterHovered = inRepositoryFooter
			if !shared.IsLeftClick(typedMsg) {
				return m, nil, true
			}
		}
		if shared.IsLeftClick(typedMsg) && inRepositoryFooter {
			return m, openRepositoryURLCmd(), true
		}
	case tea.KeyMsg:
		if typedMsg.String() == "ctrl+c" {
			return m, tea.Quit, true
		}
	}

	return m, nil, false
}
