package shared

const (
	MinExamplesCount     = 1
	MaxExamplesCount     = 50
	DefaultExamplesCount = 10
)

type TrainingSettings struct {
	Difficulty    Difficulty
	ExamplesCount int
}

func DefaultTrainingSettings() TrainingSettings {
	return TrainingSettings{
		Difficulty:    DifficultyEasy,
		ExamplesCount: DefaultExamplesCount,
	}
}

func NormalizeExamplesCount(count int) int {
	if count < MinExamplesCount {
		return MinExamplesCount
	}
	if count > MaxExamplesCount {
		return MaxExamplesCount
	}

	return count
}
