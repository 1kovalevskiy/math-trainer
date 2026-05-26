package tui

import (
	"context"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/result"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/settings"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/start"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/task"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	tea "github.com/charmbracelet/bubbletea"
)

type TrainingSettings = mathmodels.TrainingSettings

type trainingSettingsStore interface {
	SaveTrainingSettings(ctx context.Context, settings TrainingSettings) error
}

type Model struct {
	ctx            context.Context
	mathController mathController
	settingsStore  trainingSettingsStore
	screen         Screen
	settings       mathmodels.TrainingSettings
	width          int
	height         int

	repositoryFooterHovered bool
	startModel              start.Model
	settingsModel           settings.Model
	taskModel               task.Model
	resultModel             result.Model
}

func NewModel(ctx context.Context, mathController mathController, settingsStore trainingSettingsStore) Model {
	defaultSettings := mathmodels.DefaultTrainingSettings()
	if mathController != nil {
		defaultSettings = mathController.GetDefaultSettings()
	}

	return Model{
		ctx:            ctx,
		mathController: mathController,
		settingsStore:  settingsStore,
		screen:         ScreenStart,
		settings:       defaultSettings,
		startModel:     start.NewModel(),
		settingsModel:  settings.NewModel(defaultSettings, mathController),
		taskModel:      task.NewModel(nil, defaultSettings),
		resultModel:    result.NewModel(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
