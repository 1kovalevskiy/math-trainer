package settings

import "testing"

func TestSettingsFocusMovesBetweenRowsAndActions(t *testing.T) {
	t.Parallel()

	focus := newSettingsFocus()

	for range 4 {
		focus = focus.moveDown()
	}
	if !focus.isSetting(settingExamplesCount) {
		t.Fatalf("expected examples row focus, got %+v", focus)
	}

	focus = focus.moveDown()
	if !focus.isAction(actionApply) {
		t.Fatalf("expected apply action focus, got %+v", focus)
	}

	focus = focus.moveRight()
	if !focus.isAction(actionBack) {
		t.Fatalf("expected back action focus, got %+v", focus)
	}

	focus = focus.moveUp()
	if !focus.isSetting(settingExamplesCount) {
		t.Fatalf("expected examples row after moving up from actions, got %+v", focus)
	}
}
