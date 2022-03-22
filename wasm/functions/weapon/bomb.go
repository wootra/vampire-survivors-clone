package weapon

type Bomb struct {
	weaponType WeaponType
	look       string
	name       string
}

func (g Bomb) GetType() WeaponType {
	return g.weaponType
}

func (g Bomb) GetLook() string {
	return g.look
}

func (g Bomb) GetName() string {
	return g.look
}

func CreateABomb() WeaponData {
	Bomb := Bomb{BOMB, "a Bomb", "my-Bomb"} // could be many varient
	return WeaponData(Bomb)
}
