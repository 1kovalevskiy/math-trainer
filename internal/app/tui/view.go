package tui

func (m Model) View() string {
	switch m.screen {
	case ScreenStart:
		return m.startModel.View()
	case ScreenSettings:
		return m.settingsModel.View()
	case ScreenTask:
		return m.taskModel.View()
	case ScreenResult:
		return m.resultModel.View()
	default:
		return "Неизвестный экран"
	}
}
