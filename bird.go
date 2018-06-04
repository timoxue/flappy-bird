package main

import (
	"fmt"
	"strconv"

	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type bird struct {
	time       int
	textures   []*sdl.Texture
	x, y, w, h int32
}

func newBird(renderer *sdl.Renderer) (*bird, error) {
	var textures []*sdl.Texture
	for i := 1; i <= 4; i++ {
		texture, err := img.LoadTexture(renderer, "./res/imgs/bird_frame_"+strconv.Itoa(i)+".png")
		if err != nil {
			return nil, fmt.Errorf("Could not load bird image %v", err)
		}
		textures = append(textures, texture)
	}
	return &bird{textures: textures, x: 30, y: 200, w: 50, h: 43}, nil
}

func (bird *bird) render(renderer *sdl.Renderer) error {
	rect := &sdl.Rect{X: bird.x, Y: 600 - bird.y - bird.h/2, W: bird.w, H: bird.h}
	if err := renderer.Copy(bird.textures[bird.time/10%4], nil, rect); err != nil {
		return fmt.Errorf("Can not copy bird image: %v", err)
	}
	bird.time++
	return nil
}

func (bird *bird) destroy() {
	for _, texture := range bird.textures {
		texture.Destroy()
	}
}
