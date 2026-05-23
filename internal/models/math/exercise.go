package math

type Exercise struct {
	Left     int
	Right    int
	Operator Operator
}

type CurrentExercise struct {
	Order    int
	Total    int
	Exercise Exercise
}
