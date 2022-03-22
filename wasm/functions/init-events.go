package functions

import (
	"fmt"        //to fix the redline, refer .vscode/settings.json in this workspace
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	"github.com/wootra/vampire-survivors-clone/wasm/functions/characters"
)

func InitEvents(data *Data) {
	var clickByMouse = Wrapper(data, ClickByMouse)
	var setCanvas = Wrapper(data, SetCanvasFuncs)
	var keyDown = Wrapper(data, KeyDown)
	var keyUp = Wrapper(data, KeyUp)

	fmt.Println("web-asm-sample")

	js.Global().Set("clickByMouse", clickByMouse)
	js.Global().Set("setCanvas", setCanvas)
	js.Global().Set("keyDown", keyDown)
	js.Global().Set("keyUp", keyUp)
}

func InitCharacters(data *Data) {
	data.AddAnEnemy(characters.CreateNewEnemyData(characters.BAT))
}
