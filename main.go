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
	BoardOffsetX = 100
	BoardOffsetY = 470
	BlockW       = 16
)

type PixelTetris struct {
	T   *tetriscore.Tetris
	Imd *imdraw.IMDraw
}

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

func (pt *PixelTetris) Render() {
	pt.Imd.Clear()
	for i := 0; i < 200; i++ {
		value := pt.T.Board[i+20]

		pt.RenderBlock(i, value)
	}

	pt.RenderCurrentPiece()

	pt.RenderNext(
		BoardOffsetX-BlockW*4-10,
		BoardOffsetY-BlockW*4,
		pt.T.HoldPiece,
	)

	for i := 0; i < 6; i++ {
		pt.RenderNext(
			BoardOffsetX+BlockW*10+10,
			BoardOffsetY-(4+i*5)*BlockW,
			pt.T.PieceQueue[(pt.T.NextIndex+i)%14],
		)
	}
}

func (pt *PixelTetris) RenderCurrentPiece() {
	p := pt.T.CurrentPiece
	mask := tetriscore.Tetrominos[p.TetrominoType][p.State]
	color := tetriscore.TetrominoColors[p.TetrominoType]

	for i := 0; i < 16; i++ {
		if mask[i] == 1 {
			y := p.Y + i/4 - 2
			x := p.X + i%4

			if y >= 0 {
				pt.RenderBlock(x+y*10, color)
			}
		}
	}
}

func (pt *PixelTetris) RenderNext(x, y, piece int) {
	r := pixel.R(
		float64(x),
		float64(y),
		float64(x+BlockW*4),
		float64(y+BlockW*4),
	)
	pt.Imd.Color = pixel.RGB(0.2, 0.2, 0.2)

	pt.Imd.Push(r.Min, r.Max)
	pt.Imd.Rectangle(0)

	if piece < 0 {
		return
	}

	mask := tetriscore.Tetrominos[piece][1]
	pt.Imd.Color = getColor(tetriscore.TetrominoColors[piece])

	for i := 0; i < 16; i++ {
		if mask[i] == 1 {
			py := y + (3-i/4)*BlockW
			px := x + (i%4)*BlockW

			r = pixel.R(
				float64(px),
				float64(py),
				float64(px+BlockW),
				float64(py+BlockW),
			)

			pt.Imd.Push(r.Min, r.Max)
			pt.Imd.Rectangle(0)
		}
	}
}

func getColor(block tetriscore.Block) pixel.RGBA {
	switch block {
	case tetriscore.Yellow:
		return pixel.RGB(1, 0.93, 0.36)
	case tetriscore.Red:
		return pixel.RGB(255, 0, 0)
	case tetriscore.Cyan:
		return pixel.RGB(0, 255, 255)
	case tetriscore.Green:
		return pixel.RGB(0, 255, 0)
	case tetriscore.Purple:
		return pixel.RGB(204, 0, 204)
	case tetriscore.Blue:
		return pixel.RGB(0, 0, 204)
	case tetriscore.Orange:
		return pixel.RGB(255, 128, 0)
	case tetriscore.Empty:
		return pixel.RGB(0.2, 0.2, 0.2)
	default:
		return pixel.RGB(1, 0, 1)
	}
}

func (pt *PixelTetris) RenderBlock(i int, block tetriscore.Block) {
	x := i % 10
	y := i / 10
	r := pixel.R(
		float64(BoardOffsetX+BlockW*x),
		float64(BoardOffsetY-BlockW*y),
		float64(BoardOffsetX+BlockW*(x+1)),
		float64(BoardOffsetY-BlockW*(y+1)),
	)

	pt.Imd.Color = getColor(block)

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

	ia := initInputArray()

	for !win.Closed() && !tetris.T.FlagLoss {
		target := time.Now().Add(frametime)

		var is tetriscore.InputState
		for i, input := range ia {
			if win.Pressed(input) {
				is |= 1 << uint32(i)
			}
		}

		win.Clear(colornames.Aliceblue)
		tetris.T.Update(is)
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
