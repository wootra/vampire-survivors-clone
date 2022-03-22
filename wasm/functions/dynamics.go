package functions

type MovementType int64

const (
	STOP  MovementType = 0
	LEFT  MovementType = 1
	UP    MovementType = 2
	RIGHT MovementType = 3
	DOWN  MovementType = 4
)

func CalculateInATick(data *Data) interface{} {
	speedY := data.character.speed
	speedX := data.character.speed

	switch data.character.movementCode {
	case DOWN:
		data.character.posY += speedY
		if data.character.posY > 100 {
			data.character.posY = 100
		}
		break
	case UP:
		data.character.posY -= speedY
		if data.character.posY < 0 {
			data.character.posY = 0
		}
		break
	case RIGHT:
		data.character.posX += speedX
		if data.character.posX > 100 {
			data.character.posX = 100
		}
		break
	case LEFT:
		data.character.posX -= speedX
		if data.character.posX < 0 {
			data.character.posX = 0
		}
		break
	}
	return ""
}
