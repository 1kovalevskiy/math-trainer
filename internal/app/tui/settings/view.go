package settings

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	zone "github.com/lrstanley/bubblezone"
)

const (
	zoneAddPrev      = "settings:add:prev"
	zoneAddNext      = "settings:add:next"
	zoneSubtractPrev = "settings:subtract:prev"
	zoneSubtractNext = "settings:subtract:next"
	zoneMultiplyPrev = "settings:multiply:prev"
	zoneMultiplyNext = "settings:multiply:next"
	zoneDividePrev   = "settings:divide:prev"
	zoneDivideNext   = "settings:divide:next"
	zoneCountDec     = "settings:count:dec"
	zoneCountInc     = "settings:count:inc"
	zoneApply        = "settings:apply"
	zoneBack         = "settings:back"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(ui.Title.Render("Настройки тренировки") + "\n")
	b.WriteString(ui.Subtitle.Render("Параметры сессии перед стартом") + "\n\n")
	b.WriteString(m.operationLine(m.cursor == rowAddDifficulty, "Сложение", m.settings.AddDifficulty, zoneAddPrev, zoneAddNext) + "\n")
	b.WriteString(m.operationLine(m.cursor == rowSubtractDifficulty, "Вычитание", m.settings.SubtractDifficulty, zoneSubtractPrev, zoneSubtractNext) + "\n")
	b.WriteString(m.operationLine(m.cursor == rowMultiplyDifficulty, "Умножение", m.settings.MultiplyDifficulty, zoneMultiplyPrev, zoneMultiplyNext) + "\n")
	b.WriteString(m.operationLine(m.cursor == rowDivideDifficulty, "Деление", m.settings.DivideDifficulty, zoneDividePrev, zoneDivideNext) + "\n")
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

func (m Model) operationLine(active bool, label string, difficulty mathmodels.Difficulty, prevZone, nextZone string) string {
	value := fmt.Sprintf(
		"%s %s %s",
		zone.Mark(prevZone, ui.SmallButton("←", active)),
		ui.Value.Render(shared.DifficultyLabel(difficulty)),
		zone.Mark(nextZone, ui.SmallButton("→", active)),
	)
	return settingLine(active, label, value)
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
