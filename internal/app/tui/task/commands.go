package task

import (
	"math/rand"
	"time"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

func GenerateExerciseCmd(difficulty shared.Difficulty) tea.Cmd {
	return func() tea.Msg {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		min, max := shared.NumbersRange(difficulty)

		left := random.Intn(max-min+1) + min
		right := random.Intn(max-min+1) + min

		operators := []string{"+", "-"}
		operator := operators[random.Intn(len(operators))]
		if operator == "-" && right > left {
			left, right = right, left
		}

		return GeneratedMsg{Exercise: Exercise{Left: left, Right: right, Operator: operator}}
	}
}

func emit(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}
