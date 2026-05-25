package result

import (
	"fmt"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	zone "github.com/lrstanley/bubblezone"
)

func (m Model) renderActionButtons() string {
	buttons := make([]string, 0, len(m.options))
	for i, option := range m.options {
		buttons = append(buttons, zone.Mark(optionZoneID(i), ui.MenuItem(m.cursor == i, option)))
	}
	return ui.JoinInline(buttons, 1)
}

func optionZoneID(index int) string {
	return fmt.Sprintf("result:option:%d", index)
}
