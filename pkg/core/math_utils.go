package core

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

func Mat4FromPosition(position *mgl32.Vec3) mgl32.Mat4 {
	return mgl32.Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, position[0], position[1], position[2], 1}
}

func Mat4FromEuler(euler *mgl32.Vec3) mgl32.Mat4 {
	x := euler[0]
	y := euler[1]
	z := euler[2]
	sinX := math.Sin(float64(x))
	cosX := math.Cos(float64(x))
	sinY := math.Sin(float64(y))
	cosY := math.Cos(float64(y))
	sinZ := math.Sin(float64(z))
	cosZ := math.Cos(float64(z))

	return mgl32.Mat4{
		float32(cosY * cosZ),
		float32(cosX*sinZ + cosZ*sinX*sinY),
		float32(sinX*sinZ - cosX*cosZ*sinY),
		0,
		float32(-cosY * sinZ),
		float32(cosX*cosZ - sinX*sinY*sinZ),
		float32(cosZ*sinX + cosX*sinY*sinZ),
		0,
		float32(sinY),
		float32(-cosY * sinX),
		float32(cosX * cosY),
		0,
		0,
		0,
		0,
		1,
	}
}

func Mat4FromScale(scale *mgl32.Vec3) mgl32.Mat4 {
	return mgl32.Mat4{scale[0], 0, 0, 0, 0, scale[1], 0, 0, 0, 0, scale[2], 0, 0, 0, 0, 1}
}
