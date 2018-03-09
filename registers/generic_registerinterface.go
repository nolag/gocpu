package registers

import "github.com/cheekybits/genny/generic"

// TypeName the number of bits that the register represents.  Note generic.Type is used here because we do not need to use anything promised from numbers
type TypeName generic.Type

type nBits generic.Type

// nBitsBitRegister represents a register holding the described number of bits.
type nBitsBitRegister interface {
	// SetFromIntName sets the value held by this register.
	SetFromTypeName(value TypeName)

	// TypeNameValue gets the value held by this register
	TypeNameValue() TypeName
}
