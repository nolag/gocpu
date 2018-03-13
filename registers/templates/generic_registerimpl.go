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

// SetFrombackingType does nothing for the zero register.
func (register ZeroRegister) SetFrombackingType(value backingType) {
}

// ValueAsbackingType gets the value held by this register
func (register ZeroRegister) ValueAsbackingType() backingType {
	return 0
}
