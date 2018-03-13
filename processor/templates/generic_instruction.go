package templates

import "github.com/cheekybits/genny/generic"

type instructionType generic.Type

// instructioninstructionType represents an instruction
type instructioninstructionType instructionType

// InstructionRunnerinstructionType runs an InstructioninstructionpcType
type InstructionRunnerinstructionType interface {
	// RuninstructionType runs a single instrution (without incrementing the PC for its own read)
	RuninstructionType(instruction instructionType) error
}
