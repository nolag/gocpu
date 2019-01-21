package testhelper

import (
	"fmt"
	"math"
	"testing"

	"github.com/nolag/gocpu/memory"
	"github.com/stretchr/testify/assert"
)

var anyIndex = uint64(1000)
var anyMemorySize = uint64(1400)
var anyValues = []byte{2, 4, 6, 15}
var anyByte = byte(10)

// CreateMem serves as a helper to allow this test helper to create memory for testing
// Tests that use this are expected to assume 1,2, and 4 bytes are valid read/write
// but it is expected that it can be created without this boundary to allow boundary tests to work.
type CreateMem func(size uint64) memory.Memory

type boundaryTestType int

const (
	startOnBoundary boundaryTestType = iota
	endAfterBoundary
	startAfter
)

// RunReadWriteTest runs a test for reading and writing
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
	assert.Equal(t, anyByte, b, "Incorrect byte returned after write/read")
	assert.ElementsMatch(t, anyValues, data, "Wrong amount of data returned")

	if isBacked {
		data[1] = anyByte
		b, rOneErr = mem.ReadOneByte(anyIndex + 1)
		assert.NoError(t, rOneErr, "Read write memories should not return an error on byte reads")
		assert.Equal(t, anyByte, b, "If memory is backed, modifying raw reads must modify the memory")
	}
}

// RunSizeTest tests the size of a memory
func RunSizeTest(t *testing.T, create CreateMem) {
	// Given
	mem := create(anyMemorySize)

	// When
	size := mem.Size()

	// Then
	assert.Equalf(t, anyMemorySize, size, "Expected size to be %v but was %v", anyMemorySize, size)
}

// RunBoundsTests tests the memory's bounderies
func RunBoundsTests(t *testing.T, create CreateMem) {
	// Given
	mem := create(anyMemorySize)
	sboR := memory.AccessViolationError{Location: anyMemorySize, NumBytes: 1, WasRead: true}
	sboW := memory.AccessViolationError{Location: anyMemorySize, NumBytes: 1, WasRead: false}
	overflow := memory.AccessViolationError{Location: math.MaxUint64, NumBytes: 2, WasRead: true}

	// when
	// Write overflow is hard to test, I don't have enough RAM to have it get to testing that part, so just test read.
	_, _, actualOverflow := mem.ReadRaw(math.MaxUint64, 2)

	// Single byte tests
	_, actualSboR := mem.ReadOneByte(anyMemorySize)
	actualSboW := mem.WriteOneByte(anyByte, anyMemorySize)
	_, okR := mem.ReadOneByte(anyMemorySize - 1)
	okW := mem.WriteOneByte(anyByte, anyMemorySize-1)

	// When - Then
	assertBoundary(t, create, startOnBoundary, "Accessing memory after the beginning must fail")
	assertBoundary(t, create, endAfterBoundary, "Accessing memory starting in a valid location, but ending out of memory must fail")
	assertBoundary(t, create, startAfter, "Starting out of memory must fail")
	assertAccessViolation(t, sboR, actualSboR, "Reading a single byte out of memory must fail")
	assert.NoError(t, okR, "Reading the last byte must not return an error")
	assertAccessViolation(t, sboW, actualSboW, "Writing a single byte out of memory must fail")
	assert.NoError(t, okW, "Writing the last byte must not return an error")
	assertAccessViolation(t, overflow, actualOverflow, "Overflowing the uint64 must fail")
}

// RunPowTwoAllignmentFailures tests that word allignment is nessesary
func RunPowTwoAllignmentFailures(t *testing.T, maxPow int, create CreateMem) {
	mem := create(anyMemorySize)
	for p := uint(0); maxPow < maxPow; p++ {
		nb := uint64(1) << p
		for i := uint64(0); i < nb; i++ {
			startIndex := uint64(1000) + uint64(i)
			_, _, actualR := mem.ReadRaw(startIndex, nb)
			actualW := mem.WriteRaw(make([]byte, nb), startIndex)
			expectedR := memory.AccessViolationError{Location: startIndex, NumBytes: nb, WasRead: true}
			expectedW := memory.AccessViolationError{Location: startIndex, NumBytes: nb, WasRead: false}
			message := fmt.Sprint("Unaligned access on read size ", nb, " when offset by ", i)
			assertAccessViolation(t, expectedR, actualR, "Reading: "+message)
			assertAccessViolation(t, expectedW, actualW, "Writing: "+message)
		}
	}
}

// RunDisallowedSize shows that numBytes is not allowed to be read
func RunDisallowedSize(t *testing.T, numBytes uint64, create CreateMem) {
	mem := create(anyMemorySize)
	startIndex := uint64(1000)
	_, _, actualR := mem.ReadRaw(startIndex, numBytes)
	actualW := mem.WriteRaw(make([]byte, numBytes), startIndex)
	expectedR := memory.AccessViolationError{Location: startIndex, NumBytes: numBytes, WasRead: true}
	expectedW := memory.AccessViolationError{Location: startIndex, NumBytes: numBytes, WasRead: false}
	assertAccessViolation(t, expectedR, actualR, "Reading disallowed number of bytes must fail")
	assertAccessViolation(t, expectedW, actualW, "Writing disallowed number of two bytes must fail")
}

func assertBoundary(t *testing.T, create CreateMem, boundary boundaryTestType, message string) {
	var startIndex uint64
	var size uint64
	var okR error
	var okW error
	var mem memory.Memory
	switch boundary {
	case startAfter:
		mem = create(anyMemorySize)
		size = 2
		startIndex = anyMemorySize + 2
		_, _, okR = mem.ReadRaw(startIndex-4, size)
		okW = mem.WriteRaw(make([]byte, size), startIndex-4)
	case endAfterBoundary:
		mem = create(anyMemorySize - 2)
		size = 4
		startIndex = anyMemorySize - 4
		_, _, okR = mem.ReadRaw(startIndex, 2)
		okW = mem.WriteRaw(make([]byte, 2), startIndex-2)
	case startOnBoundary:
		mem = create(anyMemorySize)
		size = 1
		startIndex = anyMemorySize
		_, _, okR = mem.ReadRaw(startIndex-1, 1)
		okW = mem.WriteRaw(make([]byte, 1), startIndex-1)
	}

	// Given
	expectedR := memory.AccessViolationError{Location: startIndex, NumBytes: size, WasRead: true}
	expectedW := memory.AccessViolationError{Location: startIndex, NumBytes: size, WasRead: false}

	// When
	_, _, actualR := mem.ReadRaw(startIndex, size)
	actualW := mem.WriteRaw(make([]byte, size), startIndex)

	// Then
	assert.NoError(t, okR, "Reads within bounds must not fail for test %v", message)
	assert.NoError(t, okW, "Writes within bounds must not fail for test %v", message)
	assertAccessViolation(t, expectedR, actualR, "Reading: "+message)
	assertAccessViolation(t, expectedW, actualW, "Writing: "+message)
}

func assertAccessViolation(t *testing.T, expected memory.AccessViolationError, actual error, msg string) {
	if !assert.Error(t, actual, "%v: Had no exception", msg) {
		return
	}

	actualAsAccess, ok := actual.(*memory.AccessViolationError)
	assert.Truef(t, ok, "%v: Wrong error type returned got %v", msg, actual)
	assert.Equal(t, expected, *actualAsAccess, msg)
}
