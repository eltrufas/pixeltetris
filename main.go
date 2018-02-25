package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eltrufas/pixeltetris/context"
	"github.com/eltrufas/pixeltetris/game"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

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

	ctx := context.CreateContext(imdraw.New(nil), win)

	frametime, err := time.ParseDuration("16.66ms")

	ctx.PushState(game.CreateGame())

	ia := initInputArray()

	for !win.Closed() {
		target := time.Now().Add(frametime)

		pressed := make([]bool, 0, 32)

		for _, input := range ia {
			pressed = append(pressed, ctx.Win.Pressed(input))
		}

		win.Clear(colornames.Aliceblue)
		ctx.Update(pressed)
		ctx.Render()

		ctx.Imd.Draw(ctx.Win)
		win.Update()

		dt := target.Sub(time.Now())
		time.Sleep(dt)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pixelgl.Run(run)
}
