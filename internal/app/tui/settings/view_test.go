package settings

import (
	"os"
	"regexp"
	"strings"
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

var ansiPattern = regexp.MustCompile(`\x1b\[[0-9;?]*[A-Za-z]`)

func init() {
	zone.NewGlobal()
}

func TestViewDifficultyRowsMatchActionButtonsWidth(t *testing.T) {
	t.Parallel()

	for cursor := rowAddDifficulty; cursor <= rowBack; cursor++ {
		model := NewModel(mathmodels.DefaultTrainingSettings(), testRules{})
		model.cursor = cursor
		view := model.View()
		lines := strings.Split(view, "\n")

		actionWidth := actionLineWidth(t, lines)
		labels := []string{"Сложение", "Вычитание", "Умножение", "Деление", "Количество примеров"}
		for _, label := range labels {
			line := findLine(lines, label)
			if line == "" {
				t.Fatalf("cursor %d: line with %q not found", cursor, label)
			}
			if strings.Contains(line, "\n") {
				t.Fatalf("cursor %d: line %q contains newline", cursor, label)
			}
			if got := lipgloss.Width(line); got != actionWidth {
				t.Fatalf("cursor %d: line %q width mismatch: got %d, want %d", cursor, label, got, actionWidth)
			}
		}
	}
}

func TestViewDefaultNormalizedShape(t *testing.T) {
	t.Parallel()

	model := NewModel(mathmodels.DefaultTrainingSettings(), testRules{})
	got := normalizeView(model.View())
	wantBytes, err := os.ReadFile("testdata/settings_default.golden")
	if err != nil {
		t.Fatal(err)
	}
	want := normalizeView(string(wantBytes))

	if got != want {
		t.Fatalf("normalized settings view mismatch:\n got:\n%s\nwant:\n%s", got, want)
	}
}

func actionLineWidth(t *testing.T, lines []string) int {
	t.Helper()

	for _, line := range lines {
		if strings.Contains(line, "Применить") && strings.Contains(line, "Назад") {
			return lipgloss.Width(line)
		}
	}

	t.Fatal("action buttons line not found")
	return 0
}

func findLine(lines []string, contains string) string {
	for _, line := range lines {
		if strings.Contains(line, contains) {
			return line
		}
	}
	return ""
}

func stripANSI(value string) string {
	return ansiPattern.ReplaceAllString(value, "")
}

func normalizeView(value string) string {
	lines := strings.Split(strings.TrimRight(stripANSI(value), "\n"), "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " ")
	}

	return strings.Join(lines, "\n")
}
