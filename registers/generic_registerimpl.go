package registers

import "github.com/cheekybits/genny/generic"

// BackingType is the int type that backs the register
type BackingType generic.Number

// BackingTypeRegister is the register for the backing int
type BackingTypeRegister BackingType

// SetFromBackingType sets the 32 bit value held by this register.
func (register *BackingTypeRegister) SetFromBackingType(value BackingType) {
	*register = BackingTypeRegister(value)
}

// BackingTypeValue gets the value held by this register
func (register *BackingTypeRegister) BackingTypeValue() BackingType {
	return BackingType(*register)
}

// SetFromBackingType does nothing for the zero register.
func (register ZeroRegister) SetFromBackingType(value BackingType) {
}

// BackingTypeValue gets the value held by this register
func (register ZeroRegister) BackingTypeValue() BackingType {
	return 0
}
