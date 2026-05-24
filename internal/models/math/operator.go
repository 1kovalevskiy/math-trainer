package math

type Operator string

const (
	OperatorAdd      Operator = "add"
	OperatorSubtract Operator = "subtract"
	OperatorMultiply Operator = "multiply"
	OperatorDivide   Operator = "divide"
)

func (o Operator) String() string {
	switch o {
	case OperatorAdd, OperatorSubtract, OperatorMultiply, OperatorDivide:
		return string(o)
	default:
		return "unknown"
	}
}
