package types


type HeroName string 

const (
	GoldFish HeroName = "Mr. Gold Fish"
)


type HeroData struct {
	CharInfo CharacterInfo
	MovementCode                           MovementType
	LastMovement                           MovementType
}

func CreateAHero (charType HeroName) *HeroData {
	hero:= &HeroData{}
	hero.CharInfo.Init(CharDetails[charType], 0)
	
	hero.MovementCode =  MovementType{Up: false, Down: false, Left: false, Right: false}
	hero.LastMovement = MovementType{Up: false, Down: false, Left: false, Right: false}
	return hero
	
}

func (c *HeroData) Stop() {
	c.MovementCode.Left = false
	c.MovementCode.Up = false
	c.MovementCode.Right = false
	c.MovementCode.Down = false
}
