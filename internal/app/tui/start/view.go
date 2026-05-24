package start

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	zone "github.com/lrstanley/bubblezone"
)

func (m Model) View() string {
	var b strings.Builder
	buttonWidth := ui.MaxButtonWidth(m.options)

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

func (m Model) ViewWithSize(_, _ int) string {
	return m.View()
}

func optionZoneID(index int) string {
	return fmt.Sprintf("start:option:%d", index)
}
