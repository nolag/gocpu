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

	set := func() error { return reg.SetValue32(anyValue) }
	set2 := func() error { return reg2.SetValue32F(anyValueF) }

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

	set := func() error { return reg.SetValue32(anyValue) }
	set2 := func() error { return reg2.SetValue32F(anyValueF) }

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
	val, err := reg.Value32()
	valf, errF := reg.Value32F()

	// When - Then
	assertValue(t, val, 0)
	assertValueF(t, valf, 0.0)

	if err != nil || errF != nil {
		t.Fatalf("Supplied registers must not return an error")
	}
}

func TestZeroRegisterDoesNotSetValue(t *testing.T) {
	// Given
	reg := ZeroRegister32{}

	// When
	err := reg.SetValue32(anyValue)
	errF := reg.SetValue32F(anyValueF)

	// Then
	assertGetValue(t, &reg, 0)
	if err != nil || errF != nil {
		t.Fatalf("Supplied registers must not return an error")
	}

}

func assertValue(t *testing.T, actual uint32, expected uint32) {
	if actual != expected {
		t.Fatalf("Expected: 0x%x, Got: 0x%x", expected, actual)
	}
}

func assertValueF(t *testing.T, actual float32, expected float32) {
	if actual != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, actual)
	}
}

func assertGetValue(t *testing.T, reg Register32, expected uint32) {
	// When
	value, err := reg.Value32()
	valueF, errF := reg.Value32F()

	// Then
	assertValue(t, value, expected)
	assertValueF(t, valueF, math.Float32frombits(expected))

	if err != nil || errF != nil {
		t.Fatalf("Supplied registers must not return an error")
	}
}

type runSet func() error

func assertSet(t *testing.T, reg Register32, setToAnyValue runSet) {
	// When
	err := setToAnyValue()

	// Then
	assert.Equal(t, nil, err, "Supplied registers must not return an error")
	assertGetValue(t, reg, anyValue)
}
