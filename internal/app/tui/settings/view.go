package settings

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	zone "github.com/lrstanley/bubblezone"
)

const (
	zoneDifficultyPrev = "settings:difficulty:prev"
	zoneDifficultyNext = "settings:difficulty:next"
	zoneCountDec       = "settings:count:dec"
	zoneCountInc       = "settings:count:inc"
	zoneApply          = "settings:apply"
	zoneBack           = "settings:back"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(ui.Title.Render("Настройки тренировки") + "\n")
	b.WriteString(ui.Subtitle.Render("Параметры сессии перед стартом") + "\n\n")
	b.WriteString(
		settingLine(
			m.cursor == rowDifficulty,
			"Сложность",
			fmt.Sprintf(
				"%s %s %s",
				zone.Mark(zoneDifficultyPrev, ui.SmallButton("←", m.cursor == rowDifficulty)),
				ui.Value.Render(shared.DifficultyLabel(m.settings.Difficulty)),
				zone.Mark(zoneDifficultyNext, ui.SmallButton("→", m.cursor == rowDifficulty)),
			),
		) + "\n",
	)
	b.WriteString(
		settingLine(
			m.cursor == rowExamplesCount,
			"Количество примеров",
			fmt.Sprintf(
				"%s %s %s",
				zone.Mark(zoneCountDec, ui.SmallButton("-", m.cursor == rowExamplesCount)),
				ui.Value.Render(fmt.Sprintf("%d", m.settings.ExamplesCount)),
				zone.Mark(zoneCountInc, ui.SmallButton("+", m.cursor == rowExamplesCount)),
			),
		) + "\n\n",
	)
	b.WriteString(zone.Mark(zoneApply, ui.Button("Применить", m.cursor == rowApply)))
	b.WriteString(" ")
	b.WriteString(zone.Mark(zoneBack, ui.Button("Назад", m.cursor == rowBack)))

	return b.String()
}

func settingLine(active bool, label string, value string) string {
	marker := "  "
	if active {
		marker = ui.Accent.Render("▸ ")
	}

	content := fmt.Sprintf("%s%s: %s", marker, ui.Label.Render(label), value)
	if active {
		return ui.Button(content, true)
	}

	return ui.Button(content, false)
}
