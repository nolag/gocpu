package mock

// InstructionRunner32 mocks out an instruciton runner
import (
	"testing"

	"github.com/nolag/gocpu/processor"
	"github.com/stretchr/testify/assert"
)

// InstructionRunner32 mocks a processor.InstructionRunner32
type InstructionRunner32 struct {
	Core          processor.Core32
	ExpectedError error
	ExpectedPc    interface{}
	ExpectedValue interface{}
	NumTimesRun   int
	T             *testing.T
}

// RunInstruction32 runs a single 32 bit instrution, without incrementing the PC
func (runner *InstructionRunner32) RunInstruction32(instruction uint32) error {
	assert.Equal(runner.T, runner.ExpectedValue, instruction, "Wrong instruction passed to runner")
	assert.Equal(runner.T, runner.ExpectedPc, runner.Core.Pc.Value32(), "Wrong PC at time of")
	runner.NumTimesRun++
	return runner.ExpectedError
}
