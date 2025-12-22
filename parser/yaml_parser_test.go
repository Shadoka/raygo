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
    ambient: 0.1
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
	assert.Assert(t, desc.Materials[0].Ambient == 0.1)
	assert.Assert(t, desc.Materials[0].Diffuse == 0.2)
	assert.Assert(t, desc.Materials[0].Specular == 0.3)
	assert.Assert(t, desc.Materials[0].Shininess == 0.4)
	assert.Assert(t, desc.Materials[0].Reflective == 0.5)
	assert.Assert(t, desc.Materials[0].Transparency == 0.6)
	assert.Assert(t, desc.Materials[0].RefractiveIndex == 0.7)
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
