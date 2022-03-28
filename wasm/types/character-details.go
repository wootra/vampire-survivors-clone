package types

var CharDetails = map[HeroName]CharacterDetail{
	GoldFish: {
		Name: GoldFish,
		Speed: 2, 
		Shield: 0, 
		Armor:0, 
		Life: 50, //init values
    	ImageKey: "fish",
    	Weapon: []Weapon{
			CreateAWeapon(BUBBLE_SHOT),
		},
		Size: CHAR_SIZE,
	},
}

var EnemyDetails = map[EnemyType]CharacterDetail{
	CAT_SOLDIER: {
		Name: "Cat Soldier",
		Speed: 0.5, 
		Shield: 0, 
		Armor:0, 
		Life: 50, //init values
    	ImageKey: "cats",
    	Weapon: []Weapon{
			CreateAWeapon(HIT_BY_BODY),
		},
		Size: CHAR_SIZE,
	},
}
