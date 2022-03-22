package functions

import (
	//to fix the redline, refer .vscode/settings.json in this workspace
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	"github.com/wootra/vampire-survivors-clone/wasm/functions/characters"
)

func setCanvasFuncs(data *Data, this js.Value, args []js.Value) interface{} {
	SetCanvas(data, args[0].Int(), args[1].Int(), &args[2])
	return ""
}

func clickByMouseSync(data *Data, this js.Value, args []js.Value) interface{} {
	return ""
}

func keyDownSync(data *Data, this js.Value, args []js.Value) interface{} {
	code := args[0].String()
	KeyDown(data, code)
	return ""
}

func keyUpSync(data *Data, this js.Value, args []js.Value) interface{} {
	KeyUp(data)
	return ""
}

func InitEvents(data *Data) {
	var clickByMouse = Wrapper(data, clickByMouseSync)
	var setCanvas = Wrapper(data, setCanvasFuncs)
	var keyDown = Wrapper(data, keyDownSync)
	var keyUp = Wrapper(data, keyUpSync)

	js.Global().Set("clickByMouse", clickByMouse)
	js.Global().Set("setCanvas", setCanvas)
	js.Global().Set("keyDown", keyDown)
	js.Global().Set("keyUp", keyUp)
}

func InitCharacters(data *Data) {
	data.AddAnEnemy(characters.CreateNewEnemyData(characters.BAT))
}
