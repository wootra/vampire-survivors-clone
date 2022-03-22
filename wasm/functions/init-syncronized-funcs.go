package functions

import (
	//to fix the redline, refer .vscode/settings.json in this workspace
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	"github.com/wootra/vampire-survivors-clone/wasm/functions/characters"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func setCanvasFuncs(data *types.Data, this js.Value, args []js.Value) interface{} {
	SetCanvas(data, args[0].Int(), args[1].Int(), &args[2])
	return ""
}

func clickByMouseSync(data *types.Data, this js.Value, args []js.Value) interface{} {
	return ""
}

func keyDownSync(data *types.Data, this js.Value, args []js.Value) interface{} {
	code := args[0].String()
	KeyDown(data, code)
	return ""
}

func keyUpSync(data *types.Data, this js.Value, args []js.Value) interface{} {
	KeyUp(data)
	return ""
}

func Wrapper(data *types.Data, funcToWrap func(*types.Data, js.Value, []js.Value) interface{}) js.Func {
	ret := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return funcToWrap(data, this, args)
	})
	return ret
}

func InitEvents(data *types.Data) {
	var clickByMouse = Wrapper(data, clickByMouseSync)
	var setCanvas = Wrapper(data, setCanvasFuncs)
	var keyDown = Wrapper(data, keyDownSync)
	var keyUp = Wrapper(data, keyUpSync)

	js.Global().Set("clickByMouse", clickByMouse)
	js.Global().Set("setCanvas", setCanvas)
	js.Global().Set("keyDown", keyDown)
	js.Global().Set("keyUp", keyUp)
}

func InitCharacters(data *types.Data) {
	data.AddAnEnemy(characters.CreateNewEnemyData(types.BAT))
}
