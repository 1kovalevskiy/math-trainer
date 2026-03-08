package settings

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"

type ApplyDifficultyMsg struct {
	Difficulty shared.Difficulty
}

type BackMsg struct{}
