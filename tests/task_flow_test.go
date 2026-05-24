package e2e

import (
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestIntegrationTaskFlow_EmptySubmitDoesNotAdvance(t *testing.T) {
	s := newSession(t,
		exercise(8, 2, mathmodels.OperatorSubtract),
	)

	s.key(t, "enter")
	s.eventuallyViewContains(t, "Математическое задание")
	s.eventuallyViewContains(t, "Пример 1 из 10")

	s.key(t, "enter")
	s.eventuallyViewContains(t, "Пример 1 из 10")
	s.eventuallyViewContains(t, "8 - 2 = ?")

	state := s.requireState(t)
	if got := len(state.Results); got != 0 {
		t.Fatalf("empty submit changed results: got %d", got)
	}
}
