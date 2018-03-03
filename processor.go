// Package gocpu provides basics for building a CPU simulator
package gocpu

import (
	"encoding/binary"

	"github.com/nolag/gocpu/memory"
	"github.com/nolag/gocpu/registers"
)

// Processor is an interface for a simulated processor
type Processor interface {
	// Step runs the next instruction, returns error to indicate an unhandeled exception
	Step() error
}

// ProcessorBasics32 provides the basics in a 32 bit processor
type ProcessorBasics32 interface {
	// DecementPc32 decrements the value of the PC by amount
	DecementPc32(amount uint32) error

	// Endianness is they byte order that the processor is using
	Endianness() binary.ByteOrder

	// IncementPc32 increments the value of PC by amount
	IncementPc32(amount uint32) error

	// Memory gets the memory that this processor holds
	Memory() memory.Memory

	// Gets the value of the current PC
	Pc32() (uint32, error)

	// Registers32 gets the registers that can be accessed as Registers32 that this processor holds
	Registers32() []registers.Register32

	// SetEndianness allows you to change the endianness of the processor
	SetEndianness(endianness binary.ByteOrder)

	// SetMemory sets the backing memory
	SetMemory(memory memory.Memory)

	// SetPc the value of the PC
	SetPc(value uint32) error

	// SetRegisters sets the registers
	SetRegisters(registers []registers.Register32)
}

// InstructionRunner32 can run a 32-bit instruction
type InstructionRunner32 interface {
	// Runs a single 32 bit instrution, without incrementing the PC
	RunInstruction32(instruction uint32) error
}
