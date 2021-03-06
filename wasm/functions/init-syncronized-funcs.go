package functions

import (
	//to fix the redline, refer .vscode/settings.json in this workspace
	"fmt"
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func Wrapper(data *types.Data, funcToWrap func(*types.Data, js.Value, []js.Value) interface{}) js.Func {
	ret := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return funcToWrap(data, this, args)
	})
	return ret
}

func setCanvasFuncs(data *types.Data, this js.Value, args []js.Value) interface{} {
	SetCanvas(data, args[0].Int(), args[1].Int(), &args[2])
	return ""
}

func setGlueFunctions(data *types.Data, this js.Value, args []js.Value) interface{} {
	fmt.Printf("set glue function")
	data.GlueFunctions = &args[0]
	return ""
}

func clickByMouseSync(data *types.Data, this js.Value, args []js.Value) interface{} {
	return ""
}

func keyDownSync(data *types.Data, this js.Value, args []js.Value) interface{} {
	KeyDown(data, args[0].String())
	return ""
}

func keyUpSync(data *types.Data, this js.Value, args []js.Value) interface{} {
	KeyUp(data, args[0].String())
	return ""
}

func setBackgroundSync(data *types.Data, this js.Value, args []js.Value) interface{} {
	data.Canvas.Background = types.BitmapImage(args[0].Get("image"))
	// data.Canvas.Save()
	return ""
}

func InitEvents(data *types.Data) {

	js.Global().Set("clickByMouse", Wrapper(data, clickByMouseSync))
	js.Global().Set("setCanvas", Wrapper(data, setCanvasFuncs))
	js.Global().Set("setGlueFunctions", Wrapper(data, setGlueFunctions))
	js.Global().Set("keyDown", Wrapper(data, keyDownSync))
	js.Global().Set("keyUp", Wrapper(data, keyUpSync))
	js.Global().Set("setBackground", Wrapper(data, setBackgroundSync))
}
