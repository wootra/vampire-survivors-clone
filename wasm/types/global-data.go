package types

//to fix the redline, refer .vscode/settings.json in this workspace

type Data struct {
	Canvas    Canvas
	Character CharacterData
	Enemies   []EnemyData
}

func (data *Data) AddAnEnemy(enemy ...EnemyData) {
	data.Enemies = append(data.Enemies, enemy...)
}

func (data *Data) AddAnEnemies(enemies []EnemyData) {
	data.Enemies = append(data.Enemies, enemies...)
}
