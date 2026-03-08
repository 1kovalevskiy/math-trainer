package task

import "github.com/1kovalevskiy/math-trainer/internal/app/tui/shared"

type GeneratedMsg struct {
	Exercise Exercise
}

type SubmitMsg struct {
	Result shared.ExampleResult
}

type SkipMsg struct {
	Result shared.ExampleResult
}

type BackMsg struct{}
