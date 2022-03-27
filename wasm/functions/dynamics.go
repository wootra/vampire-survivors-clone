package functions

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/wootra/vampire-survivors-clone/wasm/types"
	"github.com/wootra/vampire-survivors-clone/wasm/utils"
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

func collisionForHero(data *types.Data, idsToFind []uint64, dirX, dirY float64) []uint64 {
	ret := []uint64{}

	for _, id := range idsToFind {
		enemy := data.Enemies[id]
		if enemy.PosX < data.Character.PosX+types.CHAR_SIZE &&
			enemy.PosX > data.Character.PosX-types.CHAR_SIZE &&
			enemy.PosY < data.Character.PosY+types.CHAR_SIZE &&
			enemy.PosY > data.Character.PosY-types.CHAR_SIZE {
			if int(rand.Float32()*100000)%int(1/enemy.Weapon.Probability) == 0 {
				data.Character.Life -= enemy.Weapon.Damage * (enemy.Weapon.Accuracy + (1-enemy.Weapon.Accuracy)*rand.Float64())
				fmt.Println("character is hit! left life:", data.Character.Life)
			}
			ret = append(ret, id)
		}

	}
	return ret
}

func CalculateHeroPos(data *types.Data, world *types.World, active int) {
	character := data.Character
	speedY := character.Speed * character.SpeedAdjust
	speedX := character.Speed * character.SpeedAdjust

	move := character.MovementCode
	nextX, nextY := character.PosX, character.PosY

	if move.Down {
		nextY += speedY
		if nextY > 50 {
			nextY = 50
		}
	} else if move.Up {
		nextY -= speedY
		if nextY < -50 {
			nextY = -50
		}
	}

	if move.Right {
		nextX += speedX
		if nextX > 50 {
			nextX = 50
		}
	} else if move.Left {
		nextX -= speedX
		if nextX < -50 {
			nextX = -50
		}
	}
	posInWorld := int((nextX+50)/10) + int((nextY+50)/10)*types.WORLD_WIDTH

	idInPos := world.Pt[(active+10-1)%10][posInWorld]

	distX := nextX - character.PosX
	distY := nextY - character.PosY

	hits := collisionForHero(data, idInPos, distX, distY)

	if len(hits) == 0 {
		character.PosX = nextX //update pos
		character.PosY = nextY
		world.Pt[(active+10-1)%10][posInWorld] = append(idInPos, character.Id)
	} else {
		nextX = character.PosX + distX*0.1 //slow down the speed because of hits
		nextY = character.PosY + distY*0.1
		// calculate the room position again
		posInWorld := int((nextX+50)/10) + int((nextY+50)/10)*types.WORLD_WIDTH
		idInPos = world.Pt[(active+10-1)%10][posInWorld]
		world.Pt[(active+10-1)%10][posInWorld] = append(idInPos, character.Id)
		// push the enemies away
		for _, enemyId := range hits {
			data.Enemies[enemyId].PushedByOthers.X -= distX * 0.5
			data.Enemies[enemyId].PushedByOthers.Y -= distY * 0.5
		}
	}

}

func collisionForEnemy(data *types.Data, enemy *types.EnemyData, idsToFind []uint64, dirX, dirY float64) []uint64 {
	ret := []uint64{}

	for _, id := range idsToFind {
		if id == data.Character.Id {
			if enemy.PosX < data.Character.PosX+types.CHAR_SIZE &&
				enemy.PosX > data.Character.PosX-types.CHAR_SIZE &&
				enemy.PosY < data.Character.PosY+types.CHAR_SIZE &&
				enemy.PosY > data.Character.PosY-types.CHAR_SIZE {
				if int(rand.Float32()*100000)%int(1/enemy.Weapon.Probability) == 0 {
					data.Character.Life -= enemy.Weapon.Damage * (enemy.Weapon.Accuracy + (1-enemy.Weapon.Accuracy)*rand.Float64())
					fmt.Println("character is hit! left life:", data.Character.Life)
				}
				fmt.Println("collision with character")
				ret = append(ret, id)
			}
		} else if id != enemy.Id {
			otherEnemy := data.Enemies[id]
			if enemy.PosX < otherEnemy.PosX+types.CHAR_SIZE &&
				enemy.PosX > otherEnemy.PosX-types.CHAR_SIZE &&
				enemy.PosY < otherEnemy.PosY+types.CHAR_SIZE &&
				enemy.PosY > otherEnemy.PosY-types.CHAR_SIZE {
				distX := (enemy.PosX - otherEnemy.PosX) / 10
				distY := (enemy.PosY - otherEnemy.PosY) / 10

				enemy.PushedByOthers.X += distX
				enemy.PushedByOthers.Y += distY

				data.Enemies[id].PushedByOthers.X -= distX //push the unit away
				data.Enemies[id].PushedByOthers.Y -= distX
				fmt.Println("collision with other enemy")
				ret = append(ret, id)
			}
		}
	}
	return ret
}

func CalculateEnemyPos(data *types.Data, enemy *types.EnemyData, world *types.World, active int) {
	character := data.Character
	dirX := character.PosX - enemy.PosX
	dirY := character.PosY - enemy.PosY

	r := math.Sqrt(dirX*dirX + dirY*dirY)

	nextX := enemy.PosX + enemy.Speed*dirX/r + enemy.PushedByOthers.X
	nextY := enemy.PosY + enemy.Speed*dirY/r + enemy.PushedByOthers.Y

	if nextX < -50 || nextY < -50 || nextX > 50 || nextY > 50 {
		//out side of the screen
		enemy.PosX = nextX
		enemy.PosY = nextY
		return
	}

	enemy.Direction = utils.Ternary(dirX > 0, 1, -1)

	posInWorld := int((nextX+50)/10) + int((nextY+50)/10)*types.WORLD_WIDTH
	fmt.Println("posInWord:", posInWorld, int((nextX+50)/10), int((nextY+50)/10))
	idInPos := world.Pt[(active+10-1)%10][posInWorld]
	hits := collisionForEnemy(data, enemy, idInPos, dirX, dirY)
	if len(hits) == 0 {
		enemy.PosX = nextX //update pos
		enemy.PosY = nextY
		enemy.PushedByOthers.X = 0 //when nobody is there, force is removed.
		enemy.PushedByOthers.Y = 0
		world.Pt[(active+10-1)%10][posInWorld] = append(idInPos, enemy.Id)
		enemy.FrameIndex = 0
	} else { //some collided units
		// do not move position
		if enemy.PosX < -50 || enemy.PosY < -50 || enemy.PosX > 50 || enemy.PosY > 50 {
			//out side of the screen
			return
		}
	
		posInWorld := int((enemy.PosX+50)/10) + int((enemy.PosY+50)/10)*types.WORLD_WIDTH
		idInPos = world.Pt[(active+10-1)%10][posInWorld]
		world.Pt[(active+10-1)%10][posInWorld] = append(idInPos, enemy.Id)
		enemy.FrameIndex = (enemy.FrameIndex)%2 + 1
	}

	// fmt.Println(math.Atan(dirY/dirX), enemy.Direction)
}
