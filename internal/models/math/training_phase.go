package math

type TrainingPhase string

const (
	TrainingPhaseIdle       TrainingPhase = "idle"
	TrainingPhaseInProgress TrainingPhase = "in_progress"
	TrainingPhaseFinished   TrainingPhase = "finished"
)

func (p TrainingPhase) String() string {
	switch p {
	case TrainingPhaseIdle, TrainingPhaseInProgress, TrainingPhaseFinished:
		return string(p)
	default:
		return "unknown"
	}
}
