package example

import (
	"bbase/internal/logger"
	"bbase/internal/task"
)

func (i *exampleInternal) Checkout(t *task.Task) error {
	logger.SetTaskStatus(t.ID, "Checking out", "Example", "Test")
	return nil
}
