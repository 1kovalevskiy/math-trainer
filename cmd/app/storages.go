package app

import (
	"context"

	mathMemory "github.com/1kovalevskiy/math-trainer/internal/storages/memory/math"
)

func (a *App) initStorages(_ context.Context) error {
	a.mathStorage = mathMemory.New()
	a.addCloser("math_memory_storage", a.mathStorage.CloseStorage)

	return nil
}
