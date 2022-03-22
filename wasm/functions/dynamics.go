package functions

import (
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func KeyDown(data *types.Data, keyCode string) {

	switch keyCode {
	case "ArrowDown":
		data.Character.MovementCode = types.DOWN
		break
	case "ArrowUp":
		data.Character.MovementCode = types.UP
		break
	case "ArrowLeft":
		data.Character.MovementCode = types.LEFT
		break
	case "ArrowRight":
		data.Character.MovementCode = types.RIGHT
		break
	}
}

func KeyUp(data *types.Data) {
	data.Character.MovementCode = types.STOP
}

func CalculateInATick(data *types.Data) interface{} {
	speedY := data.Character.Speed
	speedX := data.Character.Speed

	switch data.Character.MovementCode {
	case types.DOWN:
		data.Character.PosY += speedY
		if data.Character.PosY > 100 {
			data.Character.PosY = 100
		}
		break
	case types.UP:
		data.Character.PosY -= speedY
		if data.Character.PosY < 0 {
			data.Character.PosY = 0
		}
		break
	case types.RIGHT:
		data.Character.PosX += speedX
		if data.Character.PosX > 100 {
			data.Character.PosX = 100
		}
		break
	case types.LEFT:
		data.Character.PosX -= speedX
		if data.Character.PosX < 0 {
			data.Character.PosX = 0
		}
		break
	}
	return ""
}
