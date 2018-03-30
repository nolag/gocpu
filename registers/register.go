// Package registers provides basic register types for simulating a CPU
package registers

//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_iregister8.go gen "typeName=uint8 name=Register8"
//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_iregister16.go gen "typeName=uint16 name=Register16"

//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_iintregister32.go gen "typeName=uint32 name=IntRegister32"
//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_ifloatregister32.go gen "typeName=float32 name=FloatRegister32"

//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_iintregister64.go gen "typeName=uint64 name=IntRegister64"
//go:generate genny -pkg registers -in=templates/generic_registerinterface.go -out=gen_ifloatregister64.go gen "typeName=float64 name=FloatRegister64"

//go:generate genny -pkg registers -in=templates/generic_registerimpl.go -out=gen_registerimpl.go gen "backingType=uint8,uint16,uint32,uint64,float32,float64"

import "math"

// ZeroRegister is a register that will always return zero, but allows writes
type ZeroRegister struct{}

// CanWrite true if the register actualy does writes, or false if they are ignored
func (ZeroRegister) CanWrite() bool {
	return false
}

// ProgramCounter represents a program counter
type ProgramCounter interface {
	// InrementAsPc increments the register by value byte pc read
	InrementAsPc(value byte)

	// ReadAsPc reads the value cased to a 64 bit allowing a memory read from that location to follow
	ReadAsPc() uint64
}

// IRegister32 represents a register holding the described number of bits.
type IRegister32 interface {
	IFloatRegister32
	IIntRegister32
}

// IRegister64 represents a register holding the described number of bits.
type IRegister64 interface {
	IFloatRegister64
	IIntRegister64
}

// SetFromFloat32 sets the value held by this register.
func (register *RegisterUint32) SetFromFloat32(value float32) {
	*register = RegisterUint32(math.Float32bits(value))
}

// ValueAsFloat32 gets the value held by this register
func (register *RegisterUint32) ValueAsFloat32() float32 {
	return math.Float32frombits(uint32(*register))
}

// DecrementFloat32 decrements the register by value
func (register *RegisterUint32) DecrementFloat32(value float32) {
	register.SetFromFloat32(register.ValueAsFloat32() - value)
}

// InrementFloat32 increments the register by value
func (register *RegisterUint32) InrementFloat32(value float32) {
	register.SetFromFloat32(register.ValueAsFloat32() + value)
}

// SetFromFloat64 sets the value held by this register.
func (register *RegisterUint64) SetFromFloat64(value float64) {
	*register = RegisterUint64(math.Float64bits(value))
}

// ValueAsFloat64 gets the value held by this register
func (register *RegisterUint64) ValueAsFloat64() float64 {
	return math.Float64frombits(uint64(*register))
}

// DecrementFloat64 decrements the register by value
func (register *RegisterUint64) DecrementFloat64(value float64) {
	register.SetFromFloat64(register.ValueAsFloat64() - value)
}

// InrementFloat64 increments the register by value
func (register *RegisterUint64) InrementFloat64(value float64) {
	register.SetFromFloat64(register.ValueAsFloat64() + value)
}

// SetFromUint32 sets the value held by this register.
func (register *RegisterFloat32) SetFromUint32(value uint32) {
	*register = RegisterFloat32(math.Float32frombits(value))
}

// ValueAsUint32 gets the value held by this register
func (register *RegisterFloat32) ValueAsUint32() uint32 {
	return math.Float32bits(float32(*register))
}

// DecrementUint32 decrements the register by value
func (register *RegisterFloat32) DecrementUint32(value uint32) {
	register.SetFromUint32(register.ValueAsUint32() - value)
}

// InrementUint32 increments the register by value
func (register *RegisterFloat32) InrementUint32(value uint32) {
	register.SetFromUint32(register.ValueAsUint32() + value)
}

// SetFromUint64 sets the value held by this register.
func (register *RegisterFloat64) SetFromUint64(value uint64) {
	*register = RegisterFloat64(math.Float64frombits(value))
}

// ValueAsUint64 gets the value held by this register
func (register *RegisterFloat64) ValueAsUint64() uint64 {
	return math.Float64bits(float64(*register))
}

// DecrementUint64 decrements the register by value
func (register *RegisterFloat64) DecrementUint64(value uint64) {
	register.SetFromUint64(register.ValueAsUint64() - value)
}

// InrementUint64 increments the register by value
func (register *RegisterFloat64) InrementUint64(value uint64) {
	register.SetFromUint64(register.ValueAsUint64() + value)
}
