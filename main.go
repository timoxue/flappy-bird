package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	window     *sdl.Window
	renderer   *sdl.Renderer
	err        error
	background *bg
)

func intialize() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		_ = fmt.Errorf("Could not init SDL %v", err)
	}

	window, renderer, err = sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		_ = fmt.Errorf("Could not create window and renderer %v", err)
	}

	background = newBackground(renderer)
}

func loop() {
	background.loop(renderer)
	renderer.Present()
}

func update() {

}

func destroy() {
	sdl.Delay(3000)
	background.destory()
	window.Destroy()
}

func main() {
	intialize()
	loop()
	update()
	destroy()
}
