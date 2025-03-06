package math

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateMatrix(t *testing.T) {
	data := [][]float64{
		{1.0, 2.0, 3.0, 4.0},
		{5.5, 6.5, 7.5, 8.5},
		{9.0, 10.0, 11.0, 12.0},
		{13.5, 14.5, 15.5, 16.5},
	}
	m := CreateMatrix(data)

	assert.Assert(t, m.Get(0, 0) == 1.0)
	assert.Assert(t, m.Get(1, 1) == 6.5)
	assert.Assert(t, m.Get(2, 2) == 11.0)
	assert.Assert(t, m.Get(3, 3) == 16.5)
}

func TestCreateMatrixFlat(t *testing.T) {
	data := []float64{
		1.0, 2.0, 3.0, 4.0,
		5.5, 6.5, 7.5, 8.5,
		9.0, 10.0, 11.0, 12.0,
		13.5, 14.5, 15.5, 16.5,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, m.Get(0, 0) == 1.0)
	assert.Assert(t, m.Get(1, 1) == 6.5)
	assert.Assert(t, m.Get(2, 2) == 11.0)
	assert.Assert(t, m.Get(3, 3) == 16.5)
}

func TestCreateMatrix2By2(t *testing.T) {
	data := []float64{
		-3.0, 5.0,
		1.0, -2.0,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, m.Get(0, 0) == -3.0)
	assert.Assert(t, m.Get(0, 1) == 5.0)
	assert.Assert(t, m.Get(1, 0) == 1.0)
	assert.Assert(t, m.Get(1, 1) == -2.0)
}

func TestCreateMatrix3By3(t *testing.T) {
	data := []float64{
		-3.0, 5.0, 0.0,
		1.0, -2.0, -7.0,
		0.0, 1.0, 1.0,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, m.Get(0, 0) == -3.0)
	assert.Assert(t, m.Get(1, 1) == -2.0)
	assert.Assert(t, m.Get(2, 2) == 1.0)
}

func TestEqualsIdentical(t *testing.T) {
	data := []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, m.Equals(m))
}

func TestEqualsNotIdentical(t *testing.T) {
	data := []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	}
	m := CreateMatrixFlat(data)

	data = []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.2, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	}
	m2 := CreateMatrixFlat(data)

	assert.Assert(t, !m.Equals(m2))
}

func TestMulM(t *testing.T) {
	data := []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	}
	m := CreateMatrixFlat(data)

	data = []float64{
		-2.0, 1.0, 2.0, 3.0,
		3.0, 2.0, 1.0, -1.0,
		4.0, 3.0, 6.0, 5.0,
		1.0, 2.0, 7.0, 8.0,
	}
	m2 := CreateMatrixFlat(data)

	data = []float64{
		20.0, 22.0, 50.0, 48.0,
		44.0, 54.0, 114.0, 108.0,
		40.0, 58.0, 110.0, 102.0,
		16.0, 26.0, 46.0, 42.0,
	}
	expected := CreateMatrixFlat(data)

	actual := m.MulM(m2)

	assert.Assert(t, expected.Equals(actual))
}

func TestMulT(t *testing.T) {
	data := []float64{
		1.0, 2.0, 3.0, 4.0,
		2.0, 4.0, 4.0, 2.0,
		8.0, 6.0, 4.0, 1.0,
		0.0, 0.0, 0.0, 1.0,
	}
	m := CreateMatrixFlat(data)
	tuple := CreateTuple(1.0, 2.0, 3.0, 1.0)
	expected := CreateTuple(18.0, 24.0, 33.0, 1.0)

	actual := m.MulT(tuple)

	assert.Assert(t, expected.Equals(actual))
}

func TestMulMIdentity(t *testing.T) {
	data := []float64{
		20.0, 22.0, 50.0, 48.0,
		44.0, 54.0, 114.0, 108.0,
		40.0, 58.0, 110.0, 102.0,
		16.0, 26.0, 46.0, 42.0,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, m.Equals(m.MulM(IdentityMatrix())))
}

