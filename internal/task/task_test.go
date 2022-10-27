package task

import (
	"bbase/internal/profile"
	"bbase/internal/proxy"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	tg, err := NewTaskGroup(
		time.Duration(500*time.Millisecond),
		proxy.NewProxyGroup(),
		profile.NewProfileGroup(),
	)
	if err != nil {
		t.Error("Error creating new task group")
	}
	if tg.ID == "" {
		t.Error("ID is empty")
	}
	if tg.Mut == nil {
		t.Error("Mut is nil")
	}
	tsk := NewTask(tg)
	if tsk.ID == "" {
		t.Error("ID is empty")
	}
	if tsk.ParentGroup.ID != tg.ID {
		t.Error("ParentGroup is not the same as the one passed in")
	}
	if tsk.Ctx == nil {
		t.Error("Ctx is nil")
	}
	if tsk.CancelFunc == nil {
		t.Error("CancelFunc is nil")
	}
	if len(tg.Tasks) != 1 {
		t.Error("Task was not added to TaskGroup")
	}
}

func TestRemoveTaskFromGroup(t *testing.T) {
	tg, err := NewTaskGroup(
		time.Duration(500*time.Millisecond),
		proxy.NewProxyGroup(),
		profile.NewProfileGroup(),
	)
	if err != nil {
		t.Error("Error creating new task group")
	}
	tsk := NewTask(tg)
	if len(tg.Tasks) != 1 {
		t.Error("Task was not added to TaskGroup")
	}
	RemoveTaskFromGroup(tsk)
	if len(tg.Tasks) != 0 {
		t.Error("Task was not removed from TaskGroup")
	}
}
