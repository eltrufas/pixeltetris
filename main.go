package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/eltrufas/pixeltetris/context"
	"github.com/eltrufas/pixeltetris/game"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

import _ "net/http/pprof"

func initInputArray() []pixelgl.Button {
	inputArray := make([]pixelgl.Button, 0, 32)

	inputArray = append(inputArray, pixelgl.KeyLeft)
	inputArray = append(inputArray, pixelgl.KeyRight)
	inputArray = append(inputArray, pixelgl.KeyUp)
	inputArray = append(inputArray, pixelgl.KeyDown)
	inputArray = append(inputArray, pixelgl.KeySpace)
	inputArray = append(inputArray, pixelgl.KeyLeftShift)

	return inputArray
}

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

	ia := initInputArray()
	pressed := make([]bool, 0, 32)
	for _, input := range ia {
		pressed = append(pressed, ctx.Win.Pressed(input))
	}

	for !win.Closed() {
		target := time.Now().Add(frametime)

		start := time.Now()

		for i, input := range ia {
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
