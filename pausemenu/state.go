package pausemenu

import (
	"github.com/eltrufas/pixeltetris/context"
	"github.com/eltrufas/pixeltetris/input"
)

type State struct {
}

func (s *State) Update(ctx *context.Context, pressed []bool) bool {
	if pressed[input.Enter] {
		ctx.PopState()
	}

	return false
}

func (s *State) Render(ctx *context.Context) bool {
	return true
}
