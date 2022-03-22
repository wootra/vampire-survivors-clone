package functions

import (
	"syscall/js"

	Chars "github.com/wootra/vampire-survivors-clone/wasm/functions/characters"
)

type Data struct {
	Canvas    Canvas
	Character Chars.CharacterData
	Enemies   []Chars.EnemyData
}

func (data *Data) AddAnEnemy(enemy ...Chars.EnemyData) {
	data.Enemies = append(data.Enemies, enemy...)
}

func (data *Data) AddAnEnemies(enemies []Chars.EnemyData) {
	data.Enemies = append(data.Enemies, enemies...)
}

func SetCharacterData(data *Data, charData Chars.CharacterData) {
	data.Character = charData
}

func SetCharacterDataJs(data *Data, this js.Value, args []js.Value) interface{} {
	return ""
}

func CreateNewData() Data {
	return Data{CreateNewCanvas(), Chars.CreateNewCharacterData(), []Chars.EnemyData{}}
}
