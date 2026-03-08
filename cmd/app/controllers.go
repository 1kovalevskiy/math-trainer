package app

import (
	"context"

	mathController "github.com/1kovalevskiy/math-trainer/internal/controllers/math"
)

func (a *App) initControllers(_ context.Context) error {
	a.mathController = mathController.New()
	return nil
}
