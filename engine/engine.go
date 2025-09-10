package engine

import (
	"fmt"
	"time"
)

type (
	Engine struct {
		setup               func()
		update              func()
		framerate           float64
		currentInputMapping InputMapping
		game                Game
		renderer            RasterRenderer
	}

	Game struct {
		Name          string
		InputMappings map[string]InputMapping
		Levels        []*Level
	}

	Level struct {
		Name    string
		Player  *Player
		Objects []*Object
	}
)

var engine *Engine = nil

func NewEngine(setup, update func(), framerate float64) *Engine {
	if engine == nil {
		engine = &Engine{
			setup:     setup,
			update:    update,
			framerate: framerate,
		}
	}
	return engine
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

func Main() {
	engine := NewEngine(func() {}, nil, 60)
	engine.Start()
}

func Print(a ...any) {
	fmt.Println(a...)
}
