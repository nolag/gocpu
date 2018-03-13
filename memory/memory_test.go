package memory_test

import (
	"encoding/binary"
	"errors"
	"math"
	"testing"

	"github.com/nolag/gocpu/memory"
	"github.com/nolag/gocpu/mock"
	"github.com/stretchr/testify/assert"
)

var any8BitUint = uint8(0xFA)
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
	err := memory.AccessViolationError{Location: anyNumber, NumBytes: anyOtherNumber, WasRead: anyWasRead}

	// When
	msg := err.Error()

	// Then
	assert.Equal(t, "Cannot access memory location: 123 Num bytes: 323 Was read: True", msg, "Wrong message for acces violation")
}

func TestReadFloat32(t *testing.T) {
	readFloat32 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error) {
		return memory.ReadFloat32(mem, byteOrder, index)
	}

	mem := Setup32BitMockMemory(t)
	RunMemoryReadTest(t, mem, any32BitFloat, readFloat32, "read flaot32")
}

func TestReadFloat64(t *testing.T) {
	readFloat64 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error) {
		return memory.ReadFloat64(mem, byteOrder, index)
	}

	mem := Setup64BitMockMemory(t)
	RunMemoryReadTest(t, mem, any64BitFloat, readFloat64, "Read float64")
}

func TestReadUint18(t *testing.T) {
	readUint8 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error) {
		return memory.ReadUint8(mem, byteOrder, index)
	}

	mem := Setup8BitMockMemory(t)
	RunMemoryReadTest(t, mem, any8BitUint, readUint8, "Read uint8")
}

func TestReadUint16(t *testing.T) {
	readUint16 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error) {
		return memory.ReadUint16(mem, byteOrder, index)
	}

	mem := Setup16BitMockMemory(t)
	RunMemoryReadTest(t, mem, any16BitUint, readUint16, "Read uint16")
}

func TestReadUint32(t *testing.T) {
	readUint32 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error) {
		return memory.ReadUint32(mem, byteOrder, index)
	}

	mem := Setup32BitMockMemory(t)
	RunMemoryReadTest(t, mem, any32BitUint, readUint32, "Read uint32")
}

func TestReadUint64(t *testing.T) {
	readUint64 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error) {
		return memory.ReadUint64(mem, byteOrder, index)
	}

	mem := Setup64BitMockMemory(t)
	RunMemoryReadTest(t, mem, any64BitUint, readUint64, "Read uint64")
}

func TestWriteUint8(t *testing.T) {
	// Given
	writeUint8 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) error {
		return memory.WriteUint8(mem, byteOrder, any8BitUint, index)
	}

	mem := Setup8BitMockMemory(t)

	// When - Then
	RunMemoryWriteTest(t, mem, writeUint8, "Write uint8")
}

func TestWriteUint16(t *testing.T) {
	// Given
	writeUint16 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) error {
		return memory.WriteUint16(mem, byteOrder, any16BitUint, index)
	}

	mem := Setup16BitMockMemory(t)

	// When - Then
	RunMemoryWriteTest(t, mem, writeUint16, "Write uint16")
}

func TestWriteFloat32(t *testing.T) {
	// Given
	writeFloat32 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) error {
		return memory.WriteFloat32(mem, byteOrder, any32BitFloat, index)
	}

	mem := Setup32BitMockMemory(t)

	// When - Then
	RunMemoryWriteTest(t, mem, writeFloat32, "Write float32")
}

func TestWriteFloat64(t *testing.T) {
	// Given
	writeFloat64 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) error {
		return memory.WriteFloat64(mem, byteOrder, any64BitFloat, index)
	}

	mem := Setup64BitMockMemory(t)

	// When - Then
	RunMemoryWriteTest(t, mem, writeFloat64, "Write float64")
}

func TestWriteUint32(t *testing.T) {
	// Given
	writeUint32 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) error {
		return memory.WriteUint32(mem, byteOrder, any32BitUint, index)
	}

	mem := Setup32BitMockMemory(t)

	// When - Then
	RunMemoryWriteTest(t, mem, writeUint32, "Write uint32")
}

func TestWriteUint64(t *testing.T) {
	// Given
	writeUint64 := func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) error {
		return memory.WriteUint64(mem, byteOrder, any64BitUint, index)
	}

	mem := Setup64BitMockMemory(t)

	// When - Then
	RunMemoryWriteTest(t, mem, writeUint64, "Write uint64")
}

type TestMemRead func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) (interface{}, error)
type TestMemWrite func(mem memory.Memory, byteOrder binary.ByteOrder, index uint64) error

func RunMemoryReadTest(t *testing.T, goodMem memory.Memory, expected interface{}, test TestMemRead, testing string) {
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

func RunMemoryWriteTest(t *testing.T, goodMem *mock.Memory, test TestMemWrite, testing string) {
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

func Setup8BitMockMemory(t *testing.T) *mock.Memory {
	data := []byte{any8BitUint}
	return &mock.Memory{Data: data, ExpectedIndex: anyIndex, Fail: nil, T: t}
}

func Setup16BitMockMemory(t *testing.T) *mock.Memory {
	data := make([]byte, 2)
	anyEndianness.PutUint16(data, any16BitUint)
	return &mock.Memory{Data: data, ExpectedIndex: anyIndex, Fail: nil, T: t}
}

func Setup32BitMockMemory(t *testing.T) *mock.Memory {
	data := make([]byte, 4)
	anyEndianness.PutUint32(data, any32BitUint)
	return &mock.Memory{Data: data, ExpectedIndex: anyIndex, Fail: nil, T: t}
}

func Setup64BitMockMemory(t *testing.T) *mock.Memory {
	data := make([]byte, 8)
	anyEndianness.PutUint64(data, any64BitUint)
	return &mock.Memory{Data: data, ExpectedIndex: anyIndex, Fail: nil, T: t}
}

func SetupFailMemory(t *testing.T) *mock.Memory {
	data := make([]byte, 1)
	return &mock.Memory{Data: data, ExpectedIndex: anyIndex, Fail: errAny, T: t}
}
