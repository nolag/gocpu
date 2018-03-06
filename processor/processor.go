// Package processor provides the bascis for implmenting a processor
package processor

// Core32 provides the basics in a 32 bit processor
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

// Core32 provides a core for a 32-bit CPU
type Core32 struct {
	Endianness binary.ByteOrder
	Memory     memory.Memory
	Registers  []registers.Register32
	Pc         registers.Register32
}

// InstructionRunner32 can run a 32-bit instruction
type InstructionRunner32 interface {
	// RunInstruction32 runs a single 32 bit instrution, without incrementing the PC
	RunInstruction32(instruction uint32) error
}

// ErrorCallback allows a callback to be made when an error is returned
type ErrorCallback func(err error) error
