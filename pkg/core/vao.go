package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type VaoStore struct {
	meshIDToVao []uint32
}

func (vaoStore *VaoStore) Init() {
	vaoStore.meshIDToVao = make([]uint32, 5, 5)
}

func (vaoStore *VaoStore) RegisterMesh(mesh *Mesh) {
	// TODO: ensure capacity
	vaoStore.meshIDToVao[mesh.ID] = createMeshVao(mesh)
}

func (vaoStore *VaoStore) GetVaoByMeshID(meshID MeshID) uint32 {
	// TODO: check capacity
	return vaoStore.meshIDToVao[meshID]
}

func createMeshVao(mesh *Mesh) uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var positionVbo uint32
	gl.GenBuffers(1, &positionVbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, positionVbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(mesh.Positions)*4, gl.Ptr(mesh.Positions), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(PositionAttributeLocation)
	gl.VertexAttribPointer(PositionAttributeLocation, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	var normalVbo uint32
	gl.GenBuffers(1, &normalVbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, normalVbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(mesh.Normals)*4, gl.Ptr(mesh.Normals), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(NormalAttributeLocation)
	gl.VertexAttribPointer(NormalAttributeLocation, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	var indexVbo uint32
	gl.GenBuffers(1, &indexVbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexVbo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(mesh.Indices)*2, gl.Ptr(mesh.Indices), gl.STATIC_DRAW)

	return vao
}
