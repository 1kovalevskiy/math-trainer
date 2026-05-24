package e2e

import (
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestIntegrationResultFlow_RetryStartsFreshTraining(t *testing.T) {
	s := newSession(t,
		exercise(4, 4, mathmodels.OperatorAdd),
		exercise(9, 2, mathmodels.OperatorSubtract),
	)

	s.applyExamplesCount(t, 1)
	s.key(t, "enter")
	s.eventuallyViewContains(t, "4 + 4 = ?")
	s.typeText(t, "8")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Результаты тренировки")

	s.key(t, "enter")
	s.eventuallyViewContains(t, "Математическое задание")
	s.eventuallyViewContains(t, "Пример 1 из 1")
	s.eventuallyViewContains(t, "9 - 2 = ?")

	state := s.requireState(t)
	if got := len(state.Results); got != 0 {
		t.Fatalf("new retry session has old results: got %d", got)
	}
}
