package core

import "github.com/go-gl/mathgl/mgl32"

type System struct {
	entityStore   EntityStore
	drawSystem    DrawSystem
	physicsSystem PhysicsSystem
}

func (system *System) Init(windowWidth int, windowHeight int) error {
	if err := system.drawSystem.Init(windowWidth, windowHeight); err != nil {
		return err
	}

	cubePositions, cubeNormals, cubeIndices := buildCubeMesh()
	cubeMesh, err := system.drawSystem.MeshStore.RegisterMesh("cube", cubePositions, cubeNormals, cubeIndices)
	if err != nil {
		panic(err)
	}

	system.drawSystem.VaoStore.RegisterMesh(cubeMesh)

	system.entityStore.Init()

	entity := Entity{
		meshID: cubeMesh.ID,
		transform: Transform{
			Position: mgl32.Vec3{0, 0, 0},
			Rotation: mgl32.Vec3{0, 0, 0},
			Scale:    mgl32.Vec3{1, 1, 1},
		},
	}

	system.entityStore.RegisterEntity(entity)

	system.physicsSystem.Init()

	return nil
}

func (system *System) Run() {
	system.drawSystem.Run(&system.entityStore)
}

func buildCubeMesh() ([]float32, []float32, []uint16) {
	v0 := mgl32.Vec3{0, 0, 0}
	v1 := mgl32.Vec3{1, 0, 0}
	v2 := mgl32.Vec3{0, 1, 0}
	v3 := mgl32.Vec3{1, 1, 0}
	v4 := mgl32.Vec3{0, 0, 1}
	v5 := mgl32.Vec3{1, 0, 1}
	v6 := mgl32.Vec3{0, 1, 1}
	v7 := mgl32.Vec3{1, 1, 1}
	farN := v2.Sub(v0).Cross(v1.Sub(v0)).Normalize()
	leftN := v4.Sub(v0).Cross(v1.Sub(v2)).Normalize()
	rightN := v1.Sub(v5).Cross(v7.Sub(v5)).Normalize()
	bottomN := v1.Sub(v0).Cross(v4.Sub(v0)).Normalize()
	topN := v6.Sub(v2).Cross(v3.Sub(v2)).Normalize()
	nearN := v5.Sub(v4).Cross(v6.Sub(v4)).Normalize()
	n0 := farN.Add(leftN).Add(bottomN).Normalize()
	n1 := farN.Add(rightN).Add(bottomN).Normalize()
	n2 := farN.Add(leftN).Add(topN).Normalize()
	n3 := farN.Add(rightN).Add(topN).Normalize()
	n4 := nearN.Add(leftN).Add(bottomN).Normalize()
	n5 := nearN.Add(rightN).Add(bottomN).Normalize()
	n6 := nearN.Add(leftN).Add(topN).Normalize()
	n7 := nearN.Add(rightN).Add(topN).Normalize()

	positions := []float32{
		v0[0], v0[1], v0[2],
		v1[0], v1[1], v1[2],
		v2[0], v2[1], v2[2],
		v3[0], v3[1], v3[2],
		v4[0], v4[1], v4[2],
		v5[0], v5[1], v5[2],
		v6[0], v6[1], v6[2],
		v7[0], v7[1], v7[2],
	}

	normals := []float32{
		n0[0], n0[1], n0[2],
		n1[0], n1[1], n1[2],
		n2[0], n2[1], n2[2],
		n3[0], n3[1], n3[2],
		n4[0], n4[1], n4[2],
		n5[0], n5[1], n5[2],
		n6[0], n6[1], n6[2],
		n7[0], n7[1], n7[2],
	}

	indices := []uint16{
		// far
		0, 1, 2,
		1, 2, 3,
		// left
		0, 2, 4,
		2, 4, 6,
		// right
		1, 3, 5,
		3, 5, 7,
		// bottom
		0, 1, 4,
		1, 4, 5,
		// top
		2, 3, 6,
		3, 6, 7,
		// near
		4, 5, 6,
		5, 6, 7,
	}

	return positions, normals, indices
}
