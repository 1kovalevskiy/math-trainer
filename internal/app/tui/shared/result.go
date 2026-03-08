package shared

type ResultStatus string

const (
	ResultStatusCorrect   ResultStatus = "correct"
	ResultStatusIncorrect ResultStatus = "incorrect"
	ResultStatusSkipped   ResultStatus = "skipped"
)

type ExampleResult struct {
	Order         int
	Expression    string
	CorrectAnswer int
	UserAnswer    *int
	Status        ResultStatus
}
