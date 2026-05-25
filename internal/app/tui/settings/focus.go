package settings

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"

type focusSection int

const (
	focusSettings focusSection = iota
	focusActions
)

type settingRow int

const (
	settingAddDifficulty settingRow = iota
	settingSubtractDifficulty
	settingMultiplyDifficulty
	settingDivideDifficulty
	settingExamplesCount
)

type actionButton int

const (
	actionApply actionButton = iota
	actionBack
)

type settingsFocus struct {
	section focusSection
	setting settingRow
	action  actionButton
}

func newSettingsFocus() settingsFocus {
	return settingsFocus{section: focusSettings, setting: settingAddDifficulty}
}

func allSettingsFocuses() []settingsFocus {
	return []settingsFocus{
		{section: focusSettings, setting: settingAddDifficulty},
		{section: focusSettings, setting: settingSubtractDifficulty},
		{section: focusSettings, setting: settingMultiplyDifficulty},
		{section: focusSettings, setting: settingDivideDifficulty},
		{section: focusSettings, setting: settingExamplesCount},
		{section: focusActions, action: actionApply},
		{section: focusActions, action: actionBack},
	}
}

func (f settingsFocus) isSetting(row settingRow) bool {
	return f.section == focusSettings && f.setting == row
}

func (f settingsFocus) isAction(action actionButton) bool {
	return f.section == focusActions && f.action == action
}

func (f settingsFocus) moveUp() settingsFocus {
	if f.section == focusActions {
		return settingsFocus{section: focusSettings, setting: settingExamplesCount}
	}
	f.setting = settingRow(ui.MoveIndex(int(f.setting), -1, int(settingAddDifficulty), int(settingExamplesCount)))
	return f
}

func (f settingsFocus) moveDown() settingsFocus {
	if f.section == focusActions {
		return f
	}
	if f.setting == settingExamplesCount {
		return settingsFocus{section: focusActions, action: actionApply}
	}
	f.setting = settingRow(ui.MoveIndex(int(f.setting), 1, int(settingAddDifficulty), int(settingExamplesCount)))
	return f
}

func (f settingsFocus) moveLeft() settingsFocus {
	if f.section != focusActions {
		return f
	}
	f.action = actionButton(ui.MoveIndex(int(f.action), -1, int(actionApply), int(actionBack)))
	return f
}

func (f settingsFocus) moveRight() settingsFocus {
	if f.section != focusActions {
		return f
	}
	f.action = actionButton(ui.MoveIndex(int(f.action), 1, int(actionApply), int(actionBack)))
	return f
}
