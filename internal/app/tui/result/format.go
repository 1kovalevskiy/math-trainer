package result

import (
	"fmt"
	"strings"
	"time"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func settingsDifficultySummary(settings mathmodels.TrainingSettings) string {
	parts := []string{
		fmt.Sprintf("+ %s", shared.DifficultyLabel(settings.AddDifficulty)),
		fmt.Sprintf("- %s", shared.DifficultyLabel(settings.SubtractDifficulty)),
		fmt.Sprintf("* %s", shared.DifficultyLabel(settings.MultiplyDifficulty)),
		fmt.Sprintf("/ %s", shared.DifficultyLabel(settings.DivideDifficulty)),
	}
	return strings.Join(parts, ", ")
}

func formatElapsed(elapsed time.Duration) string {
	if elapsed < 0 {
		elapsed = 0
	}
	elapsed = elapsed.Truncate(time.Second)

	totalSeconds := int(elapsed.Seconds())
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}
