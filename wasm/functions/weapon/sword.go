package weapon

type Sword struct {
	weaponType WeaponType
	look       string
	name       string
}

func (g Sword) GetType() WeaponType {
	return g.weaponType
}

func (g Sword) GetLook() string {
	return g.look
}

func (g Sword) GetName() string {
	return g.look
}

func CreateASword() WeaponData {
	Sword := Sword{SWORD, "a Sword", "my-Sword"} // could be many varient
	return WeaponData(Sword)
}
