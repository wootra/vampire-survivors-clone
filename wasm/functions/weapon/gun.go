package weapon

type Gun struct {
	weaponType WeaponType
	look       string
	name       string
}

func (g Gun) GetType() WeaponType {
	return g.weaponType
}

func (g Gun) GetLook() string {
	return g.look
}

func (g Gun) GetName() string {
	return g.look
}

func CreateAGun() WeaponData {
	gun := Gun{GUN, "a gun", "my-gun"} // could be many varient
	return WeaponData(gun)
}
