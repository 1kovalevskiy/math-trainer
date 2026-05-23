package math

type TrainingSnapshot struct {
	Phase    TrainingPhase
	Settings TrainingSettings
	Current  *CurrentExercise
	Summary  *TrainingSummary
}
