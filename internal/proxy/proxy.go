package proxy

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Proxy struct {
	ParentGroup *ProxyGroup
	ID          string
	URL         string
}

func NewProxy(parentGroup *ProxyGroup, url string) *Proxy {
	p := &Proxy{
		ParentGroup: parentGroup,
		ID:          uuid.NewString(),
		URL:         proxyToURL(url),
	}
	parentGroup.Proxies = append(parentGroup.Proxies, p)
	return p
}

func RemoveProxyFromGroup(proxy *Proxy) error {
	proxy.ParentGroup.mut.Lock()
	defer proxy.ParentGroup.mut.Unlock()
	for i, p := range proxy.ParentGroup.Proxies {
		if p.ID == proxy.ID {
			proxy.ParentGroup.Proxies = append(proxy.ParentGroup.Proxies[:i], proxy.ParentGroup.Proxies[i+1:]...)
			return nil
		}
	}
	return errors.New("proxy not found")
}

func proxyToURL(proxy string) string {
	proxySplit := strings.Split(proxy, ":")

	if len(proxySplit) == 2 {
		return fmt.Sprintf("http://%s:%s", proxySplit[0], proxySplit[1])
	} else if len(proxySplit) == 4 {
		return fmt.Sprintf("http://%s:%s@%s:%s", proxySplit[2], proxySplit[3], proxySplit[0], proxySplit[1])
	}

	if strings.Contains(proxy, "http") {
		return proxy
	}

	return fmt.Sprintf("http://%s", proxy)
}
