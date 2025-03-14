package math

import (
	"math"
)

func Translation(x float64, y float64, z float64) Matrix {
	translationMatrix := IdentityMatrix()
	translationMatrix.data[0][3] = x
	translationMatrix.data[1][3] = y
	translationMatrix.data[2][3] = z
	return translationMatrix
}

func Scaling(xScale float64, yScale float64, zScale float64) Matrix {
	scalingMatrix := IdentityMatrix()
	scalingMatrix.data[0][0] = xScale
	scalingMatrix.data[1][1] = yScale
	scalingMatrix.data[2][2] = zScale
	return scalingMatrix
}

func Rotation_X(rotRadian float64) Matrix {
	rotationMatrix := IdentityMatrix()
	rotationMatrix.data[1][1] = math.Cos(rotRadian)
	rotationMatrix.data[1][2] = math.Sin(rotRadian) * -1.0
	rotationMatrix.data[2][1] = math.Sin(rotRadian)
	rotationMatrix.data[2][2] = math.Cos(rotRadian)
	return rotationMatrix
}

func Rotation_Y(rotRadian float64) Matrix {
	rotationMatrix := IdentityMatrix()
	rotationMatrix.data[0][0] = math.Cos(rotRadian)
	rotationMatrix.data[0][2] = math.Sin(rotRadian)
	rotationMatrix.data[2][0] = math.Sin(rotRadian) * -1.0
	rotationMatrix.data[2][2] = math.Cos(rotRadian)
	return rotationMatrix
}

func Rotation_Z(rotRadian float64) Matrix {
	rotationMatrix := IdentityMatrix()
	rotationMatrix.data[0][0] = math.Cos(rotRadian)
	rotationMatrix.data[0][1] = math.Sin(rotRadian) * -1.0
	rotationMatrix.data[1][0] = math.Sin(rotRadian)
	rotationMatrix.data[1][1] = math.Cos(rotRadian)
	return rotationMatrix
}

func Shearing(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) Matrix {
	shearingMatrix := IdentityMatrix()
	shearingMatrix.data[0][1] = xy
	shearingMatrix.data[0][2] = xz
	shearingMatrix.data[1][0] = yx
	shearingMatrix.data[1][2] = yz
	shearingMatrix.data[2][0] = zx
	shearingMatrix.data[2][1] = zy
	return shearingMatrix
}

func ViewTransform(from Point, to Point, up Vector) Matrix {
	forward := to.Subtract(from).Normalize()
	upn := up.Normalize()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)

	orientationData := []float64{
		left.X, left.Y, left.Z, 0.0,
		trueUp.X, trueUp.Y, trueUp.Z, 0.0,
		-forward.X, -forward.Y, -forward.Z, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
	orientation := CreateMatrixFlat(orientationData)

	return orientation.MulM(Translation(-from.X, -from.Y, -from.Z))
}
