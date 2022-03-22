package main

import (
	"fmt"        //to fix the redline, refer .vscode/settings.json in this workspace
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	gameLoop "github.com/kutase/go-gameloop"
	funcs "github.com/wootra/vampire-survivors-clone/wasm/functions"
)

var data funcs.Data = funcs.CreateNewData()

func main() {

	var clickByMouse = Wrapper(&data, funcs.ClickByMouse)
	var setCanvas = Wrapper(&data, funcs.SetCanvasFuncs)
	var keyDown = Wrapper(&data, funcs.KeyDown)
	var keyUp = Wrapper(&data, funcs.KeyUp)

	fmt.Println("web-asm-sample")

	js.Global().Set("clickByMouse", clickByMouse)
	js.Global().Set("setCanvas", setCanvas)
	js.Global().Set("keyDown", keyDown)
	js.Global().Set("keyUp", keyUp)

	gl := gameLoop.New(10, func(delta float64) {
		// update values
		funcs.CalculateInATick(&data)
	})

	update := gameLoop.New(100, func(delta float64) {
		// canvas update
		funcs.DrawInCanvas(&data)
	})

	gl.Start()
	update.Start()

	// Stop Game Loop:
	// gl.Stop()

	// Don't stop main goroutine
	<-make(chan bool)
}
