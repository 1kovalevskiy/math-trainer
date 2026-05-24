package tui

import (
	"errors"
	"fmt"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/result"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/settings"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/start"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/task"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
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
		m.settingsModel = settings.NewModel(m.settings, m.mathController)
		m.screen = ScreenSettings
		return m, nil
	case start.OpenTaskMsg:
		return m.startTraining()
	case start.QuitMsg:
		return m, tea.Quit
	case settings.ApplySettingsMsg:
		m.settings = m.mathController.NormalizeSettings(typedMsg.Settings)
		m.screen = ScreenStart
		return m, nil
	case settings.BackMsg:
		m.screen = ScreenStart
		return m, nil
	case task.SubmitMsg:
		return m, submitAnswerCmd(m.ctx, m.mathController, typedMsg.Answer)
	case task.SkipMsg:
		return m, skipCurrentCmd(m.ctx, m.mathController)
	case task.BackMsg:
		m.screen = ScreenStart
		return m, cancelTrainingCmd(m.ctx, m.mathController)
	case result.RetryTaskMsg:
		return m.startTraining()
	case result.OpenSettingsMsg:
		m.settingsModel = settings.NewModel(m.settings, m.mathController)
		m.screen = ScreenSettings
		return m, nil
	case result.BackToStartMsg:
		m.screen = ScreenStart
		return m, nil
	case trainingSnapshotMsg:
		if typedMsg.err != nil {
			m.taskModel = m.taskModel.WithError(errorText(typedMsg.err))
			return m, nil
		}
		return m.applySnapshot(typedMsg.snapshot), nil
	case cancelTrainingMsg:
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

func (m Model) startTraining() (Model, tea.Cmd) {
	return m, startTrainingCmd(m.ctx, m.mathController, m.settings)
}

func (m Model) applySnapshot(snapshot mathmodels.TrainingSnapshot) Model {
	m.settings = snapshot.Settings

	switch snapshot.Phase {
	case mathmodels.TrainingPhaseInProgress:
		m.taskModel = task.NewModel(snapshot.Current, snapshot.Settings)
		m.screen = ScreenTask
	case mathmodels.TrainingPhaseFinished:
		m.resultModel = result.NewModel().WithSummary(snapshot.Summary)
		m.screen = ScreenResult
	default:
		m.screen = ScreenStart
	}

	return m
}

func errorText(err error) string {
	if errors.Is(err, mathmodels.ErrInvalidAnswer) {
		return "Ответ должен быть числом"
	}

	return fmt.Sprintf("%v", err)
}
