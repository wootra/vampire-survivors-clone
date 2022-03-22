package functions

import (
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	"github.com/wootra/vampire-survivors-clone/wasm/functions/characters"
)

func KeyDown(data *Data, this js.Value, args []js.Value) interface{} {
	code := args[0].String()

	switch code {
	case "ArrowDown":
		data.character.MovementCode = characters.DOWN
		break
	case "ArrowUp":
		data.character.MovementCode = characters.UP
		break
	case "ArrowLeft":
		data.character.MovementCode = characters.LEFT
		break
	case "ArrowRight":
		data.character.MovementCode = characters.RIGHT
		break
	}
	// fmt.Println("key down")
	return ""
}

func KeyUp(data *Data, this js.Value, args []js.Value) interface{} {
	// fmt.Println("key up")
	data.character.MovementCode = characters.STOP
	return ""
}
