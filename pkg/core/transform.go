package core

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Transform struct {
	Position mgl32.Vec3
	Rotation mgl32.Vec3
	Scale    mgl32.Vec3
}

func (t *Transform) Matrix() mgl32.Mat4 {
	translation := Mat4FromPosition(&t.Position)
	rotation := Mat4FromEuler(&t.Rotation)
	scale := Mat4FromScale(&t.Scale)

	return translation.Mul4(rotation.Mul4(scale))
}
