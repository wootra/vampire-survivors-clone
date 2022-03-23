package functions

import (
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewCanvas() *types.Canvas {
	return &types.Canvas{0, 0, false, nil, types.Context2D{}, types.BitmapImage{}}
}

func DrawInCanvas(data *types.Data) interface{} {
	if data.Canvas.CanvasFuncs == nil {
		return ""
	}
	data.Canvas.Restore()
	xScale := float32(data.Canvas.Width) / 100
	yScale := float32(data.Canvas.Height) / 100
	var charSize float32 = 10
	data.Canvas.GetContext().FillRect(data.Character.PosX*xScale-charSize/2, data.Character.PosY*yScale-charSize/2, charSize, charSize, 255, 0, 0, 255)

	for _, enemy := range data.Enemies {
		if enemy.Status == types.MOVED {
			data.Canvas.GetContext().FillRect(enemy.PosX*xScale-charSize/2, enemy.PosY*yScale-charSize/2, charSize, charSize, 0, 255, 0, 255)
		} else if enemy.Status == types.IDLE {
			data.Canvas.GetContext().FillRect(enemy.PosX*xScale-charSize/2, enemy.PosY*yScale-charSize/2, charSize, charSize, 255, 255, 0, 255)
		}
	}

	// fmt.Println("draw in canvas", data.Character.PosX*xScale, data.Character.PosY*yScale)
	return ""
}

func SetCanvas(data *types.Data, width, height int, funcs *js.Value) {

	data.Canvas.Width = width
	data.Canvas.Height = height
	data.Canvas.CanvasFuncs = funcs
	data.Canvas.Context = types.Context2D(data.Canvas.CanvasFuncs.Call("getContext"))
	data.Canvas.Init = true
	data.Canvas.CanvasFuncs.Call("getBackground", "back-1")
}
