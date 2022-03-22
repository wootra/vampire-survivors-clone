package functions

import (
	"fmt"
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace
)

func ClickByMouse(data *Data, this js.Value, args []js.Value) interface{} {
	fmt.Println("click by mouse")
	return ""
}

