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
	for i := 0; i < 4; i++ {
		s.key(t, "down")
	}
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
	s.requireGeneratedSettings(t, mathmodels.TrainingSettings{
		AddDifficulty:      mathmodels.DifficultyMedium,
		SubtractDifficulty: mathmodels.DifficultyEasy,
		MultiplyDifficulty: mathmodels.DifficultyDisabled,
		DivideDifficulty:   mathmodels.DifficultyDisabled,
		ExamplesCount:      11,
	})
}

func TestIntegrationSettingsFlow_CannotDisableAllOperators(t *testing.T) {
	s := newSession(t,
		exercise(4, 2, mathmodels.OperatorAdd),
	)

	s.key(t, "down")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Настройки тренировки")

	// Add: easy -> medium -> hard -> disabled
	s.key(t, "right")
	s.key(t, "right")
	s.key(t, "right")
	// Subtract: easy -> medium -> hard -> disabled
	s.key(t, "down")
	s.key(t, "right")
	s.key(t, "right")
	s.key(t, "right")
	// Multiply stays disabled
	s.key(t, "down")
	// Divide stays disabled
	s.key(t, "down")

	s.key(t, "down")
	s.key(t, "down")
	s.key(t, "enter")

	s.key(t, "up")
	s.key(t, "enter")
	s.eventuallyViewContains(t, "Математическое задание")

	settingsUsed := s.gen.Settings()[0]
	if settingsUsed.AddDifficulty == mathmodels.DifficultyDisabled &&
		settingsUsed.SubtractDifficulty == mathmodels.DifficultyDisabled &&
		settingsUsed.MultiplyDifficulty == mathmodels.DifficultyDisabled &&
		settingsUsed.DivideDifficulty == mathmodels.DifficultyDisabled {
		t.Fatal("expected at least one enabled operator")
	}
}
