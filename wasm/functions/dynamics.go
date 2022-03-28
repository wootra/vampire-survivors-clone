package functions

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/wootra/vampire-survivors-clone/wasm/types"
	"github.com/wootra/vampire-survivors-clone/wasm/utils"
)

func AdjustSpeed(isDiagonal bool, character *types.HeroData) {
	if isDiagonal {
		character.CharInfo.SpeedAdjust = 0.7 // 1/1.414 since it is diagnonal movement
	} else {
		character.CharInfo.SpeedAdjust = 1
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
	data.Character.CharInfo.FrameOffset++
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
		collisionDist:=enemy.CharInfo.Detail.Size + data.Character.CharInfo.Detail.Size
		distX:= (enemy.CharInfo.PosX - data.Character.CharInfo.PosX)
		distY:= (enemy.CharInfo.PosY - data.Character.CharInfo.PosY)
		
		if math.Sqrt(distX*distX + distY*distY) < collisionDist {
			// enemy.CharInfo.PushedByOthers.X += dirX*0.1
			// enemy.CharInfo.PushedByOthers.Y += dirY*0.1
		// if enemy.CharInfo.PosX < data.Character.CharInfo.PosX+types.CHAR_SIZE &&
		// 	enemy.CharInfo.PosX > data.Character.CharInfo.PosX-types.CHAR_SIZE &&
		// 	enemy.CharInfo.PosY < data.Character.CharInfo.PosY+types.CHAR_SIZE &&
		// 	enemy.CharInfo.PosY > data.Character.CharInfo.PosY-types.CHAR_SIZE {
				// for _,weapon := range enemy.Weapon {
				// 	if int(rand.Float32()*100000)%int(1/weapon.Probability) == 0 {
				// 		data.Character.CharInfo.Life -= weapon.Damage * (weapon.Accuracy + (1-weapon.Accuracy)*rand.Float64())
				// 		fmt.Println("character is hit! left life:", data.Character.CharInfo.Life)
				// 	}
				// }
			
			ret = append(ret, id)
		}

	}
	return ret
}

func CalculateHeroPos(data *types.Data, world *types.World, active int) {
	character := data.Character
	speedY := character.CharInfo.Speed * character.CharInfo.SpeedAdjust
	speedX := character.CharInfo.Speed * character.CharInfo.SpeedAdjust

	move := character.MovementCode
	nextX, nextY := character.CharInfo.PosX, character.CharInfo.PosY

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
	nx := int(nextX)
	ny := int(nextY)
	idInPos := world.GetObjects((active+10-1)%10, nx, ny)

	distX := nextX - character.CharInfo.PosX
	distY := nextY - character.CharInfo.PosY
	

	hits := collisionForHero(data, idInPos, distX, distY)

	if len(hits) == 0 {
		character.CharInfo.PosX = nextX //update pos
		character.CharInfo.PosY = nextY
		world.AddObject((active+10-1)%10, nx,ny, character.CharInfo.Id)
	} else {
		nextX = character.CharInfo.PosX + distX*0.1 //slow down the speed because of hits
		nextY = character.CharInfo.PosY + distY*0.1
		nx = int(nextX)
		ny = int(nextY)
		// calculate the room position again
		world.AddObject((active+10-1)%10, nx,ny, character.CharInfo.Id)
		// push the enemies away
		for _, enemyId := range hits {
			data.Enemies[enemyId].CharInfo.PushedByOthers.X -= distX * 0.5
			data.Enemies[enemyId].CharInfo.PushedByOthers.Y -= distY * 0.5
		}
	}

}

