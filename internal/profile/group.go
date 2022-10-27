package profile

import (
	"sync"

	"github.com/google/uuid"
)

type ProfileGroup struct {
	ID       string
	Profiles []*Profile
	mut      *sync.Mutex
}

func NewProfileGroup() *ProfileGroup {
	return &ProfileGroup{
		ID:  uuid.NewString(),
		mut: &sync.Mutex{},
	}
}

func RemoveAllProfilesFromGroup(group *ProfileGroup) {
	group.mut.Lock()
	defer group.mut.Unlock()
	group.Profiles = nil
}
