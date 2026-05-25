package tui

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/result"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/settings"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/start"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/task"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) handleScreenMsg(msg tea.Msg) (Model, tea.Cmd, bool) {
	switch msg.(type) {
	case start.OpenSettingsMsg:
		m.settingsModel = settings.NewModel(m.settings, m.mathController)
		m.screen = ScreenSettings
		return m, nil, true
	case start.OpenTaskMsg:
		next, cmd := m.startTraining()
		return next, cmd, true
	case start.QuitMsg:
		return m, tea.Quit, true
	case settings.ApplySettingsMsg:
		typedMsg := msg.(settings.ApplySettingsMsg)
		normalized := m.mathController.NormalizeSettings(typedMsg.Settings)
		return m, persistSettingsCmd(m.ctx, m.settingsStore, normalized), true
	case settings.BackMsg:
		m.screen = ScreenStart
		return m, nil, true
	case task.SubmitMsg:
		typedMsg := msg.(task.SubmitMsg)
		return m, submitAnswerCmd(m.ctx, m.mathController, typedMsg.Answer), true
	case task.SkipMsg:
		return m, skipCurrentCmd(m.ctx, m.mathController), true
	case task.BackMsg:
		m.screen = ScreenStart
		return m, cancelTrainingCmd(m.ctx, m.mathController), true
	case result.RetryTaskMsg:
		next, cmd := m.startTraining()
		return next, cmd, true
	case result.OpenSettingsMsg:
		m.settingsModel = settings.NewModel(m.settings, m.mathController)
		m.screen = ScreenSettings
		return m, nil, true
	case result.BackToStartMsg:
		m.screen = ScreenStart
		return m, nil, true
	default:
		return m, nil, false
	}
}
