package functions

import (
	"math"

	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func KeyDown(data *types.Data, keyCode string) {
	switch keyCode {
	case "ArrowDown":
		data.Character.MovementCode.Down = true
		data.Character.MovementCode.Up = false
		break
	case "ArrowUp":
		data.Character.MovementCode.Up = true
		data.Character.MovementCode.Down = false
		break
	case "ArrowLeft":
		data.Character.MovementCode.Left = true
		data.Character.MovementCode.Right = false
		break
	case "ArrowRight":
		data.Character.MovementCode.Right = true
		data.Character.MovementCode.Left = false
		break
	}
	data.Character.FrameOffset++
}

func KeyUp(data *types.Data, keyCode string) {
	switch keyCode {
	case "ArrowDown":
		data.Character.MovementCode.Down = false
		break
	case "ArrowUp":
		data.Character.MovementCode.Up = false
		break
	case "ArrowLeft":
		data.Character.MovementCode.Left = false
		break
	case "ArrowRight":
		data.Character.MovementCode.Right = false
		break
	}
}

func CalculateHeroPos(character *types.CharacterData) {
	speedY := character.Speed
	speedX := character.Speed

	move := character.MovementCode

	if move.Down {
		character.PosY += speedY
		if character.PosY > 50 {
			character.PosY = 50
		}
	} else if move.Up {
		character.PosY -= speedY
		if character.PosY < -50 {
			character.PosY = -50
		}
	}

	if move.Right {
		character.PosX += speedX
		if character.PosX > 50 {
			character.PosX = 50
		}
	} else if move.Left {
		character.PosX -= speedX
		if character.PosX < -50 {
			character.PosX = -50
		}
	}
}

func CalculateEnemyPos(character *types.CharacterData, enemy *types.EnemyData) {

	dirX := float64(character.PosX - enemy.PosX)
	dirY := float64(character.PosY - enemy.PosY)

	r := math.Sqrt(dirX*dirX + dirY*dirY)

	enemy.PosX = enemy.PosX + enemy.Speed*float32(dirX/r)
	enemy.PosY = enemy.PosY + enemy.Speed*float32(dirY/r)
}
