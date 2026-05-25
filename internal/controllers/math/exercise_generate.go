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
	switch operator {
	case mathmodels.OperatorAdd:
		leftMin, leftMax, rightMin, rightMax := addRanges(difficulty)
		left := randomInRange(random, leftMin, leftMax)
		right := randomInRange(random, rightMin, rightMax)
		return mathmodels.Exercise{Left: left, Right: right, Operator: operator}
	case mathmodels.OperatorSubtract:
		leftMin, leftMax, rightMin, rightMax, allowNegative := subtractRanges(difficulty)
		left := randomInRange(random, leftMin, leftMax)
		right := randomInRange(random, rightMin, rightMax)
		if !allowNegative && right > left {
			left, right = right, left
		}
		return mathmodels.Exercise{Left: left, Right: right, Operator: operator}
	case mathmodels.OperatorMultiply:
		leftMin, leftMax, rightMin, rightMax := multiplyRanges(difficulty)
		left := randomInRange(random, leftMin, leftMax)
		right := randomInRange(random, rightMin, rightMax)
		return mathmodels.Exercise{Left: left, Right: right, Operator: operator}
	case mathmodels.OperatorDivide:
		divisorMin, divisorMax, quotientMin, quotientMax := divideRanges(difficulty)
		divisor := randomInRange(random, divisorMin, divisorMax)
		quotient := randomInRange(random, quotientMin, quotientMax)
		return mathmodels.Exercise{Left: divisor * quotient, Right: divisor, Operator: operator}
	default:
		return mathmodels.Exercise{Left: 1, Right: 1, Operator: mathmodels.OperatorAdd}
	}
}

func randomInRange(random *rand.Rand, min int, max int) int {
	return random.Intn(max-min+1) + min
}
