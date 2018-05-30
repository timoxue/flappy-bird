package main

import (
	"fmt"
	"os"

	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type bg struct {
	texture *sdl.Texture
}

func newBackground(renderer *sdl.Renderer) *bg {
	texture, err := img.LoadTexture(renderer, "./res/imgs/bg.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to load background image %s\n", err)
	}
	return &bg{texture: texture}
}

func (bg *bg) destory() {
	bg.texture.Destroy()
}

func (bg *bg) loop(renderer *sdl.Renderer) {
	renderer.Clear()
	if err := renderer.Copy(bg.texture, nil, nil); err != nil {
		fmt.Fprintf(os.Stderr, "fail to loop or copy background image %s\n", err)
	}
}
