package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func makeSummary(state mathmodels.TrainingState) *mathmodels.TrainingSummary {
	results := cloneResults(state.Results)
	correct := 0
	for _, result := range results {
		if result.Status == mathmodels.ResultStatusCorrect {
			correct++
		}
	}

	return &mathmodels.TrainingSummary{
		Settings: state.Settings,
		Results:  results,
		Correct:  correct,
		Total:    len(results),
	}
}
