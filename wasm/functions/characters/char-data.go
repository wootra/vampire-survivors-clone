package characters

import (
	"github.com/wootra/vampire-survivors-clone/wasm/functions/weapon"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewCharacterData() types.CharacterData {
	return types.CharacterData{-1, -1, 0.5, 0, 0, 100, weapon.CreateAWeapon(types.GUN), types.STOP}
}
