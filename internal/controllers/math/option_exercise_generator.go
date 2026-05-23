package mathcontroller

func WithExerciseGenerator(generator ExerciseGenerator) Option {
	return func(controller *Controller) {
		if generator != nil {
			controller.generate = generator
		}
	}
}