func TestTranspose(t *testing.T) {
	data := []float64{
		0.0, 9.0, 3.0, 0.0,
		9.0, 8.0, 0.0, 8.0,
		1.0, 8.0, 5.0, 3.0,
		0.0, 0.0, 5.0, 8.0,
	}
	m := CreateMatrixFlat(data)

	data = []float64{
		0.0, 9.0, 1.0, 0.0,
		9.0, 8.0, 8.0, 0.0,
		3.0, 0.0, 5.0, 5.0,
		0.0, 8.0, 3.0, 8.0,
	}
	expected := CreateMatrixFlat(data)

	assert.Assert(t, expected.Equals(m.Transpose()))
}

func TestTransposeIdentity(t *testing.T) {
	assert.Assert(t, IdentityMatrix().Equals(IdentityMatrix().Transpose()))
}

func TestDeterminant(t *testing.T) {
	data := []float64{
		1.0, 5.0,
		-3.0, 2.0,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, floatEquals(17, m.Determinant()))
}

func TestSubmatrix3By3(t *testing.T) {
	data := []float64{
		1.0, 5.0, 0.0,
		-3.0, 2.0, 7.0,
		0.0, 6.0, -3.0,
	}
	m := CreateMatrixFlat(data)

	data = []float64{
		-3.0, 2.0,
		0.0, 6.0,
	}
	expected := CreateMatrixFlat(data)

	assert.Assert(t, expected.Equals(m.Submatrix(0, 2)))
}

func TestSubmatrix4By4(t *testing.T) {
	data := []float64{
		-6.0, 1.0, 1.0, 6.0,
		-8.0, 5.0, 8.0, 6.0,
		-1.0, 0.0, 8.0, 2.0,
		-7.0, 1.0, -1.0, 1.0,
	}
	m := CreateMatrixFlat(data)

	data = []float64{
		-6.0, 1.0, 6.0,
		-8.0, 8.0, 6.0,
		-7.0, -1.0, 1.0,
	}
	expected := CreateMatrixFlat(data)

	assert.Assert(t, expected.Equals(m.Submatrix(2, 1)))
}

func TestMinor(t *testing.T) {
	data := []float64{
		3.0, 5.0, 0.0,
		2.0, -1.0, -7.0,
		6.0, -1.0, 5.0,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, floatEquals(25.0, m.Minor(1, 0)))
}

func TestCofactor(t *testing.T) {
	data := []float64{
		3.0, 5.0, 0.0,
		2.0, -1.0, -7.0,
		6.0, -1.0, 5.0,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, floatEquals(-12.0, m.Cofactor(0, 0)))
	assert.Assert(t, floatEquals(-25.0, m.Cofactor(1, 0)))
}

func TestDeterminant3By3(t *testing.T) {
	data := []float64{
		1.0, 2.0, 6.0,
		-5.0, 8.0, -4.0,
		2.0, 6.0, 4.0,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, floatEquals(56.0, m.Cofactor(0, 0)))
	assert.Assert(t, floatEquals(12.0, m.Cofactor(0, 1)))
	assert.Assert(t, floatEquals(-46.0, m.Cofactor(0, 2)))
	assert.Assert(t, floatEquals(-196.0, m.Determinant()))
}

func TestDeterminant4By4(t *testing.T) {
	data := []float64{
		-2.0, -8.0, 3.0, 5.0,
		-3.0, 1.0, 7.0, 3.0,
		1.0, 2.0, -9.0, 6.0,
		-6.0, 7.0, 7.0, -9.0,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, floatEquals(690.0, m.Cofactor(0, 0)))
	assert.Assert(t, floatEquals(447.0, m.Cofactor(0, 1)))
	assert.Assert(t, floatEquals(210.0, m.Cofactor(0, 2)))
	assert.Assert(t, floatEquals(51.0, m.Cofactor(0, 3)))
	assert.Assert(t, floatEquals(-4071.0, m.Determinant()))
}

