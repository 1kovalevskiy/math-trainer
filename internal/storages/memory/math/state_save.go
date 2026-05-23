package mathmemory

import (
	"context"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func (s *Storage) SaveState(ctx context.Context, state mathmodels.TrainingState) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if s == nil {
		return mathmodels.ErrNoActiveTraining
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.state = cloneState(state)
	s.hasState = true
	return nil
}
