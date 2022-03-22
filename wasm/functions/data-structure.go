package functions

type Data struct {
	canvas    Canvas
	character CharacterData
	enemies   []EnemyData
}

func CreateNewData() Data {
	return Data{CreateNewCanvas(), CreateNewCharacterData(), []EnemyData{}}
}
