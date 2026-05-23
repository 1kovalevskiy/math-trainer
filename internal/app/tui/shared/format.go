package shared

import (
	"fmt"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func DifficultyLabel(difficulty mathmodels.Difficulty) string {
	switch difficulty {
	case mathmodels.DifficultyEasy:
		return "Легко"
	case mathmodels.DifficultyMedium:
		return "Средне"
	case mathmodels.DifficultyHard:
		return "Сложно"
	default:
		return "Неизвестно"
	}
}

func ExerciseText(exercise mathmodels.Exercise) string {
	return fmt.Sprintf("%d %s %d", exercise.Left, OperatorSymbol(exercise.Operator), exercise.Right)
}

func OperatorSymbol(operator mathmodels.Operator) string {
	switch operator {
	case mathmodels.OperatorAdd:
		return "+"
	case mathmodels.OperatorSubtract:
		return "-"
	default:
		return "?"
	}
}
