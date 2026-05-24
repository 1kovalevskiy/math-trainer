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

type Model struct {
	ctx            context.Context
	mathController mathController
	screen         Screen
	settings       mathmodels.TrainingSettings
	width          int
	height         int

	startModel    start.Model
	settingsModel settings.Model
	taskModel     task.Model
	resultModel   result.Model
}

func NewModel(ctx context.Context, mathController mathController) Model {
	defaultSettings := mathmodels.TrainingSettings{
		AddDifficulty:      mathmodels.DifficultyEasy,
		SubtractDifficulty: mathmodels.DifficultyEasy,
		MultiplyDifficulty: mathmodels.DifficultyDisabled,
		DivideDifficulty:   mathmodels.DifficultyDisabled,
		ExamplesCount:      mathmodels.DefaultExamplesCount,
	}
	if mathController != nil {
		defaultSettings = mathController.GetDefaultSettings()
	}

	return Model{
		ctx:            ctx,
		mathController: mathController,
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
