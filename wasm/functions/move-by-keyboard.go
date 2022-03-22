package functions

import (
	"fmt"
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace
)

func KeyDown(data *Data, this js.Value, args []js.Value) interface{} {
	code := args[0].String()

	switch code {
	case "ArrowDown":
		data.character.movementCode = DOWN
		break
	case "ArrowUp":
		data.character.movementCode = UP
		break
	case "ArrowLeft":
		data.character.movementCode = LEFT
		break
	case "ArrowRight":
		data.character.movementCode = RIGHT
		break
	}
	fmt.Println("key down")
	return ""
}

func KeyUp(data *Data, this js.Value, args []js.Value) interface{} {
	fmt.Println("key up")
	data.character.movementCode = STOP
	return ""
}
