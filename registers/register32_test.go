package registers

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var anyValue = uint32(0xFEEDBAD1)
var anyValueF = math.Float32frombits(anyValue)

func TestUintRegesterCanWrite(t *testing.T) {
	// Given
	reg := UintRegister32(anyValue)

	// When - Then
	assert.True(t, reg.CanWrite(), "Uint registers can write")
}

func TestUintRegisterGetsValue(t *testing.T) {
	// Given
	reg := UintRegister32(anyValue)

	// When - Then
	assertGetValue(t, &reg, anyValue)
}

func TestUintRegisterSetsValue(t *testing.T) {
	// Given
	reg := UintRegister32(0)
	reg2 := UintRegister32(0)

	set := func() { reg.SetValue32(anyValue) }
	set2 := func() { reg2.SetValue32F(anyValueF) }

	// When - Then
	assertSet(t, &reg, set)
	assertSet(t, &reg2, set2)
}

func TestFloatRegesterCanWrite(t *testing.T) {
	// Given
	reg := FloatRegister32(anyValueF)

	// When - Then
	assert.True(t, reg.CanWrite(), "Float registers can write")
}

func TestFloatRegisterGetsValue(t *testing.T) {
	// Given
	reg := FloatRegister32(anyValueF)

	// When - Then
	assertGetValue(t, &reg, anyValue)
}

func TestFloatRegisterSetsValue(t *testing.T) {
	// Given
	reg := FloatRegister32(0)
	reg2 := FloatRegister32(0)

	set := func() { reg.SetValue32(anyValue) }
	set2 := func() { reg2.SetValue32F(anyValueF) }

	// When - Then
	assertSet(t, &reg, set)
	assertSet(t, &reg2, set2)
}

func TestFloatRegesterCanNotWrite(t *testing.T) {
	// Given
	reg := ZeroRegister32{}

	// When - Then
	assert.False(t, reg.CanWrite(), "Zero registers can not write")
}

func TestZeroRegisterGetsZeroValue(t *testing.T) {
	// Given
	reg := ZeroRegister32{}

	// When
	val := reg.Value32()
	valF := reg.Value32F()

	// When - Then
	assert.Equal(t, uint32(0), val, "Zero register must return 0")
	assert.Equal(t, float32(0.0), valF, "Zero register must return 0")
}

func TestZeroRegisterDoesNotSetValue(t *testing.T) {
	// Given
	reg := ZeroRegister32{}

	// When
	reg.SetValue32(anyValue)
	reg.SetValue32F(anyValueF)

	// Then
	assertGetValue(t, &reg, 0)

}

func assertGetValue(t *testing.T, reg Register32, expected uint32) {
	// When
	value := reg.Value32()
	valueF := reg.Value32F()

	// Then
	assert.Equal(t, expected, value, "Wrong value for uint32")
	assert.Equal(t, math.Float32frombits(expected), valueF, "Wrong value for float32")
}

type runSet func()

func assertSet(t *testing.T, reg Register32, setToAnyValue runSet) {
	// When
	setToAnyValue()

	// Then
	assertGetValue(t, reg, anyValue)
}
