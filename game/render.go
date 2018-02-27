package game

import (
	"github.com/eltrufas/pixeltetris/context"
	"github.com/eltrufas/tetriscore"
	"github.com/faiface/pixel"
  "fmt"
  "golang.org/x/image/colornames"
	"github.com/faiface/pixel/text"
)

func (s *State) Render(ctx *context.Context) bool {
	for i := 0; i < 200; i++ {
		color := getColor(s.T.Board[i+20])
		s.RenderBlock(ctx, i, color)
	}

	s.RenderPiece(ctx, s.T.CurrentPiece, 1)

	s.RenderPiece(ctx, s.T.GhostPiece(), 0.5)

  s.RenderScore(ctx)

	s.RenderNext(
		ctx,
		s.OffsetX-s.BlockW*4-10,
		s.OffsetY-s.BlockW*4,
		s.T.HoldPiece,
	)

	for i := 0; i < 6; i++ {
		s.RenderNext(
			ctx,
			s.OffsetX+s.BlockW*10+10,
			s.OffsetY-(4+i*5)*s.BlockW,
			s.T.PieceQueue[(s.T.NextIndex+i)%14],
		)
	}
	return true
}

func getColor(block tetriscore.Block) pixel.RGBA {
	switch block {
	case tetriscore.Yellow:
		return pixel.RGB(1, 0.93, 0.36)
	case tetriscore.Red:
		return pixel.RGB(1, 0.2, 0.2)
	case tetriscore.Cyan:
		return pixel.RGB(0, 1, 1)
	case tetriscore.Green:
		return pixel.RGB(0, 1, 0)
	case tetriscore.Purple:
		return pixel.RGB(0.75, 0.35, 0.75)
	case tetriscore.Blue:
		return pixel.RGB(0, 0, 0.8)
	case tetriscore.Orange:
		return pixel.RGB(1, 0.5, 0)
	case tetriscore.Empty:
		return pixel.RGB(0.2, 0.2, 0.2)
	default:
		return pixel.RGB(1, 0, 1)
	}
}

func (s *State) RenderPiece(ctx *context.Context, p tetriscore.Piece, alpha float64) {
	mask := tetriscore.Tetrominos[p.TetrominoType][p.State]
	color := getColor(tetriscore.TetrominoColors[p.TetrominoType])
	color = pixel.Alpha(alpha).Mul(color)

	for i := 0; i < 16; i++ {
		if mask[i] == 1 {
			y := p.Y + i/4 - 2
			x := p.X + i%4

			if y >= 0 {
				s.RenderBlock(ctx, x+y*10, color)
			}
		}
	}
}

func (s *State) RenderNext(ctx *context.Context, x, y, piece int) {
	r := pixel.R(
		float64(x),
		float64(y),
		float64(x+s.BlockW*4),
		float64(y+s.BlockW*4),
	)
	ctx.Imd.Color = pixel.RGB(0.2, 0.2, 0.2)

	ctx.Imd.Push(r.Min, r.Max)
	ctx.Imd.Rectangle(0)

	if piece < 0 {
		return
	}

	mask := tetriscore.Tetrominos[piece][1]
	ctx.Imd.Color = getColor(tetriscore.TetrominoColors[piece])

	for i := 0; i < 16; i++ {
		if mask[i] == 1 {
			py := y + (3-i/4)*s.BlockW
			px := x + (i%4)*s.BlockW

			r = pixel.R(
				float64(px),
				float64(py),
				float64(px+s.BlockW),
				float64(py+s.BlockW),
			)

			ctx.Imd.Push(r.Min, r.Max)
			ctx.Imd.Rectangle(0)
		}
	}
}

func (s *State) RenderBlock(ctx *context.Context, i int, color pixel.RGBA) {
	x := i % 10
	y := i / 10
	r := pixel.R(
		float64(s.OffsetX+s.BlockW*x),
		float64(s.OffsetY-s.BlockW*y),
		float64(s.OffsetX+s.BlockW*(x+1)),
		float64(s.OffsetY-s.BlockW*(y+1)),
	)

	ctx.Imd.Color = color
	ctx.Imd.Push(r.Min, r.Max)
	ctx.Imd.Rectangle(0)
}

func (s *State) RenderScore(ctx *context.Context){
  txt := text.New(pixel.V(float64(s.OffsetX-s.BlockW*4-10), float64(s.OffsetY-s.BlockW*16)), ctx.Atlas)
  txt.Color = colornames.Red
  fmt.Fprintln(txt, "Score: ")
  fmt.Fprintln(txt, s.T.Score)
  fmt.Fprintln(txt, "Level: ", s.T.Level)
  txt.Draw(ctx.Win, pixel.IM)
}
