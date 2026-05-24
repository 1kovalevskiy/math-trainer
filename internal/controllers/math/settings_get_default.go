package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func (c *Controller) GetDefaultSettings() mathmodels.TrainingSettings {
	if c == nil {
		return mathmodels.DefaultTrainingSettings()
	}

	return c.NormalizeSettings(c.defaultSettings)
}
