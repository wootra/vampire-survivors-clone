package functions

import "github.com/wootra/vampire-survivors-clone/wasm/functions/characters"

func KeyDown(data *Data, keyCode string) {

	switch keyCode {
	case "ArrowDown":
		data.Character.MovementCode = characters.DOWN
		break
	case "ArrowUp":
		data.Character.MovementCode = characters.UP
		break
	case "ArrowLeft":
		data.Character.MovementCode = characters.LEFT
		break
	case "ArrowRight":
		data.Character.MovementCode = characters.RIGHT
		break
	}
}

func KeyUp(data *Data) {
	data.Character.MovementCode = characters.STOP
}

func CalculateInATick(data *Data) interface{} {
	speedY := data.Character.Speed
	speedX := data.Character.Speed

	switch data.Character.MovementCode {
	case characters.DOWN:
		data.Character.PosY += speedY
		if data.Character.PosY > 100 {
			data.Character.PosY = 100
		}
		break
	case characters.UP:
		data.Character.PosY -= speedY
		if data.Character.PosY < 0 {
			data.Character.PosY = 0
		}
		break
	case characters.RIGHT:
		data.Character.PosX += speedX
		if data.Character.PosX > 100 {
			data.Character.PosX = 100
		}
		break
	case characters.LEFT:
		data.Character.PosX -= speedX
		if data.Character.PosX < 0 {
			data.Character.PosX = 0
		}
		break
	}
	return ""
}
