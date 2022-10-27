package example

import (
	"bbase/internal/logger"
	"bbase/internal/profile"
	"bbase/internal/proxy"
	"bbase/internal/task"
	"time"
)

type Example struct {
	Steps []func(*exampleInternal, *task.Task) error

	// Add any other fields you need here
	// These fields are shared between all tasks in the group
	//
	// SKUs map[string]string
	// etc ...
}

type exampleInternal struct {
	global *Example

	proxy   *proxy.Proxy
	profile *profile.Profile

	// Add any other fields you need here
	// These fields are unique to each task
	//
	// CSRFToken string
	// etc ...
}

func NewExample() *Example {
	return &Example{
		Steps: []func(*exampleInternal, *task.Task) error{
			(*exampleInternal).Monitor,
			(*exampleInternal).Checkout,
		},
	}
}

func (e *Example) Start(t *task.Task) error {
	logger.SetTaskStatus(t.ID, "Starting", "Example", "Test")

	internal := &exampleInternal{
		global:  e,
		proxy:   task.GetRandomProxyFromGroup(t.ParentGroup),
		profile: task.GetRandomProfileFromGroup(t.ParentGroup),
	}

	for i := 0; i < len(e.Steps); i++ {
		select {
		case <-t.Ctx.Done():
			logger.SetTaskStatus(t.ID, "Stopped", "Example", "Test")
			return nil
		default:
		}

		err := e.Steps[i](internal, t)
		if err != nil {
			time.Sleep(t.ParentGroup.ErrorDelay)
			i--
		}
	}

	logger.SetTaskStatus(t.ID, "Finished", "Example", "Test")

	return nil
}
