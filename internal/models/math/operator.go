package math

type Operator string

const (
	OperatorAdd      Operator = "add"
	OperatorSubtract Operator = "subtract"
)

func (o Operator) String() string {
	switch o {
	case OperatorAdd, OperatorSubtract:
		return string(o)
	default:
		return "unknown"
	}
}
