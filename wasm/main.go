package main

import (
	"math/rand"

	gameLoop "github.com/kutase/go-gameloop"
	funcs "github.com/wootra/vampire-survivors-clone/wasm/functions"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

var imageLoaded = false
var loadImageLoop *gameLoop.GameLoop = nil
var bufferHandlerLoop *gameLoop.GameLoop = nil
var gl *gameLoop.GameLoop = nil
var update *gameLoop.GameLoop = nil 

func initWorld() *types.World {
	world := &types.World{Active: 0, Apocalypse: false}
	for i := range world.Pt {
		world.Pt[i]= nil
	}
	world.BufferToDelete = make(chan int, 10)
	world.MainChannel = make(chan bool)
	return world
}

func main() {
	var data *types.Data = funcs.CreateNewData()

	var world *types.World = initWorld()

	rand.Seed(100)
	funcs.InitCharacters(data)
	funcs.InitEvents(data)

	loadImageLoop = gameLoop.New(100, func(delta float64) {
		if funcs.CheckIfImageLoaded(data) {
			imageLoaded = true
			loadImageLoop.Stop()

		}
	})
	var frame int = 0
	gl = gameLoop.New(10, func(delta float64) {
		frame++

		// update values
		if !imageLoaded {
			return
		}
		funcs.CalculateHeroPos(data, world, world.Active)
		data.Character.CharInfo.FrameIndex = (data.Character.CharInfo.FrameOffset + (frame / 10)) % 2
		for _, enemy := range data.Enemies {
			funcs.CalculateEnemyPos(data, enemy, world, world.Active)
		}
		world.Active = (world.Active + 1) % 10
		world.BufferToDelete <- (world.Active + 10 - 2) % 10 // delete unused buffer
	})

	update = gameLoop.New(200, func(delta float64) {
		if !imageLoaded {
			return
		}
		// canvas update
		funcs.DrawInCanvas(data)
	})

	bufferHandlerLoop = gameLoop.New(100, func(delta float64) {
		if world.Apocalypse {
			bufferHandlerLoop.Stop()
			gl.Stop()
			update.Stop()
			return
		}
		buf := <-world.BufferToDelete
		for xKey,map1:= range world.Pt[buf]{
			for yKey:=range map1 {
				map1[yKey] = nil
			}
			map1[xKey] = nil
		}
		world.Pt[buf] = nil
	})

	gl.Start()
	update.Start()
	loadImageLoop.Start()
	bufferHandlerLoop.Start()
	// Stop Game Loop:
	// gl.Stop()

	// Don't stop main goroutine
	<-world.MainChannel
}
