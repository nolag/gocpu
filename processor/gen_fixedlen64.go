// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package processor

import (
	"encoding/binary"

	"github.com/nolag/gocpu/memory"
	"github.com/nolag/gocpu/registers"
)

// FixedInstructionLenRunnerUint64 runs uint64 instructions one at a time, by calling InstructionRunnerUint64.
type FixedInstructionLenRunnerUint64 struct {
	memory.Memory
	binary.ByteOrder
	InstructionRunnerUint64
	Pc                        registers.ProgramCounter
	MemoryReadFailureCallback ErrorCallback
}

// Step runs the next instruction, returns error to indicate an unhandeled exception
func (cpu *FixedInstructionLenRunnerUint64) Step() error {
	i := instructionuint64(0)
	val, err := memory.ReadUint64(cpu.Memory, cpu.ByteOrder, cpu.Pc.ReadAsPc())

	if err != nil {
		callback := cpu.MemoryReadFailureCallback
		if callback != nil {
			err = cpu.MemoryReadFailureCallback(err)
		}

		return err
	}

	err = cpu.RunUint64(val)
	if err == nil {
		cpu.Pc.InrementAsPc(i.size())
	}

	return err
}
