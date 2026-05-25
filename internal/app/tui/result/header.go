package result

import (
	"fmt"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func renderSummaryHeaderLines(summary mathmodels.TrainingSummary) []string {
	return []string{
		uiTitle("Результаты тренировки"),
		resultSubtitleStyle.Render("Сводка по всем примерам"),
		"",
		resultLabelStyle.Render("Сложности: ") + resultValueStyle.Render(settingsDifficultySummary(summary.Settings)),
		resultAccentStyle.Render(fmt.Sprintf("Правильных: %d из %d", summary.Correct, summary.Total)),
		resultLabelStyle.Render("Время: ") + resultValueStyle.Render(formatElapsed(summary.Elapsed)),
		"",
	}
}

func uiTitle(text string) string {
	return resultTitleStyle.Render(text)
}
