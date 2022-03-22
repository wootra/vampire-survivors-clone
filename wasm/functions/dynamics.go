package functions

import (
	"fmt"
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
		if character.PosY > 100 {
			character.PosY = 100
		}
		break
	case types.UP:
		character.PosY -= speedY
		if character.PosY < 0 {
			character.PosY = 0
		}
		break
	case types.RIGHT:
		character.PosX += speedX
		if character.PosX > 100 {
			character.PosX = 100
		}
		break
	case types.LEFT:
		character.PosX -= speedX
		if character.PosX < 0 {
			character.PosX = 0
		}
		break
	}
}

func CalculateEnemyPos(character types.CharacterData, enemy *types.EnemyData) {

	dirX := character.PosX - enemy.PosX
	dirY := character.PosY - enemy.PosY

	angle := math.Atan(float64(dirY / dirX))

	enemy.PosX += enemy.Speed * float32(math.Cos(angle))
	enemy.PosY += enemy.Speed * float32(math.Sin(angle))

	fmt.Println(enemy.CharName, enemy.PosX, enemy.PosY, dirX, dirY, angle, enemy.Speed,
		(enemy.Speed * float32(math.Cos(angle))), (enemy.Speed * float32(math.Sin(angle))))
}
