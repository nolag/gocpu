package gocpu

import (
	"encoding/binary"

	"github.com/nolag/gocpu/memory"
	"github.com/nolag/gocpu/registers"
)

// SimpleProcessorBasics32 is an implementation of ProcessorBasics32 that uses a uint32 for the Pc
type SimpleProcessorBasics32 struct {
	BackingEndianness binary.ByteOrder
	BackingMemory     memory.Memory
	BackingRegisters  []registers.Register32
	Pc                uint32
}

// DecementPc32 decrements the value of the PC by amount
func (basics *SimpleProcessorBasics32) DecementPc32(amount uint32) error {
	basics.Pc -= amount
	return nil
}

// Endianness is they byte order that the processor is using
func (basics *SimpleProcessorBasics32) Endianness() binary.ByteOrder {
	return basics.BackingEndianness
}

// IncementPc32 increments the value of PC by amount
func (basics *SimpleProcessorBasics32) IncementPc32(amount uint32) error {
	basics.Pc += amount
	return nil
}

// Memory gets the memory that this processor holds
func (basics *SimpleProcessorBasics32) Memory() memory.Memory {
	return basics.BackingMemory
}

// Pc32 gets the value of the current PC
func (basics *SimpleProcessorBasics32) Pc32() (uint32, error) {
	return basics.Pc, nil
}

// Registers32 gets the registers that can be accessed as Registers32 that this processor holds
func (basics *SimpleProcessorBasics32) Registers32() []registers.Register32 {
	return basics.BackingRegisters
}

// SetEndianness allows you to change the endianness of the processor
func (basics *SimpleProcessorBasics32) SetEndianness(endianness binary.ByteOrder) {
	basics.BackingEndianness = endianness
}

// SetMemory sets the backing memory
func (basics *SimpleProcessorBasics32) SetMemory(memory memory.Memory) {
	basics.BackingMemory = memory
}

// SetPc the value of the PC
func (basics *SimpleProcessorBasics32) SetPc(value uint32) error {
	basics.Pc = value
	return nil
}

// SetRegisters sets the registers
func (basics *SimpleProcessorBasics32) SetRegisters(registers []registers.Register32) {
	basics.BackingRegisters = registers
}
