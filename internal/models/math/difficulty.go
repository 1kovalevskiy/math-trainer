package math

type Difficulty string

const (
	DifficultyEasy   Difficulty = "easy"
	DifficultyMedium Difficulty = "medium"
	DifficultyHard   Difficulty = "hard"
)

func (d Difficulty) String() string {
	switch d {
	case DifficultyEasy, DifficultyMedium, DifficultyHard:
		return string(d)
	default:
		return "unknown"
	}
}
