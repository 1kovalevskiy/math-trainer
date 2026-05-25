package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func (c *Controller) GetPreviousDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty {
	switch current {
	case mathmodels.DifficultyDisabled:
		return mathmodels.DifficultyExpert
	case mathmodels.DifficultyStarter:
		return mathmodels.DifficultyDisabled
	case mathmodels.DifficultyEasy:
		return mathmodels.DifficultyStarter
	case mathmodels.DifficultyMedium:
		return mathmodels.DifficultyEasy
	case mathmodels.DifficultyHard:
		return mathmodels.DifficultyMedium
	case mathmodels.DifficultyExpert:
		return mathmodels.DifficultyHard
	default:
		return mathmodels.DifficultyStarter
	}
}
