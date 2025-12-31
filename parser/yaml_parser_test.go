package parser

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestParseColors(t *testing.T) {
	yml := `
colors:
  - name: light_gray
    r: 229
    g: 229
    b: 229
  - name: black
    r: 0
    g: 0
    b: 0`

	desc := ParseYaml(yml)

	assert.Assert(t, desc != nil)
	assert.Assert(t, len(desc.Colors) == 2)
	assert.Assert(t, desc.Colors[0].Name == "light_gray")
	assert.Assert(t, desc.Colors[0].R == 229)
	assert.Assert(t, desc.Colors[0].G == 229)
	assert.Assert(t, desc.Colors[0].B == 229)
	assert.Assert(t, desc.Colors[1].Name == "black")
	assert.Assert(t, desc.Colors[1].R == 0)
	assert.Assert(t, desc.Colors[1].G == 0)
	assert.Assert(t, desc.Colors[1].B == 0)
}

func TestParseCheckerPattern(t *testing.T) {
	yml := `
patterns:
  checker:
    - name: gray_black_pattern
      colorA: light_gray
      colorB: black`

	desc := ParseYaml(yml)

	assert.Assert(t, desc != nil)
	assert.Assert(t, len(desc.Patterns.Checker) == 1)
	assert.Assert(t, len(desc.Colors) == 0)
	assert.Assert(t, desc.Patterns.Checker[0].Name == "gray_black_pattern")
	assert.Assert(t, desc.Patterns.Checker[0].ColorA == "light_gray")
	assert.Assert(t, desc.Patterns.Checker[0].ColorB == "black")
}

func TestParseCheckerPatternWithTransforms(t *testing.T) {
	yml := `
patterns:
  checker:
    - name: gray_black_pattern
      colorA: light_gray
      colorB: black
      transforms:
        - type: rotation
          x: 90.0
        - type: translation
          x: 100
          y: 23.4`

	desc := ParseYaml(yml)

	assert.Assert(t, desc != nil)
	assert.Assert(t, len(desc.Patterns.Checker) == 1)
	assert.Assert(t, len(desc.Colors) == 0)
	assert.Assert(t, len(desc.Patterns.Checker[0].Transforms) == 2)
	assert.Assert(t, desc.Patterns.Checker[0].Name == "gray_black_pattern")
	assert.Assert(t, desc.Patterns.Checker[0].ColorA == "light_gray")
	assert.Assert(t, desc.Patterns.Checker[0].ColorB == "black")
	assert.Assert(t, desc.Patterns.Checker[0].Transforms[0].Type == "rotation")
	assert.Assert(t, desc.Patterns.Checker[0].Transforms[1].X == 100.0)
}

func TestParseMaterial(t *testing.T) {
	yml := `
materials:
  - name: black_mat
    color: black
    #ambient: 0.1
    diffuse: 0.2
    specular: 0.3
    shininess: 0.4
    reflective: 0.5
    transparency: 0.6
    refractiveIndex: 0.7`

	desc := ParseYaml(yml)

	assert.Assert(t, desc != nil)
	assert.Assert(t, len(desc.Patterns.Checker) == 0)
	assert.Assert(t, len(desc.Colors) == 0)
	assert.Assert(t, len(desc.Materials) == 1)
	assert.Assert(t, desc.Materials[0].Name == "black_mat")
	assert.Assert(t, desc.Materials[0].Pattern == "")
	assert.Assert(t, desc.Materials[0].Ambient == nil)
	assert.Assert(t, *desc.Materials[0].Diffuse == 0.2)
	assert.Assert(t, *desc.Materials[0].Specular == 0.3)
	assert.Assert(t, *desc.Materials[0].Shininess == 0.4)
	assert.Assert(t, *desc.Materials[0].Reflective == 0.5)
	assert.Assert(t, *desc.Materials[0].Transparency == 0.6)
	assert.Assert(t, *desc.Materials[0].RefractiveIndex == 0.7)
}

func TestParseSceneObject(t *testing.T) {
	yml := `
scene:
  planes:
    - name: p1
      material: black_mat
      transform: rotate_right_90`

	desc := ParseYaml(yml)

	assert.Assert(t, desc != nil)
	assert.Assert(t, len(desc.Patterns.Checker) == 0)
	assert.Assert(t, len(desc.Colors) == 0)
	assert.Assert(t, len(desc.Materials) == 0)
	assert.Assert(t, len(desc.Scene.Planes) == 1)
	assert.Assert(t, desc.Scene.Planes[0].Name == "p1")
	assert.Assert(t, desc.Scene.Planes[0].Material == "black_mat")
	assert.Assert(t, desc.Scene.Planes[0].Transform == "rotate_right_90")
}

