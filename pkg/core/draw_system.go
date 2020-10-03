package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type DrawSystem struct {
	BasicProgram BasicProgram
	MeshStore    MeshStore
	VaoStore     VaoStore
}

func (drawSystem *DrawSystem) Init() error {
	if err := drawSystem.BasicProgram.Init(); err != nil {
		return err
	}
	drawSystem.MeshStore.Init()
	drawSystem.VaoStore.Init()

	return nil
}

func (drawSystem *DrawSystem) Run(entityStore *EntityStore, cameraStore *CameraStore) {
	camera := cameraStore.MainCamera
	worldToView := camera.Transform.Matrix().Inv()
	viewToProjection := camera.PerspectiveProjection.Matrix()

	for idx := range entityStore.Entities {
		entity := &entityStore.Entities[idx]
		mesh := drawSystem.MeshStore.GetMeshByID(entity.meshID)
		vao := drawSystem.VaoStore.GetVaoByMeshID(entity.meshID)

		gl.BindVertexArray(vao)

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
