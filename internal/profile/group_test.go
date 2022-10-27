package profile

import "testing"

func TestNewProfileGroup(t *testing.T) {
	pg := NewProfileGroup()
	if pg.ID == "" {
		t.Error("ID is empty")
	}
	if pg.mut == nil {
		t.Error("Mut is nil")
	}
}
