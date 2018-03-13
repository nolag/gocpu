// Package mock provides mocks to use for interfaces in gocpu sub packages
package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Memory provides a mock of memory,
// Raw reads return Data and RawWrites set data
// A single byte read/write goes to data[0].
// If fail is not nil, it will be returned instead of taking the action above
// Before any other action, an assertion is made on the index being equal to ExpectedIndex
type Memory struct {
	Data          []byte
	ExpectedIndex uint64
	Fail          error
	T             *testing.T
}

// ReadOneByte reads a byte at memory location index
func (memory *Memory) ReadOneByte(index uint64) (byte, error) {
	assert.Equal(memory.T, memory.ExpectedIndex, index, "Wrong index in read one byte")
	if memory.Fail != nil {
		return 0, memory.Fail
	}

	return memory.Data[0], nil
}

// ReadRaw allows reading from memory starting at startIndex and providing numBytes bytes
// data is the bytes read
// backed, when true means that changes made to data will impact the memory stored
// err is any error that occured
func (memory *Memory) ReadRaw(startIndex uint64, numBytes uint64) (data []byte, backed bool, err error) {
	assert.Equal(memory.T, memory.ExpectedIndex, startIndex, "Wrong index used")
	if memory.Fail != nil {
		return nil, false, memory.Fail
	}

	return memory.Data, false, nil
}

// Size in bytes this memory can represent
func (memory *Memory) Size() uint64 {
	return uint64(len(memory.Data))
}

// WriteOneByte reads a byte at memory location index
func (memory *Memory) WriteOneByte(val byte, index uint64) error {
	assert.Equal(memory.T, memory.ExpectedIndex, index, "Wrong index used")
	if memory.Fail != nil {
		return memory.Fail
	}

	memory.Data = []byte{val}
	return nil
}

// WriteRaw writes data back to memory
func (memory *Memory) WriteRaw(data []byte, startIndex uint64) error {
	assert.Equal(memory.T, memory.ExpectedIndex, startIndex, "Wrong index used")
	if memory.Fail != nil {
		return memory.Fail
	}

	memory.Data = data
	return nil
}
