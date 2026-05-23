package mathcontroller

import (
	"context"
	"fmt"
)

func (c *Controller) CancelTraining(ctx context.Context) error {
	if err := c.validate(); err != nil {
		return err
	}
	if err := c.storage.ClearState(ctx); err != nil {
		return fmt.Errorf("clear training state: %w", err)
	}

	return nil
}
