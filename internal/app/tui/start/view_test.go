package start

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func init() {
	zone.NewGlobal()
}

func TestViewRendersAllButtonsWithEqualWidth(t *testing.T) {
	t.Parallel()

	model := NewModel()
	lines := strings.Split(model.View(), "\n")

	widths := make([]int, 0, len(model.options))
	for _, line := range lines {
		if strings.Contains(line, "Начать тренировку") || strings.Contains(line, "Настройки сложности") || strings.Contains(line, "Выход") {
			widths = append(widths, lipgloss.Width(line))
		}
	}

	if len(widths) != 3 {
		t.Fatalf("expected 3 button lines, got %d", len(widths))
	}
	for i := 1; i < len(widths); i++ {
		if widths[i] != widths[0] {
			t.Fatalf("button width mismatch: got %v", widths)
		}
	}
}
