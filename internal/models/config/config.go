package configmodel

import mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"

type Config struct {
	App      App      `json:"app" yaml:"app"`
	Training Training `json:"training" yaml:"training"`
}

type App struct {
	LogLevel string `json:"log_level" yaml:"log_level" env:"APP_LOG_LEVEL" env-default:"INFO"`
}

type Training struct {
	AddDifficulty      mathmodels.Difficulty `json:"add_difficulty" yaml:"add_difficulty"`
	SubtractDifficulty mathmodels.Difficulty `json:"subtract_difficulty" yaml:"subtract_difficulty"`
	MultiplyDifficulty mathmodels.Difficulty `json:"multiply_difficulty" yaml:"multiply_difficulty"`
	DivideDifficulty   mathmodels.Difficulty `json:"divide_difficulty" yaml:"divide_difficulty"`
	ExamplesCount      int                   `json:"examples_count" yaml:"examples_count"`
}

func Default() Config {
	return Config{
		App:      App{LogLevel: "INFO"},
		Training: TrainingFromMath(mathmodels.DefaultTrainingSettings()),
	}
}

func (t Training) ToMath() mathmodels.TrainingSettings {
	return mathmodels.TrainingSettings{
		AddDifficulty:      t.AddDifficulty,
		SubtractDifficulty: t.SubtractDifficulty,
		MultiplyDifficulty: t.MultiplyDifficulty,
		DivideDifficulty:   t.DivideDifficulty,
		ExamplesCount:      t.ExamplesCount,
	}
}

func TrainingFromMath(settings mathmodels.TrainingSettings) Training {
	return Training{
		AddDifficulty:      settings.AddDifficulty,
		SubtractDifficulty: settings.SubtractDifficulty,
		MultiplyDifficulty: settings.MultiplyDifficulty,
		DivideDifficulty:   settings.DivideDifficulty,
		ExamplesCount:      settings.ExamplesCount,
	}
}
