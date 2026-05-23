package mathmemory

import (
	"context"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func (s *Storage) ClearState(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if s == nil {
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.state = mathmodels.TrainingState{}
	s.hasState = false
	return nil
}
