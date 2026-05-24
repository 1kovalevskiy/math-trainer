package tui

import (
	"context"
	"errors"
	"testing"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/settings"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
)

type persistTestController struct{}

func (persistTestController) GetDefaultSettings() mathmodels.TrainingSettings {
	return mathmodels.DefaultTrainingSettings()
}

func (persistTestController) NormalizeSettings(s mathmodels.TrainingSettings) mathmodels.TrainingSettings {
	return s
}

func (persistTestController) GetNextDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty {
	return current
}

func (persistTestController) GetPreviousDifficulty(current mathmodels.Difficulty) mathmodels.Difficulty {
	return current
}

func (persistTestController) StartTraining(context.Context, mathmodels.TrainingSettings) (mathmodels.TrainingSnapshot, error) {
	return mathmodels.TrainingSnapshot{}, nil
}

func (persistTestController) SubmitAnswer(context.Context, string) (mathmodels.TrainingSnapshot, error) {
	return mathmodels.TrainingSnapshot{}, nil
}

func (persistTestController) SkipCurrent(context.Context) (mathmodels.TrainingSnapshot, error) {
	return mathmodels.TrainingSnapshot{}, nil
}

func (persistTestController) CancelTraining(context.Context) error {
	return nil
}

type fakeSettingsStore struct {
	err   error
	saved mathmodels.TrainingSettings
}

func (s *fakeSettingsStore) SaveTrainingSettings(_ context.Context, settings mathmodels.TrainingSettings) error {
	s.saved = settings
	return s.err
}

func TestUpdateApplySettingsPersistsAndReturnsToStart(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := &fakeSettingsStore{}
	model := NewModel(ctx, persistTestController{}, store)
	model.screen = ScreenSettings

	desired := mathmodels.TrainingSettings{
		AddDifficulty:      mathmodels.DifficultyMedium,
		SubtractDifficulty: mathmodels.DifficultyEasy,
		MultiplyDifficulty: mathmodels.DifficultyDisabled,
		DivideDifficulty:   mathmodels.DifficultyDisabled,
		ExamplesCount:      12,
	}

	next, cmd := model.Update(settings.ApplySettingsMsg{Settings: desired})
	if cmd == nil {
		t.Fatal("expected persist command")
	}

	msg := cmd()
	next2, _ := next.Update(msg)
	updated, ok := next2.(Model)
	if !ok {
		t.Fatalf("model type mismatch: %T", next2)
	}

	if updated.screen != ScreenStart {
		t.Fatalf("screen mismatch: got %v", updated.screen)
	}
	if updated.settings != desired {
		t.Fatalf("settings mismatch: got %+v, want %+v", updated.settings, desired)
	}
	if store.saved != desired {
		t.Fatalf("saved settings mismatch: got %+v, want %+v", store.saved, desired)
	}
}

func TestUpdateApplySettingsKeepsSettingsScreenOnPersistError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := &fakeSettingsStore{err: errors.New("write failed")}
	model := NewModel(ctx, persistTestController{}, store)
	model.screen = ScreenSettings

	next, cmd := model.Update(settings.ApplySettingsMsg{Settings: mathmodels.DefaultTrainingSettings()})
	if cmd == nil {
		t.Fatal("expected persist command")
	}

	next2, _ := next.Update(cmd())
	updated, ok := next2.(Model)
	if !ok {
		t.Fatalf("model type mismatch: %T", next2)
	}

	if updated.screen != ScreenSettings {
		t.Fatalf("expected to stay on settings screen, got %v", updated.screen)
	}
	if updated.settings != mathmodels.DefaultTrainingSettings() {
		t.Fatalf("expected persisted settings to stay unchanged on error: got %+v", updated.settings)
	}
}
