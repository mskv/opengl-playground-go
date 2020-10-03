package core

import (
	"github.com/go-gl/mathgl/mgl32"
)

type CameraStore struct {
	MainCamera Camera
}

type Camera struct {
	Transform             Transform
	PerspectiveProjection PerspectiveProjection
}

type PerspectiveProjection struct {
	NearZ         float32
	FarZ          float32
	FovY          float32
	AspectRatioXY float32
}

func (perspectiveProjection *PerspectiveProjection) Matrix() mgl32.Mat4 {
	return mgl32.Perspective(
		perspectiveProjection.FovY,
		perspectiveProjection.AspectRatioXY,
		perspectiveProjection.NearZ,
		perspectiveProjection.FarZ,
	)
}
