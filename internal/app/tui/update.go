package tui

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/result"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/settings"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/start"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/task"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch typedMsg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = typedMsg.Width
		m.height = typedMsg.Height
	case tea.KeyMsg:
		if typedMsg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case start.OpenSettingsMsg:
		m.settingsModel = settings.NewModel(m.difficulty)
		m.screen = ScreenSettings
		return m, nil
	case start.OpenTaskMsg:
		m.taskModel = task.NewModel(m.difficulty)
		m.screen = ScreenTask
		return m, m.taskModel.Init()
	case start.QuitMsg:
		return m, tea.Quit
	case settings.ApplyDifficultyMsg:
		m.difficulty = typedMsg.Difficulty
		m.screen = ScreenStart
		return m, nil
	case settings.BackMsg:
		m.screen = ScreenStart
		return m, nil
	case task.SubmitMsg:
		m.resultModel = result.NewModel().WithOutcome(result.Outcome{
			Difficulty: typedMsg.Difficulty,
			Expression: typedMsg.Expression,
			Expected:   typedMsg.Expected,
			Answer:     typedMsg.Answer,
			Correct:    typedMsg.Correct,
		})
		m.screen = ScreenResult
		return m, nil
	case task.BackMsg:
		m.screen = ScreenStart
		return m, nil
	case result.RetryTaskMsg:
		m.taskModel = task.NewModel(m.difficulty)
		m.screen = ScreenTask
		return m, m.taskModel.Init()
	case result.OpenSettingsMsg:
		m.settingsModel = settings.NewModel(m.difficulty)
		m.screen = ScreenSettings
		return m, nil
	case result.BackToStartMsg:
		m.screen = ScreenStart
		return m, nil
	}

	var cmd tea.Cmd
	switch m.screen {
	case ScreenStart:
		m.startModel, cmd = m.startModel.Update(msg)
	case ScreenSettings:
		m.settingsModel, cmd = m.settingsModel.Update(msg)
	case ScreenTask:
		m.taskModel, cmd = m.taskModel.Update(msg)
	case ScreenResult:
		m.resultModel, cmd = m.resultModel.Update(msg)
	}

	return m, cmd
}
