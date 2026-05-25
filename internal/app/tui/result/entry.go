package result

import (
	"fmt"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func renderEntry(entry mathmodels.ExampleResult) string {
	base := fmt.Sprintf("%d) %s = ", entry.Order, shared.ExerciseText(entry.Exercise))

	switch entry.Status {
	case mathmodels.ResultStatusCorrect:
		answer := 0
		if entry.UserAnswer != nil {
			answer = *entry.UserAnswer
		}
		return correctStyle.Render(base + fmt.Sprintf("%d", answer))
	case mathmodels.ResultStatusIncorrect:
		userAnswer := "_"
		if entry.UserAnswer != nil {
			userAnswer = fmt.Sprintf("%d", *entry.UserAnswer)
		}
		return incorrectStyle.Render(base + userAnswer + fmt.Sprintf(" (ответ: %d)", entry.CorrectAnswer))
	case mathmodels.ResultStatusSkipped:
		return skippedStyle.Render(base + fmt.Sprintf("____ (ответ: %d)", entry.CorrectAnswer))
	default:
		return resultSurfaceStyle.Render(fmt.Sprintf("%d) %s", entry.Order, shared.ExerciseText(entry.Exercise)))
	}
}
