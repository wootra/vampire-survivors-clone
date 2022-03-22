package functions

import (
	"syscall/js" //to fix the redline, refer .vscode/settings.json in this workspace

	"github.com/wootra/vampire-survivors-clone/wasm/functions/weapon"
)

type CharacterData struct {
	posX, posY, speed, shield, armor, life float32
	weaponType                             weapon.WeaponType
	weaponInfo                             weapon.WeaponData
	movementCode                           MovementType
}

type EnemyData struct {
	posX, posY, speed, shield, armor, life float32
	weaponType                             weapon.WeaponType
}

func SetCharacterData(data *Data, charData CharacterData) {
	data.character = charData
}

func SetCharacterDataJs(data *Data, this js.Value, args []js.Value) interface{} {
	return ""
}

func CreateNewCharacterData() CharacterData {
	return CharacterData{-1, -1, 0.05, 0, 0, 100, weapon.GUN, weapon.CreateAGun(), STOP}
}
