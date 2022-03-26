package types

type WeaponType int64

const (
	SWORD WeaponType = 0
	GUN   WeaponType = 1
	BOMB  WeaponType = 2
)

type Weapon struct {
	WeaponType  WeaponType
	Look        string
	Name        string
	Damage      float64
	Probability float64 // if 0.3, when random(100) < 30 ==> hit
	Accuracy    float64 // if 0.6, RealDamage = Damage * (0.6 + random()*0.4)
}

var idPool uint64 = 0

func GetNewId() uint64 {
	idPool++
	return idPool
}
