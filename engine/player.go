package engine

import "fmt"

type (
	Player struct {
		Object       Object
		InputMapping string // key of a key-value pair that exists in engine.Game.InputMappings
	}
)

func NewPlayer(p *Player) *Player {
	return p
}

func (p *Player) AddAction(name string, a func(self any)) *Player {
	Listen(fmt.Sprintf("action_%s", name), a, p)
	return p
}
