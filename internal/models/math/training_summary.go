package math

import "time"

type TrainingSummary struct {
	Settings TrainingSettings
	Results  []ExampleResult
	Correct  int
	Total    int
	Elapsed  time.Duration
}
