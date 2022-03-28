package types

type WeaponType int64

const (
	HIT_BY_BODY WeaponType = 0
	BUBBLE_SHOT   WeaponType = 1
	EXPLODE  WeaponType = 2
)

type WeaponDetail struct {
	WeaponType  WeaponType
	ImageKey        string
	Name        string
	Damage      float64
	Range 		float64
	Probability float64 // if 0.3, when random(100) < 30 ==> hit
	Accuracy    float64 // if 0.6, RealDamage = Damage * (0.6 + random()*0.4)
	MaxLevel	uint
}

var WeaponDetails = map[WeaponType]WeaponDetail {
	HIT_BY_BODY: {WeaponType: HIT_BY_BODY, ImageKey: "hit-image-set", Name: "Hit by body!", MaxLevel: 10, Range:1},
	BUBBLE_SHOT: {WeaponType: BUBBLE_SHOT, ImageKey: "bubble-image-set", Name: "Hit by body!", MaxLevel: 10, Range:1},
	EXPLODE: {WeaponType: EXPLODE, ImageKey: "explode-image-set", Name: "Hit by body!", MaxLevel: 10, Range:1},

}

type Weapon struct {
	Detail  WeaponDetail //should never be changed once it is saved
	Damage      float64
	Range 	float64
	Probability float64 // if 0.3, when random(100) < 30 ==> hit
	Accuracy    float64 // if 0.6, RealDamage = Damage * (0.6 + random()*0.4)
	Level	uint
}

var idPool uint64 = 0

func GetNewId() uint64 {
	idPool++
	return idPool
}

func produceAWeapon(
	wType WeaponType,
	look, name string,
	damage, probability, accuracy float64,
	maxLevel uint,
) Weapon {
	wDetail := WeaponDetails[wType]
	return Weapon{
		Detail: wDetail,
		Damage: wDetail.Damage, 
		Probability: wDetail.Probability,
		Range: wDetail.Range,
		Accuracy: wDetail.Accuracy,
		Level: 0,
	}
}

func CreateAWeapon(weaponType WeaponType) Weapon {
	var weapon WeaponType = weaponType

	//later, more info will be added
	switch weapon {
	case EXPLODE:
		return produceAWeapon(weaponType, "a EXPLODE", "my-EXPLODE", 2, 0.3, 0.5, 2)
	case BUBBLE_SHOT:
		return produceAWeapon(weaponType, "a BUBBLE_SHOT", "my-BUBBLE_SHOT", 2, 0.3, 0.5, 5)
	case HIT_BY_BODY:
		return produceAWeapon(weaponType, "a HIT_BY_BODY", "my-HIT_BY_BODY", 2, 0.3, 0.5, 5)
	}
	return produceAWeapon(EXPLODE, "a EXPLODE", "my-EXPLODE", 2, 0.3, 0.5, 5)
}
