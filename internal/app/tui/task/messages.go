package task

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"

type GeneratedMsg struct {
	Exercise Exercise
}

type SubmitMsg struct {
	Difficulty shared.Difficulty
	Expression string
	Expected   int
	Answer     int
	Correct    bool
}

type BackMsg struct{}
