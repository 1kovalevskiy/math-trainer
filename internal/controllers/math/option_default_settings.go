package mathcontroller

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func WithDefaultSettings(settings mathmodels.TrainingSettings) Option {
	return func(c *Controller) {
		if c == nil {
			return
		}
		c.defaultSettings = settings
	}
}
