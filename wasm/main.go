package main

import (
	gameLoop "github.com/kutase/go-gameloop"
	funcs "github.com/wootra/vampire-survivors-clone/wasm/functions"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

var imageLoaded = false
var loadImageLoop *gameLoop.GameLoop = nil

func main() {
	var data *types.Data = funcs.CreateNewData()
	funcs.InitCharacters(data)
	funcs.InitEvents(data)

	loadImageLoop = gameLoop.New(100, func(delta float64) {
		if funcs.CheckIfImageLoaded(data) {
			imageLoaded = true
			loadImageLoop.Stop()
		}
	})
	var frame int = 0
	gl := gameLoop.New(10, func(delta float64) {
		frame++

		// update values
		if !imageLoaded {
			return
		}
		funcs.CalculateHeroPos(data.Character)
		data.Character.FrameIndex = (data.Character.FrameOffset + (frame / 10)) % 2
		for _, enemy := range data.Enemies {
			funcs.CalculateEnemyPos(data.Character, enemy)
		}
	})

	update := gameLoop.New(100, func(delta float64) {
		if !imageLoaded {
			return
		}
		// canvas update
		funcs.DrawInCanvas(data)
	})

	gl.Start()
	update.Start()
	loadImageLoop.Start()
	// Stop Game Loop:
	// gl.Stop()

	// Don't stop main goroutine
	<-make(chan bool)
}
