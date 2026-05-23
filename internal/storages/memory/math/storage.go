package mathmemory

import (
	"sync"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

type Storage struct {
	mu       sync.RWMutex
	hasState bool
	state    mathmodels.TrainingState
}

func New() *Storage {
	return &Storage{}
}
