package shared

type Difficulty string

const (
	DifficultyEasy   Difficulty = "easy"
	DifficultyMedium Difficulty = "medium"
	DifficultyHard   Difficulty = "hard"
)

func AllDifficulties() []Difficulty {
	return []Difficulty{DifficultyEasy, DifficultyMedium, DifficultyHard}
}

func (d Difficulty) String() string {
	switch d {
	case DifficultyEasy:
		return "Легко"
	case DifficultyMedium:
		return "Средне"
	case DifficultyHard:
		return "Сложно"
	default:
		return "Неизвестно"
	}
}

func NumbersRange(d Difficulty) (min int, max int) {
	switch d {
	case DifficultyEasy:
		return 1, 10
	case DifficultyMedium:
		return 10, 50
	case DifficultyHard:
		return 50, 100
	default:
		return 1, 10
	}
}
