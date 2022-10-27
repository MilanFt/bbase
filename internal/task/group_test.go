package task

import (
	"bbase/internal/profile"
	"bbase/internal/proxy"
	"testing"
	"time"
)

func TestNewTaskGroup(t *testing.T) {
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
}

func TestGetRandomProfileFromGroup(t *testing.T) {
	tg, err := NewTaskGroup(
		time.Duration(500*time.Millisecond),
		proxy.NewProxyGroup(),
		profile.NewProfileGroup(),
	)
	if err != nil {
		t.Error("Error creating new task group")
	}

	a := profile.NewAddress("US", "John", "Doe", "123 Main St", "New York", "NY", "12345", "1234567890")
	pm := profile.NewPaymentMethod("John Doe", "1234567890123456", "01", "2022", "123")
	w := profile.NewWeb3Wallet("0x12345678901")
	p := profile.NewProfile(tg.ProfileGroup, "test@test.com", a, a, pm, w)

	if len(tg.ProfileGroup.Profiles) != 1 {
		t.Error("Profile was not added to ProfileGroup")
	}

	p2 := GetRandomProfileFromGroup(tg)
	if p2.ID != p.ID {
		t.Error("Profile returned is not the same as the one added")
	}
}

func TestGetRandomProxyFromGroup(t *testing.T) {
	tg, err := NewTaskGroup(
		time.Duration(500*time.Millisecond),
		proxy.NewProxyGroup(),
		profile.NewProfileGroup(),
	)
	if err != nil {
		t.Error("Error creating new task group")
	}

	p := proxy.NewProxy(tg.ProxyGroup, "192.168.1.1:2000")
	if len(tg.ProxyGroup.Proxies) != 1 {
		t.Error("Proxy was not added to ProxyGroup")
	}

	p2 := GetRandomProxyFromGroup(tg)
	if p2.ID != p.ID {
		t.Error("Proxy returned is not the same as the one added")
	}
}
