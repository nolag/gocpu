package templates

import "github.com/cheekybits/genny/generic"

type backingType generic.Number

// RegisterbackingType is the register for the backing int
type RegisterbackingType backingType

// SetFrombackingType does nothing for the zero register.
func (register *RegisterbackingType) SetFrombackingType(value backingType) {
	*register = RegisterbackingType(value)
}

// ValueAsbackingType gets the value held by this register
func (register *RegisterbackingType) ValueAsbackingType() backingType {
	return backingType(*register)
}

// DecrementbackingType decrements the register by value
func (register *RegisterbackingType) DecrementbackingType(value backingType) {
	*register -= RegisterbackingType(value)
}

// InrementbackingType increments the register by value
func (register *RegisterbackingType) InrementbackingType(value backingType) {
	*register += RegisterbackingType(value)
}

// InrementAsPc increments the register by value byte pc read
func (register *RegisterbackingType) InrementAsPc(value byte) {
	*register += RegisterbackingType(value)
}

// ReadAsPc reads the value cased to a 64 bit allowing a memory read from that location to follow
func (register *RegisterbackingType) ReadAsPc() uint64 {
	return uint64(*register)
}
