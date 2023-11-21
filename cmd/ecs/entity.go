package ecs

type Entity struct {
	components map[ComponentType]ComponentTyper
}

// HasComponent returns true if the entity has the component of type cType associated
func (e *Entity) HasComponent(ctype ComponentType) bool {
	_, found := e.components[ctype]
	return found
}

// GetComponent returns the component from the entity, panics if not found
func (e *Entity) GetComponent(ctype ComponentType) ComponentTyper {
	c, found := e.components[ctype]
	if !found {
		panic("component not found")
	}
	return c
}

// AddComponent adds a component to the entity, panics if a component of the same type already exists
func (e *Entity) AddComponent(c ComponentTyper) {
	for ct := range e.components {
		if ct == c.Type() {
			panic("entity already contains this component")
		}
	}
	e.components[c.Type()] = c
}
