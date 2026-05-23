package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func makeSnapshot(state mathmodels.TrainingState) mathmodels.TrainingSnapshot {
	if state.Finished {
		return mathmodels.TrainingSnapshot{
			Phase:    mathmodels.TrainingPhaseFinished,
			Settings: state.Settings,
			Summary:  makeSummary(state),
		}
	}

	return mathmodels.TrainingSnapshot{
		Phase:    mathmodels.TrainingPhaseInProgress,
		Settings: state.Settings,
		Current: &mathmodels.CurrentExercise{
			Order:    state.CurrentOrder,
			Total:    state.Settings.ExamplesCount,
			Exercise: state.CurrentExercise,
		},
	}
}
