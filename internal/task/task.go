package task

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type Task struct {
	ParentGroup *TaskGroup
	ID          string
	Ctx         context.Context
	CancelFunc  context.CancelFunc
}

func NewTask(parentGroup *TaskGroup) *Task {
	ctx, cancel := context.WithCancel(context.Background())
	t := &Task{
		ParentGroup: parentGroup,
		ID:          uuid.NewString(),
		Ctx:         ctx,
		CancelFunc:  cancel,
	}
	parentGroup.Tasks = append(parentGroup.Tasks, t)
	return t
}

func RemoveTaskFromGroup(t *Task) error {
	t.ParentGroup.Mut.Lock()
	defer t.ParentGroup.Mut.Unlock()
	for i, t2 := range t.ParentGroup.Tasks {
		if t2.ID == t.ID {
			t.CancelFunc()
			t.ParentGroup.Tasks = append(t.ParentGroup.Tasks[:i], t.ParentGroup.Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
