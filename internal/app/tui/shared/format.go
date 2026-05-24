package shared

import (
	"fmt"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func DifficultyLabel(difficulty mathmodels.Difficulty) string {
	switch difficulty {
	case mathmodels.DifficultyDisabled:
		return "Отключено"
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
	case mathmodels.OperatorMultiply:
		return "*"
	case mathmodels.OperatorDivide:
		return "/"
	default:
		return "?"
	}
}

func DifficultyForOperator(settings mathmodels.TrainingSettings, operator mathmodels.Operator) mathmodels.Difficulty {
	switch operator {
	case mathmodels.OperatorAdd:
		return settings.AddDifficulty
	case mathmodels.OperatorSubtract:
		return settings.SubtractDifficulty
	case mathmodels.OperatorMultiply:
		return settings.MultiplyDifficulty
	case mathmodels.OperatorDivide:
		return settings.DivideDifficulty
	default:
		return mathmodels.DifficultyDisabled
	}
}
