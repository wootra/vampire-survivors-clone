package types

type WeaponType int64

const (
	SWORD WeaponType = 0
	GUN   WeaponType = 1
	BOMB  WeaponType = 2
)

type Weapon struct {
	WeaponType WeaponType
	Look       string
	Name       string
}
