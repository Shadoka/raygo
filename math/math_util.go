package math

import gomath "math"

const EPSILON = 0.00001

func floatEquals(a float64, b float64) bool {
	return gomath.Abs(a-b) < EPSILON
}
