package functions

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

func (c Context2D) clearRect() {
	js.Value(c).Call("clearRect")
}

func (c Context2D) fillRect(x, y, w, h float32, r, g, b, a uint8) {
	js.Value(c).Set("fillStyle", fmt.Sprintf("rgb(%d,%d,%d,%d)", r, g, b, a))
	js.Value(c).Call("fillRect", x, y, w, h)
}

func (c Canvas) GetContext() Context2D {
	if c.CanvasFuncs != nil {
		return Context2D(c.Context)
	}
	return Context2D{}
}

func CreateNewCanvas() Canvas {
	return Canvas{0, 0, false, nil, Context2D{}, BitmapImage{}}
}

func DrawInCanvas(data *Data) interface{} {
	if data.Canvas.CanvasFuncs == nil {
		return ""
	}
	xScale := float32(data.Canvas.Width) / 100
	yScale := float32(data.Canvas.Height) / 100
	var charSize float32 = 10
	data.Canvas.GetContext().fillRect(data.Character.PosX*xScale-charSize/2, data.Character.PosY*yScale-charSize/2, charSize, charSize, 255, 0, 0, 255)
	// fmt.Println("draw in canvas", data.Character.PosX*xScale, data.Character.PosY*yScale)
	return ""
}

func SetCanvas(data *Data, width, height int, funcs *js.Value) {

	if data.Canvas.Width == 0 && data.Canvas.Height == 0 {
		//when window size is initialized, set character's position
		data.Character.PosX = 50
		data.Character.PosY = 50
	}
	data.Canvas.Width = width
	data.Canvas.Height = height
	data.Canvas.CanvasFuncs = funcs
	data.Canvas.Context = Context2D(data.Canvas.CanvasFuncs.Call("getContext"))

	fmt.Printf("%T", data.Canvas.Context)
	data.Canvas.Init = true
	data.Canvas.Background = BitmapImage(data.Canvas.CanvasFuncs.Call("getBackground"))

	fmt.Println("width", data.Canvas.Width, "height", data.Canvas.Height)
}
