package tui

import (
	"context"
	"strings"
	"testing"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/result"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func init() {
	zone.NewGlobal()
}

func TestResultScrollUsesVisibleResultsViewport(t *testing.T) {
	model := NewModel(context.Background(), persistTestController{}, nil)
	model.width = 90
	model.height = 20
	model.screen = ScreenResult
	model.resultModel = result.NewModel().WithSummary(tuiTestSummary(15))
	model.resultModel = model.resultModelWithCurrentViewport(model.resultModel)

	before := model.View()
	if !strings.Contains(before, "1) 1 + 2") {
		t.Fatalf("expected first result before scroll, got %q", before)
	}

	next, _ := model.Update(tea.KeyMsg{Type: tea.KeyDown})
	updated, ok := next.(Model)
	if !ok {
		t.Fatalf("model type mismatch: %T", next)
	}

	after := updated.View()
	if strings.Contains(after, "1) 1 + 2") {
		t.Fatalf("expected first result to scroll out of visible viewport, got %q", after)
	}
	if !strings.Contains(after, "4) 4 + 5") {
		t.Fatalf("expected next result row after scroll, got %q", after)
	}
}

func TestResultScrollbarStaysAtRightEdgeInRootView(t *testing.T) {
	model := NewModel(context.Background(), persistTestController{}, nil)
	model.width = 126
	model.height = 36
	model.screen = ScreenResult
	model.resultModel = result.NewModel().WithSummary(tuiTestSummary(40))
	model.resultModel = model.resultModelWithCurrentViewport(model.resultModel)

	view := model.View()
	found := false
	for _, line := range strings.Split(view, "\n") {
		if !strings.Contains(line, "█") && !strings.Contains(line, "│") {
			continue
		}
		if strings.Contains(line, "Результаты") || strings.Contains(line, "Сводка") || strings.Contains(line, "Сложности") {
			continue
		}
		if strings.Count(line, "│") < 2 && !strings.Contains(line, "█") {
			continue
		}

		if strings.HasSuffix(line, "█  │") || strings.HasSuffix(line, "│  │") {
			found = true
			continue
		}
	}
	if !found {
		t.Fatalf("expected root view scrollbar at right edge, got %q", view)
	}
}

func tuiTestSummary(count int) *mathmodels.TrainingSummary {
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
