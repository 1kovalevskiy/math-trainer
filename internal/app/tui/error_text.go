package tui

import (
	"errors"
	"fmt"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func errorText(err error) string {
	if errors.Is(err, mathmodels.ErrInvalidAnswer) {
		return "Ответ должен быть числом"
	}

	return fmt.Sprintf("%v", err)
}
