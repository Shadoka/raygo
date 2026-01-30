package math

import (
	"math"
	"slices"
)

type Matrix struct {
	data      []float64
	dimension int
}

func CreateMatrixFlat(m []float64) Matrix {
	dim := int(math.Sqrt(float64(len(m))))

	return Matrix{
		data:      m,
		dimension: dim,
	}
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
	mulData := make([]float64, len(m.data))
	for y := range m.dimension {
		for x := range m.dimension {
			mulData[CalculateFlatIndex(x, y, 4)] = m.Get(x, 0)*other.Get(0, y) +
				m.Get(x, 1)*other.Get(1, y) +
				m.Get(x, 2)*other.Get(2, y) +
				m.Get(x, 3)*other.Get(3, y)
		}
	}

	return CreateMatrixFlat(mulData)
}

func (m Matrix) MulT(t Tuple) Tuple {
	return Tuple{
		X: m.data[0]*t.X + m.data[1]*t.Y + m.data[2]*t.Z + m.data[3]*t.W,
		Y: m.data[4]*t.X + m.data[5]*t.Y + m.data[6]*t.Z + m.data[7]*t.W,
		Z: m.data[8]*t.X + m.data[9]*t.Y + m.data[10]*t.Z + m.data[11]*t.W,
		W: m.data[12]*t.X + m.data[13]*t.Y + m.data[14]*t.Z + m.data[15]*t.W,
	}
}

func (m Matrix) Get(row int, column int) float64 {
	return m.data[row*m.dimension+column]
}

func CalculateFlatIndex(row int, column int, dimension int) int {
	return row*dimension + column
}

func (m Matrix) Transpose() Matrix {
	result := make([]float64, len(m.data))

	for y := range m.dimension {
		for x := range m.dimension {
			result[CalculateFlatIndex(x, y, m.dimension)] = m.Get(y, x)
		}
	}

	return CreateMatrixFlat(result)
}

func (m Matrix) Determinant() float64 {
	result := 0.0
	if m.dimension == 2 {
		result = m.Get(0, 0)*m.Get(1, 1) - m.Get(0, 1)*m.Get(1, 0)
	} else {
		for column := range m.dimension {
			result = result + m.Get(0, column)*m.Cofactor(0, column)
		}
	}
	return result
}

func (m Matrix) Submatrix(row int, column int) Matrix {
	dataCopy := cloneMatrixData(m.data)
	rowOffset := row * m.dimension
	// delete the row
	dataCopy = slices.Delete(dataCopy, rowOffset, rowOffset+m.dimension)
	// delete the column
	for i := 0; i < m.dimension-1; i++ {
		deletionIndex := column + i*m.dimension - i
		dataCopy = slices.Delete(dataCopy, deletionIndex, deletionIndex+1)
	}

	return CreateMatrixFlat(dataCopy)
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
	invertedData := make([]float64, len(m.data))
	det := m.Determinant()

	for row := range m.dimension {
		for column := range m.dimension {
			c := m.Cofactor(row, column)
			invertedData[CalculateFlatIndex(column, row, m.dimension)] = c / det
		}
	}

	return CreateMatrixFlat(invertedData)
}

func IdentityMatrix() Matrix {
	return CreateMatrixFlat([]float64{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	})
}

func cloneMatrixData(data []float64) []float64 {
	result := make([]float64, len(data))
	copy(result, data)
	return result
}
