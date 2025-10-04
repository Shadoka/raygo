package math

import (
	"math"
)

func Translation(x float64, y float64, z float64) Matrix {
	translationMatrix := IdentityMatrix()
	translationMatrix.data[3] = x
	translationMatrix.data[7] = y
	translationMatrix.data[11] = z
	return translationMatrix
}

func Scaling(xScale float64, yScale float64, zScale float64) Matrix {
	scalingMatrix := IdentityMatrix()
	scalingMatrix.data[0] = xScale
	scalingMatrix.data[5] = yScale
	scalingMatrix.data[10] = zScale
	return scalingMatrix
}

func Rotation_X(rotRadian float64) Matrix {
	rotationMatrix := IdentityMatrix()
	rotationMatrix.data[5] = math.Cos(rotRadian)
	rotationMatrix.data[6] = math.Sin(rotRadian) * -1.0
	rotationMatrix.data[9] = math.Sin(rotRadian)
	rotationMatrix.data[10] = math.Cos(rotRadian)
	return rotationMatrix
}

func Rotation_Y(rotRadian float64) Matrix {
	rotationMatrix := IdentityMatrix()
	rotationMatrix.data[0] = math.Cos(rotRadian)
	rotationMatrix.data[2] = math.Sin(rotRadian)
	rotationMatrix.data[8] = math.Sin(rotRadian) * -1.0
	rotationMatrix.data[10] = math.Cos(rotRadian)
	return rotationMatrix
}

func Rotation_Z(rotRadian float64) Matrix {
	rotationMatrix := IdentityMatrix()
	rotationMatrix.data[0] = math.Cos(rotRadian)
	rotationMatrix.data[1] = math.Sin(rotRadian) * -1.0
	rotationMatrix.data[4] = math.Sin(rotRadian)
	rotationMatrix.data[5] = math.Cos(rotRadian)
	return rotationMatrix
}

func Shearing(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) Matrix {
	shearingMatrix := IdentityMatrix()
	shearingMatrix.data[1] = xy
	shearingMatrix.data[2] = xz
	shearingMatrix.data[4] = yx
	shearingMatrix.data[6] = yz
	shearingMatrix.data[8] = zx
	shearingMatrix.data[9] = zy
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
