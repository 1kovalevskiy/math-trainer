package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func addRanges(difficulty mathmodels.Difficulty) (leftMin int, leftMax int, rightMin int, rightMax int) {
	switch difficulty {
	case mathmodels.DifficultyStarter:
		return 1, 10, 1, 10
	case mathmodels.DifficultyEasy:
		return 10, 99, 10, 99
	case mathmodels.DifficultyMedium:
		return 50, 999, 50, 999
	case mathmodels.DifficultyHard:
		return 100, 9999, 100, 9999
	case mathmodels.DifficultyExpert:
		return 1000, 9999, 1000, 9999
	default:
		return 1, 10, 1, 10
	}
}

func subtractRanges(difficulty mathmodels.Difficulty) (leftMin int, leftMax int, rightMin int, rightMax int, allowNegative bool) {
	switch difficulty {
	case mathmodels.DifficultyStarter:
		return 1, 10, 1, 10, false
	case mathmodels.DifficultyEasy:
		return 10, 99, 10, 99, false
	case mathmodels.DifficultyMedium:
		return 50, 999, 50, 999, false
	case mathmodels.DifficultyHard:
		return 100, 9999, 100, 9999, true
	case mathmodels.DifficultyExpert:
		return 1000, 9999, 1000, 9999, true
	default:
		return 1, 10, 1, 10, false
	}
}

func multiplyRanges(difficulty mathmodels.Difficulty) (leftMin int, leftMax int, rightMin int, rightMax int) {
	switch difficulty {
	case mathmodels.DifficultyStarter:
		return 1, 10, 1, 10
	case mathmodels.DifficultyEasy:
		return 2, 12, 2, 12
	case mathmodels.DifficultyMedium:
		return 10, 50, 2, 12
	case mathmodels.DifficultyHard:
		return 10, 99, 10, 20
	case mathmodels.DifficultyExpert:
		return 50, 100, 50, 100
	default:
		return 1, 10, 1, 10
	}
}

func divideRanges(difficulty mathmodels.Difficulty) (divisorMin int, divisorMax int, quotientMin int, quotientMax int) {
	switch difficulty {
	case mathmodels.DifficultyStarter:
		return 1, 10, 1, 10
	case mathmodels.DifficultyEasy:
		return 2, 20, 2, 20
	case mathmodels.DifficultyMedium:
		return 10, 50, 10, 50
	case mathmodels.DifficultyHard:
		return 20, 75, 20, 75
	case mathmodels.DifficultyExpert:
		return 50, 100, 50, 100
	default:
		return 1, 10, 1, 10
	}
}
