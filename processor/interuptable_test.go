package processor_test

import (
	"errors"
	"testing"

	"github.com/nolag/gocpu/processor"

	"github.com/nolag/gocpu/mock"
	"github.com/stretchr/testify/assert"
)

var anyQueueSize = 10

func TestInteruptableWithNoDelayedActions(t *testing.T) {
	// Given
	processorMock := mock.NewProcessor(nil)
	interuptable := processor.NewInteruptable(processorMock, anyQueueSize)

	// When
	err := interuptable.Step()

	// Then
	assert.Equal(t, 1, processorMock.NumTimesStepped, "Interuptable must always run step.")
	assert.NoError(t, err, "Error must not be returned from step if the base processor did not have an error")
}

func TestInteruptableWaitsTheCorrectNumberOfStepsToRunAnAction(t *testing.T) {
	// Given
	processorMock := mock.NewProcessor(nil)
	interuptable := processor.NewInteruptable(processorMock, anyQueueSize)
	delay0Call, delay0Verify, _ := mock.NewCallback(t, 1, "Making a callback 0")
	delay1Call, delay1Verify, delay1NVerify := mock.NewCallback(t, 1, "making a callback 1")
	delay3Call, delay3Verify, delay3NVerify := mock.NewCallback(t, 1, "making a callback 3")
	delay0 := processor.InteruptAction{Action: delay0Call, Delay: 0}
	delay1 := processor.InteruptAction{Action: delay1Call, Delay: 1}
	delay3 := processor.InteruptAction{Action: delay3Call, Delay: 3}

	// Do not order in executation order
	interuptable.Actions <- &delay1
	interuptable.Actions <- &delay0
	interuptable.Actions <- &delay3

	// When - Then
	interuptable.Step()
	delay0Verify()
	delay1NVerify()
	delay3NVerify()

	interuptable.Step()
	delay1Verify()
	delay3NVerify()

	interuptable.Step()
	delay3NVerify()

	interuptable.Step()
	delay3Verify()
}

func TestInteruptableWillRunAllActionsThatAreReady(t *testing.T) {
	// Given
	processorMock := mock.NewProcessor(nil)
	interuptable := processor.NewInteruptable(processorMock, anyQueueSize)
	delay0Call, delay0Verify, _ := mock.NewCallback(t, 1, "Making a callback 0")
	delay1Call, delay1Verify, _ := mock.NewCallback(t, 1, "making a callback 1")
	delay0 := processor.InteruptAction{Action: delay0Call, Delay: 0}
	delay1 := processor.InteruptAction{Action: delay1Call, Delay: 0}
	interuptable.Actions <- &delay0
	interuptable.Actions <- &delay1

	// When
	interuptable.Step()

	// Then
	delay0Verify()
	delay1Verify()
}

func TestInteruptableCallsStepBeforeRunningActions(t *testing.T) {
	// Given
	delay0Call, _, delay0NVerify := mock.NewCallback(t, 1, "Making a callback 0")
	delay0 := processor.InteruptAction{Action: delay0Call, Delay: 0}
	processorMock := mock.NewProcessorWithCallback(nil, delay0NVerify)
	interuptable := processor.NewInteruptable(processorMock, anyQueueSize)
	interuptable.Actions <- &delay0

	// When - Then
	interuptable.Step()
}

func TestInteruptableDoesNotTakeActionOnDelayedActionsWhenErroIsReturned(t *testing.T) {
	// Given
	anyError := errors.New("Anything")
	processorMock := mock.NewProcessor(anyError)
	interuptable := processor.NewInteruptable(processorMock, anyQueueSize)
	delay0Call, delay0Verify, delay0NVerify := mock.NewCallback(t, 1, "Making a callback 0")
	delay1Call, delay1Verify, delay1NVerify := mock.NewCallback(t, 1, "making a callback 1")
	delay0 := processor.InteruptAction{Action: delay0Call, Delay: 0}
	delay1 := processor.InteruptAction{Action: delay1Call, Delay: 1}
	interuptable.Actions <- &delay0
	interuptable.Actions <- &delay1

	// When - Then
	interuptable.Step()
	delay0NVerify()
	delay1NVerify()

	processorMock.ErrToReturn = nil
	interuptable.Step()
	delay0Verify()
	delay1NVerify()

	processorMock.ErrToReturn = anyError
	interuptable.Step()
	delay1NVerify()

	processorMock.ErrToReturn = nil
	interuptable.Step()
	delay1Verify()
}

func TestInteruptableReturnsErrorFromBaseProcessor(t *testing.T) {
	// Given
	anyError := errors.New("Anything")
	processorMock := mock.NewProcessor(anyError)
	interuptable := processor.NewInteruptable(processorMock, anyQueueSize)

	// When
	err := interuptable.Step()

	// Then
	assert.Equal(t, anyError, err, "Error returned must match the underlying processor's error")
}

func TestInteruptableSeesInterupsDuringAnInstructionBeforeNextStep(t *testing.T) {
	// Given
	delay0Call, delay0Verify, _ := mock.NewCallback(t, 1, "Making a callback 0")
	delay1Call, delay1Verify, delay1NVerify := mock.NewCallback(t, 1, "Making a callback 1")
	delay0 := processor.InteruptAction{Action: delay0Call, Delay: 0}
	delay1 := processor.InteruptAction{Action: delay1Call, Delay: 1}

	processorMock := mock.NewProcessor(nil)
	interuptable := processor.NewInteruptable(processorMock, anyQueueSize)

	callback := func() {
		interuptable.Actions <- &delay0
		interuptable.Actions <- &delay1
	}

	processorMock.Callback = callback

	// When - Then
	interuptable.Step()
	delay0Verify()
	delay1NVerify()

	processorMock.Callback = nil
	interuptable.Step()
	delay1Verify()
}

func TestNewInteruptableSetsCorrectChannelSize(t *testing.T) {
	// Given
	delay0Call, _, _ := mock.NewCallback(t, 1, "Making a callback 0")
	processorMock := mock.NewProcessor(nil)
	interuptable := processor.NewInteruptable(processorMock, anyQueueSize)
	delay0 := processor.InteruptAction{Action: delay0Call, Delay: 0}

	// When - Then
	// Not blocking proves at least that many, hard to prove exactly that many...
	for i := 0; i < anyQueueSize; i++ {
		interuptable.Actions <- &delay0
	}
}
