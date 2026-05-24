package math

const (
	MinExamplesCount     = 1
	MaxExamplesCount     = 50
	DefaultExamplesCount = 10
)

type TrainingSettings struct {
	AddDifficulty      Difficulty
	SubtractDifficulty Difficulty
	MultiplyDifficulty Difficulty
	DivideDifficulty   Difficulty
	ExamplesCount      int
}
