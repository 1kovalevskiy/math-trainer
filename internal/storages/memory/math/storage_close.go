package mathmemory

import "context"

func (s *Storage) CloseStorage(ctx context.Context) error {
	return s.ClearState(ctx)
}
