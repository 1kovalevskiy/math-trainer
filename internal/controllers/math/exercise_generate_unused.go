package mathcontroller

import (
	"fmt"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

const maxUniqueExerciseAttempts = 100

func (c *Controller) generateUnused(
	difficulty mathmodels.Difficulty,
	used map[mathmodels.Exercise]struct{},
) (mathmodels.Exercise, error) {
	for attempt := 0; attempt < maxUniqueExerciseAttempts; attempt++ {
		exercise, err := c.generate(difficulty)
		if err != nil {
			return mathmodels.Exercise{}, err
		}
		if _, exists := used[exercise]; !exists {
			return exercise, nil
		}
	}

	return mathmodels.Exercise{}, fmt.Errorf(
		"%w after %d attempts",
		mathmodels.ErrUniqueExerciseExhausted,
		maxUniqueExerciseAttempts,
	)
}
