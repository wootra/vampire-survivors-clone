package types

type MovementType struct {
	Up    bool
	Right bool
	Down  bool
	Left  bool
}

type PushedForce struct {
	X float64
	Y float64
}

type CharacterDetail struct {
	Name 	HeroName
	Speed, Shield, Armor, Life float64 //init values
	ImageKey                               string
	Weapon                                 []Weapon //only init weapons
}

type CharacterInfo struct {
	Id                                     uint64
	Detail CharacterDetail
	Level	int
	PosX, PosY, Speed, Shield, Armor, Life float64
	SpeedAdjust                            float64
	FrameIndex                             int
	FrameOffset                            int
	Weapon                                 []Weapon
	PushedByOthers                         PushedForce
	SpeedAdjustX                           float64
	SpeedAdjustY                           float64
	CurrSpeedX	float64  // based on direction, split x and y factor of the speed to figure out the bounce strength
	CurrSpeedY	float64
}

func (c *CharacterInfo) setId(id uint64) {
	c.Id = id
}

func (c *CharacterInfo) setPos(x,y float64) {
	c.PosX = x
	c.PosY = y
}

func (c *CharacterInfo) Init(detail CharacterDetail, level int) *CharacterInfo {
	c.Id = GetNewId()
	
	c.Detail = detail
	c.Level = 0

	c.Weapon = append(c.Detail.Weapon, CreateAWeapon(BUBBLE_SHOT))
	c.PosX, c.PosY = 0,0
	c.Speed, c.Shield, c.Armor, c.Life = detail.Speed, detail.Shield, detail.Armor, detail.Life
	c.PushedByOthers = PushedForce{X:0, Y:0}
	c.SpeedAdjust, c.SpeedAdjustX, c.SpeedAdjustY = 1, 0,0
	c.CurrSpeedX, c.CurrSpeedY = 0,0
	return c
}




const CHAR_SIZE = 4
