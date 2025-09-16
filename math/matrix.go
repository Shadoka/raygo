package math

import (
	"fmt"
	"math"
	"slices"
)

type Matrix struct {
	data              [][]float64
	dimension         int
	cachedDeterminant *float64
}

func CreateMatrix(m [][]float64) Matrix {
	return Matrix{
		data:              m,
		dimension:         len(m),
		cachedDeterminant: nil,
	}
}

func CreateMatrixFlat(m []float64) Matrix {
	dim := int(math.Sqrt(float64(len(m))))

	nestedMatrix := create2dArray(dim)

	y := 0
	for x, d := range m {
		y = x / dim
		nestedMatrix[y][x%dim] = d
	}

	return Matrix{
		data:      nestedMatrix,
		dimension: dim,
	}
}

func (m Matrix) Get(x int, y int) float64 {
	return m.data[x][y]
}

func (m Matrix) Equals(other Matrix) bool {
	if m.dimension != other.dimension {
		return false
	}

	result := true
	for y := range m.dimension {
		for x := range m.dimension {
			result = result && floatEquals(m.Get(x, y), other.Get(x, y))
		}
	}

	return result
}

func (m Matrix) MulM(other Matrix) Matrix {
	if m.dimension != other.dimension {
		panic(fmt.Sprintf("matrices have unequal dimensions: %v & %v", m.dimension, other.dimension))
	}

	if m.dimension != 4 {
		panic("matrix multiplication only implemented for matrices of dim 4")
	}

	mulData := create2dArray(m.dimension)
	for y := range m.dimension {
		for x := range m.dimension {
			mulData[x][y] = m.data[x][0]*other.data[0][y] +
				m.data[x][1]*other.data[1][y] +
				m.data[x][2]*other.data[2][y] +
				m.data[x][3]*other.data[3][y]
		}
	}

	return CreateMatrix(mulData)
}

func (m Matrix) MulT(t Tuple) Tuple {
	if m.dimension != 4 {
		panic("matrix/tuple multiplication only implemented for matrices of dim 4")
	}

	return Tuple{
		X: m.data[0][0]*t.X + m.data[0][1]*t.Y + m.data[0][2]*t.Z + m.data[0][3]*t.W,
		Y: m.data[1][0]*t.X + m.data[1][1]*t.Y + m.data[1][2]*t.Z + m.data[1][3]*t.W,
		Z: m.data[2][0]*t.X + m.data[2][1]*t.Y + m.data[2][2]*t.Z + m.data[2][3]*t.W,
		W: m.data[3][0]*t.X + m.data[3][1]*t.Y + m.data[3][2]*t.Z + m.data[3][3]*t.W,
	}
}

func (m Matrix) Transpose() Matrix {
	result := create2dArray(m.dimension)

	for y := range m.dimension {
		for x := range m.dimension {
			result[x][y] = m.data[y][x]
		}
	}

	return CreateMatrix(result)
}

func (m *Matrix) Determinant() float64 {
	if m.cachedDeterminant != nil {
		return *m.cachedDeterminant
	}
	result := 0.0
	if m.dimension == 2 {
		result = m.data[0][0]*m.data[1][1] - m.data[0][1]*m.data[1][0]
	} else {
		for column := range m.dimension {
			result = result + m.data[0][column]*m.Cofactor(0, column)
		}
	}
	m.cachedDeterminant = &result
	return result
}

func (m Matrix) Submatrix(row int, column int) Matrix {
	if int(math.Max(float64(row), float64(column))) > m.dimension-1 {
		panic(fmt.Sprintf("trying to create a submatrix with wrong parameters. dim: %v, row: %v, col: %v", m.dimension, row, column))
	}

	dataCopy := cloneMatrixData(m.data)
	subData := slices.Delete(dataCopy, row, row+1)
	for currentRow := range m.dimension - 1 {
		subData[currentRow] = slices.Delete(subData[currentRow], column, column+1)
	}

	return CreateMatrix(subData)
}

func (m Matrix) Minor(row int, column int) float64 {
	submatrix := m.Submatrix(row, column)
	return submatrix.Determinant()
}

func (m Matrix) Cofactor(row int, column int) float64 {
	minor := m.Minor(row, column)

	if (row+column)%2 == 1 {
		minor = minor * -1
	}

	return minor
}

func (m Matrix) IsInvertible() bool {
	return !floatEquals(0.0, m.Determinant())
}

func (m Matrix) Inverse() Matrix {
	if !m.IsInvertible() {
		panic("trying to invert a non invertible matrix")
	}

	invertedData := create2dArray(m.dimension)
	det := m.Determinant()

	for row := range m.dimension {
		for column := range m.dimension {
			c := m.Cofactor(row, column)
			invertedData[column][row] = c / det
		}
	}

	return CreateMatrix(invertedData)
}

func IdentityMatrix() Matrix {
	return CreateMatrix([][]float64{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	})
}

func create2dArray(d int) [][]float64 {
	nestedMatrix := make([][]float64, d)
	for i := range d {
		nestedMatrix[i] = make([]float64, d)
	}
	return nestedMatrix
}

func cloneMatrixData(data [][]float64) [][]float64 {
	result := make([][]float64, len(data))
	for i := range data {
		result[i] = make([]float64, len(data[i]))
		copy(result[i], data[i])
	}
	return result
}
