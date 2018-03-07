package processor_test

import (
	"encoding/binary"
	"errors"
	"testing"

	"github.com/nolag/gocpu/mock"
	"github.com/nolag/gocpu/processor"
	"github.com/nolag/gocpu/registers"
	"github.com/stretchr/testify/assert"
)

func TestFixedLen32IncrementsPc(t *testing.T) {
	// Given
	processor, _ := CreateTestFixedLenProcessor32(t)
	pcBefore := processor.Pc.Value32()

	// When
	err := processor.Step()

	// Then
	assert.Equal(t, pcBefore+4, processor.Pc.Value32(), "PC not incremented")
	assert.NoError(t, err, "Step must not return an error if none was returned to it")
}

func TestFixedLen32MakesCallWithCorrectValue(t *testing.T) {
	// Given
	processor, runnerMock := CreateTestFixedLenProcessor32(t)

	// When
	processor.Step()

	// Then
	assert.Equal(t, 1, runnerMock.NumTimesRun, "Callback must happen exactly once")
}

func TestFixedLen32ErrorReturnedFromMemoryWithCallback(t *testing.T) {
	// Given
	errExpected := errors.New("Wrong failure to return")
	errFail := errors.New("Anything")
	numTimesCallbackRun := 0
	callback := func(err error) error {
		assert.Equal(t, errFail, err, "Callback made with wrong error")
		numTimesCallbackRun++
		return errExpected
	}

	// When - Then
	RunBadMemoryTest(t, errFail, errExpected, callback)
	assert.Equal(t, 1, numTimesCallbackRun, "Callback for memory errors must happen exactly once")
}

func TestFixedLen32ErrorReturnedFromMemoryWithNoCallback(t *testing.T) {
	// Given
	errFail := errors.New("Anything")

	// When - Then
	RunBadMemoryTest(t, errFail, errFail, nil)
}

func RunBadMemoryTest(t *testing.T, errFail error, errExpected error, callback processor.ErrorCallback) {
	// Given
	processor, runner := CreateTestFixedLenProcessor32(t)
	memory := &mock.Memory{Data: nil, ExpectedIndex: uint64(processor.Pc.Value32()), Fail: errFail, T: t}
	processor.Memory = memory
	processor.MemoryReadFailureCallback = callback
	processor.InstructionRunner32 = mock.NewUnexpectedInstructionRunner32Callback(
		t,
		"To instruciton runner when error is returned from memory")

	// When
	err := processor.Step()

	// Then
	assert.Equal(t, errExpected, err, "Wrong error returned")
	assert.Equal(t, 0, runner.NumTimesRun, "Wrong number of times to run the callback")
}

func CreateTestFixedLenProcessor32(t *testing.T) (processor.FixedLen32, *mock.InstructionRunner32) {
	anyValue := uint32(500)
	anyOtherValue := uint32(0xF00DBEEF)
	anyPc := registers.UintRegister32(anyValue)
	data := make([]byte, 4)
	anyEndianness := binary.BigEndian
	anyEndianness.PutUint32(data, anyOtherValue)
	anyMemory := &mock.Memory{Data: data, ExpectedIndex: uint64(anyValue), Fail: nil, T: t}
	anyRegisters := []registers.Register32{&anyPc}
	anyCore := processor.Core32{Endianness: anyEndianness, Memory: anyMemory, Registers: anyRegisters, Pc: &anyPc}
	notCalledErrorCallback := mock.NewUnexpectedCallback(t, "running instruction with no callback")
	anyRunner := &mock.InstructionRunner32{Core: anyCore, ExpectedError: nil, ExpectedPc: anyValue + 4, ExpectedValue: anyOtherValue, T: t}
	processor := processor.FixedLen32{Core32: anyCore, InstructionRunner32: anyRunner, MemoryReadFailureCallback: notCalledErrorCallback}
	return processor, anyRunner
}
