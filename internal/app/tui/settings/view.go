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

var activeRowMarkerStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("221")).
	Background(lipgloss.Color("62"))

var (
	settingsRowActiveStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("230")).
				Background(lipgloss.Color("62"))
	settingsRowInactiveStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("252")).
					Background(lipgloss.Color("236"))
)

type settingsLayout struct {
	labelWidth int
	valueWidth int
	rowWidth   int
	applyWidth int
	backWidth  int
}

func (m Model) View() string {
	var b strings.Builder
	layout := m.makeSettingsLayout()
	actionsLine := zone.Mark(zoneApply, ui.ButtonFixed("Применить", m.cursor == rowApply, layout.applyWidth)) + " " +
		zone.Mark(zoneBack, ui.ButtonFixed("Назад", m.cursor == rowBack, layout.backWidth))

	b.WriteString(ui.Title.Render("Настройки тренировки") + "\n")
	b.WriteString(ui.Subtitle.Render("Параметры сессии перед стартом") + "\n\n")
	if m.errText != "" {
		b.WriteString(ui.Error.Render(m.errText) + "\n\n")
	}
	b.WriteString(m.operationLine(m.cursor == rowAddDifficulty, "Сложение", m.settings.AddDifficulty, zoneAddPrev, zoneAddNext, layout) + "\n")
	b.WriteString(m.operationLine(m.cursor == rowSubtractDifficulty, "Вычитание", m.settings.SubtractDifficulty, zoneSubtractPrev, zoneSubtractNext, layout) + "\n")
	b.WriteString(m.operationLine(m.cursor == rowMultiplyDifficulty, "Умножение", m.settings.MultiplyDifficulty, zoneMultiplyPrev, zoneMultiplyNext, layout) + "\n")
	b.WriteString(m.operationLine(m.cursor == rowDivideDifficulty, "Деление", m.settings.DivideDifficulty, zoneDividePrev, zoneDivideNext, layout) + "\n")
	b.WriteString(
		settingLine(
			m.cursor == rowExamplesCount,
			"Количество примеров",
			m.countValue(m.cursor == rowExamplesCount),
			layout,
		) + "\n\n",
	)
	b.WriteString(actionsLine)

	return b.String()
}

func (m Model) operationLine(active bool, label string, difficulty mathmodels.Difficulty, prevZone, nextZone string, layout settingsLayout) string {
	return settingLine(active, label, m.operationValue(active, difficulty, prevZone, nextZone), layout)
}

func (m Model) operationValue(active bool, difficulty mathmodels.Difficulty, prevZone, nextZone string) string {
	rowStyle := settingsRowStyle(active)
	return fmt.Sprintf(
		"%s%s%s%s%s",
		zone.Mark(prevZone, ui.SmallButton("←", active)),
		rowStyle.Render(" "),
		rowStyle.Render(shared.DifficultyLabel(difficulty)),
		rowStyle.Render(" "),
		zone.Mark(nextZone, ui.SmallButton("→", active)),
	)
}

func (m Model) countValue(active bool) string {
	rowStyle := settingsRowStyle(active)
	return fmt.Sprintf(
		"%s%s%s%s%s",
		zone.Mark(zoneCountDec, ui.SmallButton("-", active)),
		rowStyle.Render(" "),
		rowStyle.Render(fmt.Sprintf("%d", m.settings.ExamplesCount)),
		rowStyle.Render(" "),
		zone.Mark(zoneCountInc, ui.SmallButton("+", active)),
	)
}

