package result

import (
	"strings"
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestRenderResultEntriesPreservesResultOrder(t *testing.T) {
	t.Parallel()

	first := 10
	second := 20
	entries := renderResultEntries([]mathmodels.ExampleResult{
		{
			Order:      1,
			Exercise:   mathmodels.Exercise{Left: 1, Right: 2, Operator: mathmodels.OperatorAdd},
			UserAnswer: &first,
			Status:     mathmodels.ResultStatusCorrect,
		},
		{
			Order:         2,
			Exercise:      mathmodels.Exercise{Left: 3, Right: 4, Operator: mathmodels.OperatorAdd},
			CorrectAnswer: 7,
			UserAnswer:    &second,
			Status:        mathmodels.ResultStatusIncorrect,
		},
	})

	joined := strings.Join(entries, "\n")
	if !strings.Contains(joined, "1) 1 + 2") || !strings.Contains(joined, "2) 3 + 4") {
		t.Fatalf("expected ordered rendered entries, got %q", joined)
	}
}
