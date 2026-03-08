package tui

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/result"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/settings"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
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
		m.settingsModel = settings.NewModel(m.settings)
		m.screen = ScreenSettings
		return m, nil
	case start.OpenTaskMsg:
		return m.startTraining()
	case start.QuitMsg:
		return m, tea.Quit
	case settings.ApplySettingsMsg:
		m.settings = typedMsg.Settings
		m.screen = ScreenStart
		return m, nil
	case settings.BackMsg:
		m.screen = ScreenStart
		return m, nil
	case task.SubmitMsg:
		m.session.results = append(m.session.results, typedMsg.Result)
		return m.nextExerciseOrFinish()
	case task.SkipMsg:
		m.session.results = append(m.session.results, typedMsg.Result)
		return m.nextExerciseOrFinish()
	case task.BackMsg:
		m.screen = ScreenStart
		return m, nil
	case result.RetryTaskMsg:
		return m.startTraining()
	case result.OpenSettingsMsg:
		m.settingsModel = settings.NewModel(m.settings)
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

func (m Model) startTraining() (Model, tea.Cmd) {
	total := shared.NormalizeExamplesCount(m.settings.ExamplesCount)
	m.settings.ExamplesCount = total
	m.session = trainingSession{
		total:   total,
		results: make([]shared.ExampleResult, 0, total),
	}
	m.taskModel = task.NewModel(m.settings.Difficulty, 1, total)
	m.screen = ScreenTask

	return m, m.taskModel.Init()
}

func (m Model) nextExerciseOrFinish() (Model, tea.Cmd) {
	if len(m.session.results) >= m.session.total {
		m.resultModel = result.NewModel().WithSummary(result.Summary{
			Difficulty: m.settings.Difficulty,
			Results:    m.session.results,
			Correct:    countCorrect(m.session.results),
		})
		m.screen = ScreenResult
		return m, nil
	}

	nextOrder := len(m.session.results) + 1
	m.taskModel = task.NewModel(m.settings.Difficulty, nextOrder, m.session.total)
	m.screen = ScreenTask

	return m, m.taskModel.Init()
}

func countCorrect(results []shared.ExampleResult) int {
	total := 0
	for _, entry := range results {
		if entry.Status == shared.ResultStatusCorrect {
			total++
		}
	}

	return total
}
