package tui

import (
	"errors"
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestErrorTextMapsInvalidAnswerForUI(t *testing.T) {
	t.Parallel()

	if got, want := errorText(mathmodels.ErrInvalidAnswer), "Ответ должен быть числом"; got != want {
		t.Fatalf("error text mismatch: got %q, want %q", got, want)
	}
}

func TestErrorTextFallsBackToErrorMessage(t *testing.T) {
	t.Parallel()

	if got, want := errorText(errors.New("boom")), "boom"; got != want {
		t.Fatalf("error text mismatch: got %q, want %q", got, want)
	}
}
