package types

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

type EnemyName string

const (
	BAT EnemyName = "BAT"
	BUG EnemyName = "BUG"
)

type CharacterData struct {
	PosX, PosY, Speed, Shield, Armor, Life float32
	Weapon                                 Weapon
	MovementCode                           MovementType
}

type EnemyData struct {
	CharName                               EnemyName
	PosX, PosY, Speed, Shield, Armor, Life float32
	Weapon                                 Weapon
	SwormType                              SwormType
}
