package weapon

import (
	Types "github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateAWeapon(weaponType Types.WeaponType) Types.Weapon {
	var weapon Types.WeaponType = weaponType

	//later, more info will be added
	switch weapon {
	case Types.BOMB:
		return Types.Weapon{weaponType, "a Bomb", "my-Bomb"}
	case Types.GUN:
		return Types.Weapon{weaponType, "a Gun", "my-gun"}
	case Types.SWORD:
		return Types.Weapon{weaponType, "a Sword", "my-Sword"}

	}
	return Types.Weapon{Types.BOMB, "a Bomb", "my-Bomb"}
}
