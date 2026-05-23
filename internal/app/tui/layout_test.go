package tui

import (
	"strings"
	"testing"
)

func TestRenderScreenContentPlacesContentInCenterAndHintsAtBottom(t *testing.T) {
	t.Parallel()

	rendered := renderScreenContent("Main", []string{"Hint"}, 20, 9)
	lines := strings.Split(rendered, "\n")

	if got, want := len(lines), 9; got != want {
		t.Fatalf("rendered height mismatch: got %d, want %d", got, want)
	}
	if got := lineIndexContaining(lines, "Main"); got < 2 || got > 3 {
		t.Fatalf("content line mismatch: got index %d, want centered line 2 or 3", got)
	}
	if line := lines[7]; !strings.Contains(line, "Hint") {
		t.Fatalf("hint line mismatch: got %q, want line with Hint", line)
	}
}

func TestRenderScreenContentKeepsHintsAtSameBottomOffset(t *testing.T) {
	t.Parallel()

	short := strings.Split(renderScreenContent("Short", []string{"Hint"}, 20, 9), "\n")
	tall := strings.Split(renderScreenContent("Tall\nContent\nBlock", []string{"First hint", "Hint"}, 20, 9), "\n")

	if short[7] != tall[7] {
		t.Fatalf("hint line should be stable: short %q, tall %q", short[7], tall[7])
	}
}

func TestRenderScreenContentFitsTallContent(t *testing.T) {
	t.Parallel()

	rendered := renderScreenContent(strings.Repeat("line\n", 20), []string{"Hint"}, 20, 9)
	lines := strings.Split(rendered, "\n")

	if got, want := len(lines), 9; got != want {
		t.Fatalf("rendered height mismatch: got %d, want %d", got, want)
	}
	if got := lineIndexContaining(lines, "…"); got == -1 {
		t.Fatalf("expected overflow marker in fitted content, got %q", rendered)
	}
}

func lineIndexContaining(lines []string, needle string) int {
	for i, line := range lines {
		if strings.Contains(line, needle) {
			return i
		}
	}

	return -1
}
