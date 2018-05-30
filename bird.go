package main

import (
	"fmt"
	"os"
	"strconv"

	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var time = 1

type bird struct {
}

func newBird() *bird {
	return &bird{}
}

func (*bird) fly(renderer *sdl.Renderer) {
	if time <= 4 {
		time++
	} else {
		time = 1
	}

	birdimg, err := img.Load("./res/imgs/bird_frame_" + strconv.Itoa(time))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to load bird image %s", err)
	}
	defer birdimg.Free()

}
