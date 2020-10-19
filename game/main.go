package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(
		"Gaming In Go",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL,
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		err = window.Destroy()
		if err != nil {
			log.Fatal("Error Destroying Window", err)
		}
	}()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		err = renderer.Destroy()
		if err != nil {
			log.Fatal("Error Destroying Renderer", err)
		}
	}()

	p, err := newPlayer(renderer)
	if err != nil {
		log.Fatal("Error Creating Player :", err)
	}
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		err = renderer.SetDrawColor(255, 255, 255, 255)
		err = renderer.Clear()
		if err != nil {
			log.Fatal(err)
		}
		err = renderer.Copy(p.tex, &sdl.Rect{X: 0, Y: 0, W: 105, H: 105}, &sdl.Rect{X: 0, Y: 0, W: 105, H: 105})
		if err != nil {
			log.Fatal("Error Copying texture to  rect", err)
		}
		renderer.Present()
	}
}
