package gocpu

// FixedLenProcessor32 runs a sinlge 32 bit instruction one at a time, by calling InstructionRunner32.
type FixedLenProcessor32 struct {
	ProcessorCore32
	InstructionRunner32
	dealyedActions []Callback
}

// Step runs the next instruction, returns error to indicate an unhandeled exception
func (processor *FixedLenProcessor32) Step() error {
	return nil
}
