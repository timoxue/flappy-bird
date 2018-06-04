package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	window     *sdl.Window
	renderer   *sdl.Renderer
	err        error
	background *bg
	flybird    *bird
	events     chan sdl.Event
)

func intialize() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		_ = fmt.Errorf("Could not init SDL %v", err)
	}

	window, renderer, err = sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create window and renderer %v", err)
	}

	background = newBackground(renderer)
	flybird, err = newBird(renderer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create window and renderer %v", err)
	}
	events = make(chan sdl.Event)
	go func(events chan sdl.Event) {
		for {
			select {
			case events <- sdl.WaitEvent():
			default:
			}
		}
	}(events)
}

func loop() {
	tick := time.Tick(10 * time.Millisecond)
	for {
		select {
		case e := <-events:
			switch e.(type) {
			case *sdl.QuitEvent:
				return
			}
		case <-tick:
			background.loop(renderer)
			flybird.render(renderer)
			renderer.Present()
		}
	}

}

func update(events <-chan sdl.Event, renderer *sdl.Renderer) {

}

func destroy() {
	sdl.Delay(3000)
	flybird.destroy()
	background.destory()
	window.Destroy()
}

func main() {
	intialize()
	loop()
	destroy()
}
