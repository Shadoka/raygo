package app

import "testing"

func TestRun(t *testing.T) {
	Run([]string{"-f", "../local/penguin-scene.yaml", "-o", "penguin"})
}
