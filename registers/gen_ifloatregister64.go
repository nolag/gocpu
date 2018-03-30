// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package registers

// IFloatRegister64 represents a register holding the described number of bits.
type IFloatRegister64 interface {
	// DecrementnBits decrements the register by value
	DecrementFloat64(value float64)

	// InrementnBits increments the register by value
	InrementFloat64(value float64)

	// SetFromFloat64 sets the value held by this register.
	SetFromFloat64(value float64)

	// ValueAsFloat64 gets the value held by this register
	ValueAsFloat64() float64
}

// SetFromFloat64 does nothing for the zero register.
func (register ZeroRegister) SetFromFloat64(value float64) {
}

// ValueAsFloat64 gets the value held by this register
func (register ZeroRegister) ValueAsFloat64() float64 {
	return 0
}

// DecrementFloat64 decrements the register by value
func (register ZeroRegister) DecrementFloat64(value float64) {}

// InrementFloat64 increments the register by value
func (register ZeroRegister) InrementFloat64(value float64) {}
