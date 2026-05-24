package tui

import (
	"context"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	tea "github.com/charmbracelet/bubbletea"
)

type trainingSnapshotMsg struct {
	snapshot mathmodels.TrainingSnapshot
	err      error
}

type cancelTrainingMsg struct {
	err error
}

type persistSettingsMsg struct {
	settings mathmodels.TrainingSettings
	err      error
}

func startTrainingCmd(
	ctx context.Context,
	controller mathController,
	settings mathmodels.TrainingSettings,
) tea.Cmd {
	return func() tea.Msg {
		snapshot, err := controller.StartTraining(ctx, settings)
		return trainingSnapshotMsg{snapshot: snapshot, err: err}
	}
}

func submitAnswerCmd(ctx context.Context, controller mathController, answer string) tea.Cmd {
	return func() tea.Msg {
		snapshot, err := controller.SubmitAnswer(ctx, answer)
		return trainingSnapshotMsg{snapshot: snapshot, err: err}
	}
}

func skipCurrentCmd(ctx context.Context, controller mathController) tea.Cmd {
	return func() tea.Msg {
		snapshot, err := controller.SkipCurrent(ctx)
		return trainingSnapshotMsg{snapshot: snapshot, err: err}
	}
}

func cancelTrainingCmd(ctx context.Context, controller mathController) tea.Cmd {
	return func() tea.Msg {
		return cancelTrainingMsg{err: controller.CancelTraining(ctx)}
	}
}

func persistSettingsCmd(ctx context.Context, store trainingSettingsStore, settings mathmodels.TrainingSettings) tea.Cmd {
	return func() tea.Msg {
		if store == nil {
			return persistSettingsMsg{settings: settings}
		}

		return persistSettingsMsg{settings: settings, err: store.SaveTrainingSettings(ctx, settings)}
	}
}
