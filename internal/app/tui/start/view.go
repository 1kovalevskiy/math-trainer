package start

import (
	"fmt"
	"math"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

func (m Model) View() string {
	var b strings.Builder
	buttonWidth := 1

	for _, option := range m.options {
		width := lipgloss.Width(ui.MenuItem(false, option))
		buttonWidth = int(math.Max(float64(buttonWidth), float64(width)))
	}

	b.WriteString(ui.Title.Render("Математический тренажер") + "\n")
	b.WriteString(ui.Subtitle.Render("Тренировка устного счета") + "\n\n")
	for i, option := range m.options {
		if i > 0 {
			b.WriteString("\n")
		}
		button := ui.MenuItemFixed(m.cursor == i, option, buttonWidth)
		b.WriteString(zone.Mark(optionZoneID(i), button) + "\n")
	}

	return b.String()
}

func optionZoneID(index int) string {
	return fmt.Sprintf("start:option:%d", index)
}
