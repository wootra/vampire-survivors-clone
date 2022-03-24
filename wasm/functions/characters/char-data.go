package characters

import (
	"github.com/wootra/vampire-survivors-clone/wasm/functions/weapon"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewCharacterData() *types.CharacterData {
	return &types.CharacterData{
		PosX:         0,
		PosY:         0,
		Speed:        5,
		Shield:       0,
		FrameIndex:   0,
		FrameOffset:  0,
		Armor:        0,
		Life:         100,
		Weapon:       weapon.CreateAWeapon(types.GUN),
		MovementCode: types.STOP}
}
