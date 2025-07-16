package engine

type (
	Scene struct {
		Player      Player
		Environment any
		Objects     []Object
	}
)

func (s *Scene) Render() {

}
