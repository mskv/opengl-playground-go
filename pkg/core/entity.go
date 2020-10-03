package core

type EntityStore struct {
	Entities []Entity
}

type Entity struct {
	ID        int
	transform Transform
	movement  Movement
	meshID    MeshID
}

func (entityStore *EntityStore) Init() {
	entityStore.Entities = make([]Entity, 0, 5)
}

func (entityStore *EntityStore) GetEntityByID(entityID int) *Entity {
	// TODO: handle capacity
	return &entityStore.Entities[entityID]
}

func (entityStore *EntityStore) RegisterEntity(entity Entity) *Entity {
	entity.ID = len(entityStore.Entities)
	entityStore.Entities = append(entityStore.Entities, entity)
	return &entityStore.Entities[entity.ID]
}
