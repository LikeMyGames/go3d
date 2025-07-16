package engine

import (
	"time"
)

type (
	Engine struct {
		setup        func()
		update       func()
		framerate    float64
		currentLevel Scene
	}
)

func NewEngine(setup, update func(), framerate float64) Engine {
	return Engine{
		setup:        setup,
		update:       update,
		framerate:    framerate,
		currentLevel: Scene{},
	}
}

func (e *Engine) Start() {
	e.setup()
	go ReadInput()
	t := 0
	for {
		time.Sleep(time.Second / time.Duration(60))
		// fmt.Println(t)
		t++
	}
}

func main() {
	engine := NewEngine(func() {}, nil, 60)
	engine.Start()
}
