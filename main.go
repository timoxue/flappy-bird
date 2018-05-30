package main

import (
	"fmt"
	"os"

	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		_ = fmt.Errorf("Could not init SDL %v", err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		_ = fmt.Errorf("Could not create window and renderer %v", err)
	}
	defer window.Destroy()

	_ = renderer

	img, err := img.Load("./res/imgs/bg.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to load image %s\n", err)
	}
	defer img.Free()

	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to create texture; %s\n", err)
	}
	defer texture.Destroy()

	renderer.Clear()
	renderer.Copy(texture, nil, nil)
	renderer.Present()

	sdl.Delay(3000)
}
