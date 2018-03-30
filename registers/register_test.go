package registers_test

import (
	"math"
	"testing"

	"github.com/nolag/gocpu/registers"
	"github.com/stretchr/testify/assert"
)

var anyValue = uint32(103)
var anyValueF = math.Float32frombits(anyValue)
var anyBigValue = uint64(1234)
var anyBigValueF = math.Float64frombits(anyBigValue)
var anotherValue = uint32(10)
var anotherValueF = float32(1000030.3)
var anotherBigValue = uint64(110)
var anotherBigValueF = float64(132423141.323)

func TestRegisterUint32SetsFloatCorectly(t *testing.T) {
	// Given
	register := registers.RegisterUint32(0)

	// When
	register.SetFromFloat32(anyValueF)

	// Then
	assert.Equal(t, anyValue, uint32(register), "Float not set correctly")
}

func TestRegisterUint32GetsFloatCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterUint32(anyValue)

	// When
	actual := register.ValueAsFloat32()

	// Then
	assert.Equal(t, anyValueF, actual, "Float not gotten correctly")
}

func TestRegisterUint32IncrementsFloatCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterUint32(anyValue)

	// When
	register.InrementFloat32(anotherValueF)
	actual := math.Float32frombits(register.ValueAsUint32())

	// Then
	assert.Equal(t, anyValueF+anotherValueF, actual, "Float not incremented correctly")
}

func TestRegisterUint32DecrementsFloatCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterUint32(anyValue)

	// When
	register.DecrementFloat32(anotherValueF)
	actual := math.Float32frombits(register.ValueAsUint32())

	// Then
	assert.Equal(t, anyValueF-anotherValueF, actual, "Float not incremented correctly")
}

func TestRegisterUint64SetsFloatCorectly(t *testing.T) {
	// Given
	register := registers.RegisterUint64(0)

	// When
	register.SetFromFloat64(anyBigValueF)

	// Then
	assert.Equal(t, anyBigValue, uint64(register), "Float not set correctly")
}

func TestRegisterUint64GetsFloatCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterUint64(anyBigValue)

	// When
	actual := register.ValueAsFloat64()

	// Then
	assert.Equal(t, anyBigValueF, actual, "Float gotten set correctly")
}

func TestRegisterUint64IncrementsFloatCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterUint64(anyBigValue)

	// When
	register.InrementFloat64(anotherBigValueF)
	actual := math.Float64frombits(register.ValueAsUint64())

	// Then
	assert.Equal(t, anyBigValueF+anotherBigValueF, actual, "Float not incremented correctly")
}

func TestRegisterUint64DecrementsFloatCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterUint64(anyBigValue)

	// When
	register.DecrementFloat64(anotherBigValueF)
	actual := math.Float64frombits(register.ValueAsUint64())

	// Then
	assert.Equal(t, anyBigValueF-anotherBigValueF, actual, "Float not incremented correctly")
}

func TestRegisterFloat32SetsUintCorectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat32(0)

	// When
	register.SetFromUint32(anyValue)

	// Then
	assert.Equal(t, anyValueF, float32(register), "Uint not set correctly")
}

func TestRegisterFloat32GetsUintCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat32(anyValueF)

	// When
	actual := register.ValueAsUint32()

	// Then
	assert.Equal(t, anyValue, actual, "Uint not gotten correctly")
}

func TestRegisterFloat32IncrementsUintCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat32(anyValueF)

	// When
	register.InrementUint32(anotherValue)
	actual := math.Float32bits(register.ValueAsFloat32())

	// Then
	assert.Equal(t, anyValue+anotherValue, actual, "Uint not incremented correctly")
}

func TestRegisterFloat32DecrementsUintCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat32(anyValueF)

	// When
	register.DecrementUint32(anotherValue)

	actual := math.Float32bits(register.ValueAsFloat32())

	// Then
	assert.Equal(t, anyValue-anotherValue, actual, "Uint not incremented correctly")
}

func TestRegisterFloat64SetsUintCorectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat64(0)

	// When
	register.SetFromUint64(anyBigValue)

	// Then
	assert.Equal(t, anyBigValueF, float64(register), "Uint not set correctly")
}

func TestRegisterFloat64GetsUintCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat64(anyBigValueF)

	// When
	actual := register.ValueAsUint64()

	// Then
	assert.Equal(t, anyBigValue, actual, "Uint not gotten correctly")
}

func TestRegisterFloat64IncrementsUintCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat64(anyBigValueF)

	// When
	register.InrementUint64(anotherBigValue)
	actual := math.Float64bits(register.ValueAsFloat64())

	// Then
	assert.Equal(t, anyBigValue+anotherBigValue, actual, "Uint not incremented correctly")
}

