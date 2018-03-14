package mock

// Processor is a mock processor.  use NewProcessor to create one
type Processor struct {
	NumTimesStepped int
	ErrToReturn     error
	Callback        func()
}

// NewProcessor creates a new processor
func NewProcessor(errToReturn error) *Processor {
	return &Processor{0, errToReturn, nil}
}

// NewProcessorWithCallback creates a new processor that will make a callback on step
func NewProcessorWithCallback(errToReturn error, callback func()) *Processor {
	return &Processor{0, errToReturn, callback}
}

// Step runs the next instruction, returns error to indicate an unhandeled exception
func (cpu *Processor) Step() error {
	cpu.NumTimesStepped++
	if cpu.Callback != nil {
		cpu.Callback()
	}

	return cpu.ErrToReturn
}
