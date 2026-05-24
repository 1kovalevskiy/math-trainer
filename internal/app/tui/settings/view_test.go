package settings

import (
	"strings"
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func init() {
	zone.NewGlobal()
}

func TestViewDifficultyRowsMatchActionButtonsWidth(t *testing.T) {
	t.Parallel()

	model := NewModel(mathmodels.DefaultTrainingSettings(), testRules{})
	view := model.View()
	lines := strings.Split(view, "\n")

	var actionWidth int
	for _, line := range lines {
		if strings.Contains(line, "Применить") && strings.Contains(line, "Назад") {
			actionWidth = lipgloss.Width(line)
			break
		}
	}
	if actionWidth == 0 {
		t.Fatal("action buttons line not found")
	}

	labels := []string{"Сложение", "Вычитание", "Умножение", "Деление"}
	for _, label := range labels {
		line := findLine(lines, label)
		if line == "" {
			t.Fatalf("line with %q not found", label)
		}
		if got := lipgloss.Width(line); got != actionWidth {
			t.Fatalf("line %q width mismatch: got %d, want %d", label, got, actionWidth)
		}
	}

	countLine := findLine(lines, "Количество примеров")
	if countLine == "" {
		t.Fatal("count line not found")
	}
	if got := lipgloss.Width(countLine); got != actionWidth {
		t.Fatalf("count line width mismatch: got %d, want %d", got, actionWidth)
	}
}

func findLine(lines []string, contains string) string {
	for _, line := range lines {
		if strings.Contains(line, contains) {
			return line
		}
	}
	return ""
}
