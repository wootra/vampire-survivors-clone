package functions

import (
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace
)

func Wrapper(data *Data, funcToWrap func(*Data, js.Value, []js.Value) interface{}) js.Func {
	ret := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return funcToWrap(data, this, args)
	})
	return ret
}
