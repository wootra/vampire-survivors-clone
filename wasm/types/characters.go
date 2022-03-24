package types

type MovementType struct {
	Up    bool
	Left  bool
	Right bool
	Down  bool
}

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

type EnemyStatus int

const (
	IDLE    EnemyStatus = 0
	MOVED   EnemyStatus = 1
	BLOCKED EnemyStatus = 2
	DIED    EnemyStatus = 3
	HIT     EnemyStatus = 4
)

type CharacterData struct {
	PosX, PosY, Speed, Shield, Armor, Life float32
	FrameIndex                             int
	FrameOffset                            int
	Weapon                                 Weapon
	MovementCode                           MovementType
}

func (c *CharacterData) Stop() {
	c.MovementCode.Left = false
	c.MovementCode.Up = false
	c.MovementCode.Right = false
	c.MovementCode.Down = false
}

type EnemyData struct {
	CharName                               EnemyName
	PosX, PosY, Speed, Shield, Armor, Life float32
	FrameIndex                             int
	Weapon                                 Weapon
	SwormType                              SwormType
	Status                                 EnemyStatus
}
