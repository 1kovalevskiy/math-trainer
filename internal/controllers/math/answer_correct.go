package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func correctAnswer(exercise mathmodels.Exercise) int {
	switch exercise.Operator {
	case mathmodels.OperatorAdd:
		return exercise.Left + exercise.Right
	case mathmodels.OperatorSubtract:
		return exercise.Left - exercise.Right
	case mathmodels.OperatorMultiply:
		return exercise.Left * exercise.Right
	case mathmodels.OperatorDivide:
		if exercise.Right == 0 {
			return 0
		}
		return exercise.Left / exercise.Right
	default:
		return 0
	}
}
