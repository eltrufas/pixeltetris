package game

import (
	"github.com/eltrufas/pixeltetris/context"
	"github.com/eltrufas/pixeltetris/input"
	"github.com/eltrufas/pixeltetris/pausemenu"
	"github.com/eltrufas/tetriscore"
  "github.com/eltrufas/rltetris"
  "fmt"
)

func (s *State) Update(ctx *context.Context, ia []bool) bool {
	if ia[input.Escape] {
		ctx.PushState(&pausemenu.State{})
	}
	var is tetriscore.InputState
  is = tetriscore.InputState(rltetris.GetEGreedyAction(s.Weights, s.Game.GetState(), s.Game.LegalAction()))

  s.Action = uint32(is)

	if s.Game.Tetris.FlagLoss {
		s.Game = rltetris.CreateTetris()
    s.T = s.Game.Tetris
    fmt.Println(s.Weights)
    rltetris.Sarsa(s.Weights, 20000, 0.00001, 0.5)
	}


  s.ShouldPress = !s.ShouldPress

	s.T.Update(is)
	return true
}
