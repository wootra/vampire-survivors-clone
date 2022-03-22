package characters

import (
	"math/rand"
	//to fix the redline, refer .vscode/settings.json in this workspace
	"github.com/wootra/vampire-survivors-clone/wasm/functions/weapon"
)

type EnemyName string

const (
	BAT EnemyName = "BAT"
)

func CreateNewEnemyData(enemyName EnemyName) EnemyData {
	rand.Seed(100)
	switch enemyName {
	case BAT:
		return EnemyData{enemyName, -rand.Float32(), -rand.Float32(), 0.5, 0, 0, 100, weapon.GUN, weapon.CreateAGun(), FLY}
	}
	return EnemyData{enemyName, -rand.Float32(), -rand.Float32(), 0.5, 0, 0, 100, weapon.GUN, weapon.CreateAGun(), FLY}
}
