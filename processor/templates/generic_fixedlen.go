package templates

import (
	"encoding/binary"

	"github.com/cheekybits/genny/generic"
	memory "github.com/nolag/gocpu/processor/templates/memhack"
	"github.com/nolag/gocpu/registers"
)

type runnerType generic.Number
type runnerCapsType generic.Number

// FixedInstructionLenRunnerrunnerType runs runnerType instructions one at a time, by calling InstructionRunnerrunnerType.
type FixedInstructionLenRunnerrunnerType struct {
	memory.Memory
	binary.ByteOrder
	InstructionRunnerrunnerType
	Pc                        registers.ProgramCounter
	MemoryReadFailureCallback ErrorCallback
}

// Step runs the next instruction, returns error to indicate an unhandeled exception
func (cpu *FixedInstructionLenRunnerrunnerType) Step() error {
	i := instructionrunnerType(0)
	val, err := memory.ReadrunnerCapsType(cpu.Memory, cpu.ByteOrder, cpu.Pc.ReadAsPc())

	if err != nil {
		callback := cpu.MemoryReadFailureCallback
		if callback != nil {
			err = cpu.MemoryReadFailureCallback(err)
		}

		return err
	}

	err = cpu.RunrunnerCapsType(val)
	if err == nil {
		cpu.Pc.InrementAsPc(i.size())
	}

	return err
}
