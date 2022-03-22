package main

import (
	gameLoop "github.com/kutase/go-gameloop"
	funcs "github.com/wootra/vampire-survivors-clone/wasm/functions"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func main() {
	var data types.Data = funcs.CreateNewData()
	funcs.InitCharacters(&data)
	funcs.InitEvents(&data)

	gl := gameLoop.New(10, func(delta float64) {
		// update values
		funcs.CalculateInATick(&data)
	})

	update := gameLoop.New(16, func(delta float64) {
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
