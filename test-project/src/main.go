package game

import (
	eng "github.com/likemygames/go3D/engine"
)

func Game() eng.Game {
	return eng.Game{
		Name: "Game Name",
		Levels: []eng.Level{
			{
				Name:   "Level 1",
				Player: eng.Player{},
				Objects: []eng.Object{
					{
						Name:    "Default Cube",
						Variant: eng.MeshType,
						Mesh: eng.Mesh{
							File: "cube.obj",
							Name: "Cube",
						},
						Children: []eng.Object{},
					},
				},
			},
		},
	}
}
