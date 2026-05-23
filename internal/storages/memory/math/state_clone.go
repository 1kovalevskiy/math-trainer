package mathmemory

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func cloneState(state mathmodels.TrainingState) mathmodels.TrainingState {
	state.Results = cloneResults(state.Results)
	return state
}
