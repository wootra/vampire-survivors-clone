package types

type MovementType struct {
	Up    bool
	Right bool
	Down  bool
	Left  bool
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

type PushedForce struct {
	X float64
	Y float64
}

type CharacterData struct {
	Id                                     uint64
	PosX, PosY, Speed, Shield, Armor, Life float64
	SpeedAdjust                            float64
	FrameIndex                             int
	FrameOffset                            int
	Weapon                                 Weapon
	MovementCode                           MovementType
	LastMovement                           MovementType
	ImageKey                               string
	PushedByOthers                         PushedForce
}

func (c *CharacterData) Stop() {
	c.MovementCode.Left = false
	c.MovementCode.Up = false
	c.MovementCode.Right = false
	c.MovementCode.Down = false
}

type EnemyData struct {
	Id                                     uint64
	CharName                               EnemyName
	PosX, PosY, Speed, Shield, Armor, Life float64
	SpeedAdjustX                           float64
	SpeedAdjustY                           float64
	FrameIndex                             int
	Weapon                                 Weapon
	SwormType                              SwormType
	Status                                 EnemyStatus
	ImageKey                               string
	Direction                              float64
	PushedByOthers                         PushedForce
}

const CHAR_SIZE = 4
