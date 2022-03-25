package functions

import (
	"math"

	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func AdjustSpeed(isDiagonal bool, character *types.CharacterData) {
	if isDiagonal {
		character.SpeedAdjust = 0.7 // 1/1.414 since it is diagnonal movement
	} else {
		character.SpeedAdjust = 1
	}
}

func KeyDown(data *types.Data, keyCode string) {
	movement := &data.Character.MovementCode
	lastMove := &data.Character.LastMovement
	switch keyCode {
	case "ArrowDown":
		movement.Down = true
		movement.Up = false
		lastMove.Down = true
		lastMove.Up = false
		AdjustSpeed(movement.Left || movement.Right, data.Character)
		break
	case "ArrowUp":
		movement.Up = true
		movement.Down = false
		lastMove.Up = true
		lastMove.Down = false
		AdjustSpeed(movement.Left || movement.Right, data.Character)
		break
	case "ArrowLeft":
		movement.Left = true
		movement.Right = false
		lastMove.Left = true
		lastMove.Right = false
		AdjustSpeed(movement.Up || movement.Down, data.Character)
		break
	case "ArrowRight":
		movement.Right = true
		movement.Left = false
		lastMove.Right = true
		lastMove.Left = false
		AdjustSpeed(movement.Up || movement.Down, data.Character)
		break
	}

	if !movement.Right {
		lastMove.Right = false
	}
	if !movement.Left {
		lastMove.Left = false
	}
	if !movement.Up {
		lastMove.Up = false
	}
	if !movement.Down {
		lastMove.Down = false
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
	speedY := character.Speed * character.SpeedAdjust
	speedX := character.Speed * character.SpeedAdjust

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

func ternary(test bool, a, b float32) float32 {
	if test {
		return a
	} else {
		return b
	}
}

func CalculateEnemyPos(character *types.CharacterData, enemy *types.EnemyData) {

	dirX := float64(character.PosX - enemy.PosX)
	dirY := float64(character.PosY - enemy.PosY)

	r := math.Sqrt(dirX*dirX + dirY*dirY)

	enemy.PosX = enemy.PosX + enemy.Speed*float32(dirX/r)
	enemy.PosY = enemy.PosY + enemy.Speed*float32(dirY/r)
	if dirX > 0 {
		enemy.Direction = 1
	}
	enemy.Direction = ternary(dirX > 0, 1, -1)
	// fmt.Println(math.Atan(dirY/dirX), enemy.Direction)
}
