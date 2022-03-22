package functions

import (
	"fmt"
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewCanvas() types.Canvas {
	return types.Canvas{0, 0, false, nil, types.Context2D{}, types.BitmapImage{}}
}

func DrawInCanvas(data *types.Data) interface{} {
	if data.Canvas.CanvasFuncs == nil {
		return ""
	}
	xScale := float32(data.Canvas.Width) / 100
	yScale := float32(data.Canvas.Height) / 100
	var charSize float32 = 10
	data.Canvas.GetContext().FillRect(data.Character.PosX*xScale-charSize/2, data.Character.PosY*yScale-charSize/2, charSize, charSize, 255, 0, 0, 255)
	// fmt.Println("draw in canvas", data.Character.PosX*xScale, data.Character.PosY*yScale)
	return ""
}

func SetCanvas(data *types.Data, width, height int, funcs *js.Value) {

	if data.Canvas.Width == 0 && data.Canvas.Height == 0 {
		//when window size is initialized, set character's position
		data.Character.PosX = 50
		data.Character.PosY = 50
	}
	data.Canvas.Width = width
	data.Canvas.Height = height
	data.Canvas.CanvasFuncs = funcs
	data.Canvas.Context = types.Context2D(data.Canvas.CanvasFuncs.Call("getContext"))

	fmt.Printf("%T", data.Canvas.Context)
	data.Canvas.Init = true
	data.Canvas.CanvasFuncs.Call("getBackground", "back-1")
	// data.Canvas.Background = types.BitmapImage(data.Canvas.CanvasFuncs.Call("getBackground"))
	// data.Canvas.DrawBackground()
	fmt.Println("width", data.Canvas.Width, "height", data.Canvas.Height)
}
