package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func (c *Controller) GetNextDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty {
	switch current {
	case mathmodels.DifficultyDisabled:
		return mathmodels.DifficultyEasy
	case mathmodels.DifficultyEasy:
		return mathmodels.DifficultyMedium
	case mathmodels.DifficultyMedium:
		return mathmodels.DifficultyHard
	case mathmodels.DifficultyHard:
		return mathmodels.DifficultyDisabled
	default:
		return mathmodels.DifficultyEasy
	}
}
