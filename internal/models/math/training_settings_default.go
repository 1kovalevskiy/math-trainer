package math

func DefaultTrainingSettings() TrainingSettings {
	return TrainingSettings{
		AddDifficulty:      DifficultyStarter,
		SubtractDifficulty: DifficultyStarter,
		MultiplyDifficulty: DifficultyDisabled,
		DivideDifficulty:   DifficultyDisabled,
		ExamplesCount:      DefaultExamplesCount,
	}
}
