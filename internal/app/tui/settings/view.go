package settings

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(ui.Title.Render("Настройки тренировки") + "\n")
	b.WriteString(ui.Subtitle.Render("Параметры сессии перед стартом") + "\n\n")
	b.WriteString(settingLine(m.cursor == rowDifficulty, "Сложность", m.settings.Difficulty.String()) + "\n")
	b.WriteString(settingLine(m.cursor == rowExamplesCount, "Количество примеров", fmt.Sprintf("%d", m.settings.ExamplesCount)) + "\n\n")
	b.WriteString(ui.MenuItem(m.cursor == rowApply, "Применить") + "\n")
	b.WriteString(ui.MenuItem(m.cursor == rowBack, "Назад") + "\n")
	b.WriteString("\n" + ui.Hint.Render("↑/↓ - выбор пункта"))
	b.WriteString("\n" + ui.Hint.Render("←/→ - изменить значение"))
	b.WriteString("\n" + ui.Hint.Render("Enter - подтвердить, Esc - назад"))

	return b.String()
}

func settingLine(active bool, label string, value string) string {
	content := fmt.Sprintf("%s: %s", ui.Label.Render(label), ui.Value.Render(value))
	if active {
		return ui.MenuActive.Render("▸ " + content)
	}

	return ui.MenuInactive.Render("  " + content)
}
