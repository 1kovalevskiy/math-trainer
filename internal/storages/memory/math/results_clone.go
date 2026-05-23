package mathmemory

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func cloneResults(results []mathmodels.ExampleResult) []mathmodels.ExampleResult {
	if results == nil {
		return nil
	}

	cloned := make([]mathmodels.ExampleResult, len(results))
	for i, result := range results {
		cloned[i] = result
		if result.UserAnswer != nil {
			answer := *result.UserAnswer
			cloned[i].UserAnswer = &answer
		}
	}

	return cloned
}
