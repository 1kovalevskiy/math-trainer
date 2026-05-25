package result

import (
	"strings"
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestRenderSummaryHeaderLinesIncludesScoreAndElapsedTime(t *testing.T) {
	t.Parallel()

	summary := mathmodels.TrainingSummary{
		Settings: mathmodels.DefaultTrainingSettings(),
		Correct:  2,
		Total:    3,
		Elapsed:  125000000000,
	}

	header := strings.Join(renderSummaryHeaderLines(summary), "\n")

	if !strings.Contains(header, "Правильных: 2 из 3") {
		t.Fatalf("expected score in header, got %q", header)
	}
	if !strings.Contains(header, "Время: ") || !strings.Contains(header, "2:05") {
		t.Fatalf("expected elapsed time in header, got %q", header)
	}
}
