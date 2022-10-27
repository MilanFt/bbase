package proxy

import (
	"sync"

	"github.com/google/uuid"
)

type ProxyGroup struct {
	ID      string
	Proxies []*Proxy
	mut     *sync.Mutex
}

func NewProxyGroup() *ProxyGroup {
	return &ProxyGroup{
		ID:  uuid.NewString(),
		mut: &sync.Mutex{},
	}
}

func RemoveAllProxiesFromGroup(group *ProxyGroup) {
	group.mut.Lock()
	defer group.mut.Unlock()
	group.Proxies = nil
}
