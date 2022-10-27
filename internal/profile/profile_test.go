package profile

import "testing"

func TestNewAddress(t *testing.T) {
	a := NewAddress("US", "John", "Doe", "123 Main St", "New York", "NY", "12345", "1234567890")
	if a.Country != "US" {
		t.Error("Country is not US")
	}
	if a.FirstName != "John" {
		t.Error("FirstName is not John")
	}
	if a.LastName != "Doe" {
		t.Error("LastName is not Doe")
	}
	if a.Street != "123 Main St" {
		t.Error("Street is not 123 Main St")
	}
	if a.City != "New York" {
		t.Error("City is not New York")
	}
	if a.State != "NY" {
		t.Error("State is not NY")
	}
	if a.Zip != "12345" {
		t.Error("Zip is not 12345")
	}
	if a.Phone != "1234567890" {
		t.Error("Phone is not 1234567890")
	}
}

func TestNewPaymentMethod(t *testing.T) {
	pm := NewPaymentMethod("John Doe", "1234567890123456", "01", "2022", "123")
	if pm.Name != "John Doe" {
		t.Error("Name is not John Doe")
	}
	if pm.CardNumber != "1234567890123456" {
		t.Error("CardNumber is not 1234567890123456")
	}
	if pm.ExpMonth != "01" {
		t.Error("ExpMonth is not 01")
	}
	if pm.ExpYear != "2022" {
		t.Error("ExpYear is not 2022")
	}
	if pm.CVV != "123" {
		t.Error("CVV is not 123")
	}
}

func TestNewWeb3Wallet(t *testing.T) {
	w := NewWeb3Wallet("0x12345678901")
	if w.PrivateKey != "0x12345678901" {
		t.Error("PrivateKey is not 0x12345678901")
	}
}

func TestNewProfile(t *testing.T) {
	pg := NewProfileGroup()

	a := NewAddress("US", "John", "Doe", "123 Main St", "New York", "NY", "12345", "1234567890")
	pm := NewPaymentMethod("John Doe", "1234567890123456", "01", "2022", "123")
	w := NewWeb3Wallet("0x12345678901")
	p := NewProfile(pg, "test@test.com", a, a, pm, w)

	if p.ParentGroup.ID != pg.ID {
		t.Error("ParentGroup ID is not equal to pg.ID")
	}
	if p.ID == "" {
		t.Error("ID is empty")
	}
	if p.Email != "test@test.com" {
		t.Error("Email is not test@test.com")
	}
	if p.ShippingAddress != p.BillingAddress {
		t.Error("ShippingAddress is not equal to BillingAddress")
	}
	if p.PaymentMethod != pm {
		t.Error("PaymentMethod is not equal to pm")
	}
	if p.Web3Wallet != w {
		t.Error("Web3Wallet is not equal to w")
	}
}

func TestRemoveProfileFromGroup(t *testing.T) {
	pg := NewProfileGroup()

	a := NewAddress("US", "John", "Doe", "123 Main St", "New York", "NY", "12345", "1234567890")
	pm := NewPaymentMethod("John Doe", "1234567890123456", "01", "2022", "123")
	w := NewWeb3Wallet("0x12345678901")
	p := NewProfile(pg, "test@test.com", a, a, pm, w)

	if len(pg.Profiles) != 1 {
		t.Error("Profiles length is not 1")
	}

	RemoveProfileFromGroup(p)

	if len(pg.Profiles) != 0 {
		t.Error("Profiles length is not 0")
	}
}
