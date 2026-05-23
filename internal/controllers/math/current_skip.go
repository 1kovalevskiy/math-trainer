package mathcontroller

import (
	"context"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func (c *Controller) SkipCurrent(ctx context.Context) (mathmodels.TrainingSnapshot, error) {
	return c.completeCurrent(ctx, nil, mathmodels.ResultStatusSkipped)
}
