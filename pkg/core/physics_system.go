package core

const PhysicsFps int = 60
const PhysicsFrameLengthS float32 = 1 / float32(PhysicsFps)
const PhysicsFrameLengthMs float32 = PhysicsFrameLengthS * 1000

type PhysicsSystem struct {
	simulatedMs float32
}

func (physicsSystem *PhysicsSystem) Init() {
	physicsSystem.simulatedMs = 0
}

func (physicsSystem *PhysicsSystem) Run(entityStore *EntityStore, elapsedMs float32) {
	for isFrameComplete(physicsSystem.simulatedMs, elapsedMs) {
		for idx := range entityStore.Entities {
			entity := &entityStore.Entities[idx]
			entity.transform.Position = entity.transform.Position.Add(entity.movement.DPosition.Mul(PhysicsFrameLengthS))
			entity.transform.Rotation = entity.transform.Rotation.Add(entity.movement.DRotation.Mul(PhysicsFrameLengthS))
		}

		physicsSystem.simulatedMs += PhysicsFrameLengthMs
	}
}

func isFrameComplete(simulatedMs float32, elapsedMs float32) bool {
	return elapsedMs-simulatedMs >= PhysicsFrameLengthMs
}