func collisionForEnemy(data *types.Data, enemy *types.EnemyData, idsToFind []uint64, dirX, dirY float64) []uint64 {
	ret := []uint64{}

	for _, id := range idsToFind {
		if id == data.Character.CharInfo.Id {
			collisionDist:=enemy.CharInfo.Detail.Size/2 + data.Character.CharInfo.Detail.Size/2;
			distX:= (enemy.CharInfo.PosX - data.Character.CharInfo.PosX)
			distY:= (enemy.CharInfo.PosY - data.Character.CharInfo.PosY)
			dist:=math.Sqrt(distX*distX + distY*distY)
			if dist <= collisionDist {
				fmt.Println("distx", distX, "distY", distY, "<",  collisionDist, "dist", dist)
			// if enemy.CharInfo.PosX < data.Character.CharInfo.PosX+types.CHAR_SIZE &&
			// 	enemy.CharInfo.PosX > data.Character.CharInfo.PosX-types.CHAR_SIZE &&
			// 	enemy.CharInfo.PosY < data.Character.CharInfo.PosY+types.CHAR_SIZE &&
			// 	enemy.CharInfo.PosY > data.Character.CharInfo.PosY-types.CHAR_SIZE {
					for _,weapon := range enemy.CharInfo.Weapon {
						if int(rand.Float32()*100000)%int(1/weapon.Probability) == 0 {
							data.Character.CharInfo.Life -= weapon.Damage * (weapon.Accuracy + (1-weapon.Accuracy)*rand.Float64())
							fmt.Println("character is hit! left life:", data.Character.CharInfo.Life)
						}
					}
				
				fmt.Println("collision with character")
				ret = append(ret, id)
			}
		} else if id != enemy.CharInfo.Id {
			otherEnemy := data.Enemies[id]
			collisionDist:=enemy.CharInfo.Detail.Size + otherEnemy.CharInfo.Detail.Size;
			distX:= enemy.CharInfo.PosX - otherEnemy.CharInfo.PosX
			distY:= enemy.CharInfo.PosY - otherEnemy.CharInfo.PosY
			if math.Sqrt(distX*distX + distY*distY) < collisionDist {
			// if enemy.CharInfo.PosX < otherEnemy.CharInfo.PosX+types.CHAR_SIZE &&
			// 	enemy.CharInfo.PosX > otherEnemy.CharInfo.PosX-types.CHAR_SIZE &&
			// 	enemy.CharInfo.PosY < otherEnemy.CharInfo.PosY+types.CHAR_SIZE &&
			// 	enemy.CharInfo.PosY > otherEnemy.CharInfo.PosY-types.CHAR_SIZE {
				distX := (enemy.CharInfo.PosX - otherEnemy.CharInfo.PosX) / 10
				distY := (enemy.CharInfo.PosY - otherEnemy.CharInfo.PosY) / 10

				enemy.CharInfo.PushedByOthers.X += distX
				enemy.CharInfo.PushedByOthers.Y += distY

				data.Enemies[id].CharInfo.PushedByOthers.X -= distX //push the unit away
				data.Enemies[id].CharInfo.PushedByOthers.Y -= distX
				fmt.Println("collision with other enemy")
				ret = append(ret, id)
			}
		}
	}
	return ret
}

func CalculateEnemyPos(data *types.Data, enemy *types.EnemyData, world *types.World, active int) {
	character := data.Character
	dirX := character.CharInfo.PosX - enemy.CharInfo.PosX
	dirY := character.CharInfo.PosY - enemy.CharInfo.PosY

	r := math.Sqrt(dirX*dirX + dirY*dirY)

	nextX := enemy.CharInfo.PosX + enemy.CharInfo.Speed*dirX/r + enemy.CharInfo.PushedByOthers.X
	nextY := enemy.CharInfo.PosY + enemy.CharInfo.Speed*dirY/r + enemy.CharInfo.PushedByOthers.Y
	nx := int(nextX)
	ny := int(nextY)
	if math.Abs(dirX)>100 || math.Abs(dirY)>100 {
		return //don't even move. they can't see the hero.
	}else if math.Abs(nextX)>50 || math.Abs(nextY)>50 {
		//out side of the screen, but not too far..
		//move, but do not interact each other.
		enemy.CharInfo.PosX = nextX
		enemy.CharInfo.PosY = nextY
		return
	}

	enemy.Direction = utils.Ternary(dirX > 0, 1, -1)

	// fmt.Println("posInWord:", posInWorld, int((nextX+50)/10), int((nextY+50)/10))
	idInPos := world.GetObjects((active+10-1)%10, nx, ny)
	hits := collisionForEnemy(data, enemy, idInPos, dirX, dirY)
	if len(hits) == 0 {
		enemy.CharInfo.PosX = nextX //update pos
		enemy.CharInfo.PosY = nextY
		enemy.CharInfo.PushedByOthers.X = 0 //when nobody is there, force is removed.
		enemy.CharInfo.PushedByOthers.Y = 0
		world.AddObject((active+10-1)%10, nx, ny, enemy.CharInfo.Id)
		enemy.CharInfo.FrameIndex = 0
	} else { //some collided units
		// do not move position
		if enemy.CharInfo.PosX < -50 || enemy.CharInfo.PosY < -50 || enemy.CharInfo.PosX > 50 || enemy.CharInfo.PosY > 50 {
			//out side of the screen, should not affect to each other
			return
		}
	
		nx,ny = int(enemy.CharInfo.PosX), int(enemy.CharInfo.PosY)
		world.AddObject((active+10-1)%10, nx, ny, enemy.CharInfo.Id)
		enemy.CharInfo.FrameIndex = (enemy.CharInfo.FrameIndex)%2 + 1
	}

	// fmt.Println(math.Atan(dirY/dirX), enemy.Direction)
}
