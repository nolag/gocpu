// Package registers provides basic register types for simulating a CPU
package registers

import (
	"math"
)

// Register32 represents a 32 bit register
type Register32 interface {
	// CanWrite true if the register actualy does writes, or false if they are ignored
	CanWrite() bool

	// SetValue32 sets the 32 bit value held by this register.
	SetValue32(uint32) error

	// Returns gets the 32 bit value held by this register
	Value32F() (float32, error)

	// Value32 gets the 32 bit value held by this register
	Value32() (uint32, error)

	// SetValue32 sets the 32 bit value held by this register.
	SetValue32F(float32) error
}

// UintRegister32 is a register that is backed by a uint32
type UintRegister32 uint32

// FloatRegister32 is a register that is backed by a float32
type FloatRegister32 float32

// ZeroRegister32 is a register that will always return zero, but allows writes
type ZeroRegister32 struct{}

// CanWrite true if the register actualy does writes, or false if they are ignored
func (*UintRegister32) CanWrite() bool {
	return true
}

// SetValue32 sets the 32 bit value held by this register.
func (register *UintRegister32) SetValue32(value uint32) error {
	*register = UintRegister32(value)
	return nil
}

// SetValue32F sets the 32 bit value held by this register.
func (register *UintRegister32) SetValue32F(value float32) error {
	*register = UintRegister32(math.Float32bits(value))
	return nil
}

// Value32 gets the 32 bit value held by this register
func (register *UintRegister32) Value32() (uint32, error) {
	return uint32(*register), nil
}

// Value32F gets the 32 bit value held by this register
func (register *UintRegister32) Value32F() (float32, error) {
	return math.Float32frombits(uint32(*register)), nil
}

// CanWrite true if the register actualy does writes, or false if they are ignored
func (*FloatRegister32) CanWrite() bool {
	return true
}

// SetValue32F sets the 32 bit value held by this register.
func (register *FloatRegister32) SetValue32F(value float32) error {
	*register = FloatRegister32(value)
	return nil
}

// SetValue32 sets the 32 bit value held by this register.
func (register *FloatRegister32) SetValue32(value uint32) error {
	*register = FloatRegister32(math.Float32frombits(value))
	return nil
}

// Value32 gets the 32 bit value held by this register
func (register *FloatRegister32) Value32() (uint32, error) {
	return math.Float32bits(float32(*register)), nil
}

// Value32F gets the 32 bit value held by this register
func (register *FloatRegister32) Value32F() (float32, error) {
	return float32(*register), nil
}

// CanWrite true if the register actualy does writes, or false if they are ignored
func (ZeroRegister32) CanWrite() bool {
	return false
}

// SetValue32 returns false
func (ZeroRegister32) SetValue32(uint32) error {
	return nil
}

// SetValue32F returns false
func (ZeroRegister32) SetValue32F(float32) error {
	return nil
}

// Value32 gets the 32 bit value (0) held by this register
func (ZeroRegister32) Value32() (uint32, error) {
	return 0, nil
}

// Value32F gets the 32 bit value (0) held by this register
func (ZeroRegister32) Value32F() (float32, error) {
	return 0.0, nil
}
