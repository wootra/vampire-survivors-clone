package types

import (
	"fmt"
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

func (c Context2D) ClearRect() {
	js.Value(c).Call("clearRect")
}

func (c Context2D) FillRect(x, y, w, h float32, r, g, b, a uint8) {
	js.Value(c).Set("fillStyle", fmt.Sprintf("rgb(%d,%d,%d,%d)", r, g, b, a))
	js.Value(c).Call("fillRect", x, y, w, h)
}

func (c Context2D) Translate(dx, dy int) {
	js.Value(c).Call("translate", dx, dy)
}

func (c *Canvas) DrawBackground() {
	if !js.Value(c.Background).IsUndefined() {
		js.Value(c.Context).Call("drawImage", c.Background, 20, 20, 185, 175, 50, 50, 185, 175)
	} else {
		fmt.Println("background image is undefined")
	}
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
