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

	for _, focus := range allSettingsFocuses() {
		model := NewModel(mathmodels.DefaultTrainingSettings(), testRules{})
		model.focus = focus
		view := model.View()
		lines := strings.Split(view, "\n")

		actionWidth := actionLineWidth(t, lines)
		labels := []string{"Сложение", "Вычитание", "Умножение", "Деление", "Количество примеров"}
		for _, label := range labels {
			line := findLine(lines, label)
			if line == "" {
				t.Fatalf("focus %+v: line with %q not found", focus, label)
			}
			if strings.Contains(line, "\n") {
				t.Fatalf("focus %+v: line %q contains newline", focus, label)
			}
			if got := lipgloss.Width(line); got != actionWidth {
				t.Fatalf("focus %+v: line %q width mismatch: got %d, want %d", focus, label, got, actionWidth)
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

func TestSettingsLayoutWidthUsesWidestDifficultyLabel(t *testing.T) {
	t.Parallel()

	difficulties := []mathmodels.Difficulty{
		mathmodels.DifficultyDisabled,
		mathmodels.DifficultyStarter,
		mathmodels.DifficultyEasy,
		mathmodels.DifficultyMedium,
		mathmodels.DifficultyHard,
		mathmodels.DifficultyExpert,
	}

	var wantRowWidth int
	var wantValueWidth int
	for _, difficulty := range difficulties {
		model := NewModel(mathmodels.TrainingSettings{
			AddDifficulty:      difficulty,
			SubtractDifficulty: difficulty,
			MultiplyDifficulty: difficulty,
			DivideDifficulty:   difficulty,
			ExamplesCount:      mathmodels.DefaultExamplesCount,
		}, testRules{})
		layout := model.makeSettingsLayout()

		if wantRowWidth == 0 {
			wantRowWidth = layout.row.RowWidth
			wantValueWidth = layout.row.ValueWidth
			continue
		}
		if got := layout.row.RowWidth; got != wantRowWidth {
			t.Fatalf("row width changed for %q: got %d, want %d", difficulty, got, wantRowWidth)
		}
		if got := layout.row.ValueWidth; got != wantValueWidth {
			t.Fatalf("value width changed for %q: got %d, want %d", difficulty, got, wantValueWidth)
		}
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
