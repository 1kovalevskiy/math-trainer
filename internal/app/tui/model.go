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
	screen     Screen
	difficulty shared.Difficulty
	width      int
	height     int

	startModel    start.Model
	settingsModel settings.Model
	taskModel     task.Model
	resultModel   result.Model
}

func NewModel() Model {
	difficulty := shared.DifficultyEasy

	return Model{
		screen:        ScreenStart,
		difficulty:    difficulty,
		startModel:    start.NewModel(),
		settingsModel: settings.NewModel(difficulty),
		taskModel:     task.NewModel(difficulty),
		resultModel:   result.NewModel(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
