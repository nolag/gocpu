package memory

// Slice represents memory backed by a []byte.
// Use NewSlice to create one.
type Slice struct {
	data []byte
}

// NewSlice creates a new Slice backed by a []byte with lenght of size
func NewSlice(size uint64) Memory {
	return &Slice{make([]byte, size)}
}

// ReadOneByte reads a byte at memory location index
func (memory *Slice) ReadOneByte(index uint64) (byte, error) {
	if index >= uint64(len(memory.data)) {
		return 0, &AccessViolationError{index, 1, true}
	}

	return memory.data[index], nil
}

// ReadRaw allows reading from memory starting at startIndex and providing numBytes bytes
// data is the bytes read
// backed, when true means that changes made to data will impact the memory stored
// err is any error that occured
func (memory *Slice) ReadRaw(startIndex uint64, numBytes uint64) (data []byte, backed bool, err error) {
	maxIndex := startIndex + numBytes
	if maxIndex > uint64(len(memory.data)) || maxIndex < startIndex {
		return nil, false, &AccessViolationError{startIndex, numBytes, true}
	}

	return memory.data[startIndex:maxIndex], true, nil
}

// Size in bytes this memory can represent
func (memory *Slice) Size() uint64 {
	return uint64(len(memory.data))
}

// WriteOneByte reads a byte at memory location index
func (memory *Slice) WriteOneByte(val byte, index uint64) error {
	if index >= uint64(len(memory.data)) {
		return &AccessViolationError{index, 1, false}
	}

	memory.data[index] = val
	return nil
}

// WriteRaw writes data back to memory
func (memory *Slice) WriteRaw(data []byte, startIndex uint64) error {
	numBytes := uint64(len(data))
	maxIndex := startIndex + numBytes
	if maxIndex > uint64(len(memory.data)) || maxIndex < startIndex {
		return &AccessViolationError{startIndex, numBytes, false}
	}

	copy(memory.data[startIndex:], data)
	return nil
}
