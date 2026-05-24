package math

func DefaultTrainingSettings() TrainingSettings {
	return TrainingSettings{
		AddDifficulty:      DifficultyEasy,
		SubtractDifficulty: DifficultyEasy,
		MultiplyDifficulty: DifficultyDisabled,
		DivideDifficulty:   DifficultyDisabled,
		ExamplesCount:      DefaultExamplesCount,
	}
}
