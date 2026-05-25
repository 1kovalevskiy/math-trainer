package e2e

import (
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestIntegrationTrainingFlow_DivisionOnly(t *testing.T) {
	s := newSession(t,
		exercise(12, 3, mathmodels.OperatorDivide),
	)

	s.key(t, "down")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Настройки тренировки")

	// Add: starter -> easy -> medium -> hard -> expert -> disabled
	for i := 0; i < 5; i++ {
		s.key(t, "right")
	}
	// Subtract: starter -> easy -> medium -> hard -> expert -> disabled
	s.key(t, "down")
	for i := 0; i < 5; i++ {
		s.key(t, "right")
	}
	// Multiply remains disabled, move to divide and enable starter
	s.key(t, "down")
	s.key(t, "down")
	s.key(t, "right")
	// Examples count: 10 -> 1
	s.key(t, "down")
	for i := 0; i < 9; i++ {
		s.key(t, "left")
	}
	// Apply
	s.key(t, "down")
	s.key(t, "down")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Математический тренажер")

	s.key(t, "up")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Математическое задание")
	s.eventuallyViewContains(t, "12 / 3 = ?")
	s.eventuallyViewContains(t, "Сложность: Начальный")

	s.typeText(t, "4")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Результаты тренировки")
	s.eventuallyViewContains(t, "1) 12 / 3 = 4")
}
