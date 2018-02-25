package game

import (
	"github.com/eltrufas/pixeltetris/context"
	"github.com/eltrufas/pixeltetris/input"
	"github.com/eltrufas/pixeltetris/pausemenu"
	"github.com/eltrufas/tetriscore"
)

func (s *State) Update(ctx *context.Context, ia []bool) bool {
	if ia[input.Escape] {
		ctx.PushState(&pausemenu.State{})
	}
	var is tetriscore.InputState
	for i, input := range ia {
		if input {
			is |= 1 << uint32(i)
		}
	}

	if s.T.FlagLoss {
		ctx.PopState()
	}

	s.T.Update(is)
	return true
}
