package functions

import "github.com/wootra/vampire-survivors-clone/wasm/functions/characters"

func CalculateInATick(data *Data) interface{} {
	speedY := data.character.Speed
	speedX := data.character.Speed

	switch data.character.MovementCode {
	case characters.DOWN:
		data.character.PosY += speedY
		if data.character.PosY > 100 {
			data.character.PosY = 100
		}
		break
	case characters.UP:
		data.character.PosY -= speedY
		if data.character.PosY < 0 {
			data.character.PosY = 0
		}
		break
	case characters.RIGHT:
		data.character.PosX += speedX
		if data.character.PosX > 100 {
			data.character.PosX = 100
		}
		break
	case characters.LEFT:
		data.character.PosX -= speedX
		if data.character.PosX < 0 {
			data.character.PosX = 0
		}
		break
	}
	return ""
}
