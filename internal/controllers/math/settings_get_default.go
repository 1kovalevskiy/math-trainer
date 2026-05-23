package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func (c *Controller) GetDefaultSettings() mathmodels.TrainingSettings {
	return mathmodels.TrainingSettings{
		Difficulty:    mathmodels.DifficultyEasy,
		ExamplesCount: mathmodels.DefaultExamplesCount,
	}
}
