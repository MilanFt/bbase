package profile

import (
	"errors"

	"github.com/google/uuid"
)

type Profile struct {
	ParentGroup     *ProfileGroup
	ID              string
	Email           string
	ShippingAddress *Address
	BillingAddress  *Address
	PaymentMethod   *PaymentMethod
	Web3Wallet      *Web3Wallet
}

type Address struct {
	Country   string
	FirstName string
	LastName  string
	Street    string
	City      string
	State     string
	Zip       string
	Phone     string
}

type PaymentMethod struct {
	Name       string
	CardNumber string
	ExpMonth   string
	ExpYear    string
	CVV        string
}

type Web3Wallet struct {
	PrivateKey string
}

func NewProfile(
	parentGroup *ProfileGroup,
	email string,
	shippingAddress *Address,
	billingAddress *Address,
	paymentMethod *PaymentMethod,
	web3Wallet *Web3Wallet,
) *Profile {
	p := &Profile{
		ParentGroup:     parentGroup,
		ID:              uuid.NewString(),
		Email:           email,
		ShippingAddress: shippingAddress,
		BillingAddress:  billingAddress,
		PaymentMethod:   paymentMethod,
		Web3Wallet:      web3Wallet,
	}
	parentGroup.Profiles = append(parentGroup.Profiles, p)
	return p
}

func NewAddress(
	country string,
	firstName string,
	lastName string,
	street string,
	city string,
	state string,
	zip string,
	phone string,
) *Address {
	return &Address{
		Country:   country,
		FirstName: firstName,
		LastName:  lastName,
		Street:    street,
		City:      city,
		State:     state,
		Zip:       zip,
		Phone:     phone,
	}
}

func NewPaymentMethod(
	name string,
	cardNumber string,
	expMonth string,
	expYear string,
	cvv string,
) *PaymentMethod {
	return &PaymentMethod{
		Name:       name,
		CardNumber: cardNumber,
		ExpMonth:   expMonth,
		ExpYear:    expYear,
		CVV:        cvv,
	}
}

func NewWeb3Wallet(privateKey string) *Web3Wallet {
	return &Web3Wallet{
		PrivateKey: privateKey,
	}
}

func RemoveProfileFromGroup(profile *Profile) error {
	profile.ParentGroup.mut.Lock()
	defer profile.ParentGroup.mut.Unlock()
	for i, p := range profile.ParentGroup.Profiles {
		if p.ID == profile.ID {
			profile.ParentGroup.Profiles = append(profile.ParentGroup.Profiles[:i], profile.ParentGroup.Profiles[i+1:]...)
			return nil
		}
	}
	return errors.New("profile not found")
}
