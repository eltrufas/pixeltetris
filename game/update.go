package game

import (
	"github.com/eltrufas/pixeltetris/context"
	"github.com/eltrufas/pixeltetris/input"
	"github.com/eltrufas/pixeltetris/pausemenu"
	"github.com/eltrufas/tetriscore"
	"github.com/faiface/pixel/pixelgl"
)

var actions []uint32 = []uint32{0, 1, 2, 4, 8, 16, 32, 64}

func (s *State) Update(ctx *context.Context, ia []bool) bool {
	if ia[input.Escape] {
		ctx.PushState(&pausemenu.State{})
	}

	if s.T.FlagLoss {
		s.T = tetriscore.CreateTetris()
		s.lastAction = s.player.GetAction(s.T)

	}

	if s.T.Update(s.lastAction) {
		s.lastAction = s.player.GetAction(s.T)
	}
	return true
}
