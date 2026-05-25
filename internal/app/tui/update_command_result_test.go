package tui

import (
	"context"
	"testing"

	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

func TestHandleCommandResultMsgAppliesTrainingSnapshot(t *testing.T) {
	t.Parallel()

	model := NewModel(context.Background(), persistTestController{}, nil)
	snapshot := mathmodels.TrainingSnapshot{
		Phase:    mathmodels.TrainingPhaseInProgress,
		Settings: mathmodels.DefaultTrainingSettings(),
		Current: &mathmodels.CurrentExercise{
			Order: 1,
			Total: 1,
			Exercise: mathmodels.Exercise{
				Left:     1,
				Right:    2,
				Operator: mathmodels.OperatorAdd,
			},
		},
	}

	updated, cmd, handled := model.handleCommandResultMsg(trainingSnapshotMsg{snapshot: snapshot})

	if !handled {
		t.Fatal("expected training snapshot message to be handled")
	}
	if cmd != nil {
		t.Fatalf("expected no command, got %T", cmd())
	}
	if updated.screen != ScreenTask {
		t.Fatalf("screen mismatch: got %v, want %v", updated.screen, ScreenTask)
	}
}
