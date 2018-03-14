package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"

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

// NewCallback creates a callback (callback), an assertion (verify) it was called expectedNumCalls,
// and an assertion it wasn't called expectedNumCalls (nverify)
func NewCallback(t *testing.T, expectedNumCalls uint, when string) (callback func(), verify func(), nverify func()) {
	numTimesCalled := uint(0)
	callback = func() {
		numTimesCalled++
		assert.Falsef(t, numTimesCalled < expectedNumCalls, "Callback for %v was called back %v times.", when, numTimesCalled)
	}

	verify = func() {
		assert.Equal(t, expectedNumCalls, numTimesCalled, "Callback for %v", when)
	}

	nverify = func() {
		assert.NotEqual(t, expectedNumCalls, numTimesCalled, "Callback for %v called before expected", when)
	}

	return
}
