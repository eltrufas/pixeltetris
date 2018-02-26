package context

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"time"

	"golang.org/x/image/colornames"
)

type Context struct {
	Imd        *imdraw.IMDraw
	Win        *pixelgl.Window
	stateStack []State
	Atlas      *text.Atlas
	startTime  time.Time
	times      [120]float64
	timer      int
	avgTime    float64
}

type State interface {
	Update(*Context, []bool) bool
	Render(*Context) bool
}

func (ctx *Context) PushState(s State) {
	ctx.stateStack = append(ctx.stateStack, s)
}

func (ctx *Context) PopState() {
	ctx.stateStack = ctx.stateStack[:len(ctx.stateStack)-1]
}

func (ctx *Context) NotEmpty() bool {
	return len(ctx.stateStack) > 0
}

func (ctx *Context) Update(pressed []bool) {
	for i := len(ctx.stateStack) - 1; i >= 0; i-- {
		if !ctx.stateStack[i].Update(ctx, pressed) {
			break
		}
	}
}

func (ctx *Context) StartTimer() {
	ctx.startTime = time.Now()
}

func (ctx *Context) StopTimer() {
	ctx.times[ctx.timer] = time.Now().Sub(ctx.startTime).Seconds()

	ctx.avgTime = 0
	for _, t := range ctx.times {
		ctx.avgTime += t
	}

	ctx.avgTime /= 120
	ctx.avgTime *= 1000
}

func (ctx *Context) renderFPS() {
	txt := text.New(pixel.V(0, 470), ctx.Atlas)
	txt.Color = colornames.Red
	fmt.Fprintf(txt, "Tiempo por cuadro: %5.2vms", ctx.avgTime)
	txt.Draw(ctx.Win, pixel.IM)
}

func (ctx *Context) Render() {
	ctx.Win.Clear(colornames.Aliceblue)
	ctx.Imd = imdraw.New(nil)
	for _, state := range ctx.stateStack {
		if !state.Render(ctx) {
			break
		}
	}
	ctx.renderFPS()
	ctx.Imd.Draw(ctx.Win)
	ctx.Win.Update()
}

func CreateContext(imd *imdraw.IMDraw, win *pixelgl.Window) *Context {
	c := Context{
		Imd:        imd,
		Win:        win,
		stateStack: make([]State, 0, 16),
		Atlas:      text.NewAtlas(basicfont.Face7x13, text.ASCII),
	}
	return &c
}
