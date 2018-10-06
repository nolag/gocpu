// Package processor provides the bascis for implmenting a processor
package processor

//go:generate genny -pkg processor -in=templates/generic_instructionrunner.go -out=gen_instructionrunner.go gen "instructionType=uint8,uint16,uint32,uint64"
//go:generate genny -pkg processor -in=templates/generic_fixedlen.go -out=gen_fixedlen.go gen "runnerType=uint8,uint16,uint32,uint64"

// Processor is an interface for a simulated processor
type Processor interface {
	// Step runs the next instruction, returns error to indicate an unhandeled exception
	Step() error
}

// ErrorCallback allows a callback to be made when an error is returned
type ErrorCallback func(err error) error

func (instructionUint8) size() uint8 {
	return 1
}

func (instructionUint16) size() uint8 {
	return 2
}

func (instructionUint32) size() uint8 {
	return 4
}

func (instructionUint64) size() uint8 {
	return 8
}
