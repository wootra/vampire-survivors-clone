package functions

import (
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewData() *types.Data {
	return &types.Data{
		Canvas:        CreateNewCanvas(),
		Character:     &types.HeroData{},
		Enemies:       map[uint64]*types.EnemyData{},
		GlueFunctions: nil,
	}
}

func InitCharacters(data *types.Data) {
	data.Character = types.CreateAHero(types.GoldFish)
	for i := 0; i < 100; i++ {
		data.AddAnEnemy(types.CreateAnEnemy(types.CAT_SOLDIER))
		// data.AddAnEnemy(characters.CreateNewEnemyData(data, types.CAT_SOLDIER))
		// data.AddAnEnemy(characters.CreateNewEnemyData(data, types.CAT_SOLDIER))
		// data.AddAnEnemy(characters.CreateNewEnemyData(data, types.CAT_SOLDIER))
	}

}