func settingLine(active bool, label string, value string, layout settingsLayout) string {
	rowStyle := settingsRowStyle(active)
	leftMarker := rowStyle.Render(" ")
	rightMarker := rowStyle.Render(" ")
	if active {
		leftMarker = activeRowMarkerStyle.Render("▸")
		rightMarker = activeRowMarkerStyle.Render("◂")
	}

	labelText := label + ":"
	labelPadding := max(0, layout.labelWidth-lipgloss.Width(labelText))
	valuePadding := max(0, layout.valueWidth-lipgloss.Width(value))
	valuePadLeft := valuePadding / 2
	valuePadRight := valuePadding - valuePadLeft
	content := rowStyle.Render("  ") +
		leftMarker +
		rowStyle.Render(labelText) +
		rowStyle.Render(strings.Repeat(" ", labelPadding)) +
		rowStyle.Render(" ") +
		rowStyle.Render(strings.Repeat(" ", valuePadLeft)) +
		value +
		rowStyle.Render(strings.Repeat(" ", valuePadRight)) +
		rightMarker +
		rowStyle.Render("  ")
	pad := max(0, layout.rowWidth-lipgloss.Width(content))
	return content + rowStyle.Render(strings.Repeat(" ", pad))
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

func (m Model) settingsValueWidth(countValue string) int {
	valueWidth := lipgloss.Width(m.countValue(false))
	valueWidth = max(valueWidth, lipgloss.Width(m.countValue(true)))
	diffRows := []struct {
		difficulty mathmodels.Difficulty
	}{
		{difficulty: m.settings.AddDifficulty},
		{difficulty: m.settings.SubtractDifficulty},
		{difficulty: m.settings.MultiplyDifficulty},
		{difficulty: m.settings.DivideDifficulty},
	}

	for _, row := range diffRows {
		valueWidth = max(valueWidth, lipgloss.Width(m.operationValue(false, row.difficulty, "", "")))
		valueWidth = max(valueWidth, lipgloss.Width(m.operationValue(true, row.difficulty, "", "")))
	}

	return valueWidth
}

func settingsRowWidth(labelWidth int, valueWidth int) int {
	return 2 + 1 + labelWidth + 1 + valueWidth + 1 + 2
}

func (m Model) settingsVisualRowWidth(layout settingsLayout) int {
	width := settingsRowWidth(layout.labelWidth, layout.valueWidth)
	width = max(width, lipgloss.Width(m.operationLine(m.cursor == rowAddDifficulty, "Сложение", m.settings.AddDifficulty, zoneAddPrev, zoneAddNext, layout)))
	width = max(width, lipgloss.Width(m.operationLine(m.cursor == rowSubtractDifficulty, "Вычитание", m.settings.SubtractDifficulty, zoneSubtractPrev, zoneSubtractNext, layout)))
	width = max(width, lipgloss.Width(m.operationLine(m.cursor == rowMultiplyDifficulty, "Умножение", m.settings.MultiplyDifficulty, zoneMultiplyPrev, zoneMultiplyNext, layout)))
	width = max(width, lipgloss.Width(m.operationLine(m.cursor == rowDivideDifficulty, "Деление", m.settings.DivideDifficulty, zoneDividePrev, zoneDivideNext, layout)))
	width = max(width, lipgloss.Width(settingLine(m.cursor == rowExamplesCount, "Количество примеров", m.countValue(m.cursor == rowExamplesCount), layout)))
	return width
}

func (m Model) makeSettingsLayout() settingsLayout {
	layout := settingsLayout{
		labelWidth: settingsLabelWidth(),
		valueWidth: m.settingsValueWidth(""),
	}
	layout.rowWidth = m.settingsVisualRowWidth(layout)

	minActionsWidth := lipgloss.Width(ui.Button("Применить", false) + " " + ui.Button("Назад", false))
	if layout.rowWidth < minActionsWidth {
		layout.rowWidth = minActionsWidth
	}

	layout.applyWidth, layout.backWidth = actionContentWidths(
		layout.rowWidth,
		"Применить",
		m.cursor == rowApply,
		"Назад",
		m.cursor == rowBack,
	)

	return layout
}

func settingsRowStyle(active bool) lipgloss.Style {
	if active {
		return settingsRowActiveStyle
	}
	return settingsRowInactiveStyle
}

func actionContentWidths(totalLineWidth int, applyText string, applyActive bool, backText string, backActive bool) (int, int) {
	applyMin := lipgloss.Width(applyText)
	backMin := lipgloss.Width(backText)

	applyWidth := applyMin
	backWidth := backMin
	for {
		line := ui.ButtonFixed(applyText, applyActive, applyWidth) + " " + ui.ButtonFixed(backText, backActive, backWidth)
		current := lipgloss.Width(line)
		if current >= totalLineWidth {
			break
		}
		if applyWidth <= backWidth {
			applyWidth++
		} else {
			backWidth++
		}
	}

	return applyWidth, backWidth
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