func TestParseGroup(t *testing.T) {
	yml := `
scene:
  planes:
    - name: p1
      material: black_mat
      transform: rotate_right_90
  cubes:
    - name: c1
      material: red_mat
      transform: rotate_right_90
  groups:
    - name: g1
      children:
        - p1
    - name: g2
      children:
        - g1
        - c1`

	desc := ParseYaml(yml)

	assert.Assert(t, desc != nil)
	assert.Assert(t, len(desc.Patterns.Checker) == 0)
	assert.Assert(t, len(desc.Colors) == 0)
	assert.Assert(t, len(desc.Materials) == 0)
	assert.Assert(t, len(desc.Scene.Planes) == 1)
	assert.Assert(t, desc.Scene.Planes[0].Name == "p1")
	assert.Assert(t, desc.Scene.Planes[0].Material == "black_mat")
	assert.Assert(t, desc.Scene.Planes[0].Transform == "rotate_right_90")
	assert.Assert(t, len(desc.Scene.Groups) == 2)
	assert.Assert(t, len(desc.Scene.Groups[0].Children) == 1)
	assert.Assert(t, len(desc.Scene.Groups[1].Children) == 2)
	assert.Assert(t, desc.Scene.Groups[1].Children[0] == "g1")
}

func TestParseTriangle(t *testing.T) {
	yml := `
scene:
  triangles:
    - name: t1
      material: black_mat
      transform: rotate_right_90
      p1:
        x: 5.0
        y: 7.5
        z: 10.0
      p2:
        x: 1.0
        y: 2.0
        z: 3.0
      p3:
        x: 10.0
        y: 20.0
        z: 30.0`

	desc := ParseYaml(yml)

	assert.Assert(t, desc != nil)
	assert.Assert(t, len(desc.Patterns.Checker) == 0)
	assert.Assert(t, len(desc.Colors) == 0)
	assert.Assert(t, len(desc.Materials) == 0)
	assert.Assert(t, len(desc.Scene.Triangles) == 1)
	assert.Assert(t, desc.Scene.Triangles[0].Name == "t1")
	assert.Assert(t, desc.Scene.Triangles[0].Material == "black_mat")
	assert.Assert(t, desc.Scene.Triangles[0].Transform == "rotate_right_90")
	assert.Assert(t, desc.Scene.Triangles[0].P1.X == 5.0)
	assert.Assert(t, desc.Scene.Triangles[0].P2.Y == 2.0)
	assert.Assert(t, desc.Scene.Triangles[0].P3.Z == 30.0)
}

func TestParseCone(t *testing.T) {
	yml := `
scene:
  cones:
    - name: c1
      material: black_mat
      transform: rotate_right_90
      min: 80.0
      max: 100.0
      closed: false`

	desc := ParseYaml(yml)

	assert.Assert(t, desc != nil)
	assert.Assert(t, len(desc.Patterns.Checker) == 0)
	assert.Assert(t, len(desc.Colors) == 0)
	assert.Assert(t, len(desc.Materials) == 0)
	assert.Assert(t, len(desc.Scene.Triangles) == 0)
	assert.Assert(t, len(desc.Scene.Cones) == 1)
	assert.Assert(t, desc.Scene.Cones[0].Name == "c1")
	assert.Assert(t, desc.Scene.Cones[0].Material == "black_mat")
	assert.Assert(t, desc.Scene.Cones[0].Transform == "rotate_right_90")
	assert.Assert(t, *desc.Scene.Cones[0].Minimum == 80.0)
	assert.Assert(t, *desc.Scene.Cones[0].Maximum == 100.0)
	assert.Assert(t, desc.Scene.Cones[0].Closed == false)
}

func TestParseObject(t *testing.T) {
	yml := `
scene:
  objects:
    - name: o1
      material: black_mat
      transform: rotate_right_90
      file: ../obj/test.obj`

	desc := ParseYaml(yml)

	assert.Assert(t, desc != nil)
	assert.Assert(t, len(desc.Patterns.Checker) == 0)
	assert.Assert(t, len(desc.Colors) == 0)
	assert.Assert(t, len(desc.Materials) == 0)
	assert.Assert(t, len(desc.Scene.Triangles) == 0)
	assert.Assert(t, len(desc.Scene.Objects) == 1)
	assert.Assert(t, desc.Scene.Objects[0].Name == "o1")
	assert.Assert(t, desc.Scene.Objects[0].Material == "black_mat")
	assert.Assert(t, desc.Scene.Objects[0].Transform == "rotate_right_90")
	assert.Assert(t, desc.Scene.Objects[0].File == "../obj/test.obj")
}

func TestParseCamera(t *testing.T) {
	yml := `
width: 400
height: 200
camera:
  from:
    x: 0.0
    y: 0.0
    z: 0.0
  to:
    x: 1.0
    y: 1.0
    z: 1.0
  up:
    x: 0.0
    y: 1.0
    z: 0.0`

	desc := ParseYaml(yml)

	assert.Assert(t, desc != nil)
	assert.Assert(t, desc.Width == 400)
	assert.Assert(t, desc.Height == 200)
	assert.Assert(t, desc.Camera.From.X == 0.0)
	assert.Assert(t, desc.Camera.To.Y == 1.0)
	assert.Assert(t, desc.Camera.Up.Y == 1.0)
	assert.Assert(t, desc.Camera.LookAt == "")
	assert.Assert(t, desc.Camera.Animation == nil)
}
