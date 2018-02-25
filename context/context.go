package context

import (
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Context struct {
	Imd        *imdraw.IMDraw
	Win        *pixelgl.Window
	stateStack []State
}

type State interface {
	Update([]bool) bool
	Render(*Context) bool
}

func (ctx *Context) PushState(s State) {
	ctx.stateStack = append(ctx.stateStack, s)
}

func (ctx *Context) PopState() {
	ctx.stateStack = ctx.stateStack[:len(ctx.stateStack)-1]
}

func (ctx *Context) Update(pressed []bool) {
	for i := len(ctx.stateStack) - 1; i >= 0; i-- {
		if !ctx.stateStack[i].Update(pressed) {
			break
		}
	}
}

func (ctx *Context) Render() {
	for _, state := range ctx.stateStack {
		if !state.Render(ctx) {
			break
		}
	}
}

func CreateContext(imd *imdraw.IMDraw, win *pixelgl.Window) *Context {
	c := Context{
		Imd:        imd,
		Win:        win,
		stateStack: make([]State, 0, 16),
	}
	return &c
}
