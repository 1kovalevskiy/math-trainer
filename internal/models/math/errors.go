package math

import "errors"

var (
	ErrNoActiveTraining        = errors.New("no active training")
	ErrTrainingFinished        = errors.New("training is finished")
	ErrInvalidAnswer           = errors.New("invalid answer")
	ErrUniqueExerciseExhausted = errors.New("unique exercise exhausted")
)
