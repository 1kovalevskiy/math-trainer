package e2e

import (
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestIntegrationSettingsFlow_ApplyAffectsNextTraining(t *testing.T) {
	s := newSession(t,
		exercise(11, 4, mathmodels.OperatorAdd),
	)

	s.key(t, "down")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Настройки тренировки")

	s.key(t, "right")
	s.key(t, "down")
	s.key(t, "right")
	s.key(t, "down")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Математический тренажер")

	s.key(t, "up")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Математическое задание")
	s.eventuallyViewContains(t, "Пример 1 из 11")
	s.eventuallyViewContains(t, "Сложность: Средне")
	s.eventuallyViewContains(t, "11 + 4 = ?")
	s.requireGeneratedDifficulties(t, mathmodels.DifficultyMedium)
}
