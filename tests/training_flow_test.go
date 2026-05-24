package e2e

import (
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestIntegrationTrainingFlow_SubmitAndSkipShowsResult(t *testing.T) {
	s := newSession(t,
		exercise(2, 3, mathmodels.OperatorAdd),
		exercise(7, 4, mathmodels.OperatorSubtract),
	)

	s.applyExamplesCount(t, 2)
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Математическое задание")
	s.eventuallyViewContains(t, "Пример 1 из 2")
	s.eventuallyViewContains(t, "2 + 3 = ?")

	s.typeText(t, "5")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Пример 2 из 2")
	s.eventuallyViewContains(t, "7 - 4 = ?")

	s.key(t, "right")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Результаты тренировки")
	s.eventuallyViewContains(t, "Правильных: 1 из 2")
	s.eventuallyViewContains(t, "1) 2 + 3 = 5")
	s.eventuallyViewContains(t, "2) 7 - 4 = ____ (ответ: 3)")

	state := s.requireState(t)
	if got, want := len(state.Results), 2; got != want {
		t.Fatalf("results count mismatch: got %d, want %d", got, want)
	}
	if got, want := state.Results[0].Status, mathmodels.ResultStatusCorrect; got != want {
		t.Fatalf("first result status mismatch: got %q, want %q", got, want)
	}
	if got, want := state.Results[1].Status, mathmodels.ResultStatusSkipped; got != want {
		t.Fatalf("second result status mismatch: got %q, want %q", got, want)
	}
}
