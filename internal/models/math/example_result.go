package math

type ExampleResult struct {
	Order         int
	Exercise      Exercise
	CorrectAnswer int
	UserAnswer    *int
	Status        ResultStatus
}
