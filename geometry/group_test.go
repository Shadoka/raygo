package geometry

import (
	gomath "math"
	"raygo/math"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateEmptyGroup(t *testing.T) {
	g := EmptyGroup()

	assert.Assert(t, g.Size() == 0)
}

func TestCreateTreeWithEntries(t *testing.T) {
	g := EmptyGroup()
	root := CreateCube()
	s1 := CreateSphere()
	g2 := EmptyGroup()
	s2 := CreateGlassSphere()
	p := CreatePlane()

	g2.AddChild(s2).AddChild(p)
	g.AddChild(s1).AddChild(root).AddChild(g2)

	assert.Assert(t, g.Size() == 4)
	assert.Assert(t, g2.Size() == 2)
}

func TestNavigateTree(t *testing.T) {
	g := EmptyGroup()
	c := CreateCube()
	s1 := CreateSphere()
	g2 := EmptyGroup()
	s2 := CreateGlassSphere()
	p := CreatePlane()

	g2.AddChild(s2).AddChild(p)
	g.AddChild(s1).AddChild(c).AddChild(g2)

	selectedGroup := g.Children[2].(*Group)
	element := selectedGroup.Children[0]

	assert.Assert(t, element.Equals(s2))
	assert.Assert(t, g.Equals(p.Parent.Parent))
}

func TestLocalIntersectEmptyGroup(t *testing.T) {
	g := EmptyGroup()
	r := CreateRay(math.CreatePoint(0.0, 0.0, 0.0), math.CreateVector(0.0, 0.0, 1.0))

	actual := g.localIntersect(r)

	assert.Assert(t, len(actual) == 0)
}

func TestLocalIntersectNonEmptyGroup(t *testing.T) {
	g := EmptyGroup()
	s1 := CreateSphere()
	s2 := CreateSphere()
	s3 := CreateSphere()
	s2.SetTransform(math.Translation(0.0, 0.0, -3.0))
	s3.SetTransform(math.Translation(5.0, 0.0, 0.0))
	g.AddChild(s1).AddChild(s2).AddChild(s3)

	r := CreateRay(math.CreatePoint(0.0, 0.0, -5.0), math.CreateVector(0.0, 0.0, 1.0))

	actual := g.localIntersect(r)

	assert.Assert(t, len(actual) == 4)
	assert.Assert(t, actual[0].Object.Equals(s2))
	assert.Assert(t, actual[1].Object.Equals(s2))
	assert.Assert(t, actual[2].Object.Equals(s1))
	assert.Assert(t, actual[3].Object.Equals(s1))
}

func TestNonEmptyGroupBounds(t *testing.T) {
	g := EmptyGroup()
	s1 := CreateSphere()
	g.SetTransform(math.Scaling(2.0, 2.0, 2.0))
	s1.SetTransform(math.Translation(5.0, 0.0, 0.0))
	g.AddChild(s1)
	expected := Bounds{
		Minimum: math.CreatePoint(4.0, -1.0, -1.0),
		Maximum: math.CreatePoint(6.0, 1.0, 1.0),
	}

	b := g.Bounds()

	assert.Assert(t, expected.Equals(b))
}

func TestIntersectGroupWithTransformations(t *testing.T) {
	g := EmptyGroup()
	s1 := CreateSphere()
	g.SetTransform(math.Scaling(2.0, 2.0, 2.0))
	s1.SetTransform(math.Translation(5.0, 0.0, 0.0))
	g.AddChild(s1)

	r := CreateRay(math.CreatePoint(10.0, 0.0, -10.0), math.CreateVector(0.0, 0.0, 1.0))

	actual := g.Intersect(r)

	assert.Assert(t, len(actual) == 2)
}

func TestWorldToObject(t *testing.T) {
	g1 := EmptyGroup()
	g1.SetTransform(math.Rotation_Y(gomath.Pi / 2.0))
	g2 := EmptyGroup()
	g2.SetTransform(math.Scaling(2.0, 2.0, 2.0))
	g1.AddChild(g2)
	s := CreateSphere()
	s.SetTransform(math.Translation(5.0, 0.0, 0.0))
	g2.AddChild(s)
	expected := math.CreatePoint(0.0, 0.0, -1.0)

	actual := WorldToObject(s, math.CreatePoint(-2.0, 0.0, -10.0))

	assert.Assert(t, expected.Equals(actual))
}

func TestNormalToWorld(t *testing.T) {
	g1 := EmptyGroup()
	g1.SetTransform(math.Rotation_Y(gomath.Pi / 2.0))
	g2 := EmptyGroup()
	g2.SetTransform(math.Scaling(1.0, 2.0, 3.0))
	g1.AddChild(g2)
	s := CreateSphere()
	s.SetTransform(math.Translation(5.0, 0.0, 0.0))
	g2.AddChild(s)
	expected := math.CreateVector(0.28571, 0.42857, -0.85714)

	actual := NormalToWorld(s, math.CreateVector(gomath.Sqrt(3)/3.0, gomath.Sqrt(3)/3.0, gomath.Sqrt(3)/3.0))

	assert.Assert(t, expected.Equals(actual))
}

func TestGroupNormalAt(t *testing.T) {
	g1 := EmptyGroup()
	g1.SetTransform(math.Rotation_Y(gomath.Pi / 2.0))
	g2 := EmptyGroup()
	g2.SetTransform(math.Scaling(1.0, 2.0, 3.0))
	g1.AddChild(g2)
	s := CreateSphere()
	s.SetTransform(math.Translation(5.0, 0.0, 0.0))
	g2.AddChild(s)
	expected := math.CreateVector(0.28570, 0.42854, -0.85716)

	actual := s.NormalAt(math.CreatePoint(1.7321, 1.1547, -5.5774))

	assert.Assert(t, expected.Equals(actual))
}