func TestIsInvertibleTrue(t *testing.T) {
	data := []float64{
		6.0, 4.0, 4.0, 4.0,
		5.0, 5.0, 7.0, 6.0,
		4.0, -9.0, 3.0, -7.0,
		9.0, 1.0, 7.0, -6.0,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, m.IsInvertible())
}

func TestIsInvertibleFalse(t *testing.T) {
	data := []float64{
		-4.0, 2.0, -2.0, -3.0,
		9.0, 6.0, 2.0, 6.0,
		0.0, -5.0, 1.0, -5.0,
		0.0, 0.0, 0.0, 0.0,
	}
	m := CreateMatrixFlat(data)

	assert.Assert(t, !m.IsInvertible())
}

func TestInverse1(t *testing.T) {
	data := []float64{
		-5.0, 2.0, 6.0, -8.0,
		1.0, -5.0, 1.0, 8.0,
		7.0, 7.0, -6.0, -7.0,
		1.0, -3.0, 7.0, 4.0,
	}
	m := CreateMatrixFlat(data)

	data = []float64{
		0.21805, 0.45113, 0.24060, -0.04511,
		-0.80827, -1.45677, -0.44361, 0.52068,
		-0.07895, -0.22368, -0.05263, 0.19737,
		-0.52256, -0.81391, -0.30075, 0.30639,
	}
	expected := CreateMatrixFlat(data)

	assert.Assert(t, expected.Equals(m.Inverse()))
}

func TestInverse2(t *testing.T) {
	data := []float64{
		8.0, -5.0, 9.0, 2.0,
		7.0, 5.0, 6.0, 1.0,
		-6.0, 0.0, 9.0, 6.0,
		-3.0, 0.0, -9.0, -4.0,
	}
	m := CreateMatrixFlat(data)

	data = []float64{
		-0.15385, -0.15385, -0.28205, -0.53846,
		-0.07692, 0.12308, 0.02564, 0.03077,
		0.35897, 0.35897, 0.43590, 0.92308,
		-0.69231, -0.69231, -0.76923, -1.92308,
	}
	expected := CreateMatrixFlat(data)

	assert.Assert(t, expected.Equals(m.Inverse()))
}

func TestInverse3(t *testing.T) {
	data := []float64{
		9.0, 3.0, 0.0, 9.0,
		-5.0, -2.0, -6.0, -3.0,
		-4.0, 9.0, 6.0, 4.0,
		-7.0, 6.0, 6.0, 2.0,
	}
	m := CreateMatrixFlat(data)

	data = []float64{
		-0.04074, -0.07778, 0.14444, -0.22222,
		-0.07778, 0.03333, 0.36667, -0.33333,
		-0.02901, -0.14630, -0.10926, 0.12963,
		0.17778, 0.06667, -0.26667, 0.33333,
	}
	expected := CreateMatrixFlat(data)

	assert.Assert(t, expected.Equals(m.Inverse()))
}

func TestMulWithInverse(t *testing.T) {
	data := []float64{
		3.0, -9.0, 7.0, 3.0,
		3.0, -8.0, 2.0, -9.0,
		-4.0, 4.0, 4.0, 1.0,
		-6.0, 5.0, -1.0, 1.0,
	}
	m1 := CreateMatrixFlat(data)

	data = []float64{
		8.0, 2.0, 2.0, 2.0,
		3.0, -1.0, 7.0, 0.0,
		7.0, 0.0, 5.0, 4.0,
		6.0, -2.0, 0.0, 5.0,
	}
	m2 := CreateMatrixFlat(data)

	m3 := m1.MulM(m2)
	inverseM2 := m2.Inverse()

	assert.Assert(t, m1.Equals(m3.MulM(inverseM2)))
}
