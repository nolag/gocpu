package gocpu

import (
	"encoding/binary"
	"testing"

	"github.com/nolag/gocpu/registers"

	"github.com/nolag/gocpu/memory"
	"github.com/stretchr/testify/assert"
)

var anyEndianness = binary.BigEndian
var anyPc = uint32(0xABADBABE)
var anyMemory = memory.NewSlice(100)
var anyAmount = uint32(102)
var anyReg = make([]registers.Register32, 5)

func TestSimpleProcessorCore32PcValues(t *testing.T) {
	// Given
	basics := SimpleProcessorCore32{anyEndianness, anyMemory, anyReg, anyPc}
	anyOtherPc := uint32(0xBEEFF00D)

	// When
	actualPc, err1 := basics.Pc32()
	err2 := basics.IncementPc32(anyAmount)
	pcInc := basics.Pc
	basics.Pc = anyPc
	err3 := basics.DecementPc32(anyAmount)
	pcDec := basics.Pc
	err4 := basics.SetPc(anyOtherPc)

	// Then
	assert.NoError(t, err1, "No error must be returned to get pc")
	assert.NoError(t, err2, "No error must be returned when incrementing pc")
	assert.NoError(t, err3, "No error must be returned when decrementing")
	assert.NoError(t, err4, "No error must be returned when setting pc")
	assert.Equal(t, actualPc, anyPc, "Pc not gotten from correct source")
	assert.Equal(t, actualPc+anyAmount, pcInc, "Pc not incremented after increment call")
	assert.Equal(t, actualPc-anyAmount, pcDec, "Pc not decremented after decrement call")
	assert.Equal(t, anyOtherPc, basics.Pc, "Pc not setting")
}

func TestSimpleProcessorCore32SimpleGetSetValues(t *testing.T) {
	// Given
	basics := SimpleProcessorCore32{anyEndianness, anyMemory, anyReg, anyPc}
	otherEndianness := binary.LittleEndian
	otherMem := memory.NewSlice(1)
	otherReg := make([]registers.Register32, 1)

	// When
	actualEndianness := basics.Endianness()
	actualMem := basics.Memory()
	acutalReg := basics.Registers32()
	basics.SetEndianness(otherEndianness)
	basics.SetMemory(otherMem)
	basics.SetRegisters(otherReg)

	// Then
	assert.Equal(t, anyEndianness, actualEndianness, "Wrong enianness returned")
	assert.Equal(t, anyMemory, actualMem, "Wrong memory returned")
	assert.Equal(t, anyReg, acutalReg, "Wrong registers returned")
	assert.Equal(t, otherEndianness, basics.BackingEndianness, "Endianness not set")
	assert.Equal(t, otherMem, basics.BackingMemory, "Memory not set")
	assert.Equal(t, otherReg, basics.BackingRegisters, "Registers not set")
}
