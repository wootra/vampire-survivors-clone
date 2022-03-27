package types

var CharDetails = map[HeroName]CharacterDetail{
	GoldFish: CharacterDetail{
		Name: GoldFish,
		Speed: 2, 
		Shield: 0, 
		Armor:0, 
		Life: 50, //init values
    	ImageKey: "fish",
    	Weapon: []Weapon{
			CreateAWeapon(BUBBLE_SHOT),
		},
	},
}

var EnemyDetails = map[EnemyType]CharacterDetail{
	CAT_SOLDIER: CharacterDetail{
		Name: "Cat Soldier",
		Speed: 1, 
		Shield: 0, 
		Armor:0, 
		Life: 50, //init values
    	ImageKey: "cats",
    	Weapon: []Weapon{
			CreateAWeapon(HIT_BY_BODY),
		},
	},
}
