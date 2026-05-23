package mathmemory

import (
	"context"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func (s *Storage) GetState(ctx context.Context) (mathmodels.TrainingState, error) {
	if err := ctx.Err(); err != nil {
		return mathmodels.TrainingState{}, err
	}
	if s == nil {
		return mathmodels.TrainingState{}, mathmodels.ErrNoActiveTraining
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	if !s.hasState {
		return mathmodels.TrainingState{}, mathmodels.ErrNoActiveTraining
	}

	return cloneState(s.state), nil
}
