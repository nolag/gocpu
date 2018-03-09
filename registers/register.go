// Package registers provides basic register types for simulating a CPU
package registers

//go:generate genny -pkg registers -in=generic_registerinterface.go -out=gen_eightbitregister.go gen "TypeName=uint8 nBits=Eight"
//go:generate genny -pkg registers -in=generic_registerinterface.go -out=gen_sixteenbitregister.go gen "TypeName=uint16 nBits=Sixteen"

//go:generate genny -pkg registers -in=generic_registerinterface.go -out=gen_thirtytwoibitregister.go gen "TypeName=uint32 nBits=thirtyTwoi"
//go:generate genny -pkg registers -in=generic_registerinterface.go -out=gen_sixtyfouribitregister.go gen "TypeName=uint64 nBits=sixtyFouri"
//go:generate genny -pkg registers -in=generic_registerinterface.go -out=gen_thirtytwofbitregister.go gen "TypeName=float32 nBits=thirtyTwof"
//go:generate genny -pkg registers -in=generic_registerinterface.go -out=gen_sixtyfourfbitregister.go gen "TypeName=float64 nBits=sixtyFourf"

//go:generate genny -pkg registers -in=generic_registerimpl.go -out=gen_registerimpl.go gen "BackingType=uint8,uint16,uint32,uint64,float32,float64"

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
func (register *Uint32Register) SetFromFloat32(value float32) {
	*register = Uint32Register(math.Float32bits(value))
}

// Float32Value gets the value held by this register
func (register *Uint32Register) Float32Value() float32 {
	return math.Float32frombits(uint32(*register))
}

// SetFromFloat64 sets the value held by this register.
func (register *Uint64Register) SetFromFloat64(value float64) {
	*register = Uint64Register(math.Float64bits(value))
}

// Float64Value gets the value held by this register
func (register *Uint64Register) Float64Value() float64 {
	return math.Float64frombits(uint64(*register))
}

// SetFromUint32 sets the value held by this register.
func (register *Float32Register) SetFromUint32(value uint32) {
	*register = Float32Register(math.Float32frombits(value))
}

// Uint32Value gets the value held by this register
func (register *Float32Register) Uint32Value() uint32 {
	return math.Float32bits(float32(*register))
}

// SetFromUint64 sets the value held by this register.
func (register *Float64Register) SetFromUint64(value uint64) {
	*register = Float64Register(math.Float64frombits(value))
}

// Uint64Value gets the value held by this register
func (register *Float64Register) Uint64Value() uint64 {
	return math.Float64bits(float64(*register))
}
