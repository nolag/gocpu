package processor

import (
	"github.com/nolag/gocpu/memory"
)

// FixedLen32 runs a sinlge 32 bit instruction one at a time, by calling InstructionRunner32.
type FixedLen32 struct {
	Core32
	InstructionRunner32
	MemoryReadFailureCallback ErrorCallback
}

// Step runs the next instruction, returns error to indicate an unhandeled exception
func (processor *FixedLen32) Step() error {
	val, err := memory.ReadUint32(processor.Memory, processor.Endianness, uint64(processor.Pc.Value32()))
	if err != nil {
		callback := processor.MemoryReadFailureCallback
		if callback != nil {
			err = processor.MemoryReadFailureCallback(err)
		}

		return err
	}

	processor.Pc.SetValue32(processor.Pc.Value32() + 4)
	processor.RunInstruction32(val)
	return nil
}
