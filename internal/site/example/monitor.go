package example

import (
	"bbase/internal/logger"
	"bbase/internal/task"
	"errors"
)

func (i *exampleInternal) Monitor(t *task.Task) error {
	logger.SetTaskStatus(t.ID, "Monitoring", "Example", "Test")

	// Enable this to simulate an error
	if false {
		return errors.New("example monitor failed")
	}

	return nil
}
