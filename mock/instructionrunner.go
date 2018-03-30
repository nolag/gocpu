package mock

// InstructionRunner32 mocks out an instruciton runner
import (
	"testing"

	"github.com/nolag/gocpu/registers"

	"github.com/stretchr/testify/assert"
)

// InstructionRunner32 mocks a processor.InstructionRunnerUint32
type InstructionRunner32 struct {
	ExpectedError error
	ExpectedPc    uint64
	ExpectedValue interface{}
	NumTimesRun   int
	Pc            registers.ProgramCounter
	T             *testing.T
}

// RunUint32 runs a single 32 bit instrution, without incrementing the PC
func (runner *InstructionRunner32) RunUint32(instruction uint32) error {
	assert.Equal(runner.T, runner.ExpectedValue, instruction, "Wrong instruction passed to runner")
	assert.Equal(runner.T, runner.ExpectedPc, runner.Pc.ReadAsPc(), "Wrong PC at time of run")
	runner.NumTimesRun++
	return runner.ExpectedError
}
