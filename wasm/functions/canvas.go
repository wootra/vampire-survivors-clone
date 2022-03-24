package functions

import (
	"fmt"
	"strconv"
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	"github.com/wootra/vampire-survivors-clone/wasm/functions/characters"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewCanvas() *types.Canvas {
	return &types.Canvas{0, 0, false, nil, types.Context2D{}, types.BitmapImage{}}
}

func CheckIfImageLoaded(data *types.Data) bool {
	if data.GlueFunctions == nil {
		return false
	}
	total := data.GlueFunctions.Call("getLoadStatus").Get("total").Int()
	loaded := data.GlueFunctions.Call("getLoadStatus").Get("loaded").Int()
	if loaded != total {
		fmt.Println("Image is loading... (" + strconv.Itoa(loaded) + "/" + strconv.Itoa(total) + ")")
		return false
	}
	return true
}

func DrawInCanvas(data *types.Data) {
	if data.Canvas.CanvasFuncs == nil {
		return
	}

	data.Canvas.CanvasFuncs.Call("getBackground", "back-1")

	xScale := float32(data.Canvas.Width) / 100
	yScale := float32(data.Canvas.Height) / 100
	var charSize float32 = 10 * xScale

	// data.Canvas.GetContext().FillRect(data.Character.PosX*xScale-charSize/2, data.Character.PosY*yScale-charSize/2, charSize, charSize, 255, 0, 0, 255)
	characters.DrawCharacter(data, xScale, yScale, charSize)

	for enemyNo := range data.Enemies {
		characters.DrawEnemy(data, enemyNo, xScale, yScale, charSize)
		// if enemy.Status == types.MOVED {
		// 	data.Canvas.CanvasFuncs.Call("getCharacterImage", "cat", enemy.FrameIndex, enemy.PosX*xScale-charSize/2, enemy.PosY*yScale-charSize/2, charSize, charSize)
		// } else if enemy.Status == types.IDLE {
		// 	data.Canvas.CanvasFuncs.Call("getCharacterImage", "cat", enemy.FrameIndex, enemy.PosX*xScale-charSize/2, enemy.PosY*yScale-charSize/2, charSize, charSize)
		// }
	}

	// fmt.Println("draw in canvas", data.Character.PosX*xScale, data.Character.PosY*yScale)
}

func SetCanvas(data *types.Data, width, height int, funcs *js.Value) {

	data.Canvas.Width = width
	data.Canvas.Height = height
	data.Canvas.CanvasFuncs = funcs
	data.Canvas.Context = types.Context2D(data.Canvas.CanvasFuncs.Call("getContext"))
	data.Canvas.Init = true
	data.Canvas.CanvasFuncs.Call("getBackground", "back-1")
}