func TestRegisterFloat64DecrementsUintCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat64(anyBigValueF)

	// When
	register.DecrementUint64(anotherBigValue)
	actual := math.Float64bits(register.ValueAsFloat64())

	// Then
	assert.Equal(t, anyBigValue-anotherBigValue, actual, "Uint not incremented correctly")
}

func Test32BitTypesFitsInterface(t *testing.T) {
	// sixteen bit was chosen because 32 and 64 bit have extra and use another interface
	x := func(r registers.IRegister32) {}
	uregister := registers.RegisterUint32(0)
	fregister := registers.RegisterFloat32(0)
	x(&uregister)
	x(&fregister)
}

func Test64BitTypesFitsInterface(t *testing.T) {
	// sixteen bit was chosen because 64 and 64 bit have extra and use another interface
	x := func(r registers.IRegister64) {}
	uregister := registers.RegisterUint64(0)
	fregister := registers.RegisterFloat64(0)
	x(&uregister)
	x(&fregister)
}

/***************************************
Tests below here test generated methods*
***************************************/

// This method compiling proves it works for all types
func TestGeneratedTypesFitGeneratedInterfaces(t *testing.T) {
	// sixteen bit was chosen because 32 and 64 bit have extra and use another interface
	x := func(r registers.IRegister16) {}
	register := registers.RegisterUint16(0)
	x(&register)
}

func TestGeneratedRegisterGetsValue(t *testing.T) {
	// Given
	reg := registers.RegisterUint32(anyValue)

	// When
	val := reg.ValueAsUint32()

	// Then
	assert.Equal(t, anyValue, val, "Wrong value returned from get")
}

func TestGeneratedRegisterSetsValue(t *testing.T) {
	// Given
	reg := registers.RegisterUint32(0)

	// When
	reg.SetFromUint32(anyValue)

	// Then
	assert.Equal(t, anyValue, uint32(reg), "Wrong value returned after setting")
}

func TestGeneratedRegisterIncrementsValue(t *testing.T) {
	// Given
	reg := registers.RegisterUint32(anyValue)

	// When
	reg.InrementUint32(anotherValue)
	val := reg.ValueAsUint32()

	// Then
	assert.Equal(t, anyValue+anotherValue, val, "Wrong value returned after increment")
}

func TestGeneratedRegisterIncrementsAsPc(t *testing.T) {
	// Given
	reg := registers.RegisterUint32(anyValue)

	// When
	reg.InrementAsPc(byte(anotherValue))
	val := reg.ValueAsUint32()

	// Then
	assert.Equal(t, anyValue+anotherValue, val, "Wrong value returned from read as pc")
}

func TestGeneratedRegisterDecrementsValue(t *testing.T) {
	// Given
	reg := registers.RegisterUint32(anyValue)

	// When
	reg.DecrementUint32(anotherValue)
	val := reg.ValueAsUint32()

	// Then
	assert.Equal(t, anyValue-anotherValue, val, "Wrong value returned after decrement")
}

func TestGeneratedRegisterReadsAsPc(t *testing.T) {
	// Given
	reg := registers.RegisterUint32(anyValue)

	// When
	val := reg.ReadAsPc()

	// Then
	assert.Equal(t, uint64(anyValue), val, "Wrong value returned from read as pc")
}

func TestZeroRegisterGetsZeroValue(t *testing.T) {
	// Given
	reg := registers.ZeroRegister{}

	// When
	val := reg.ValueAsUint32()

	// When - Then
	assert.Equal(t, uint32(0), val, "Zero register must return 0")
}

func TestZeroRegisterDoesNotSetValue(t *testing.T) {
	// Given
	reg := registers.ZeroRegister{}

	// When
	reg.SetFromUint32(anyValue)

	// Then
	assert.Equal(t, uint32(0), reg.ValueAsUint32(), "Zero register must not be settable")
}

func TestZeroRegisterDoesNotIncrement(t *testing.T) {
	// Given
	reg := registers.ZeroRegister{}

	// When
	reg.InrementUint32(1)

	// Then
	assert.Equal(t, uint32(0), reg.ValueAsUint32(), "Zero register must not be incrementable")
}

func TestZeroRegisterDoesNotDecrement(t *testing.T) {
	// Given
	reg := registers.ZeroRegister{}

	// When
	reg.DecrementUint32(1)

	// Then
	assert.Equal(t, uint32(0), reg.ValueAsUint32(), "Zero register must not be incrementable")
}
