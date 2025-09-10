package engine

import "fmt"

type (
	Object struct {
		Name            string
		Variant         ObjectType
		Transform       Transform
		Mesh            *Mesh
		Children        []*Object
		ObjectFunctions ObjectFunctions
	}

	Transform struct {
		Location Vector3D
		Rotation Vector3D
		Scale    Vector3D
	}

	ObjectFunctions struct {
		OnBeginPlay func()
		OnTick      func(t uint64)
	}

	ObjectType int
)

func NewObject(obj *Object) *Object {
	return obj
}

func (o *Object) OnBeginPlay(a func(self any)) *Object {
	Listen(fmt.Sprintf("object_%p_beginPlay", o), a, o)
	return o
}

func (o *Object) OnTick(a func(self any)) *Object {
	Listen(fmt.Sprintf("object_%p_tick", o), a, o)
	return o
}

const (
	MeshType ObjectType = iota
	ColliderType
)
