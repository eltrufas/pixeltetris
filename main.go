package main

import (
	"fmt"
	"github.com/eltrufas/tetriscore"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	BoardOffsetX = 10
	BoardOffsetY = 470
	BlockW       = 16
	BlockH       = 16
)

type PixelTetris struct {
	T   *tetriscore.Tetris
	Imd *imdraw.IMDraw
}

func (pt *PixelTetris) Render() {
	pt.Imd.Clear()
	for i := 0; i < 200; i++ {
		value := pt.T.Board[i+20]

		pt.RenderBlock(i, value)
	}

	pt.RenderCurrentPiece()
}

func (pt *PixelTetris) RenderCurrentPiece() {
	p := pt.T.CurrentPiece
	mask := tetriscore.Tetrominos[p.TetrominoType][p.State]
	color := tetriscore.TetrominoColors[p.TetrominoType]

	for i := 0; i < 16; i++ {
		if mask[i] == 1 {
			y := p.Y + i/4 - 2
			x := p.X + i%4

			pt.RenderBlock(x+y*10, color)
		}
	}
}

func (pt *PixelTetris) RenderNext() {

}

func (pt *PixelTetris) RenderBlock(i int, block tetriscore.Block) {
	x := i % 10
	y := i / 10
	r := pixel.R(
		float64(BoardOffsetX+BlockW*x),
		float64(BoardOffsetY-BlockH*y),
		float64(BoardOffsetX+BlockW*(x+1)),
		float64(BoardOffsetY-BlockH*(y+1)),
	)

	switch block {
	case tetriscore.Yellow:
		pt.Imd.Color = pixel.RGB(1, 0.93, 0.36)
	case tetriscore.Red:
		pt.Imd.Color = pixel.RGB(255, 0, 0)
	case tetriscore.Cyan:
		pt.Imd.Color = pixel.RGB(0, 255, 255)
	case tetriscore.Green:
		pt.Imd.Color = pixel.RGB(0, 255, 0)
	case tetriscore.Purple:
		pt.Imd.Color = pixel.RGB(204, 0, 204)
	case tetriscore.Blue:
		pt.Imd.Color = pixel.RGB(0, 0, 204)
	case tetriscore.Orange:
		pt.Imd.Color = pixel.RGB(255, 128, 0)
	case tetriscore.Empty:
		if i >= 190 {
			pt.Imd.Color = pixel.RGB(0.8, 0.8, 0.8)
		} else if i >= 180 {
			pt.Imd.Color = pixel.RGB(0.6, 0.6, 0.6)
		} else {
			pt.Imd.Color = pixel.RGB(0.2, 0.2, 0.2)
		}
	default:
		pt.Imd.Color = pixel.RGB(1, 0, 1)
	}

	pt.Imd.Push(r.Min, r.Max)
	pt.Imd.Rectangle(0)
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

	var tetris PixelTetris
	tetris.T = tetriscore.CreateTetris()
	tetris.Imd = imdraw.New(nil)

	frametime, err := time.ParseDuration("16.66ms")

	var is tetriscore.InputState

	for !win.Closed() && !tetris.T.FlagLoss {
		target := time.Now().Add(frametime)

		is.Left = win.Pressed(pixelgl.KeyLeft)
		is.Right = win.Pressed(pixelgl.KeyRight)
		is.Down = win.Pressed(pixelgl.KeyDown)
		is.Up = win.Pressed(pixelgl.KeyUp) //Rotate piece
		is.Space = win.Pressed(pixelgl.KeySpace) //Instant placement
		is.Shift = win.Pressed(pixelgl.KeyLeftShift) //Hold a piece
		is.Enter = win.Pressed(pixelgl.KeyEnter) //Pause

		win.Clear(colornames.Aliceblue)
		if tetris.T.It.Enter == 0 {
			tetris.T.Update(is)
		}
		tetris.Render()
		tetris.Imd.Draw(win)
		win.Update()


		dt := target.Sub(time.Now())
		time.Sleep(dt)
	}
}

func main() {
	pixelgl.Run(run)
}
