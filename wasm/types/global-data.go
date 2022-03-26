package types

import "syscall/js"

//to fix the redline, refer .vscode/settings.json in this workspace

type Data struct {
	Canvas        *Canvas
	Character     *CharacterData
	GlueFunctions *js.Value
	Enemies       map[uint64]*EnemyData
}

const WORLD_WIDTH = 10
const WORLD_HEIGHT = 10

const WORLD_SIZE = WORLD_WIDTH * WORLD_HEIGHT // 10x10

type World struct {
	Pt             [10][WORLD_SIZE][]uint64 //10 x grid of id arrays.
	Active         int                      // current active grid
	Apocalypse     bool                     //if true, end the game
	MainChannel    chan bool
	BufferToDelete chan int
}

func (data *Data) AddAnEnemy(enemies ...*EnemyData) {
	for _, enemy := range enemies {
		data.Enemies[enemy.Id] = enemy
	}
}

func (data *Data) AddAnEnemies(enemies []*EnemyData) {
	for _, enemy := range enemies {
		data.Enemies[enemy.Id] = enemy
	}
}
