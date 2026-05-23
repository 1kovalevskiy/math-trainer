package math

type ResultStatus string

const (
	ResultStatusCorrect   ResultStatus = "correct"
	ResultStatusIncorrect ResultStatus = "incorrect"
	ResultStatusSkipped   ResultStatus = "skipped"
)

func (s ResultStatus) String() string {
	switch s {
	case ResultStatusCorrect, ResultStatusIncorrect, ResultStatusSkipped:
		return string(s)
	default:
		return "unknown"
	}
}
