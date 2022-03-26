package functions

import (
	"github.com/wootra/vampire-survivors-clone/wasm/functions/characters"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewData() *types.Data {
	return &types.Data{
		Canvas:        CreateNewCanvas(),
		Character:     &types.CharacterData{},
		Enemies:       map[uint64]*types.EnemyData{},
		GlueFunctions: nil,
	}
}

func InitCharacters(data *types.Data) {
	data.Character = characters.CreateNewCharacterData()
	for i := 0; i < 100; i++ {
		data.AddAnEnemy(characters.CreateNewEnemyData(data, types.BAT))
		// data.AddAnEnemy(characters.CreateNewEnemyData(data, types.BAT))
		// data.AddAnEnemy(characters.CreateNewEnemyData(data, types.BAT))
		// data.AddAnEnemy(characters.CreateNewEnemyData(data, types.BAT))
	}

}
