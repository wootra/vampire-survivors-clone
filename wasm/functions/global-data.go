package functions

import (
	"syscall/js"

	"github.com/wootra/vampire-survivors-clone/wasm/functions/characters"
	"github.com/wootra/vampire-survivors-clone/wasm/types"
)

func SetCharacterData(data *types.Data, charData types.CharacterData) {
	data.Character = charData
}

func SetCharacterDataJs(data *types.Data, this js.Value, args []js.Value) interface{} {
	return ""
}

func CreateNewData() types.Data {
	return types.Data{CreateNewCanvas(), characters.CreateNewCharacterData(), []types.EnemyData{}}
}
