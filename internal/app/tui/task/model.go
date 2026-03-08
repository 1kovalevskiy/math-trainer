package task

import (
	"fmt"

	"github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

type Exercise struct {
	Left     int
	Right    int
	Operator string
}

func (e Exercise) Expression() string {
	return fmt.Sprintf("%d %s %d", e.Left, e.Operator, e.Right)
}

type Model struct {
	difficulty shared.Difficulty
	index      int
	total      int
	exercise   Exercise
	input      string
	errText    string
}

func NewModel(difficulty shared.Difficulty, index int, total int) Model {
	return Model{
		difficulty: difficulty,
		index:      index,
		total:      total,
	}
}

func (m Model) Init() tea.Cmd {
	return GenerateExerciseCmd(m.difficulty)
}

func (m Model) expectedAnswer() int {
	switch m.exercise.Operator {
	case "+":
		return m.exercise.Left + m.exercise.Right
	case "-":
		return m.exercise.Left - m.exercise.Right
	default:
		return 0
	}
}
