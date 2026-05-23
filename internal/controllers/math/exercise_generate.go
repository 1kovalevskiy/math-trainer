package mathcontroller

import (
	"math/rand"
	"time"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func generateExercise(difficulty mathmodels.Difficulty) (mathmodels.Exercise, error) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	min, max := numbersRange(difficulty)

	left := random.Intn(max-min+1) + min
	right := random.Intn(max-min+1) + min

	operators := []mathmodels.Operator{mathmodels.OperatorAdd, mathmodels.OperatorSubtract}
	operator := operators[random.Intn(len(operators))]
	if operator == mathmodels.OperatorSubtract && right > left {
		left, right = right, left
	}

	return mathmodels.Exercise{
		Left:     left,
		Right:    right,
		Operator: operator,
	}, nil
}
