package characters

import (
	"fmt"
	"math/rand"

	//to fix the redline, refer .vscode/settings.json in this workspace
	"github.com/wootra/vampire-survivors-clone/wasm/functions/weapon"
	types "github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewEnemyData(data *types.Data, CharName types.EnemyName) *types.EnemyData {

	PosX := rand.Float32()*100 - 50
	PosY := rand.Float32()*100 - 50

	dir := rand.Intn(256) % 8 // probability will be distributed more
	fmt.Println("enemy is added at:", PosX, PosY, "dir:", dir)
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
	var Speed, Armor, Shield, Life float32
	var SwormType types.SwormType
	var Weapon types.Weapon
	var Status = types.IDLE

	switch CharName {
	case types.BAT:
		Speed = float32(0.5)
		Armor = float32(0)
		Shield = float32(0)
		Life = float32(100)
		SwormType = types.FLY
		Weapon = weapon.CreateAWeapon(types.GUN)
	default:
		Speed = float32(0.5)
		Armor = float32(0)
		Shield = float32(0)
		Life = float32(100)
		SwormType = types.FLY
		Weapon = weapon.CreateAWeapon(types.GUN)
	}
	return &types.EnemyData{
		CharName:     CharName,
		PosX:         PosX,
		PosY:         PosY,
		Speed:        Speed,
		Shield:       Shield,
		FrameIndex:   0,
		Armor:        Armor,
		Life:         Life,
		Weapon:       Weapon,
		SwormType:    SwormType,
		Status:       Status,
		SpeedAdjustX: 1,
		SpeedAdjustY: 1,
		Direction:    0,
	}
}
