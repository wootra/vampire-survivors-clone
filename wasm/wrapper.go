package main

import (
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	funcs "github.com/wootra/vampire-survivors-clone/wasm/functions"
)

func Wrapper(data *funcs.Data, funcToWrap func(*funcs.Data, js.Value, []js.Value) interface{}) js.Func {
	ret := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return funcToWrap(data, this, args)
	})
	return ret
}
