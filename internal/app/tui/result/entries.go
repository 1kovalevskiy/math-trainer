package result

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

func renderResultEntries(results []mathmodels.ExampleResult) []string {
	entries := make([]string, 0, len(results))
	for _, entry := range results {
		entries = append(entries, renderEntry(entry))
	}

	return entries
}
