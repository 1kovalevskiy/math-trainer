package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func (c *Controller) GetPreviousDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty {
	switch current {
	case mathmodels.DifficultyEasy:
		return mathmodels.DifficultyHard
	case mathmodels.DifficultyMedium:
		return mathmodels.DifficultyEasy
	case mathmodels.DifficultyHard:
		return mathmodels.DifficultyMedium
	default:
		return mathmodels.DifficultyEasy
	}
}
