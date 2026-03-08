package tui

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/result"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/settings"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/start"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/task"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	screen   Screen
	settings shared.TrainingSettings
	width    int
	height   int
	session  trainingSession

	startModel    start.Model
	settingsModel settings.Model
	taskModel     task.Model
	resultModel   result.Model
}

type trainingSession struct {
	total   int
	results []shared.ExampleResult
}

func NewModel() Model {
	defaultSettings := shared.DefaultTrainingSettings()

	return Model{
		screen:        ScreenStart,
		settings:      defaultSettings,
		session:       trainingSession{total: defaultSettings.ExamplesCount},
		startModel:    start.NewModel(),
		settingsModel: settings.NewModel(defaultSettings),
		taskModel:     task.NewModel(defaultSettings.Difficulty, 1, defaultSettings.ExamplesCount),
		resultModel:   result.NewModel(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
