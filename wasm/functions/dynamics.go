package functions

import (
	"math"

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
	data.Character.FrameOffset++
}

func KeyUp(data *types.Data) {
	data.Character.MovementCode = types.STOP
}

func CalculateHeroPos(character *types.CharacterData) {
	speedY := character.Speed
	speedX := character.Speed

	switch character.MovementCode {
	case types.DOWN:
		character.PosY += speedY
		if character.PosY > 50 {
			character.PosY = 50
		}
		break
	case types.UP:
		character.PosY -= speedY
		if character.PosY < -50 {
			character.PosY = -50
		}
		break
	case types.RIGHT:
		character.PosX += speedX
		if character.PosX > 50 {
			character.PosX = 50
		}
		break
	case types.LEFT:
		character.PosX -= speedX
		if character.PosX < -50 {
			character.PosX = -50
		}
		break
	}

	// fmt.Println("hero movement:", character.PosX, character.PosY, speedX, speedY)
}

func CalculateEnemyPos(character *types.CharacterData, enemy *types.EnemyData) {

	dirX := float64(character.PosX - enemy.PosX)
	dirY := float64(character.PosY - enemy.PosY)

	r := math.Sqrt(dirX*dirX + dirY*dirY)

	enemy.PosX = enemy.PosX + enemy.Speed*float32(dirX/r)
	enemy.PosY = enemy.PosY + enemy.Speed*float32(dirY/r)
}
