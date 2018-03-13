// Package processor provides the bascis for implmenting a processor
package processor

//go:generate genny -pkg processor -in=templates/generic_instruction.go -out=gen_instruction_impl.go gen "instructionType=uint8,uint16,uint32,uint64"
//go:generate genny -pkg processor -in=templates/generic_fixedlen.go -out=gen_fixedlen8.go gen "pcType=uint8,uint16,uint32,uint64 runnerType=uint8 runnerCapsType=Uint8"
//go:generate genny -pkg processor -in=templates/generic_fixedlen.go -out=gen_fixedlen16.go gen "pcType=uint8,uint16,uint32,uint64 runnerType=uint16 runnerCapsType=Uint16"
//go:generate genny -pkg processor -in=templates/generic_fixedlen.go -out=gen_fixedlen32.go gen "pcType=uint8,uint16,uint32,uint64 runnerType=uint32 runnerCapsType=Uint32"
//go:generate genny -pkg processor -in=templates/generic_fixedlen.go -out=gen_fixedlen64.go gen "pcType=uint8,uint16,uint32,uint64 runnerType=uint64 runnerCapsType=Uint64"

// Processor is an interface for a simulated processor
type Processor interface {
	// Step runs the next instruction, returns error to indicate an unhandeled exception
	Step() error
}

// ErrorCallback allows a callback to be made when an error is returned
type ErrorCallback func(err error) error

func (instructionuint8) size() uint8 {
	return 1
}

func (instructionuint16) size() uint8 {
	return 2
}

func (instructionuint32) size() uint8 {
	return 4
}

func (instructionuint64) size() uint8 {
	return 8
}
