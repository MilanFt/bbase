package task

import (
	"bbase/internal/profile"
	"bbase/internal/proxy"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

type TaskGroup struct {
	ID           string
	ErrorDelay   time.Duration
	Tasks        []*Task
	ProxyGroup   *proxy.ProxyGroup
	ProfileGroup *profile.ProfileGroup
	Mut          *sync.Mutex
}

// Global rand.Source to use for random profiles and proxies.
var randSource = rand.NewSource(time.Now().UnixNano())

func NewTaskGroup(
	errorDelay time.Duration,
	proxyGroup *proxy.ProxyGroup,
	profileGroup *profile.ProfileGroup,
) (*TaskGroup, error) {
	return &TaskGroup{
		ID:           uuid.NewString(),
		ErrorDelay:   errorDelay,
		ProxyGroup:   proxyGroup,
		ProfileGroup: profileGroup,
		Mut:          &sync.Mutex{},
	}, nil
}

func RemoveAllTasksFromGroup(taskGroup *TaskGroup) {
	taskGroup.Mut.Lock()
	defer taskGroup.Mut.Unlock()
	for _, task := range taskGroup.Tasks {
		task.CancelFunc()
	}
	taskGroup.Tasks = nil
}

func CancelAllTasksInGroup(taskGroup *TaskGroup) {
	for _, task := range taskGroup.Tasks {
		task.CancelFunc()
	}
}

func GetRandomProfileFromGroup(taskGroup *TaskGroup) *profile.Profile {
	if taskGroup.ProfileGroup == nil ||
		len(taskGroup.ProfileGroup.Profiles) == 0 {
		return nil
	}

	taskGroup.Mut.Lock()
	defer taskGroup.Mut.Unlock()

	r := rand.New(randSource)
	n := r.Intn(len(taskGroup.ProfileGroup.Profiles))

	return taskGroup.ProfileGroup.Profiles[n]
}

func GetRandomProxyFromGroup(taskGroup *TaskGroup) *proxy.Proxy {
	if taskGroup.ProxyGroup == nil ||
		len(taskGroup.ProxyGroup.Proxies) == 0 {
		return nil
	}

	taskGroup.Mut.Lock()
	defer taskGroup.Mut.Unlock()

	r := rand.New(randSource)
	n := r.Intn(len(taskGroup.ProxyGroup.Proxies))

	return taskGroup.ProxyGroup.Proxies[n]
}
