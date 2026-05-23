package math

type TrainingSummary struct {
	Settings TrainingSettings
	Results  []ExampleResult
	Correct  int
	Total    int
}
