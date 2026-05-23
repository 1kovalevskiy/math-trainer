package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func usedExercises(state mathmodels.TrainingState) map[mathmodels.Exercise]struct{} {
	used := make(map[mathmodels.Exercise]struct{}, len(state.Results))
	for _, result := range state.Results {
		used[result.Exercise] = struct{}{}
	}

	return used
}
