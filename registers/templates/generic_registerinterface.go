package templates

import "github.com/cheekybits/genny/generic"

type typeName generic.Type
type nBits generic.Type

// nBitsBitRegister represents a register holding the described number of bits.
type nBitsBitRegister interface {
	// SetFromtypeName sets the value held by this register.
	SetFromtypeName(value typeName)

	// ValueAstypeName gets the value held by this register
	ValueAstypeName() typeName
}
