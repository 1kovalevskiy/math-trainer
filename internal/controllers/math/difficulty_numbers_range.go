package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func numbersRange(difficulty mathmodels.Difficulty) (min int, max int) {
	switch difficulty {
	case mathmodels.DifficultyEasy:
		return 1, 10
	case mathmodels.DifficultyMedium:
		return 10, 50
	case mathmodels.DifficultyHard:
		return 50, 100
	default:
		return 1, 10
	}
}
