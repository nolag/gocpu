// Package registers provides basic register types for simulating a CPU
package registers

//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_eightbitregister.go gen "typeName=uint8 nBits=Eight"
//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_sixteenbitregister.go gen "typeName=uint16 nBits=Sixteen"

//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_thirtytwoibitregister.go gen "typeName=uint32 nBits=thirtyTwoi"
//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_sixtyfouribitregister.go gen "typeName=uint64 nBits=sixtyFouri"
//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_thirtytwofbitregister.go gen "typeName=float32 nBits=thirtyTwof"
//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_sixtyfourfbitregister.go gen "typeName=float64 nBits=sixtyFourf"

//go:generate genny -pkg registers -in=templates/generic_registerimpl.go -out=gen_registerimpl.go gen "backingType=uint8,uint16,uint32,uint64,float32,float64"

import "math"

// ZeroRegister is a register that will always return zero, but allows writes
type ZeroRegister struct{}

// CanWrite true if the register actualy does writes, or false if they are ignored
func (ZeroRegister) CanWrite() bool {
	return false
}

// ThirtyTwoBitRegister represents a register holding the described number of bits.
type ThirtyTwoBitRegister interface {
	thirtyTwofBitRegister
	thirtyTwoiBitRegister
}

// SixtyFourBitRegister represents a register holding the described number of bits.
type SixtyFourBitRegister interface {
	sixtyFourfBitRegister
	sixtyFouriBitRegister
}

// SetFromFloat32 sets the value held by this register.
func (register *RegisterUint32) SetFromFloat32(value float32) {
	*register = RegisterUint32(math.Float32bits(value))
}

// Float32Value gets the value held by this register
func (register *RegisterUint32) Float32Value() float32 {
	return math.Float32frombits(uint32(*register))
}

// SetFromFloat64 sets the value held by this register.
func (register *RegisterUint64) SetFromFloat64(value float64) {
	*register = RegisterUint64(math.Float64bits(value))
}

// Float64Value gets the value held by this register
func (register *RegisterUint64) Float64Value() float64 {
	return math.Float64frombits(uint64(*register))
}

// SetFromUint32 sets the value held by this register.
func (register *RegisterFloat32) SetFromUint32(value uint32) {
	*register = RegisterFloat32(math.Float32frombits(value))
}

// Uint32Value gets the value held by this register
func (register *RegisterFloat32) Uint32Value() uint32 {
	return math.Float32bits(float32(*register))
}

// SetFromUint64 sets the value held by this register.
func (register *RegisterFloat64) SetFromUint64(value uint64) {
	*register = RegisterFloat64(math.Float64frombits(value))
}

// Uint64Value gets the value held by this register
func (register *RegisterFloat64) Uint64Value() uint64 {
	return math.Float64bits(float64(*register))
}
