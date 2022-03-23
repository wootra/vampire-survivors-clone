package functions

import (
	//to fix the redline, refer .vscode/settings.json in this workspace
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	"github.com/wootra/vampire-survivors-clone/wasm/functions/characters"
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

func setBackgroundSync(data *types.Data, this js.Value, args []js.Value) interface{} {
	data.Canvas.Background = types.BitmapImage(args[0].Get("image"))
	data.Canvas.Save()
	return ""
}

func InitEvents(data *types.Data) {

	js.Global().Set("clickByMouse", Wrapper(data, clickByMouseSync))
	js.Global().Set("setCanvas", Wrapper(data, setCanvasFuncs))
	js.Global().Set("keyDown", Wrapper(data, keyDownSync))
	js.Global().Set("keyUp", Wrapper(data, keyUpSync))
	js.Global().Set("setBackground", Wrapper(data, setBackgroundSync))
}

func InitCharacters(data *types.Data) {
	data.Character = characters.CreateNewCharacterData()
	data.AddAnEnemy(characters.CreateNewEnemyData(data, types.BAT))
}
