package math

type Difficulty string

const (
	DifficultyDisabled Difficulty = "disabled"
	DifficultyStarter  Difficulty = "starter"
	DifficultyEasy     Difficulty = "easy"
	DifficultyMedium   Difficulty = "medium"
	DifficultyHard     Difficulty = "hard"
	DifficultyExpert   Difficulty = "expert"
)

func (d Difficulty) String() string {
	switch d {
	case DifficultyDisabled, DifficultyStarter, DifficultyEasy, DifficultyMedium, DifficultyHard, DifficultyExpert:
		return string(d)
	default:
		return "unknown"
	}
}
