package app

import (
	"context"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func (a *App) initPrograms() error {
	model := tui.NewModel()
	a.program = tea.NewProgram(model, tea.WithAltScreen())
	a.addCloser("bubbletea_program", func(_ context.Context) error {
		if a.program != nil {
			a.program.Quit()
		}

		return nil
	})

	return nil
}
