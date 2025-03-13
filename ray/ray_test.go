package ray

import (
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateRay(t *testing.T) {
	o := math.CreatePoint(1.0, 2.0, 3.0)
	d := math.CreateVector(4.0, 5.0, 6.0)
	ray := CreateRay(o, d)

	assert.Assert(t, o.Equals(ray.Origin))
	assert.Assert(t, d.Equals(ray.Direction))

	o = math.CreatePoint(2.0, 3.0, 4.0)
	assert.Assert(t, !o.Equals(ray.Origin))
}

func TestPosition(t *testing.T) {
	r := CreateRay(math.CreatePoint(2.0, 3.0, 4.0), math.CreateVector(1.0, 0.0, 0.0))
	expected1 := math.CreatePoint(2.0, 3.0, 4.0)
	expected2 := math.CreatePoint(3.0, 3.0, 4.0)
	expected3 := math.CreatePoint(1.0, 3.0, 4.0)
	expected4 := math.CreatePoint(4.5, 3.0, 4.0)

	assert.Assert(t, expected1.Equals(r.Position(0.0)))
	assert.Assert(t, expected2.Equals(r.Position(1.0)))
	assert.Assert(t, expected3.Equals(r.Position(-1.0)))
	assert.Assert(t, expected4.Equals(r.Position(2.5)))
}
