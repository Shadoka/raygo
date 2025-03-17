package ray

import (
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestDefaultMaterial(t *testing.T) {
	m := DefaultMaterial()
	expectedColor := math.CreateColor(1.0, 1.0, 1.0)

	assert.Assert(t, expectedColor.Equals(m.Color))
	assert.Assert(t, m.Ambient == 0.1)
	assert.Assert(t, m.Diffuse == 0.9)
	assert.Assert(t, m.Specular == 0.9)
	assert.Assert(t, m.Shininess == 200.0)
}
