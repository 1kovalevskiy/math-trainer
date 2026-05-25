package result

import (
	"fmt"
	"strings"
	"testing"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	"github.com/charmbracelet/lipgloss"
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

func TestViewWithSizePlacesScrollbarAtRightEdge(t *testing.T) {
	t.Parallel()

	width := 70
	view := NewModel().WithSummary(testSummary(40)).ViewWithSize(width, 10)

	found := false
	for _, line := range strings.Split(view, "\n") {
		if !strings.Contains(line, "█") && !strings.Contains(line, "│") {
			continue
		}
		found = true
		if lineWidth := ui.Width(line); lineWidth != width {
			t.Fatalf("scrollbar line width mismatch: got %d, want %d, line=%q", lineWidth, width, line)
		}
		trimmed := strings.TrimRight(line, " ")
		if !strings.HasSuffix(trimmed, "█") && !strings.HasSuffix(trimmed, "│") {
			t.Fatalf("expected scrollbar at right edge, got %q", line)
		}
	}
	if !found {
		t.Fatal("expected scrollbar line")
	}
}

func TestViewWithSizeDoesNotOverflowRequestedWidth(t *testing.T) {
	t.Parallel()

	width := 80
	view := NewModel().WithSummary(testSummaryWithWideEntries(10)).ViewWithSize(width, 24)

	for _, line := range strings.Split(view, "\n") {
		if lineWidth := ui.Width(line); lineWidth > width {
			t.Fatalf("line overflows requested width: got %d, want <= %d, line=%q", lineWidth, width, line)
		}
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

func TestRenderEntryAppliesStatusBackgroundToWholeResultText(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		entry mathmodels.ExampleResult
		style lipgloss.Style
		text  string
	}{
		{
			name: "correct",
			entry: mathmodels.ExampleResult{
				Order:         1,
				Exercise:      mathmodels.Exercise{Left: 2, Right: 3, Operator: mathmodels.OperatorAdd},
				CorrectAnswer: 5,
				UserAnswer:    intPtr(5),
				Status:        mathmodels.ResultStatusCorrect,
			},
			style: correctStyle,
			text:  "1) 2 + 3 = 5",
		},
		{
			name: "incorrect",
			entry: mathmodels.ExampleResult{
				Order:         2,
				Exercise:      mathmodels.Exercise{Left: 7, Right: 4, Operator: mathmodels.OperatorSubtract},
				CorrectAnswer: 3,
				UserAnswer:    intPtr(4),
				Status:        mathmodels.ResultStatusIncorrect,
			},
			style: incorrectStyle,
			text:  "2) 7 - 4 = 4 (ответ: 3)",
		},
		{
			name: "skipped",
			entry: mathmodels.ExampleResult{
				Order:         3,
				Exercise:      mathmodels.Exercise{Left: 6, Right: 2, Operator: mathmodels.OperatorMultiply},
				CorrectAnswer: 12,
				Status:        mathmodels.ResultStatusSkipped,
			},
			style: skippedStyle,
			text:  "3) 6 * 2 = ____ (ответ: 12)",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got, want := renderEntry(tc.entry), tc.style.Render(tc.text); got != want {
				t.Fatalf("rendered entry mismatch:\ngot:  %q\nwant: %q", got, want)
			}
		})
	}
}

func TestResultStatusStylesUseContrastingBackgroundsWithoutIncorrectStrikethrough(t *testing.T) {
	t.Parallel()

	if got, want := correctStyle.GetBackground(), lipgloss.Color("151"); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("correct background mismatch: got %v, want %v", got, want)
	}
	if got, want := correctStyle.GetForeground(), lipgloss.Color("235"); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("correct foreground mismatch: got %v, want %v", got, want)
	}
	if got, want := incorrectStyle.GetBackground(), lipgloss.Color("217"); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("incorrect background mismatch: got %v, want %v", got, want)
	}
	if got, want := incorrectStyle.GetForeground(), lipgloss.Color("235"); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("incorrect foreground mismatch: got %v, want %v", got, want)
	}
	if incorrectStyle.GetStrikethrough() {
		t.Fatal("incorrect result style must not be strikethrough")
	}
	if got, want := skippedStyle.GetBackground(), lipgloss.Color("240"); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("skipped background mismatch: got %v, want %v", got, want)
	}
	if got, want := skippedStyle.GetForeground(), lipgloss.Color("230"); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("skipped foreground mismatch: got %v, want %v", got, want)
	}
}

