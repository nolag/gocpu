package registers_test

import (
	"math"
	"testing"

	"github.com/nolag/gocpu/registers"
	"github.com/stretchr/testify/assert"
)

var anyValue = uint32(0xFEEDBAD1)
var anyValueF = math.Float32frombits(anyValue)
var anyBigValue = uint64(0xF00DBEED12345678)
var anyBigValueF = math.Float64frombits(anyBigValue)

func TestRegisterUint32SetsFloatCorectly(t *testing.T) {
	// Given
	register := registers.RegisterUint32(0)

	// When
	register.SetFromFloat32(anyValueF)

	// Then
	assert.Equal(t, anyValue, uint32(register), "All 32 bit representations must be implementeds")
}

func TestRegisterUint32GetsFloatCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterUint32(anyValue)

	// When
	actual := register.Float32Value()

	// Then
	assert.Equal(t, anyValueF, actual, "All 32 bit representations must be implementeds")
}

func TestRegisterUint64SetsFloatCorectly(t *testing.T) {
	// Given
	register := registers.RegisterUint64(0)

	// When
	register.SetFromFloat64(anyBigValueF)

	// Then
	assert.Equal(t, anyBigValue, uint64(register), "All 64 bit representations must be implementeds")
}

func TestRegisterUint64GetsFloatCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterUint64(anyBigValue)

	// When
	actual := register.Float64Value()

	// Then
	assert.Equal(t, anyBigValueF, actual, "All 64 bit representations must be implementeds")
}

func TestRegisterFloat32SetsUintCorectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat32(0)

	// When
	register.SetFromUint32(anyValue)

	// Then
	assert.Equal(t, anyValueF, float32(register), "All 32 bit representations must be implementeds")
}

func TestRegisterFloat32GetsUintCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat32(anyValueF)

	// When
	actual := register.Uint32Value()

	// Then
	assert.Equal(t, anyValue, actual, "All 32 bit representations must be implementeds")
}

func TestRegisterFloat64SetsUintCorectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat64(0)

	// When
	register.SetFromUint64(anyBigValue)

	// Then
	assert.Equal(t, anyBigValueF, float64(register), "All 64 bit representations must be implementeds")
}

func TestRegisterFloat64GetsUintCorrectly(t *testing.T) {
	// Given
	register := registers.RegisterFloat64(anyBigValueF)

	// When
	actual := register.Uint64Value()

	// Then
	assert.Equal(t, anyBigValue, actual, "All 64 bit representations must be implementeds")
}

/***************************************
Tests below here test generated methods*
***************************************/

// This method compiling proves it works for all types
func TestGeneratedTypesFitGeneratedInterfaces(t *testing.T) {
	// sixteen bit was chosen because 32 and 64 bit have extra and use another interface
	x := func(r registers.SixteenBitRegister) {}
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
	reg.SetFromUint32(anyValue)

	// Then
	assert.Equal(t, uint32(0), reg.ValueAsUint32(), "Zero register must not be settable")
}
