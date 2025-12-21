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