func TestViewWithSizeVerticallyCentersResultsWhenTheyFit(t *testing.T) {
	t.Parallel()

	model := NewModel().WithSummary(testSummary(3))
	view := model.ViewWithSize(100, 30)
	lines := strings.Split(view, "\n")

	resultLine := -1
	for i, line := range lines {
		if strings.Contains(line, "1)") && strings.Contains(line, "2)") && strings.Contains(line, "3)") {
			resultLine = i
			break
		}
	}
	if resultLine == -1 {
		t.Fatalf("result row not found in view: %q", view)
	}
	if resultLine <= 10 {
		t.Fatalf("expected results to be vertically centered lower in viewport, got line %d", resultLine)
	}
}

func TestRenderGridRowsAddsSpaceBetweenResultRows(t *testing.T) {
	t.Parallel()

	rows := NewModel().renderGridRows([]string{"1", "2", "3", "4", "5", "6"}, 3, 2, 1, 9)
	if len(rows) != 3 {
		t.Fatalf("rows length mismatch: got %d, want 3", len(rows))
	}
	if rows[1] != "" {
		t.Fatalf("expected blank row gap, got %q", rows[1])
	}
}

func TestRenderGridRowsCentersSingleItemLastRow(t *testing.T) {
	t.Parallel()

	contentWidth := 12
	rows := NewModel().renderGridRows([]string{"1", "2", "3", "4", "5", "6", "7"}, 3, 3, 1, contentWidth)
	last := rows[len(rows)-1]

	if strings.TrimSpace(last) != "7" {
		t.Fatalf("last row mismatch: got %q", last)
	}
	if left, right := sidePadding(last); abs(left-right) > 1 {
		t.Fatalf("expected centered last row, got left=%d right=%d line=%q", left, right, last)
	}
}

func TestRenderGridRowsCentersTwoItemLastRowAsGroup(t *testing.T) {
	t.Parallel()

	contentWidth := 12
	rows := NewModel().renderGridRows([]string{"1", "2", "3", "4", "5", "6", "7", "8"}, 3, 3, 1, contentWidth)
	last := rows[len(rows)-1]

	if !strings.Contains(last, "7") || !strings.Contains(last, "8") {
		t.Fatalf("expected last row to contain 7 and 8, got %q", last)
	}
	if left, right := sidePadding(last); abs(left-right) > 1 {
		t.Fatalf("expected centered two-item group, got left=%d right=%d line=%q", left, right, last)
	}
}

func TestRenderGridRowsCentersSingleItemLastRowInTwoColumnGrid(t *testing.T) {
	t.Parallel()

	contentWidth := 8
	rows := NewModel().renderGridRows([]string{"1", "2", "3"}, 2, 2, 1, contentWidth)
	last := rows[len(rows)-1]

	if strings.TrimSpace(last) != "3" {
		t.Fatalf("last row mismatch: got %q", last)
	}
	if left, right := sidePadding(last); abs(left-right) > 1 {
		t.Fatalf("expected centered single item in two-column grid, got left=%d right=%d line=%q", left, right, last)
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

func intPtr(value int) *int {
	return &value
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

func sidePadding(line string) (int, int) {
	left := len(line) - len(strings.TrimLeft(line, " "))
	right := ui.Width(line) - ui.Width(strings.TrimRight(line, " "))
	return left, right
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
