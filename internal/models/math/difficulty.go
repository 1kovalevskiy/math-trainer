package math

type Difficulty string

const (
	DifficultyDisabled Difficulty = "disabled"
	DifficultyEasy     Difficulty = "easy"
	DifficultyMedium   Difficulty = "medium"
	DifficultyHard     Difficulty = "hard"
)

func (d Difficulty) String() string {
	switch d {
	case DifficultyDisabled, DifficultyEasy, DifficultyMedium, DifficultyHard:
		return string(d)
	default:
		return "unknown"
	}
}
