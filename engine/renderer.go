package engine

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type (
	Scene struct {
		Player      Player
		Environment any
		Objects     []Object
	}

	RasterRenderer struct {
		Scene    *Scene
		Window   *sdl.Window
		Renderer *sdl.Renderer
	}
)

func CreateRenderer(s *Scene, winName string) RasterRenderer {
	if err := sdl.Init(uint32(sdl.INIT_EVERYTHING)); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Go SDL2 Pixel Example",
		int32(sdl.WINDOWPOS_UNDEFINED), int32(sdl.WINDOWPOS_UNDEFINED),
		800, 600, uint32(sdl.WINDOW_SHOWN))
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, uint32(sdl.RENDERER_ACCELERATED))
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	return RasterRenderer{
		Scene:    s,
		Window:   window,
		Renderer: renderer,
	}
}

func (r *RasterRenderer) Render() {
	r.Renderer.Clear()
	fmt.Println(r.Renderer.GetLogicalSize())
}

func (r *RasterRenderer) Close() {
	sdl.Quit()
	r.Window.Destroy()
	r.Renderer.Destroy()
}
