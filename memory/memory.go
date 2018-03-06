// Package memory provides classes to simulate Random Access Memory and Memory mapped I/O
package memory

import (
	"encoding/binary"
	"fmt"
	"math"
)

// Memory reprents Random Access Memory or Memory mapped I/O
type Memory interface {
	// ReadOneByte reads a byte at memory location index
	ReadOneByte(index uint64) (byte, error)

	// ReadRaw allows reading from memory starting at startIndex and providing numBytes bytes
	// data is the bytes read
	// backed, when true means that changes made to data will impact the memory stored
	// err is any error that occured
	ReadRaw(startIndex uint64, numBytes uint64) (data []byte, backed bool, err error)

	// Size in bytes this memory can represent
	Size() uint64

	// WriteOneByte reads a byte at memory location index
	WriteOneByte(val byte, index uint64) error

	// WriteRaw writes data back to memory
	WriteRaw(data []byte, startIndex uint64) error
}

// AccessViolationError is returned by Memory when an access violation occurs
type AccessViolationError struct {
	Location uint64
	NumBytes uint64
	WasRead  bool
}

func (err *AccessViolationError) Error() string {
	wasRead := " Was read: "
	if err.WasRead {
		wasRead += "True"
	} else {
		wasRead += "False"
	}

	return fmt.Sprint("Cannot access memory location: ", err.Location, " Num bytes: ", err.NumBytes, wasRead)
}

// ReadFloat32 reads a float32 at memory location index
func ReadFloat32(memory Memory, byteOrder binary.ByteOrder, index uint64) (data float32, err error) {
	val, err := ReadUint32(memory, byteOrder, index)
	return math.Float32frombits(val), err
}

// ReadFloat64 reads a float64 at memory location index
func ReadFloat64(memory Memory, byteOrder binary.ByteOrder, index uint64) (data float64, err error) {
	val, err := ReadUint64(memory, byteOrder, index)
	return math.Float64frombits(val), err
}

// ReadUint16 reads a uint32 at memory location index
func ReadUint16(memory Memory, byteOrder binary.ByteOrder, index uint64) (data uint16, err error) {
	raw, _, err := memory.ReadRaw(index, 2)
	if err != nil {
		return 0, err
	}

	return byteOrder.Uint16(raw), nil
}

// ReadUint32 reads a uint32 at memory location index
func ReadUint32(memory Memory, byteOrder binary.ByteOrder, index uint64) (data uint32, err error) {
	raw, _, err := memory.ReadRaw(index, 4)
	if err != nil {
		return 0, err
	}

	return byteOrder.Uint32(raw), nil
}

// ReadUint64 reads a uint64 at memory location index
func ReadUint64(memory Memory, byteOrder binary.ByteOrder, index uint64) (data uint64, err error) {
	raw, _, err := memory.ReadRaw(index, 8)
	if err != nil {
		return 0, err
	}

	return byteOrder.Uint64(raw), nil
}

// WriteFloat32 reads a float32 at memory location index
func WriteFloat32(memory Memory, byteOrder binary.ByteOrder, val float32, index uint64) error {
	return WriteUint32(memory, byteOrder, math.Float32bits(val), index)
}

// WriteFloat64 reads a float64 at memory location index
func WriteFloat64(memory Memory, byteOrder binary.ByteOrder, val float64, index uint64) error {
	return WriteUint64(memory, byteOrder, math.Float64bits(val), index)
}

// WriteUint16 reads a uint32 at memory location index
func WriteUint16(memory Memory, byteOrder binary.ByteOrder, val uint16, index uint64) error {
	data := make([]byte, 2)
	byteOrder.PutUint16(data, val)
	return memory.WriteRaw(data, index)
}

// WriteUint32 reads a uint32 at memory location index
func WriteUint32(memory Memory, byteOrder binary.ByteOrder, val uint32, index uint64) error {
	data := make([]byte, 4)
	byteOrder.PutUint32(data, val)
	return memory.WriteRaw(data, index)
}

// WriteUint64 reads a uint64 at memory location index
func WriteUint64(memory Memory, byteOrder binary.ByteOrder, val uint64, index uint64) error {
	data := make([]byte, 8)
	byteOrder.PutUint64(data, val)
	return memory.WriteRaw(data, index)
}
