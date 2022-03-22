package characters

import (
	"github.com/wootra/vampire-survivors-clone/wasm/functions/weapon"
)

type MovementType int64

const (
	STOP  MovementType = 0
	LEFT  MovementType = 1
	UP    MovementType = 2
	RIGHT MovementType = 3
	DOWN  MovementType = 4
)

type SwormType int64

const (
	WALK  SwormType = 0
	FLY   SwormType = 1
	JUMP  SwormType = 2
	GROUP SwormType = 3
)

type CharacterData struct {
	PosX, PosY, Speed, Shield, Armor, Life float32
	WeaponType                             weapon.WeaponType
	WeaponInfo                             weapon.WeaponData
	MovementCode                           MovementType
}

type EnemyData struct {
	charName                               EnemyName
	posX, posY, Speed, Shield, Armor, Life float32
	WeaponType                             weapon.WeaponType
	WeaponInfo                             weapon.WeaponData
	SwormType                              SwormType
}

func CreateNewCharacterData() CharacterData {
	return CharacterData{-1, -1, 0.5, 0, 0, 100, weapon.GUN, weapon.CreateAGun(), STOP}
}
