package engine

type (
	Mesh struct {
		File         string
		Name         string
		Faces        []Face
		Location     Vector3D
		CenterOffset Vector3D
		Verticies    []Vertex
		Normals      []Vector3D
	}

	MeshFileType int

	Vertex struct {
		XYZ   Vector3D
		W     *float64
		Color *Vector3D
	}

	Face struct {
		Verticies []uint
		Normal    *Vector3D
	}

	Vector3D struct {
		X float64
		Y float64
		Z float64
	}
)
