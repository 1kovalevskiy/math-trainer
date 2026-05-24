package result

import (
	"strings"
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	zone "github.com/lrstanley/bubblezone"
)

func init() {
	zone.NewGlobal()
}

func TestViewWithSizeUsesThreeColumnsByDefaultWhenWidthAllows(t *testing.T) {
	t.Parallel()

	model := NewModel().WithSummary(testSummary(9))
	view := model.ViewWithSize(120, 20)

	if strings.Contains(view, "█") || strings.Contains(view, "│") {
		t.Fatalf("unexpected scrollbar for tall viewport: %q", view)
	}
	row := findLineWithAll(view, []string{"1)", "2)", "3)"})
	if row == "" {
		t.Fatalf("expected default 3-column first row (1,2,3), got %q", view)
	}
}

func TestViewWithSizeUsesMultipleColumnsWhenHeightSmall(t *testing.T) {
	t.Parallel()

	model := NewModel().WithSummary(testSummaryWithWideEntries(8))
	view := model.ViewWithSize(50, 10)

	row := findLineWithAll(view, []string{"1)", "2)"})
	if row == "" {
		t.Fatalf("expected row-major 2-column order (1 and 2 in same row), got %q", view)
	}
	if strings.Contains(row, "3)") {
		t.Fatalf("unexpected third item in 2-column row: %q", row)
	}
}

func TestViewWithSizeUsesThreeColumnsWithRowMajorOrder(t *testing.T) {
	t.Parallel()

	model := NewModel().WithSummary(testSummary(9))
	view := model.ViewWithSize(120, 10)

	row := findLineWithAll(view, []string{"1)", "2)", "3)"})
	if row == "" {
		t.Fatalf("expected row-major 3-column order (1,2,3 in same row), got %q", view)
	}
}

func TestViewWithSizeShowsScrollbarWhenStillOverflowing(t *testing.T) {
	t.Parallel()

	model := NewModel().WithSummary(testSummary(40)).WithViewport(70, 6)
	model.scrollOffset = 5
	view := model.ViewWithSize(70, 12)

	if !strings.Contains(view, "█") {
		t.Fatalf("expected scrollbar thumb, got %q", view)
	}
	if !strings.Contains(view, "Результаты тренировки") || !strings.Contains(view, "Решить еще один пример") {
		t.Fatalf("expected fixed header and actions, got %q", view)
	}
}

func TestViewWithSizeShowsEmptyFallback(t *testing.T) {
	t.Parallel()

	model := NewModel().WithSummary(&mathmodels.TrainingSummary{})
	view := model.ViewWithSize(60, 10)
	if !strings.Contains(view, "Нет ответов") {
		t.Fatalf("expected empty fallback, got %q", view)
	}
}

func testSummary(count int) *mathmodels.TrainingSummary {
	results := make([]mathmodels.ExampleResult, 0, count)
	for i := 1; i <= count; i++ {
		answer := i
		results = append(results, mathmodels.ExampleResult{
			Order:         i,
			Exercise:      mathmodels.Exercise{Left: i, Right: i + 1, Operator: mathmodels.OperatorAdd},
			CorrectAnswer: i,
			UserAnswer:    &answer,
			Status:        mathmodels.ResultStatusCorrect,
		})
	}
	return &mathmodels.TrainingSummary{
		Settings: mathmodels.DefaultTrainingSettings(),
		Results:  results,
		Correct:  count,
		Total:    count,
	}
}

func testSummaryWithWideEntries(count int) *mathmodels.TrainingSummary {
	results := make([]mathmodels.ExampleResult, 0, count)
	for i := 1; i <= count; i++ {
		answer := i
		results = append(results, mathmodels.ExampleResult{
			Order:         i,
			Exercise:      mathmodels.Exercise{Left: 100 + i, Right: 200 + i, Operator: mathmodels.OperatorAdd},
			CorrectAnswer: i,
			UserAnswer:    &answer,
			Status:        mathmodels.ResultStatusCorrect,
		})
	}
	return &mathmodels.TrainingSummary{
		Settings: mathmodels.DefaultTrainingSettings(),
		Results:  results,
		Correct:  count,
		Total:    count,
	}
}

func findLineWithAll(view string, parts []string) string {
	lines := strings.Split(view, "\n")
	for _, line := range lines {
		matched := true
		for _, part := range parts {
			if !strings.Contains(line, part) {
				matched = false
				break
			}
		}
		if matched {
			return line
		}
	}
	return ""
}
