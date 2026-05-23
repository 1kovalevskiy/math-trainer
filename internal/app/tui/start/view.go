package start

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	zone "github.com/lrstanley/bubblezone"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(ui.Title.Render("Математический тренажер") + "\n")
	b.WriteString(ui.Subtitle.Render("Тренировка устного счета") + "\n\n")
	for i, option := range m.options {
		if i > 0 {
			b.WriteString("\n")
		}
		button := ui.MenuItem(m.cursor == i, option)
		b.WriteString(zone.Mark(optionZoneID(i), button) + "\n")
	}

	return b.String()
}

func optionZoneID(index int) string {
	return fmt.Sprintf("start:option:%d", index)
}
