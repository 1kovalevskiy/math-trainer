package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func (c *Controller) NormalizeSettings(settings mathmodels.TrainingSettings) mathmodels.TrainingSettings {
	if !isKnownDifficulty(settings.Difficulty) {
		settings.Difficulty = mathmodels.DifficultyEasy
	}
	if settings.ExamplesCount < mathmodels.MinExamplesCount {
		settings.ExamplesCount = mathmodels.MinExamplesCount
	}
	if settings.ExamplesCount > mathmodels.MaxExamplesCount {
		settings.ExamplesCount = mathmodels.MaxExamplesCount
	}

	return settings
}
