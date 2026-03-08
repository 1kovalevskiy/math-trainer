package task

import (
	"strconv"
	"unicode"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch typedMsg := msg.(type) {
	case GeneratedMsg:
		m.exercise = typedMsg.Exercise
		m.input = ""
		m.errText = ""
		return m, nil
	case tea.MouseMsg:
		if !shared.IsLeftClick(typedMsg) {
			return m, nil
		}

		switch {
		case shared.InZone(zoneSubmit, typedMsg):
			if !m.canSubmit() {
				return m, nil
			}
			return m.submitAnswer()
		case shared.InZone(zoneSkip, typedMsg):
			return m.skipCurrent()
		case shared.InZone(zoneBack, typedMsg):
			return m, emit(BackMsg{})
		default:
			return m, nil
		}
	case tea.KeyMsg:
		switch typedMsg.String() {
		case "esc":
			return m, emit(BackMsg{})
		case "s":
			return m.skipCurrent()
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
			m.errText = ""
			return m, nil
		case "enter":
			if !m.canSubmit() {
				return m, nil
			}
			return m.submitAnswer()
		default:
			if len(typedMsg.Runes) == 1 && unicode.IsDigit(typedMsg.Runes[0]) {
				m.input += string(typedMsg.Runes)
				m.errText = ""
			}
			return m, nil
		}
	}

	return m, nil
}

func (m Model) submitAnswer() (Model, tea.Cmd) {
	answer, err := strconv.Atoi(m.input)
	if err != nil {
		m.errText = "Ответ должен быть числом"
		return m, nil
	}

	expected := m.expectedAnswer()
	status := shared.ResultStatusIncorrect
	if answer == expected {
		status = shared.ResultStatusCorrect
	}
	answerCopy := answer

	return m, emit(SubmitMsg{
		Result: shared.ExampleResult{
			Order:         m.index,
			Expression:    m.exercise.Expression(),
			CorrectAnswer: expected,
			UserAnswer:    &answerCopy,
			Status:        status,
		},
	})
}

func (m Model) skipCurrent() (Model, tea.Cmd) {
	return m, emit(SkipMsg{
		Result: shared.ExampleResult{
			Order:         m.index,
			Expression:    m.exercise.Expression(),
			CorrectAnswer: m.expectedAnswer(),
			UserAnswer:    nil,
			Status:        shared.ResultStatusSkipped,
		},
	})
}
