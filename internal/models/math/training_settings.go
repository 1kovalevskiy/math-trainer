package math

const (
	MinExamplesCount     = 1
	MaxExamplesCount     = 50
	DefaultExamplesCount = 10
)

type TrainingSettings struct {
	Difficulty    Difficulty
	ExamplesCount int
}
