package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eltrufas/pixeltetris/context"
	"github.com/eltrufas/pixeltetris/game"
	"github.com/eltrufas/pixeltetris/input"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	fmt.Println("Lesgo")
	cfg := pixelgl.WindowConfig{
		Title:  "Tetris",
		Bounds: pixel.R(0, 0, 640, 480),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	ctx := context.CreateContext(nil, win)

	//frametime, err := time.ParseDuration("0.1ms")

	ctx.PushState(game.CreateGame())

	pressed := make([]bool, 0, 32)
	for _, input := range input.InputArray {
		pressed = append(pressed, ctx.Win.Pressed(input))
	}

	for !win.Closed() && ctx.NotEmpty() {
		//target := time.Now().Add(frametime)

		for i, input := range input.InputArray {
			if ctx.Win.JustPressed(input) {
				pressed[i] = true
			}

			if ctx.Win.JustReleased(input) {
				pressed[i] = false
			}
		}

		ctx.StartTimer()
		ctx.Update(pressed)
		ctx.Render()
		ctx.StopTimer()

		//dt := target.Sub(time.Now())
		//time.Sleep(dt)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pixelgl.Run(run)
}
