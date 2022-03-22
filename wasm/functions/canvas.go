package functions

import (
	"fmt"
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace
)

type Context2D interface {
	clearRect()
	fillRect()
}

type Context2DJsValue js.Value

type Canvas struct {
	width       int
	height      int
	init        bool
	canvasFuncs *js.Value
	context     Context2DJsValue
}

func (c Context2DJsValue) clearRect() {
	js.Value(c).Call("clearRect")
}

func (c Context2DJsValue) fillRect(x, y, w, h float32, r, g, b, a uint8) {
	js.Value(c).Set("fillStyle", fmt.Sprintf("rgb(%d,%d,%d,%d)", r, g, b, a))
	js.Value(c).Call("fillRect", x, y, w, h)
}

func (c Canvas) GetContext() Context2DJsValue {
	if c.canvasFuncs != nil {
		return c.context
	}
	return Context2DJsValue{}
}

func CreateNewCanvas() Canvas {
	return Canvas{0, 0, false, nil, Context2DJsValue{}}
}

func SetCanvasFuncs(data *Data, this js.Value, args []js.Value) interface{} {

	if data.canvas.width == 0 && data.canvas.height == 0 {
		//when window size is initialized, set character's position
		data.character.PosX = 50
		data.character.PosY = 50
	}
	data.canvas.width = args[0].Int()
	data.canvas.height = args[1].Int()
	data.canvas.canvasFuncs = &args[2]
	data.canvas.context = Context2DJsValue(data.canvas.canvasFuncs.Call("getContext"))
	fmt.Printf("%T", data.canvas.context)
	data.canvas.init = true

	fmt.Println("width", data.canvas.width, "height", data.canvas.height)
	return ""
}

func DrawInCanvas(data *Data) interface{} {
	if data.canvas.canvasFuncs == nil {
		return ""
	}
	xScale := float32(data.canvas.width) / 100
	yScale := float32(data.canvas.height) / 100
	var charSize float32 = 10
	data.canvas.GetContext().fillRect(data.character.PosX*xScale-charSize/2, data.character.PosY*yScale-charSize/2, charSize, charSize, 255, 0, 0, 255)
	// fmt.Println("draw in canvas", data.character.PosX*xScale, data.character.PosY*yScale)
	return ""
}
