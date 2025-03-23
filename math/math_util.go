package math

import gomath "math"

const EPSILON = 0.00001

func floatEquals(a float64, b float64) bool {
	return gomath.Abs(a-b) < EPSILON
}

func ClampToByte(n float64) uint64 {
	if n > 255.0 {
		return 255
	}
	if n < 0.0 {
		return 0
	}

	return uint64(n)
}

// color representation byte to float conversion
func BToF(c int) float64 {
	return float64(c) / 255.0
}
