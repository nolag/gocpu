package mock

import (
	"testing"

	"github.com/nolag/gocpu/processor"
)

// CreateUnexpectedCallback makes a callback that fails when called
func CreateUnexpectedCallback(t *testing.T, when string) processor.ErrorCallback {
	return func(err error) error {
		t.Fatalf("Unexpected callback on %v.", when)
		return nil
	}
}

// CreateUnexpectedInstructionRunner32Callback makes a callback that fails when called
func CreateUnexpectedInstructionRunner32Callback(t *testing.T, when string) processor.InstructionRunner32 {
	return unexpectedInstructionRunner32Callback{t, when}
}

type unexpectedInstructionRunner32Callback struct {
	t    *testing.T
	when string
}

func (runner unexpectedInstructionRunner32Callback) RunInstruction32(instruction uint32) error {
	runner.t.Fatalf("Unexpected callback on %v.", runner.when)
	return nil
}
