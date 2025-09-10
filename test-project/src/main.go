package main

import (
	"game/config"
	eng "go3d/engine"
	"time"
)

func Game() eng.Game {
	return eng.Game{
		InputMappings: config.InputMappings,
		Levels: []*eng.Level{
			{
				Name: "Level 1",
				Player: (&eng.Player{
					InputMapping: "player",
				}).AddAction("jump_forward", func(self any) {
					eng.Print("jumping forward")
				}),
				Objects: []*eng.Object{
					(&eng.Object{
						Name:    "Default Cube",
						Variant: eng.MeshType,
						Mesh: &eng.Mesh{
							File: "cube.obj",
							Name: "Cube",
						},
						Children: nil, // []*eng.Object{}
					}).OnTick(func(self any) {
						selfRef := self.(*eng.Object)
						selfRef.Transform.Rotation.Z += 0.01
					}),
				},
			},
		},
	}
}

func main() {
	eng.Print("hello world")
	time.Sleep(time.Second * 10)
}
