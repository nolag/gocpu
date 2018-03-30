package processor_test

import (
	"encoding/binary"
	"errors"
	"testing"

	"github.com/nolag/gocpu/registers"

	"github.com/nolag/gocpu/mock"
	"github.com/nolag/gocpu/processor"
	"github.com/stretchr/testify/assert"
)

/***************************************
Tests below here test generated methods*
***************************************/

func TestFixedGeneratedIncrementsPc(t *testing.T) {
	// Given
	processor, _ := CreateTestFixedLenProcessor32(t)
	pcBefore := processor.Pc.ReadAsPc()

	// When
	err := processor.Step()

	// Then
	assert.Equal(t, pcBefore+4, processor.Pc.ReadAsPc(), "PC not incremented correctly")
	assert.NoError(t, err, "Step must not return an error if none was returned to it")
}

func TestFixedGeneratedMakesCallWithCorrectValue(t *testing.T) {
	// Given
	processor, runnerMock := CreateTestFixedLenProcessor32(t)

	// When
	processor.Step()

	// Then
	assert.Equal(t, 1, runnerMock.NumTimesRun, "Callback must happen exactly once")
}

func TestFixedGeneratedErrorReturnedFromMemoryWithCallback(t *testing.T) {
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

func TestFixedGeneratedErrorReturnedFromMemoryWithNoCallback(t *testing.T) {
	// Given
	errFail := errors.New("Anything")

	// When - Then
	RunBadMemoryTest(t, errFail, errFail, nil)
}

func RunBadMemoryTest(t *testing.T, errFail error, errExpected error, callback processor.ErrorCallback) {
	// Given
	processor, runner := CreateTestFixedLenProcessor32(t)
	memory := &mock.Memory{Data: nil, ExpectedIndex: processor.Pc.ReadAsPc(), Fail: errFail, T: t}
	processor.Memory = memory
	processor.MemoryReadFailureCallback = callback
	processor.InstructionRunnerUint32 = mock.NewUnexpectedInstructionRunner32Callback(
		t,
		"To instruciton runner when error is returned from memory")

	// When
	err := processor.Step()

	// Then
	assert.Equal(t, errExpected, err, "Wrong error returned")
	assert.Equal(t, 0, runner.NumTimesRun, "Wrong number of times to run the callback")
}

func CreateTestFixedLenProcessor32(t *testing.T) (*processor.FixedInstructionLenRunnerUint32, *mock.InstructionRunner32) {
	anyValue := uint64(500)
	anyPc := registers.RegisterUint64(anyValue)
	anyOtherValue := uint32(0xF00DBEEF)
	data := make([]byte, 4)
	anyEndianness := binary.BigEndian
	anyEndianness.PutUint32(data, anyOtherValue)
	anyMemory := &mock.Memory{Data: data, ExpectedIndex: uint64(anyValue), Fail: nil, T: t}

	notCalledErrorCallback := mock.NewUnexpectedCallback(t, "running instruction with no callback")
	anyRunner := &mock.InstructionRunner32{ExpectedError: nil, ExpectedPc: anyValue + 4, ExpectedValue: anyOtherValue, Pc: &anyPc, T: t}

	processor := processor.FixedInstructionLenRunnerUint32{
		ByteOrder:               anyEndianness,
		Memory:                  anyMemory,
		InstructionRunnerUint32: anyRunner,
		Pc: &anyPc,
		MemoryReadFailureCallback: notCalledErrorCallback}

	return &processor, anyRunner
}
