package mock

// InstructionRunner32 mocks out an instruciton runner
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// InstructionRunner32 mocks a processor.InstructionRunnerUint32
type InstructionRunner32 struct {
	ExpectedError error
	ExpectedPc    interface{}
	ExpectedValue interface{}
	NumTimesRun   int
	PcGetter      func() interface{}
	T             *testing.T
}

// RunUint32 runs a single 32 bit instrution, without incrementing the PC
func (runner *InstructionRunner32) RunUint32(instruction uint32) error {
	assert.Equal(runner.T, runner.ExpectedValue, instruction, "Wrong instruction passed to runner")
	assert.Equal(runner.T, runner.ExpectedPc, runner.PcGetter(), "Wrong PC at time of")
	runner.NumTimesRun++
	return runner.ExpectedError
}
