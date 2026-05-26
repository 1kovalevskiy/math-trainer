package task

import (
	"unicode"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch typedMsg := msg.(type) {
	case tea.MouseMsg:
		if !shared.IsLeftClick(typedMsg) {
			return m, nil
		}

		switch {
		case shared.InZone(zoneSubmit, typedMsg):
			m = m.withButtonCursor(buttonSubmit)
			if !m.canSubmit() {
				return m, nil
			}
			return m.submitAnswer()
		case shared.InZone(zoneSkip, typedMsg):
			m = m.withButtonCursor(buttonSkip)
			return m.skipCurrent()
		case shared.InZone(zoneBack, typedMsg):
			m = m.withButtonCursor(buttonBack)
			return m, emit(BackMsg{})
		default:
			return m, nil
		}
	case tea.KeyMsg:
		switch typedMsg.String() {
		case "esc":
			return m, emit(BackMsg{})
		case "s":
			m = m.withButtonCursor(buttonSkip)
			return m.skipCurrent()
		case "left", "h":
			m = m.moveButtonCursor(-1)
			return m, nil
		case "right", "l":
			m = m.moveButtonCursor(1)
			return m, nil
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
			m.errText = ""
			return m, nil
		case "enter":
			return m.selectCurrent()
		default:
			if len(typedMsg.Runes) != 1 {
				return m, nil
			}

			inputRune := typedMsg.Runes[0]
			if unicode.IsDigit(inputRune) || inputRune == '-' && m.input == "" {
				m.input += string(typedMsg.Runes)
				m.errText = ""
			}
			return m, nil
		}
	}

	return m, nil
}

func (m Model) selectCurrent() (Model, tea.Cmd) {
	switch m.buttonCursor {
	case buttonSubmit:
		if !m.canSubmit() {
			return m, nil
		}
		return m.submitAnswer()
	case buttonSkip:
		return m.skipCurrent()
	case buttonBack:
		return m, emit(BackMsg{})
	default:
		return m, nil
	}
}

func (m Model) submitAnswer() (Model, tea.Cmd) {
	return m, emit(SubmitMsg{Answer: m.input})
}

func (m Model) skipCurrent() (Model, tea.Cmd) {
	return m, emit(SkipMsg{})
}
