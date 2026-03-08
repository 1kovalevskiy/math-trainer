package app

import (
	"context"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (a *App) initPrograms() error {
	zone.NewGlobal()
	a.addCloser("bubblezone", func(_ context.Context) error {
		zone.Close()
		return nil
	})

	model := tui.NewModel()
	a.program = tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	a.addCloser("bubbletea_program", func(_ context.Context) error {
		if a.program != nil {
			a.program.Quit()
		}

		return nil
	})

	return nil
}
