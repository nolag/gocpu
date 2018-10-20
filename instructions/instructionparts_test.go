package instructions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test a random int/Uint for each base

func TestInt8Base(t *testing.T) {
	minVal := Int4(-8)
	maxVal := Int4(7)
	otherGoodVal := Int4(3)
	minTooHigh := maxVal + Int4(1)
	maxTooLow := minVal - Int4(1)
	anotherTooHigh := minTooHigh + Int4(35)
	anotherTooLow := maxTooLow - Int4(12)

	assert.True(t, minVal.Verify(), "The minimum value must verify")
	assert.True(t, maxVal.Verify(), "The maximum value must verify")
	assert.True(t, otherGoodVal.Verify(), "Other good values must verify")
	assert.False(t, minTooHigh.Verify(), "Value that is 1 too high must not verify")
	assert.False(t, maxTooLow.Verify(), "Value that is 1 too low must not verify")
	assert.False(t, anotherTooHigh.Verify(), "Value that is too high must not verify")
	assert.False(t, anotherTooLow.Verify(), "Value that is too low must not verify")
}

func TestUint8Base(t *testing.T) {
	maxVal := Uint4(15)
	otherGoodVal := Uint4(3)
	minTooHigh := maxVal + Uint4(1)
	anotherTooHigh := maxVal + Uint4(12)

	assert.True(t, maxVal.Verify(), "The maximum value must verify")
	assert.True(t, otherGoodVal.Verify(), "Other good values must verify")
	assert.False(t, minTooHigh.Verify(), "Value that is 1 too high must not verify")
	assert.False(t, anotherTooHigh.Verify(), "Value that is too high must not verify")
}

func TestInt16Base(t *testing.T) {
	minVal := Int12(-2048)
	maxVal := Int12(2047)
	otherGoodVal := Int12(3)
	minTooHigh := maxVal + Int12(1)
	maxTooLow := minVal - Int12(1)
	anotherTooHigh := minTooHigh + Int12(35)
	anotherTooLow := maxTooLow - Int12(12)

	assert.True(t, minVal.Verify(), "The minimum value must verify")
	assert.True(t, maxVal.Verify(), "The maximum value must verify")
	assert.True(t, otherGoodVal.Verify(), "Other good values must verify")
	assert.False(t, minTooHigh.Verify(), "Value that is 1 too high must not verify")
	assert.False(t, maxTooLow.Verify(), "Value that is 1 too low must not verify")
	assert.False(t, anotherTooHigh.Verify(), "Value that is too high must not verify")
	assert.False(t, anotherTooLow.Verify(), "Value that is too low must not verify")
}

func TestUint16Base(t *testing.T) {
	maxVal := Uint12(4095)
	otherGoodVal := Uint12(3)
	minTooHigh := maxVal + Uint12(1)
	anotherTooHigh := maxVal + Uint12(12)

	assert.True(t, maxVal.Verify(), "The maximum value must verify")
	assert.True(t, otherGoodVal.Verify(), "Other good values must verify")
	assert.False(t, minTooHigh.Verify(), "Value that is 1 too high must not verify")
	assert.False(t, anotherTooHigh.Verify(), "Value that is too high must not verify")
}

func TestInt32Base(t *testing.T) {
	minVal := Int23(-4194304)
	maxVal := Int23(4194303)
	otherGoodVal := Int23(3)
	minTooHigh := maxVal + Int23(1)
	maxTooLow := minVal - Int23(1)
	anotherTooHigh := minTooHigh + Int23(35)
	anotherTooLow := maxTooLow - Int23(12)

	assert.True(t, minVal.Verify(), "The minimum value must verify")
	assert.True(t, maxVal.Verify(), "The maximum value must verify")
	assert.True(t, otherGoodVal.Verify(), "Other good values must verify")
	assert.False(t, minTooHigh.Verify(), "Value that is 1 too high must not verify")
	assert.False(t, maxTooLow.Verify(), "Value that is 1 too low must not verify")
	assert.False(t, anotherTooHigh.Verify(), "Value that is too high must not verify")
	assert.False(t, anotherTooLow.Verify(), "Value that is too low must not verify")
}

func TestUint32Base(t *testing.T) {
	maxVal := Uint23(8388607)
	otherGoodVal := Uint23(3)
	minTooHigh := maxVal + Uint23(1)
	anotherTooHigh := maxVal + Uint23(12)

	assert.True(t, maxVal.Verify(), "The maximum value must verify")
	assert.True(t, otherGoodVal.Verify(), "Other good values must verify")
	assert.False(t, minTooHigh.Verify(), "Value that is 1 too high must not verify")
	assert.False(t, anotherTooHigh.Verify(), "Value that is too high must not verify")
}

func TestInt64Base(t *testing.T) {
	minVal := Int45(-17592186044416)
	maxVal := Int45(17592186044415)
	otherGoodVal := Int45(3)
	minTooHigh := maxVal + Int45(1)
	maxTooLow := minVal - Int45(1)
	anotherTooHigh := minTooHigh + Int45(35)
	anotherTooLow := maxTooLow - Int45(12)

	assert.True(t, minVal.Verify(), "The minimum value must verify")
	assert.True(t, maxVal.Verify(), "The maximum value must verify")
	assert.True(t, otherGoodVal.Verify(), "Other good values must verify")
	assert.False(t, minTooHigh.Verify(), "Value that is 1 too high must not verify")
	assert.False(t, maxTooLow.Verify(), "Value that is 1 too low must not verify")
	assert.False(t, anotherTooHigh.Verify(), "Value that is too high must not verify")
	assert.False(t, anotherTooLow.Verify(), "Value that is too low must not verify")
}

func TestUint64Base(t *testing.T) {
	maxVal := Uint45(35184372088831)
	otherGoodVal := Uint45(3)
	minTooHigh := maxVal + Uint45(1)
	anotherTooHigh := maxVal + Uint45(12)

	assert.True(t, maxVal.Verify(), "The maximum value must verify")
	assert.True(t, otherGoodVal.Verify(), "Other good values must verify")
	assert.False(t, minTooHigh.Verify(), "Value that is 1 too high must not verify")
	assert.False(t, anotherTooHigh.Verify(), "Value that is too high must not verify")
}
