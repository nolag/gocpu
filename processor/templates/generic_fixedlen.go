package templates

import (
	"encoding/binary"

	"github.com/cheekybits/genny/generic"
	memory "github.com/nolag/gocpu/processor/templates/memhack"
)

type pcType generic.Number
type runnerType generic.Number
type runnerCapsType generic.Number

// FixedInstructionLenPcpcTypeRunnerrunnerType runs a sinlge 32 bit instruction one at a time, by calling InstructionRunner32.
type FixedInstructionLenPcpcTypeRunnerrunnerType struct {
	memory.Memory
	binary.ByteOrder
	InstructionRunnerrunnerType
	Pc                        pcType
	MemoryReadFailureCallback ErrorCallback
}

// Step runs the next instruction, returns error to indicate an unhandeled exception
func (cpu *FixedInstructionLenPcpcTypeRunnerrunnerType) Step() error {
	i := instructionrunnerType(0)
	val, err := memory.ReadrunnerCapsType(cpu.Memory, cpu.ByteOrder, uint64(cpu.Pc))

	if err != nil {
		callback := cpu.MemoryReadFailureCallback
		if callback != nil {
			err = cpu.MemoryReadFailureCallback(err)
		}

		return err
	}
	cpu.Pc += pcType(i.size())
	cpu.RunrunnerCapsType(val)
	return nil
}
