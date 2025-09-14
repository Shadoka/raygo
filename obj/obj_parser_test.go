package obj

import (
	"raygo/geometry"
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

var p = math.CreatePoint

func TestParseDataGibberish(t *testing.T) {
	input := `There was a young lady named Bright
who traveled much faster than light.
She set out one day
in a relative way,
and came back the previous night.`
	objData := CreateObjData()
	ParseData(objData, input)

	assert.Assert(t, len(objData.Vertices) == 0)
	assert.Assert(t, objData.IgnoredLines == 5)
}

func TestParseDataVertices(t *testing.T) {
	input := `
v -1 1 0
v -1.0000 0.5000 0.0000
v 1 0 0
v 1 1 0
`
	objData := CreateObjData()
	ParseData(objData, input)

	assert.Assert(t, len(objData.Vertices) == 4)
	assert.Assert(t, objData.IgnoredLines == 2)
	assert.Assert(t, objData.GetV(1).Equals(p(-1.0, 1.0, 0.0)))
	assert.Assert(t, objData.GetV(2).Equals(p(-1.0, 0.5, 0.0)))
	assert.Assert(t, objData.GetV(3).Equals(p(1.0, 0.0, 0.0)))
	assert.Assert(t, objData.GetV(4).Equals(p(1.0, 1.0, 0.0)))
}

func TestParseDataVerticesWithAdditionalBlank(t *testing.T) {
	input := `
 
v  7.0000 0.0000 12.0000
v  -1.0000 0.5000 0.0000
v  1 0 0
v  1 1 0
`
	objData := CreateObjData()
	ParseData(objData, input)

	assert.Assert(t, len(objData.Vertices) == 4)
	assert.Assert(t, objData.IgnoredLines == 3)
	assert.Assert(t, objData.GetV(1).Equals(p(7.0, 0.0, 12.0)))
	assert.Assert(t, objData.GetV(2).Equals(p(-1.0, 0.5, 0.0)))
	assert.Assert(t, objData.GetV(3).Equals(p(1.0, 0.0, 0.0)))
	assert.Assert(t, objData.GetV(4).Equals(p(1.0, 1.0, 0.0)))
}

func TestParseDataFaces(t *testing.T) {
	input := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0

f 1 2 3
f 1 3 4
`
	objData := CreateObjData()
	ParseData(objData, input)

	assert.Assert(t, len(objData.Vertices) == 4)
	assert.Assert(t, len(objData.Faces) == 2)
	assert.Assert(t, objData.IgnoredLines == 3)

	object := objData.ToGroup()

	assert.Assert(t, len(object.Children) == 2)
	t1 := object.Children[0].(*geometry.Triangle)
	t2 := object.Children[1].(*geometry.Triangle)
	assert.Assert(t, t1.P1.Equals(objData.GetV(1)))
	assert.Assert(t, t1.P2.Equals(objData.GetV(2)))
	assert.Assert(t, t1.P3.Equals(objData.GetV(3)))
	assert.Assert(t, t2.P1.Equals(objData.GetV(1)))
	assert.Assert(t, t2.P2.Equals(objData.GetV(3)))
	assert.Assert(t, t2.P3.Equals(objData.GetV(4)))
}

func TestParseDataFacesExtendedFormat(t *testing.T) {
	input := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0

f 1/2/3 2/3/4 3/4/5
f 1/2/3 3/4/5 4/5/6
`
	objData := CreateObjData()
	ParseData(objData, input)

	assert.Assert(t, len(objData.Vertices) == 4)
	assert.Assert(t, len(objData.Faces) == 2)
	assert.Assert(t, objData.IgnoredLines == 3)

	object := objData.ToGroup()

	assert.Assert(t, len(object.Children) == 2)
	t1 := object.Children[0].(*geometry.Triangle)
	t2 := object.Children[1].(*geometry.Triangle)
	assert.Assert(t, t1.P1.Equals(objData.GetV(1)))
	assert.Assert(t, t1.P2.Equals(objData.GetV(2)))
	assert.Assert(t, t1.P3.Equals(objData.GetV(3)))
	assert.Assert(t, t2.P1.Equals(objData.GetV(1)))
	assert.Assert(t, t2.P2.Equals(objData.GetV(3)))
	assert.Assert(t, t2.P3.Equals(objData.GetV(4)))
}

func TestParseDataTriangulation(t *testing.T) {
	input := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0
v 0 2 0

f 1 2 3 4 5
`
	objData := CreateObjData()
	ParseData(objData, input)

	assert.Assert(t, len(objData.Vertices) == 5)
	assert.Assert(t, len(objData.Faces) == 1)
	assert.Assert(t, objData.IgnoredLines == 3)

	object := objData.ToGroup()

	assert.Assert(t, len(object.Children) == 3)
	t1 := object.Children[0].(*geometry.Triangle)
	t2 := object.Children[1].(*geometry.Triangle)
	t3 := object.Children[2].(*geometry.Triangle)
	assert.Assert(t, t1.P1.Equals(objData.GetV(1)))
	assert.Assert(t, t1.P2.Equals(objData.GetV(2)))
	assert.Assert(t, t1.P3.Equals(objData.GetV(3)))
	assert.Assert(t, t2.P1.Equals(objData.GetV(1)))
	assert.Assert(t, t2.P2.Equals(objData.GetV(3)))
	assert.Assert(t, t2.P3.Equals(objData.GetV(4)))
	assert.Assert(t, t3.P1.Equals(objData.GetV(1)))
	assert.Assert(t, t3.P2.Equals(objData.GetV(4)))
	assert.Assert(t, t3.P3.Equals(objData.GetV(5)))
}

func TestParseDataGroups(t *testing.T) {
	input := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0

f 1 2 4
g First group
f 1 2 3
g Second group
f 1 3 4
`
	objData := CreateObjData()
	ParseData(objData, input)

	assert.Assert(t, len(objData.Vertices) == 4)
	assert.Assert(t, len(objData.Faces) == 1)
	assert.Assert(t, len(objData.Groups) == 2)
	assert.Assert(t, objData.IgnoredLines == 3)

	object := objData.ToGroup()

	assert.Assert(t, len(object.Children) == 3)
	t1 := object.Children[0].(*geometry.Triangle)
	g1 := object.Children[1].(*geometry.Group)
	g2 := object.Children[2].(*geometry.Group)
	assert.Assert(t, t1.P1.Equals(objData.GetV(1)))
	assert.Assert(t, t1.P2.Equals(objData.GetV(2)))
	assert.Assert(t, t1.P3.Equals(objData.GetV(4)))
	t2 := g1.Children[0].(*geometry.Triangle)
	assert.Assert(t, t2.P1.Equals(objData.GetV(1)))
	assert.Assert(t, t2.P2.Equals(objData.GetV(2)))
	assert.Assert(t, t2.P3.Equals(objData.GetV(3)))
	t3 := g2.Children[0].(*geometry.Triangle)
	assert.Assert(t, t3.P1.Equals(objData.GetV(1)))
	assert.Assert(t, t3.P2.Equals(objData.GetV(3)))
	assert.Assert(t, t3.P3.Equals(objData.GetV(4)))
}
