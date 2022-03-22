package characters

import (
	"math/rand"

	//to fix the redline, refer .vscode/settings.json in this workspace
	"github.com/wootra/vampire-survivors-clone/wasm/functions/weapon"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewEnemyData(enemyName types.EnemyName) types.EnemyData {
	rand.Seed(100)
	switch enemyName {
	case types.BAT:
		return types.EnemyData{enemyName, -rand.Float32(), -rand.Float32(), 0.5, 0, 0, 100, weapon.CreateAWeapon(types.GUN), types.FLY}
	}
	return types.EnemyData{types.BUG, -rand.Float32(), -rand.Float32(), 0.5, 0, 0, 100, weapon.CreateAWeapon(types.GUN), types.FLY}
}
