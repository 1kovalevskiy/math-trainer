package settings

import (
	"fmt"
	"strings"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	"github.com/charmbracelet/lipgloss"
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

type settingsLayout struct {
	row        ui.KeyValueRowLayout
	applyWidth int
	backWidth  int
}

func (m Model) View() string {
	var b strings.Builder
	layout := m.makeSettingsLayout()
	actionsLine := ui.JoinInline([]string{
		zone.Mark(zoneApply, ui.ButtonFixed("Применить", m.focus.isAction(actionApply), layout.applyWidth)),
		zone.Mark(zoneBack, ui.ButtonFixed("Назад", m.focus.isAction(actionBack), layout.backWidth)),
	}, 1)

	b.WriteString(ui.Title.Render("Настройки тренировки") + "\n")
	b.WriteString(ui.Subtitle.Render("Параметры сессии перед стартом") + "\n\n")
	if m.errText != "" {
		b.WriteString(ui.Error.Render(m.errText) + "\n\n")
	}
	b.WriteString(m.operationLine(m.focus.isSetting(settingAddDifficulty), "Сложение", m.settings.AddDifficulty, zoneAddPrev, zoneAddNext, layout) + "\n")
	b.WriteString(m.operationLine(m.focus.isSetting(settingSubtractDifficulty), "Вычитание", m.settings.SubtractDifficulty, zoneSubtractPrev, zoneSubtractNext, layout) + "\n")
	b.WriteString(m.operationLine(m.focus.isSetting(settingMultiplyDifficulty), "Умножение", m.settings.MultiplyDifficulty, zoneMultiplyPrev, zoneMultiplyNext, layout) + "\n")
	b.WriteString(m.operationLine(m.focus.isSetting(settingDivideDifficulty), "Деление", m.settings.DivideDifficulty, zoneDividePrev, zoneDivideNext, layout) + "\n")
	b.WriteString(
		settingLine(
			m.focus.isSetting(settingExamplesCount),
			"Количество примеров",
			m.countValue(m.focus.isSetting(settingExamplesCount)),
			layout,
		) + "\n\n",
	)
	b.WriteString(actionsLine)

	return b.String()
}

func (m Model) ViewWithSize(_, _ int) string {
	return m.View()
}

func (m Model) operationLine(active bool, label string, difficulty mathmodels.Difficulty, prevZone, nextZone string, layout settingsLayout) string {
	return settingLine(active, label, m.operationValue(active, difficulty, prevZone, nextZone), layout)
}

func (m Model) operationValue(active bool, difficulty mathmodels.Difficulty, prevZone, nextZone string) []ui.Segment {
	rowStyle := ui.SettingRowStyle(active)
	return []ui.Segment{
		{Text: zone.Mark(prevZone, ui.SmallButton("←", active))},
		{Text: " ", Style: rowStyle},
		{Text: shared.DifficultyLabel(difficulty), Style: rowStyle},
		{Text: " ", Style: rowStyle},
		{Text: zone.Mark(nextZone, ui.SmallButton("→", active))},
	}
}

func (m Model) countValue(active bool) []ui.Segment {
	rowStyle := ui.SettingRowStyle(active)
	return []ui.Segment{
		{Text: zone.Mark(zoneCountDec, ui.SmallButton("-", active))},
		{Text: " ", Style: rowStyle},
		{Text: fmt.Sprintf("%d", m.settings.ExamplesCount), Style: rowStyle},
		{Text: " ", Style: rowStyle},
		{Text: zone.Mark(zoneCountInc, ui.SmallButton("+", active))},
	}
}

func settingLine(active bool, label string, value []ui.Segment, layout settingsLayout) string {
	return ui.RenderKeyValueRow(ui.KeyValueRow{
		Active: active,
		Label:  label,
		Value:  value,
		Layout: layout.row,
	})
}

func settingsLabelWidth() int {
	return max(
		lipgloss.Width("Сложение:"),
		lipgloss.Width("Вычитание:"),
		lipgloss.Width("Умножение:"),
		lipgloss.Width("Деление:"),
		lipgloss.Width("Количество примеров:"),
	)
}

func (m Model) settingsValueWidth() int {
	valueWidth := lipgloss.Width(ui.RenderSegments(m.countValue(false)...))
	valueWidth = max(valueWidth, lipgloss.Width(ui.RenderSegments(m.countValue(true)...)))
	difficulties := []mathmodels.Difficulty{
		mathmodels.DifficultyDisabled,
		mathmodels.DifficultyStarter,
		mathmodels.DifficultyEasy,
		mathmodels.DifficultyMedium,
		mathmodels.DifficultyHard,
		mathmodels.DifficultyExpert,
	}

	for _, difficulty := range difficulties {
		valueWidth = max(valueWidth, lipgloss.Width(ui.RenderSegments(m.operationValue(false, difficulty, "", "")...)))
		valueWidth = max(valueWidth, lipgloss.Width(ui.RenderSegments(m.operationValue(true, difficulty, "", "")...)))
	}

	return valueWidth
}

func settingsRowWidth(labelWidth int, valueWidth int) int {
	return ui.KeyValueBaseRowWidth(labelWidth, valueWidth)
}

func (m Model) settingsVisualRowWidth(layout settingsLayout) int {
	width := settingsRowWidth(layout.row.LabelWidth, layout.row.ValueWidth)
	width = max(width, lipgloss.Width(m.operationLine(m.focus.isSetting(settingAddDifficulty), "Сложение", m.settings.AddDifficulty, zoneAddPrev, zoneAddNext, layout)))
	width = max(width, lipgloss.Width(m.operationLine(m.focus.isSetting(settingSubtractDifficulty), "Вычитание", m.settings.SubtractDifficulty, zoneSubtractPrev, zoneSubtractNext, layout)))
	width = max(width, lipgloss.Width(m.operationLine(m.focus.isSetting(settingMultiplyDifficulty), "Умножение", m.settings.MultiplyDifficulty, zoneMultiplyPrev, zoneMultiplyNext, layout)))
	width = max(width, lipgloss.Width(m.operationLine(m.focus.isSetting(settingDivideDifficulty), "Деление", m.settings.DivideDifficulty, zoneDividePrev, zoneDivideNext, layout)))
	width = max(width, lipgloss.Width(settingLine(m.focus.isSetting(settingExamplesCount), "Количество примеров", m.countValue(m.focus.isSetting(settingExamplesCount)), layout)))
	return width
}

func (m Model) makeSettingsLayout() settingsLayout {
	layout := settingsLayout{
		row: ui.KeyValueRowLayout{
			LabelWidth: settingsLabelWidth(),
			ValueWidth: m.settingsValueWidth(),
		},
	}
	layout.row.RowWidth = m.settingsVisualRowWidth(layout)

	minActionsWidth := lipgloss.Width(ui.Button("Применить", false) + " " + ui.Button("Назад", false))
	if layout.row.RowWidth < minActionsWidth {
		layout.row.RowWidth = minActionsWidth
	}

	layout.applyWidth, layout.backWidth = ui.StretchTwoButtonWidths(
		layout.row.RowWidth,
		"Применить",
		m.focus.isAction(actionApply),
		"Назад",
		m.focus.isAction(actionBack),
	)

	return layout
}

func max(values ...int) int {
	current := 0
	for _, value := range values {
		if value > current {
			current = value
		}
	}
	return current
}
