package proxy

import "testing"

func TestNewProxyGroup(t *testing.T) {
	pg := NewProxyGroup()
	if pg.ID == "" {
		t.Error("ID is empty")
	}
	if pg.mut == nil {
		t.Error("Mut is nil")
	}
}
