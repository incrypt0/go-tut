package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex *sdl.Texture
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {

	img, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil {
		return player{}, fmt.Errorf("loading player sprite : %v", err)

	}
	defer img.Free()
	p.tex, err = renderer.CreateTextureFromSurface(img)

	if err != nil {
		return player{}, fmt.Errorf("Error creating surface : %v", err)
	}
	return p, nil
}
