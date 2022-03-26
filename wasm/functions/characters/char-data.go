package characters

import (
	"github.com/wootra/vampire-survivors-clone/wasm/functions/weapon"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewCharacterData() *types.CharacterData {
	return &types.CharacterData{
		Id:           types.GetNewId(),
		PosX:         0,
		PosY:         0,
		Speed:        2,
		SpeedAdjust:  1,
		Shield:       0,
		FrameIndex:   0,
		FrameOffset:  0,
		Armor:        0,
		Life:         100,
		Weapon:       weapon.CreateAWeapon(types.GUN),
		MovementCode: types.MovementType{Up: false, Down: false, Left: false, Right: false},
		LastMovement: types.MovementType{Up: false, Down: false, Left: false, Right: false},
		ImageKey:     "fish",
	}
}
