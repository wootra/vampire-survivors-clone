package types

import (
	"fmt"
	"math"
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace
)

type Context2D js.Value

type BitmapImage js.Value

type Canvas struct {
	Width       int
	Height      int
	Init        bool
	CanvasFuncs *js.Value
	Context     Context2D
	Background  BitmapImage
}

func (c Canvas) GetObjectScale() float64 {
	xScale := float64(c.Width) / 100
	yScale := float64(c.Height) / 100

	if xScale < yScale { return xScale } else { return yScale }
}

func (c Context2D) ClearRect() {
	js.Value(c).Call("clearRect")
}

func (c Context2D) FillRect(x, y, w, h float64, r, g, b, a uint8) {
	js.Value(c).Set("fillStyle", fmt.Sprintf("rgb(%d,%d,%d,%d)", r, g, b, a))
	js.Value(c).Call("fillRect", float32(x), float32(y), float32(w), float32(h))
}

func (c Context2D) StrokeRect(x, y, w, h float64, r, g, b, a uint8) {
	js.Value(c).Set("strokeStyle", fmt.Sprintf("rgb(%d,%d,%d,%d)", r, g, b, a))
	js.Value(c).Call("strokeRect", float32(x), float32(y), float32(w), float32(h))
}

func (c Context2D) BeginPath(){
	js.Value(c).Call("beginPath")
}

func (c Context2D) Stroke(){
	js.Value(c).Call("stroke")
}


func (c Context2D) StrokeEllipse(x, y, w, h, rotation, startAngle, endAngle float64, r, g, b, a uint8) {
	js.Value(c).Set("strokeStyle", fmt.Sprintf("rgb(%d,%d,%d,%d)", r, g, b, a))
	c.BeginPath()
	js.Value(c).Call("ellipse", float32(x), float32(y), float32(w), float32(h), float32(rotation), float32(startAngle), float32(endAngle))
	c.Stroke()
}


func (c Context2D) StrokeCircle(x, y, radious float64, r, g, b, a uint8) {
	js.Value(c).Set("strokeStyle", fmt.Sprintf("rgb(%d,%d,%d,%d)", r, g, b, a))
	c.BeginPath()
	js.Value(c).Call("ellipse", float32(x), float32(y), float32(radious), float32(radious), float32(0), float32(0), float32(math.Pi*2))
	c.Stroke()
}


func (c Context2D) Translate(dx, dy int) {
	js.Value(c).Call("translate", dx, dy)
}

func (c Context2D) Scale(dx, dy float64) {
	js.Value(c).Call("scale", float32(dx), float32(dy))
}

func (ctx Context2D) Transform(a, b, c, d, e, f float64) {
	js.Value(ctx).Call("transform", float32(a), float32(b), float32(c), float32(d), float32(e), float32(f))
}

func (ctx Context2D) Rotate(angle float64) {
	js.Value(ctx).Call("rotate", float32(angle))
}

func (c Canvas) Restore() {
	js.Value(c.Context).Call("restore")
}

func (c Canvas) Save() {
	js.Value(c.Context).Call("save")
}

func (c Canvas) GetContext() Context2D {
	if c.CanvasFuncs != nil {
		return Context2D(c.Context)
	}
	return Context2D{}
}
