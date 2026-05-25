package tui

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/result"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/task"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if updated, cmd, handled := m.handleSystemMsg(msg); handled {
		return updated, cmd
	}
	if updated, cmd, handled := m.handleScreenMsg(msg); handled {
		return updated, cmd
	}
	if updated, cmd, handled := m.handleCommandResultMsg(msg); handled {
		return updated, cmd
	}

	if updated, cmd, handled := m.updateCurrentScreen(msg); handled {
		return updated, cmd
	}

	return m, nil
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
		m.resultModel = m.resultModelWithCurrentViewport(m.resultModel)
		m.screen = ScreenResult
	default:
		m.screen = ScreenStart
	}

	return m
}

func (m Model) resultModelWithCurrentViewport(resultModel result.Model) result.Model {
	frame := newPanelFrame(m.width, m.height)
	if !frame.hasWindowSize {
		return resultModel
	}

	contentHeight := contentHeightForScreen(screenHints(ScreenResult), frame.contentPanelHeight)
	return resultModel.WithContentSize(frame.contentWidth, contentHeight)
}
