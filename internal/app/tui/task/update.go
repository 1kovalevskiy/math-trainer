package task

import (
	"strconv"
	"unicode"

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
		case "r":
			return m, GenerateExerciseCmd(m.difficulty)
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
			return m, emit(SubmitMsg{
				Difficulty: m.difficulty,
				Expression: m.exercise.Expression(),
				Expected:   expected,
				Answer:     answer,
				Correct:    answer == expected,
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
