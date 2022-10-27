package proxy

import "testing"

func TestNewProxy(t *testing.T) {
	pg := NewProxyGroup()
	p := NewProxy(pg, "192.168.1.1:2000")
	p2 := NewProxy(pg, "192.168.1.1:2000:test:test")
	p3 := NewProxy(pg, "http://192.168.1.1:2000")
	p4 := NewProxy(pg, "192.168.1.1")
	if p.ParentGroup.ID != pg.ID {
		t.Error("ParentGroup ID is not equal to pg.ID")
	}
	if p.ID == "" {
		t.Error("ID is empty")
	}
	if p.URL != "http://192.168.1.1:2000" {
		t.Error("IP is not http://192.168.1.1:2000, got ", p.URL)
	}
	if p2.URL != "http://test:test@192.168.1.1:2000" {
		t.Error("IP is not http://test:test@192.168.1.1:2000, got ", p2.URL)
	}
	if p3.URL != "http://192.168.1.1:2000" {
		t.Error("IP is not http://192.168.1.1:2000, got ", p3.URL)
	}
	if p4.URL != "http://192.168.1.1" {
		t.Error("IP is not http://192.168.1.1, got ", p4.URL)
	}
	if len(pg.Proxies) != 4 {
		t.Error("Proxy was not added to ProxyGroup")
	}
}

func TestRemoveProxyFromGroup(t *testing.T) {
	pg := NewProxyGroup()
	p := NewProxy(pg, "192.168.1.1:2000")
	if len(pg.Proxies) != 1 {
		t.Error("Proxy was not added to ProxyGroup")
	}
	err := RemoveProxyFromGroup(p)
	if err != nil {
		t.Error("Error removing proxy from group")
	}
	if len(pg.Proxies) != 0 {
		t.Error("Proxy was not removed from ProxyGroup")
	}
}
