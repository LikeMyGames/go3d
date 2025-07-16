package engine

type (
	Player struct {
		object Object
	}
)

func (p *Player) GetObject() Object {
	return p.object
}
