package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/eltrufas/pixeltetris/context"
	"github.com/eltrufas/pixeltetris/game"
	"github.com/eltrufas/pixeltetris/input"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

import _ "net/http/pprof"

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

	frametime, err := time.ParseDuration("16.66ms")

	ctx.PushState(game.CreateGame())

	pressed := make([]bool, 0, 32)
	for _, input := range input.InputArray {
		pressed = append(pressed, ctx.Win.Pressed(input))
	}

	for !win.Closed() && ctx.NotEmpty() {
		target := time.Now().Add(frametime)

		start := time.Now()

		for i, input := range input.InputArray {
			if ctx.Win.JustPressed(input) {
				pressed[i] = true
			}

			if ctx.Win.JustReleased(input) {
				pressed[i] = false
			}
		}

		ctx.Update(pressed)
		ctx.Render()

		fmt.Println(1 / time.Now().Sub(start).Seconds())

		dt := target.Sub(time.Now())
		time.Sleep(dt)
	}
}

func main() {
	// we need a webserver to get the pprof webserver
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	rand.Seed(time.Now().UTC().UnixNano())
	pixelgl.Run(run)
}
