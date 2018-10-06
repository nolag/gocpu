package memhack

import (
	"encoding/binary"
)

type Memory interface{}

func ReadrunnerType(Memory, binary.ByteOrder, uint64) (uint64, error) {
	return 0, nil
}

type RegisterpcType interface {
	IncrementpcType(value interface{})
	ValueAspcType() uint64
}
