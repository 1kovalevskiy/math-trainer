package math

import "time"

type TrainingState struct {
	Settings        TrainingSettings
	CurrentOrder    int
	CurrentExercise Exercise
	Results         []ExampleResult
	Finished        bool
	StartedAt       time.Time
	FinishedAt      time.Time
}
