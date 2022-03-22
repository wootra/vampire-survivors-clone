package functions

import (
	"syscall/js"

	Chars "github.com/wootra/vampire-survivors-clone/wasm/functions/characters"
)

type Data struct {
	canvas    Canvas
	character Chars.CharacterData
	enemies   []Chars.EnemyData
}

func (data *Data) AddAnEnemy(enemy ...Chars.EnemyData) {
	data.enemies = append(data.enemies, enemy...)
}

func (data *Data) AddAnEnemies(enemies []Chars.EnemyData) {
	data.enemies = append(data.enemies, enemies...)
}

func SetCharacterData(data *Data, charData Chars.CharacterData) {
	data.character = charData
}

func SetCharacterDataJs(data *Data, this js.Value, args []js.Value) interface{} {
	return ""
}

func CreateNewData() Data {
	return Data{CreateNewCanvas(), Chars.CreateNewCharacterData(), []Chars.EnemyData{}}
}
