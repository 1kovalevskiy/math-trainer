package app

import (
	"context"
	"errors"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (a *App) initPrograms(ctx context.Context) error {
	if a.mathController == nil {
		return errors.New("math controller is not initialized")
	}

	zone.NewGlobal()
	a.addCloser("bubblezone", func(_ context.Context) error {
		zone.Close()
		return nil
	})

	model := tui.NewModel(ctx, a.mathController)
	a.program = tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	a.addCloser("bubbletea_program", func(_ context.Context) error {
		if a.program != nil {
			a.program.Quit()
		}

		return nil
	})

	return nil
}
