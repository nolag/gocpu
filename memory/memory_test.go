package memory

import (
	"encoding/binary"
	"errors"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var any16BitUint = uint16(0xBEEF)
var any32BitUint = uint32(0xC001BABE)
var any32BitFloat = math.Float32frombits(any32BitUint)
var any64BitUint = uint64(0xC001BEEFABADBABE)
var any64BitFloat = math.Float64frombits(any64BitUint)
var anyEndianness = binary.BigEndian
var errAny = errors.New("Anything")

func TestAccessViolationErrorMessage(t *testing.T) {
	// Given
	anyNumber := uint64(123)
	anyOtherNumber := uint64(323)
	anyWasRead := true
	err := AccessViolationError{anyNumber, anyOtherNumber, anyWasRead}

	// When
	msg := err.Error()

	// Then
	assert.Equal(t, "Cannot access memory location: 123 Num bytes: 323 Was read: True", msg, "Wrong message for acces violation")
}

func TestReadFloat32(t *testing.T) {
	readFloat32 := func(memory Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error) {
		return ReadFloat32(memory, byteOrder, index)
	}

	memory := Setup32BitMockMemory(t)
	RunMemoryReadTest(t, memory, any32BitFloat, readFloat32, "read flaot32")
}

func TestReadFloat64(t *testing.T) {
	readFloat64 := func(memory Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error) {
		return ReadFloat64(memory, byteOrder, index)
	}

	memory := Setup64BitMockMemory(t)
	RunMemoryReadTest(t, memory, any64BitFloat, readFloat64, "Read float64")
}

func TestReadUint16(t *testing.T) {
	readUint16 := func(memory Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error) {
		return ReadUint16(memory, byteOrder, index)
	}

	memory := Setup16BitMockMemory(t)
	RunMemoryReadTest(t, memory, any16BitUint, readUint16, "Read uint16")
}

func TestReadUint32(t *testing.T) {
	readUint32 := func(memory Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error) {
		return ReadUint32(memory, byteOrder, index)
	}

	memory := Setup32BitMockMemory(t)
	RunMemoryReadTest(t, memory, any32BitUint, readUint32, "Read uint32")
}

func TestReadUint64(t *testing.T) {
	readUint64 := func(memory Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error) {
		return ReadUint64(memory, byteOrder, index)
	}

	memory := Setup64BitMockMemory(t)
	RunMemoryReadTest(t, memory, any64BitUint, readUint64, "Read uint64")
}

func TestWriteUint16(t *testing.T) {
	// Given
	writeUint16 := func(memory Memory, byteOrder binary.ByteOrder, index uint64) error {
		return WriteUint16(memory, byteOrder, any16BitUint, index)
	}

	memory := Setup16BitMockMemory(t)

	// When - Then
	RunMemoryWriteTest(t, memory, writeUint16, "Write uint16")
}

func TestWriteFloat32(t *testing.T) {
	// Given
	writeFloat32 := func(memory Memory, byteOrder binary.ByteOrder, index uint64) error {
		return WriteFloat32(memory, byteOrder, any32BitFloat, index)
	}

	memory := Setup32BitMockMemory(t)

	// When - Then
	RunMemoryWriteTest(t, memory, writeFloat32, "Write float32")
}

func TestWriteFloat64(t *testing.T) {
	// Given
	writeFloat64 := func(memory Memory, byteOrder binary.ByteOrder, index uint64) error {
		return WriteFloat64(memory, byteOrder, any64BitFloat, index)
	}

	memory := Setup64BitMockMemory(t)

	// When - Then
	RunMemoryWriteTest(t, memory, writeFloat64, "Write float64")
}

func TestWriteUint32(t *testing.T) {
	// Given
	writeUint32 := func(memory Memory, byteOrder binary.ByteOrder, index uint64) error {
		return WriteUint32(memory, byteOrder, any32BitUint, index)
	}

	memory := Setup32BitMockMemory(t)

	// When - Then
	RunMemoryWriteTest(t, memory, writeUint32, "Write uint32")
}

func TestWriteUint64(t *testing.T) {
	// Given
	writeUint64 := func(memory Memory, byteOrder binary.ByteOrder, index uint64) error {
		return WriteUint64(memory, byteOrder, any64BitUint, index)
	}

	memory := Setup64BitMockMemory(t)

	// When - Then
	RunMemoryWriteTest(t, memory, writeUint64, "Write uint64")
}

type TestMemRead func(memory Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error)
type TestMemWrite func(memory Memory, byteOrder binary.ByteOrder, index uint64) error

func RunMemoryReadTest(t *testing.T, goodMem Memory, expected interface{}, test TestMemRead, testing string) {
	// Given
	badMem := SetupFailMemory(t)

	// When
	val, ok := test(goodMem, anyEndianness, anyIndex)
	_, err := test(badMem, anyEndianness, anyIndex)

	// Then
	assert.Equal(t, expected, val, "Wrong value on ", testing)
	assert.Equal(t, errAny, err, "Errors must be passed down from memory on ", testing)
	assert.NoError(t, ok, "Unexpected error on ", testing)
}

func RunMemoryWriteTest(t *testing.T, goodMem *MockMemory, test TestMemWrite, testing string) {
	// Given
	expected := goodMem.Data
	goodMem.Data = nil
	badMem := SetupFailMemory(t)

	// When
	ok := test(goodMem, anyEndianness, anyIndex)
	err := test(badMem, anyEndianness, anyIndex)

	// Then
	assert.Equal(t, errAny, err, "Errors must be passed down from memory on %v", testing)
	assert.NoError(t, ok, "Unexpected error on %v", testing)
	assert.Len(t, goodMem.Data, len(expected), "Wrong bytes set %v", testing)
	for index, val := range expected {
		assert.Equal(t, val, goodMem.Data[index], "Wong value at index %v %v", index, testing)
	}
}

func Setup16BitMockMemory(t *testing.T) *MockMemory {
	data := make([]byte, 2)
	anyEndianness.PutUint16(data, any16BitUint)
	return &MockMemory{t, data, false}
}

func Setup32BitMockMemory(t *testing.T) *MockMemory {
	data := make([]byte, 4)
	anyEndianness.PutUint32(data, any32BitUint)
	return &MockMemory{t, data, false}
}

func Setup64BitMockMemory(t *testing.T) *MockMemory {
	data := make([]byte, 8)
	anyEndianness.PutUint64(data, any64BitUint)
	return &MockMemory{t, data, false}
}

func SetupFailMemory(t *testing.T) *MockMemory {
	data := make([]byte, 1)
	return &MockMemory{t, data, true}
}

type MockMemory struct {
	T    *testing.T
	Data []byte
	Fail bool
}

func (memory *MockMemory) ReadOneByte(index uint64) (byte, error) {
	assert.Equal(memory.T, anyIndex, index, "Wrong index in read one byte")
	if memory.Fail {
		return 0, errAny
	}

	return memory.Data[0], nil
}

func (memory *MockMemory) ReadRaw(startIndex uint64, numBytes uint64) (data []byte, backed bool, err error) {
	assert.Equal(memory.T, anyIndex, startIndex, "Wrong index used")
	if memory.Fail {
		return nil, false, errAny
	}

	return memory.Data, false, nil
}

func (memory *MockMemory) Size() uint64 {
	return uint64(len(memory.Data))
}

func (memory *MockMemory) WriteOneByte(val byte, index uint64) error {
	assert.Equal(memory.T, anyIndex, index, "Wrong index used")
	if memory.Fail {
		return errAny
	}

	memory.Data[0] = val
	return nil
}

func (memory *MockMemory) WriteRaw(data []byte, startIndex uint64) error {
	assert.Equal(memory.T, anyIndex, startIndex, "Wrong index used")
	if memory.Fail {
		return errAny
	}

	memory.Data = data
	return nil
}
