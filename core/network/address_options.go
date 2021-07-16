// Copyright 2021 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package network

// AddressMutator describes setter methods for an address.
type AddressMutator interface {
	// SetScope sets the scope property of the address.
	SetScope(Scope)

	// SetCIDR sets the CIDR property of the address.
	SetCIDR(string)

	// SetSecondary indicates whether this address is not the
	// primary address of the device it is associated with.
	SetSecondary(bool)

	// SetConfigType indicates how this address was configured.
	SetConfigType(AddressConfigType)
}

// SetScope (AddressMutator) sets the input
// scope on the address receiver.
func (a *MachineAddress) SetScope(scope Scope) {
	a.Scope = scope
}

// SetCIDR (AddressMutator) sets the input
// CIDR on the address receiver.
func (a *MachineAddress) SetCIDR(cidr string) {
	a.CIDR = cidr
}

// SetSecondary (AddressMutator) sets the IsSecondary
// member to true on the address receiver.
func (a *MachineAddress) SetSecondary(isSecondary bool) {
	a.IsSecondary = isSecondary
}

// SetConfigType (AddressMutator) sets the input
// AddressConfigType on the address receiver.
func (a *MachineAddress) SetConfigType(configType AddressConfigType) {
	a.ConfigType = configType
}

// WithScope returns a functional option that can
// be used to set the input scope on an address.
func WithScope(scope Scope) func(AddressMutator) {
	return func(a AddressMutator) {
		a.SetScope(scope)
	}
}

// WithCIDR returns a functional option that can
// be used to set the input CIDR on an address.
func WithCIDR(cidr string) func(AddressMutator) {
	return func(a AddressMutator) {
		a.SetCIDR(cidr)
	}
}

// WithSecondary returns a functional option that can be used to
// indicate whether an address is not the primary for its NIC.
func WithSecondary(isSecondary bool) func(AddressMutator) {
	return func(a AddressMutator) {
		a.SetSecondary(isSecondary)
	}
}

func WithConfigType(configType AddressConfigType) func(AddressMutator) {
	return func(a AddressMutator) {
		a.SetConfigType(configType)
	}
}
