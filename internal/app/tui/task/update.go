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
	case tea.KeyMsg:
		switch typedMsg.String() {
		case "esc":
			return m, emit(BackMsg{})
		case "s":
			return m, emit(SkipMsg{
				Result: shared.ExampleResult{
					Order:         m.index,
					Expression:    m.exercise.Expression(),
					CorrectAnswer: m.expectedAnswer(),
					UserAnswer:    nil,
					Status:        shared.ResultStatusSkipped,
				},
			})
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
			m.errText = ""
			return m, nil
		case "enter":
			if m.input == "" {
				m.errText = "Введите ответ"
				return m, nil
			}

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
