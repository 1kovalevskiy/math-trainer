package app

import (
	"context"
	"errors"

	mathController "github.com/1kovalevskiy/math-trainer/internal/controllers/math"
)

func (a *App) initControllers(_ context.Context) error {
	if a.mathStorage == nil {
		return errors.New("math storage is not initialized")
	}
	if a.cfg == nil {
		return errors.New("config is not initialized")
	}

	a.mathController = mathController.New(
		a.mathStorage,
		mathController.WithDefaultSettings(a.cfg.Training.ToMath()),
	)
	return nil
}
