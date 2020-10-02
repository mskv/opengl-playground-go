package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type DrawSystem struct {
	WindowWidth  int
	WindowHeight int

	BasicProgram BasicProgram
	MeshStore    MeshStore
	VaoStore     VaoStore
}

func (drawSystem *DrawSystem) Init(windowWidth int, windowHeight int) error {
	drawSystem.WindowWidth = windowWidth
	drawSystem.WindowHeight = windowHeight

	if err := drawSystem.BasicProgram.Init(); err != nil {
		return err
	}
	drawSystem.MeshStore.Init()
	drawSystem.VaoStore.Init()

	return nil
}

func (drawSystem *DrawSystem) Run(entityStore *EntityStore) {
	for _, entity := range entityStore.Entities {
		mesh := drawSystem.MeshStore.GetMeshByID(entity.meshID)
		vao := drawSystem.VaoStore.GetVaoByMeshID(entity.meshID)

		gl.BindVertexArray(vao)

		worldToView := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
		viewToProjection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(drawSystem.WindowWidth)/float32(drawSystem.WindowHeight), 0.1, 10.0)
		modelToWorld := entity.transform.Matrix()
		modelToView := worldToView.Mul4(modelToWorld)
		modelToProjection := viewToProjection.Mul4(modelToView)
		modelToViewInverseTranspose := modelToView.Inv().Transpose()

		drawSystem.BasicProgram.Use()
		drawSystem.BasicProgram.SetUniformModelToProjection(&modelToProjection)
		drawSystem.BasicProgram.SetUniformModelToViewInverseTranspose(&modelToViewInverseTranspose)

		gl.DrawElements(gl.TRIANGLES, int32(len(mesh.Indices)), gl.UNSIGNED_SHORT, gl.PtrOffset(0))
	}
}
