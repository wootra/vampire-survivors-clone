package types

import "math/rand"


type SwormType int64
type EnemyType int64
type EnemyStatus int

const (
	WALK  SwormType = 0
	FLY   SwormType = 1
	JUMP  SwormType = 2
	GROUP SwormType = 3
)

const (
	CAT_SOLDIER EnemyType = 0
)

const (
	IDLE    EnemyStatus = 0
	MOVED   EnemyStatus = 1
	BLOCKED EnemyStatus = 2
	DIED    EnemyStatus = 3
	HIT     EnemyStatus = 4
)

type EnemyData struct {
	CharInfo CharacterInfo
	Weapon  []Weapon
	EnemyType                               EnemyType
	SwormType                              SwormType
	Status                                 EnemyStatus
	Direction                              float64
	Level	uint64
}

func getRandomPos() (float64,float64) {
	PosX := rand.Float64()*100 - 50
	PosY := rand.Float64()*100 - 50

	dir := rand.Intn(256) % 8 // probability will be distributed more
	// fmt.Println("enemy is added at:", PosX, PosY, "dir:", dir)
	if dir == 0 { // left-top
		PosX -= 100
		PosY -= 100
	} else if dir == 1 { //left
		PosX -= 100
	} else if dir == 2 { //left-bottom
		PosX -= 100
		PosY += 100
	} else if dir == 3 { //bottom
		PosY += 100
	} else if dir == 4 { //right-bottom
		PosX += 100
		PosY += 100
	} else if dir == 5 { //right
		PosX += 100
	} else if dir == 6 { //right-top
		PosX += 100
		PosY -= 100
	} else if dir == 7 { //top
		PosY -= 100
	}
	return PosX,PosY
}

func CreateAnEnemy (charType EnemyType) *EnemyData {
	enemy:= &EnemyData{}
	enemy.CharInfo.Init(EnemyDetails[charType], 0)
	enemy.CharInfo.PosX, enemy.CharInfo.PosY = getRandomPos()
	return enemy
	
}
