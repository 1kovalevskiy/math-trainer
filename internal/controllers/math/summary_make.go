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

	elapsed := state.FinishedAt.Sub(state.StartedAt)
	if state.StartedAt.IsZero() || state.FinishedAt.IsZero() || elapsed < 0 {
		elapsed = 0
	}

	return &mathmodels.TrainingSummary{
		Settings: state.Settings,
		Results:  results,
		Correct:  correct,
		Total:    len(results),
		Elapsed:  elapsed,
	}
}
