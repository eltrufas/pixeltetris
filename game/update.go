package game

import (
  "github.com/faiface/pixel/pixelgl"
	"github.com/eltrufas/pixeltetris/context"
	"github.com/eltrufas/pixeltetris/input"
	"github.com/eltrufas/pixeltetris/pausemenu"
	"github.com/eltrufas/tetriscore"
  "github.com/eltrufas/rltetris"
  "fmt"
)

var actions []uint32 = []uint32{0, 1, 2, 4, 8, 16, 32, 64}

func (s *State) Update(ctx *context.Context, ia []bool) bool {
	if ia[input.Escape] {
		ctx.PushState(&pausemenu.State{})
	}
  if ctx.Win.Pressed(pixelgl.KeyR) {
    s.Game = rltetris.CreateTetris()
    s.T = s.Game.Tetris
    fmt.Println(s.Weights)
    rltetris.Sarsa(s.Weights, 2000, 0.001, 1)
  }
	var is tetriscore.InputState
  is = tetriscore.InputState(rltetris.GetGreedyAction(s.Weights, s.Game.GetState(), s.Game.LegalAction()))

  s.Action = uint32(is)
  for i, a := range actions {
    s.Q[i] = rltetris.Q(s.Game.GetState(), a, s.Weights[a])
  }

	if s.Game.Tetris.FlagLoss {
		s.Game = rltetris.CreateTetris()
    s.T = s.Game.Tetris
    fmt.Println(s.Weights)
    rltetris.Sarsa(s.Weights, 2000, 0.001, 1)
	}


  s.ShouldPress = !s.ShouldPress

	s.T.Update(is)
	return true
}
