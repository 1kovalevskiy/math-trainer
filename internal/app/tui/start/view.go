package start

import (
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(ui.Title.Render("Математический тренажер") + "\n")
	b.WriteString(ui.Subtitle.Render("Тренировка устного счета") + "\n\n")
	for i, option := range m.options {
		b.WriteString(ui.MenuItem(m.cursor == i, option) + "\n")
	}

	b.WriteString("\n" + ui.Hint.Render("↑/↓ - выбор, Enter - подтвердить, Ctrl+C - выход"))

	return b.String()
}
