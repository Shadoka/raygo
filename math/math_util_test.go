package math

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestFloatEqualsSuccess(t *testing.T) {
	a := 1.4353445
	b := 1.4353445

	assert.Assert(t, floatEquals(a, b))
}

func TestFloatEqualsFail(t *testing.T) {
	a := 1.4353445
	b := 1.4354445

	assert.Assert(t, !floatEquals(a, b))
}

func TestClampToByte(t *testing.T) {
	n := 255.0
	expected := uint64(255)

	assert.Assert(t, expected == ClampToByte(n))
}

func TestClampToByte2(t *testing.T) {
	n := 0.0
	expected := uint64(0)

	assert.Assert(t, expected == ClampToByte(n))
}

func TestClampToByteTooHigh(t *testing.T) {
	n := 300.0
	expected := uint64(255)

	assert.Assert(t, expected == ClampToByte(n))
}

func TestClampToByteTooLow(t *testing.T) {
	n := -2.0
	expected := uint64(0)

	assert.Assert(t, expected == ClampToByte(n))
}

func TestClampToByteBetween(t *testing.T) {
	n := 133.0
	expected := uint64(133)

	assert.Assert(t, expected == ClampToByte(n))
}
