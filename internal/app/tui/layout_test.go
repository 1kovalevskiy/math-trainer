package tui

import (
	"strings"
	"testing"
	"time"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	zone "github.com/lrstanley/bubblezone"
)

func init() {
	zone.NewGlobal()
}

func TestRenderScreenContentPlacesContentInCenterAndHintsAtBottom(t *testing.T) {
	t.Parallel()

	rendered := renderScreenContent("Main", []string{"Hint"}, 20, 9, false, false)
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

func TestRenderScreenContentPlacesRepositoryFooterAtBottom(t *testing.T) {
	t.Parallel()

	rendered := renderScreenContent("Main", []string{"Hint"}, 42, 9, true, false)
	lines := strings.Split(rendered, "\n")

	if got, want := len(lines), 9; got != want {
		t.Fatalf("rendered height mismatch: got %d, want %d", got, want)
	}
	if line := lines[8]; !strings.Contains(line, "github.com/1kovalevskiy/math-trainer") {
		t.Fatalf("footer line mismatch: got %q, want repository footer at bottom", line)
	}
}

func TestRenderScreenContentMarksRepositoryFooterZone(t *testing.T) {
	zone.Scan(renderScreenContent("Main", []string{"Hint"}, 42, 9, true, false))

	if waitForZone(t, repositoryFooterZoneID) == nil {
		t.Fatal("expected repository footer zone to be registered")
	}
}

func TestRenderFooterAreaUnderlinesRepositoryFooterOnHover(t *testing.T) {
	t.Parallel()

	if !ui.FooterHover.GetUnderline() {
		t.Fatal("expected hovered repository footer style to be underlined")
	}

	rendered := renderFooterArea(42, true)
	if strings.Contains(rendered, "[4m") || strings.Contains(rendered, "[24m") {
		t.Fatalf("expected hovered repository footer to avoid literal ANSI fragments, got %q", rendered)
	}
}

func TestRenderScreenContentOmitsRepositoryFooter(t *testing.T) {
	t.Parallel()

	rendered := renderScreenContent("Main", []string{"Hint"}, 42, 9, false, false)

	if strings.Contains(rendered, repositoryFooterText) {
		t.Fatalf("expected repository footer to be omitted, got %q", rendered)
	}
}

func TestScreenChromeShowsRepositoryFooterOnlyOnStartScreen(t *testing.T) {
	t.Parallel()

	if !screenChrome(ScreenStart).footer {
		t.Fatal("expected start screen to show repository footer")
	}

	for _, screen := range []Screen{ScreenSettings, ScreenTask, ScreenResult} {
		if screenChrome(screen).footer {
			t.Fatalf("expected %v screen to omit repository footer", screen)
		}
	}
}

func TestRenderScreenContentKeepsHintsAtSameBottomOffset(t *testing.T) {
	t.Parallel()

	short := strings.Split(renderScreenContent("Short", []string{"Hint"}, 20, 9, false, false), "\n")
	tall := strings.Split(renderScreenContent("Tall\nContent\nBlock", []string{"First hint", "Hint"}, 20, 9, false, false), "\n")

	if short[7] != tall[7] {
		t.Fatalf("hint line should be stable: short %q, tall %q", short[7], tall[7])
	}
}

func TestRenderScreenContentFitsTallContent(t *testing.T) {
	t.Parallel()

	rendered := renderScreenContent(strings.Repeat("line\n", 20), []string{"Hint"}, 20, 9, false, false)
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

func waitForZone(t *testing.T, id string) *zone.ZoneInfo {
	t.Helper()

	deadline := time.Now().Add(500 * time.Millisecond)
	for time.Now().Before(deadline) {
		if zoneInfo := zone.Get(id); zoneInfo != nil {
			return zoneInfo
		}
		time.Sleep(time.Millisecond)
	}

	return nil
}
