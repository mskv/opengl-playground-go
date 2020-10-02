package core

import (
	"fmt"
)

type MeshID = int

type MeshStore struct {
	meshes []Mesh
}

type Mesh struct {
	ID        MeshID
	Name      string
	Positions []float32
	Normals   []float32
	Indices   []uint16
}

func (meshStore *MeshStore) Init() {
	meshStore.meshes = make([]Mesh, 0, 5)
}

func (meshStore *MeshStore) RegisterMesh(name string, positions []float32, normals []float32, indices []uint16) (*Mesh, error) {
	existent := meshStore.GetMeshByName(name)
	if existent != nil {
		return nil, fmt.Errorf("Mesh %s already registered", name)
	}

	meshID := len(meshStore.meshes)

	mesh := Mesh{
		ID:        meshID,
		Name:      name,
		Positions: positions,
		Normals:   normals,
		Indices:   indices,
	}

	meshStore.meshes = append(meshStore.meshes, mesh)

	return &mesh, nil
}

func (meshStore *MeshStore) GetMeshByID(meshID MeshID) *Mesh {
	// TODO: handle capacity
	return &meshStore.meshes[meshID]
}

func (meshStore *MeshStore) GetMeshByName(name string) *Mesh {
	for _, mesh := range meshStore.meshes {
		if mesh.Name == name {
			return &mesh
		}
	}
	return nil
}
