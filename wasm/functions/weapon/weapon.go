package weapon

type WeaponType int64

const (
	SWORD WeaponType = 0
	GUN   WeaponType = 1
	BOMB  WeaponType = 2
)

type WeaponData interface {
	GetType() WeaponType
	GetLook() string
	GetName() string
}
