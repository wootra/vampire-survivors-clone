package functions

import (
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func CreateNewData() *types.Data {
	return &types.Data{
		Canvas:    CreateNewCanvas(),
		Character: &types.CharacterData{},
		Enemies:   []*types.EnemyData{}}
}
