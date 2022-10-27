package site

import (
	"bbase/internal/task"
	"sync"
)

type ISite interface {
	Start(*task.Task) error
}

func RunTasks(site ISite, taskGroup *task.TaskGroup) error {
	wg := &sync.WaitGroup{}

	for _, t := range taskGroup.Tasks {
		wg.Add(1)
		go func(t *task.Task) {
			site.Start(t)
			wg.Done()
		}(t)
	}

	wg.Wait()

	return nil
}
