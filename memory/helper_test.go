package memory_test

import (
	"math"
	"testing"

	"github.com/nolag/gocpu/memory"
	"github.com/stretchr/testify/assert"
)

var anyIndex = uint64(1003)
var anyMemorySize = uint64(4000)
var anyValues = []byte{2, 4, 6, 15}
var anyByte = byte(10)

type CreateMem func(size uint64) memory.Memory

func RunReadWriteTest(t *testing.T, isBacked bool, create CreateMem) {
	// Given
	mem := create(anyMemorySize)
	valueLen := uint64(len(anyValues))

	// When
	wErr := mem.WriteRaw(anyValues, anyIndex)
	data, backed, rErr := mem.ReadRaw(anyIndex, valueLen)
	wOneErr := mem.WriteOneByte(anyByte, anyIndex+valueLen)
	b, rOneErr := mem.ReadOneByte(anyIndex + valueLen)

	// Then
	assert.NoError(t, wErr, "Read write memories should not return an error on write")
	assert.NoError(t, rErr, "Read write memories should not return an error on read")
	assert.NoError(t, wOneErr, "Read write memories should not return an error on byte writes")
	assert.NoError(t, rOneErr, "Read write memories should not return an error on byte reads")
	assert.Equal(t, isBacked, backed, "mem.Memory backing does not match expected")
	assert.Equal(t, len(anyValues), len(data), "Wrong amount of data returned")
	assert.Equal(t, anyByte, b, "Incorrect byte returned after write/read")
	for index, element := range anyValues {
		assert.Equal(t, element, data[index], "Wrong element at index ", index)
	}

	if isBacked {
		data[1] = anyByte
		b, rOneErr = mem.ReadOneByte(anyIndex + 1)
		assert.NoError(t, rOneErr, "Read write memories should not return an error on byte reads")
		assert.Equal(t, anyByte, b, "If memory is backed, modifying raw reads must modify the memory")
	}
}

func RunSizeTest(t *testing.T, create CreateMem) {
	// Given
	mem := create(anyMemorySize)

	// When
	size := mem.Size()

	// Then
	assert.Equalf(t, anyMemorySize, size, "Expected size to be %v but was %v", anyMemorySize, size)
}

func RunBoundsTests(t *testing.T, create CreateMem) {
	// Given
	mem := create(anyMemorySize)
	sboR := memory.AccessViolationError{anyMemorySize, 1, true}
	sboW := memory.AccessViolationError{anyMemorySize, 1, false}
	overflow := memory.AccessViolationError{math.MaxUint64, 2, true}

	// when
	// Write overflow is hard to test, I don't have enough RAM to have it get to testing that part, so just test read.
	_, _, actualOverflow := mem.ReadRaw(math.MaxUint64, 2)

	// Single byte tests
	_, actualSboR := mem.ReadOneByte(anyMemorySize)
	actualSboW := mem.WriteOneByte(anyByte, anyMemorySize)
	_, okR := mem.ReadOneByte(anyMemorySize - 1)
	okW := mem.WriteOneByte(anyByte, anyMemorySize-1)

	// When - Then
	assertBoundary(t, mem, 0, anyMemorySize+1, false, "Accessing more memory than is allocated must fail")
	assertBoundary(t, mem, anyIndex, anyMemorySize-anyIndex+1, true,
		"Accessing memory starting in a valid location, but ending out of memory must fail")
	assertBoundary(t, mem, anyMemorySize, 1, true, "Starting out of memory must fail")
	assertAccessViolation(t, sboR, actualSboR, "Reading a single byte out of memory must fail")
	assert.NoError(t, okR, "Reading the last byte must not return an error")
	assertAccessViolation(t, sboW, actualSboW, "Writing a single byte out of memory must fail")
	assert.NoError(t, okW, "Writing the last byte must not return an error")
	assertAccessViolation(t, overflow, actualOverflow, "Overflowing the uint64 must fail")
}

// todo need to decide what bound (or both) to test
func assertBoundary(t *testing.T, mem memory.Memory, startIndex uint64, notAllowedAmount uint64, fixByOffset bool, message string) {
	// Given
	expectedR := memory.AccessViolationError{startIndex, notAllowedAmount, true}
	expectedW := memory.AccessViolationError{startIndex, notAllowedAmount, false}

	// When
	_, _, actualR := mem.ReadRaw(startIndex, notAllowedAmount)
	actualW := mem.WriteRaw(make([]byte, notAllowedAmount), startIndex)

	var okR error
	var okW error
	if fixByOffset {
		_, _, okR = mem.ReadRaw(startIndex-1, notAllowedAmount)
		okW = mem.WriteRaw(make([]byte, notAllowedAmount), startIndex-1)
	} else {
		_, _, okR = mem.ReadRaw(startIndex, notAllowedAmount-1)
		okW = mem.WriteRaw(make([]byte, notAllowedAmount-1), startIndex)
	}

	// Then
	assert.NoError(t, okR, "Reads within bounds must not fail for test ", message)
	assert.NoError(t, okW, "Writes within bounds must not fail for test ", message)
	assertAccessViolation(t, expectedR, actualR, "Reading: "+message)
	assertAccessViolation(t, expectedW, actualW, "Writing: "+message)
}

func assertAccessViolation(t *testing.T, expected memory.AccessViolationError, actual error, msg string) {
	actualAsAccess, ok := actual.(*memory.AccessViolationError)
	assert.Truef(t, ok, "%v: Wrong error type returned got %v", msg, actual)
	assert.Equal(t, expected, *actualAsAccess, msg)
}
