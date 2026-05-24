package mathcontroller

import (
	"fmt"
	"math/rand"
	"time"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func generateExercise(settings mathmodels.TrainingSettings) (mathmodels.Exercise, error) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	enabled := enabledOperatorDifficulties(settings)
	if len(enabled) == 0 {
		return mathmodels.Exercise{}, fmt.Errorf("no enabled operators in settings: %s", formatSettingsDifficulties(settings))
	}

	choice := enabled[random.Intn(len(enabled))]
	return generateExerciseByOperator(random, choice.operator, choice.difficulty), nil
}

func generateExerciseByOperator(
	random *rand.Rand,
	operator mathmodels.Operator,
	difficulty mathmodels.Difficulty,
) mathmodels.Exercise {
	min, max := numbersRange(difficulty)

	switch operator {
	case mathmodels.OperatorAdd:
		left := random.Intn(max-min+1) + min
		right := random.Intn(max-min+1) + min
		return mathmodels.Exercise{Left: left, Right: right, Operator: operator}
	case mathmodels.OperatorSubtract:
		left := random.Intn(max-min+1) + min
		right := random.Intn(max-min+1) + min
		if right > left {
			left, right = right, left
		}
		return mathmodels.Exercise{Left: left, Right: right, Operator: operator}
	case mathmodels.OperatorMultiply:
		left := random.Intn(max-min+1) + min
		right := random.Intn(max-min+1) + min
		return mathmodels.Exercise{Left: left, Right: right, Operator: operator}
	case mathmodels.OperatorDivide:
		divisor := random.Intn(max-min+1) + min
		quotient := random.Intn(max-min+1) + min
		return mathmodels.Exercise{Left: divisor * quotient, Right: divisor, Operator: operator}
	default:
		return mathmodels.Exercise{Left: min, Right: min, Operator: mathmodels.OperatorAdd}
	}
}
