package engine

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

type (
	InputKey     string
	InputAction  string
	InputMapping map[InputKey]InputAction
	InputConfig  struct {
	}
)

// var (
// 	pressed_keys = []string{}
// )

const (
	SpacebarPress   = "input_spacebar_press"
	SpacebarRelease = "input_spacebar_release"
)

func InputSetup() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
}

func ReadInput() []string {
	// keys := []string{}
	for {
		event, err := keyboard.GetKeys(10)
		if err != nil {
			panic(err)
		}
		for {
			k := <-event
			fmt.Println(k)
		}
	}
}

func SetInputMapping(newMap string) {
	engine.currentInputMapping = engine.game.InputMappings[newMap]
}
