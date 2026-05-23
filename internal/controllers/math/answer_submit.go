package mathcontroller

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func (c *Controller) SubmitAnswer(ctx context.Context, rawAnswer string) (mathmodels.TrainingSnapshot, error) {
	answer, err := strconv.Atoi(strings.TrimSpace(rawAnswer))
	if err != nil {
		return mathmodels.TrainingSnapshot{}, fmt.Errorf("%w: %q", mathmodels.ErrInvalidAnswer, rawAnswer)
	}

	answerCopy := answer
	return c.completeCurrent(ctx, &answerCopy, mathmodels.ResultStatusIncorrect)
}
