package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func isKnownDifficulty(difficulty mathmodels.Difficulty) bool {
	switch difficulty {
	case mathmodels.DifficultyDisabled,
		mathmodels.DifficultyStarter,
		mathmodels.DifficultyEasy,
		mathmodels.DifficultyMedium,
		mathmodels.DifficultyHard,
		mathmodels.DifficultyExpert:
		return true
	default:
		return false
	}
}
