package math

type TrainingState struct {
	Settings        TrainingSettings
	CurrentOrder    int
	CurrentExercise Exercise
	Results         []ExampleResult
	Finished        bool
}
