package app

import (
	"context"
	"errors"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui"
	config "github.com/1kovalevskiy/math-trainer/internal/configs"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

type configTrainingSettingsStore struct {
	path string
}

func (s configTrainingSettingsStore) SaveTrainingSettings(ctx context.Context, settings tui.TrainingSettings) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return config.SaveTrainingSettings(s.path, settings)
}

func (a *App) initPrograms(ctx context.Context) error {
	if a.mathController == nil {
		return errors.New("math controller is not initialized")
	}

	zone.NewGlobal()
	a.addCloser("bubblezone", func(_ context.Context) error {
		zone.Close()
		return nil
	})

	store := configTrainingSettingsStore{path: a.configPath}
	model := tui.NewModel(ctx, a.mathController, store)
	a.program = tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	a.addCloser("bubbletea_program", func(_ context.Context) error {
		if a.program != nil {
			a.program.Quit()
		}

		return nil
	})

	return nil
}
