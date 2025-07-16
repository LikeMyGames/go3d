package engine

import (
	"fmt"
)

type (
	Object struct {
		name     string
		variant  ObjectType
		mesh     Mesh
		children []*Object
		ObjectFunctions
	}

	ObjectFunctions struct {
		OnBeginPlay *func()
		OnTick      *func(t uint64)
	}

	ObjectType int
)

const (
	MeshType ObjectType = iota
	ColliderType
)

func (o *Object) GetName() string {
	return o.name
}

func (o *Object) GetType() ObjectType {
	return o.variant
}

func (o *Object) GetMesh() Mesh {
	return o.mesh
}

func (o *Object) GetChildren() []*Object {
	return o.children
}

func NewObject() {
	fmt.Println(Object{
		name:    "test",
		variant: ColliderType,
	})
}
