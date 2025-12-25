package app

import (
	"testing"
)

func TestRun(t *testing.T) {
	Run([]string{"-f", "../teapot-scene.yaml"})
}
