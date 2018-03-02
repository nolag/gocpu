package memory

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var anyIndex = uint64(1003)
var anyMemorySize = uint64(4000)
var anyValues = []byte{2, 4, 6, 15}
var anyByte = byte(10)

type CreateMem func(size uint64) Memory

// todo add tests for AccessViolationError

func RunReadWriteTest(t *testing.T, isBacked bool, create CreateMem) {
	// Given
	memory := create(anyMemorySize)
	valueLen := uint64(len(anyValues))

	// When
	wErr := memory.WriteRaw(anyValues, anyIndex)
	data, backed, rErr := memory.ReadRaw(anyIndex, valueLen)
	wOneErr := memory.WriteOneByte(anyByte, anyIndex + valueLen)
	b, rOneErr := memory.ReadOneByte(anyIndex + valueLen)

	// Then
	assert.NoError(t, wErr, "Read write memories should not return an error on write")
	assert.NoError(t, rErr, "Read write memories should not return an error on read")
	assert.NoError(t, wOneErr, "Read write memories should not return an error on byte writes")
	assert.NoError(t, rOneErr, "Read write memories should not return an error on byte reads")
	assert.Equal(t, isBacked, backed, "Memory backing does not match expected")
	assert.Equal(t, len(anyValues), len(data), "Wrong amount of data returned")
	assert.Equal(t, anyByte, b, "Incorrect byte returned after write/read")
	for index, element := range anyValues {
		assert.Equal(t, element, data[index], "Wrong element at index ", index)
	}

	if isBacked {
		data[1] = anyByte
		b, rOneErr = memory.ReadOneByte(anyIndex + 1)
		assert.NoError(t, rOneErr, "Read write memories should not return an error on byte reads")
		assert.Equal(t, anyByte, b, "If memory is backed, modifying raw reads must modify the memory")
	}
}

func RunSizeTest(t *testing.T, create CreateMem) {
	// Given
	memory := create(anyMemorySize)

	// When
	size := memory.Size()

	// Then
	assert.Equalf(t, anyMemorySize, size, "Expected size to be %v but was %v", anyMemorySize, size)
}

func RunBoundsTests(t *testing.T, create CreateMem) {
	// Given
	memory := create(anyMemorySize)
	sboR := AccessViolationError{anyMemorySize, 1, true}
	sboW := AccessViolationError{anyMemorySize, 1, false}
	overflow := AccessViolationError{math.MaxUint64, 2, true}

	// when
	// Write overflow is hard to test, I don't have enough RAM to have it get to testing that part, so just test read.
	_, _, actualOverflow := memory.ReadRaw(math.MaxUint64, 2)
	
	// Single byte tests
	_, actualSboR := memory.ReadOneByte(anyMemorySize)
	actualSboW := memory.WriteOneByte(anyByte, anyMemorySize)
	_, okR := memory.ReadOneByte(anyMemorySize - 1)
	okW := memory.WriteOneByte(anyByte, anyMemorySize - 1)

	// When - Then
	assertBoundary(t, memory, 0, anyMemorySize + 1, false, "Accessing more memory than is allocated must fail")
	assertBoundary(t, memory, anyIndex, anyMemorySize - anyIndex + 1, true,
		"Accessing memory starting in a valid location, but ending out of memory must fail")
	assertBoundary(t, memory, anyMemorySize, 1, true, "Starting out of memory must fail")
	assertAccessViolation(t, sboR, actualSboR, "Reading a single byte out of memory must fail")
	assert.NoError(t, okR, "Reading the last byte must not return an error")
	assertAccessViolation(t, sboW, actualSboW, "Writing a single byte out of memory must fail")
	assert.NoError(t, okW, "Writing the last byte must not return an error")
	assertAccessViolation(t, overflow, actualOverflow, "Overflowing the uint64 must fail")
}


// todo need to decide what bound (or both) to test
func assertBoundary(t *testing.T, memory Memory, startIndex uint64, notAllowedAmount uint64, fixByOffset bool, message string) {
	// Given
	expectedR := AccessViolationError{startIndex, notAllowedAmount, true}
	expectedW := AccessViolationError{startIndex, notAllowedAmount, false}

	// When
	_, _, actualR := memory.ReadRaw(startIndex, notAllowedAmount)
	actualW := memory.WriteRaw(make([]byte, notAllowedAmount), startIndex)

	var okR error
	var okW error
	if fixByOffset {
		_, _, okR = memory.ReadRaw(startIndex -1, notAllowedAmount)
		okW = memory.WriteRaw(make([]byte, notAllowedAmount), startIndex - 1)
	} else {
		_, _, okR = memory.ReadRaw(startIndex , notAllowedAmount - 1)
		okW = memory.WriteRaw(make([]byte, notAllowedAmount - 1), startIndex)
	}
	
	
	// Then
	assert.NoError(t, okR, "Reads within bounds must not fail for test ", message)
	assert.NoError(t, okW, "Writes within bounds must not fail for test ", message)
	assertAccessViolation(t, expectedR, actualR, "Reading: " + message)
	assertAccessViolation(t, expectedW, actualW, "Writing: " + message)
}

func assertAccessViolation(t* testing.T, expected AccessViolationError, actual error, msg string) {
	actualAsAccess, ok := actual.(*AccessViolationError)
	assert.Truef(t, ok, "%v: Wrong error type returned got %v", msg, actual)
	assert.Equal(t, expected, *actualAsAccess, msg)
}
