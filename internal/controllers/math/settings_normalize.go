package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func (c *Controller) NormalizeSettings(settings mathmodels.TrainingSettings) mathmodels.TrainingSettings {
	if !isKnownDifficulty(settings.AddDifficulty) {
		settings.AddDifficulty = mathmodels.DifficultyStarter
	}
	if !isKnownDifficulty(settings.SubtractDifficulty) {
		settings.SubtractDifficulty = mathmodels.DifficultyStarter
	}
	if !isKnownDifficulty(settings.MultiplyDifficulty) {
		settings.MultiplyDifficulty = mathmodels.DifficultyDisabled
	}
	if !isKnownDifficulty(settings.DivideDifficulty) {
		settings.DivideDifficulty = mathmodels.DifficultyDisabled
	}
	if settings.ExamplesCount < mathmodels.MinExamplesCount {
		settings.ExamplesCount = mathmodels.MinExamplesCount
	}
	if settings.ExamplesCount > mathmodels.MaxExamplesCount {
		settings.ExamplesCount = mathmodels.MaxExamplesCount
	}

	if len(enabledOperatorDifficulties(settings)) == 0 {
		settings.AddDifficulty = mathmodels.DifficultyStarter
	}

	return settings
}
