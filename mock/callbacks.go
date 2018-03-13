package mock

import (
	"testing"

	"github.com/nolag/gocpu/processor"
)

// NewUnexpectedCallback makes a callback that fails when called
func NewUnexpectedCallback(t *testing.T, when string) processor.ErrorCallback {
	return func(err error) error {
		t.Fatalf("Unexpected callback on %v.", when)
		return nil
	}
}

// NewUnexpectedInstructionRunner32Callback makes a callback that fails when called
func NewUnexpectedInstructionRunner32Callback(t *testing.T, when string) processor.InstructionRunnerUint32 {
	return unexpectedInstructionRunner32Callback{t, when}
}

type unexpectedInstructionRunner32Callback struct {
	t    *testing.T
	when string
}

func (runner unexpectedInstructionRunner32Callback) RunUint32(instruction uint32) error {
	runner.t.Fatalf("Unexpected callback on %v.", runner.when)
	return nil
}
