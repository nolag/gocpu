package templates

import "github.com/cheekybits/genny/generic"

type typeName generic.Type
type name generic.Type

// Iname represents a register holding the described number of bits.
type Iname interface {
	// DecrementnBits decrements the register by value
	DecrementtypeName(value typeName)

	// InrementnBits increments the register by value
	InrementtypeName(value typeName)

	// SetFromtypeName sets the value held by this register.
	SetFromtypeName(value typeName)

	// ValueAstypeName gets the value held by this register
	ValueAstypeName() typeName
}

// SetFromtypeName does nothing for the zero register.
func (register ZeroRegister) SetFromtypeName(value typeName) {
}

// ValueAstypeName gets the value held by this register
func (register ZeroRegister) ValueAstypeName() typeName {
	return 0
}

// DecrementtypeName decrements the register by value
func (register ZeroRegister) DecrementtypeName(value typeName) {}

// InrementtypeName increments the register by value
func (register ZeroRegister) InrementtypeName(value typeName) {}
