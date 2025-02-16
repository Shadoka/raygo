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
