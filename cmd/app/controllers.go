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

	a.mathController = mathController.New(a.mathStorage)
	return nil
}
