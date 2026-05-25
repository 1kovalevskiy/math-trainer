package mathcontroller

func WithClock(clock Clock) Option {
	return func(controller *Controller) {
		if clock != nil {
			controller.clock = clock
		}
	}
}
