package weapon

import (
	Types "github.com/wootra/vampire-survivors-clone/wasm/types"
)

func produceAWeapon(
	WeaponType Types.WeaponType,
	Look, Name string,
	Damage, Probability, Accuracy float64,
) Types.Weapon {
	return Types.Weapon{
		WeaponType: WeaponType,
		Look:       Look, Name: Name,
		Damage: Damage, Probability: Probability, Accuracy: Accuracy,
	}
}

func CreateAWeapon(weaponType Types.WeaponType) Types.Weapon {
	var weapon Types.WeaponType = weaponType

	//later, more info will be added
	switch weapon {
	case Types.BOMB:
		return produceAWeapon(weaponType, "a Bomb", "my-Bomb", 2, 0.3, 0.5)
	case Types.GUN:
		return produceAWeapon(weaponType, "a Gun", "my-gun", 2, 0.3, 0.5)
	case Types.SWORD:
		return produceAWeapon(weaponType, "a Sword", "my-Sword", 2, 0.3, 0.5)
	}
	return produceAWeapon(Types.BOMB, "a Bomb", "my-Bomb", 2, 0.3, 0.5)
}
