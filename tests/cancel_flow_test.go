package e2e

import (
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestIntegrationCancelFlow_BackFromTaskClearsTraining(t *testing.T) {
	s := newSession(t,
		exercise(6, 1, mathmodels.OperatorSubtract),
	)

	s.key(t, "enter")
	s.eventuallyViewContains(t, "Математическое задание")

	s.key(t, "esc")
	s.eventuallyViewContains(t, "Математический тренажер")
	s.requireNoActiveTraining(t)
}
